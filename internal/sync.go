package internal

import (
	"fmt"
	"log"
	"slices"
	"strings"

	"github.com/fatih/color"
)

// SchemaSync 配置文件
type SchemaSync struct {
	Config   *Config
	SourceDb *MyDb
	DestDb   *MyDb
}

// NewSchemaSync 对一个配置进行同步
func NewSchemaSync(config *Config) *SchemaSync {
	s := new(SchemaSync)
	s.Config = config
	s.SourceDb = NewMyDb(config.SourceDSN, dbTypeSource)
	s.DestDb = NewMyDb(config.DestDSN, dbTypeDest)
	return s
}

// GetNewTableNames 获取所有新增加的表名
func (sc *SchemaSync) GetNewTableNames() []string {
	sourceTables := sc.SourceDb.GetTableNames()
	destTables := sc.DestDb.GetTableNames()

	var newTables []string

	for _, name := range sourceTables {
		if !inStringSlice(name, destTables) {
			newTables = append(newTables, name)
		}
	}
	return newTables
}

// AllDBTables 合并源数据库和目标数据库的表名
func (sc *SchemaSync) AllDBTables() []string {
	sourceTables := sc.SourceDb.GetTableNames()
	destTables := sc.DestDb.GetTableNames()
	tables := slices.Clone(destTables)
	for _, name := range sourceTables {
		if !inStringSlice(name, tables) {
			tables = append(tables, name)
		}
	}
	return tables
}

// RemoveTableSchemaConfig 删除表创建引擎信息，编码信息，分区信息，已修复同步表结构遇到分区表异常退出问题，
// 对于分区表，只会同步字段，索引，主键，外键的变更
func RemoveTableSchemaConfig(schema string) string {
	return strings.Split(schema, "ENGINE")[0]
}

func (sc *SchemaSync) getAlterDataByTable(table string, cfg *Config) *TableAlterData {
	sSchema := sc.SourceDb.GetTableSchema(table)
	dSchema := sc.DestDb.GetTableSchema(table)
	return sc.getAlterDataBySchema(table, sSchema, dSchema, cfg)
}

func (sc *SchemaSync) getAlterDataBySchema(table string, sSchema string, dSchema string, cfg *Config) *TableAlterData {
	alter := new(TableAlterData)
	alter.Table = table
	alter.Type = alterTypeNo

	// Try to get structured field information from INFORMATION_SCHEMA.COLUMNS
	sourceFields, sourceFieldsErr := sc.SourceDb.TableFieldsFromInformationSchema(table)
	destFields, destFieldsErr := sc.DestDb.TableFieldsFromInformationSchema(table)

	// If we can get structured field information from both databases, use it for precise comparison
	if sourceFieldsErr == nil && destFieldsErr == nil {
		log.Printf("[Debug] Using structured field comparison for table %q", table)
		alter.SchemaDiff = NewSchemaDiffWithFieldInfos(table, RemoveTableSchemaConfig(sSchema), RemoveTableSchemaConfig(dSchema), sourceFields, destFields)
	} else {
		// Fallback to legacy text-based comparison
		if sourceFieldsErr != nil {
			log.Printf("[Debug] Failed to get source fields for table %q: %s", table, errString(sourceFieldsErr))
		}
		if destFieldsErr != nil {
			log.Printf("[Debug] Failed to get dest fields for table %q: %s", table, errString(destFieldsErr))
		}
		log.Printf("[Debug] Using legacy text-based comparison for table %q", table)
		alter.SchemaDiff = newSchemaDiff(table, RemoveTableSchemaConfig(sSchema), RemoveTableSchemaConfig(dSchema))
	}

	if sSchema == dSchema {
		return alter
	}
	if len(sSchema) == 0 {
		alter.Type = alterTypeDropTable
		alter.Comment = "源数据库不存在，删除目标数据库多余的表"
		alter.SQL = append(alter.SQL, fmt.Sprintf("drop table `%s`;", table))
		return alter
	}
	if len(dSchema) == 0 {
		alter.Type = alterTypeCreate
		alter.Comment = "目标数据库不存在，创建"
		alter.SQL = append(alter.SQL, fmtTableCreateSQL(sSchema)+";")
		return alter
	}

	diffLines := sc.getSchemaDiff(alter)
	if len(diffLines) == 0 {
		return alter
	}
	alter.Type = alterTypeAlter
	if cfg.SingleSchemaChange {
		for _, line := range diffLines {
			ns := fmt.Sprintf("ALTER TABLE `%s`\n%s;", table, line)
			alter.SQL = append(alter.SQL, ns)
		}
	} else {
		ns := fmt.Sprintf("ALTER TABLE `%s`\n%s;", table, strings.Join(diffLines, ",\n"))
		alter.SQL = append(alter.SQL, ns)
	}

	return alter
}

func (sc *SchemaSync) getSchemaDiff(alter *TableAlterData) []string {
	sourceMyS := alter.SchemaDiff.Source
	destMyS := alter.SchemaDiff.Dest
	table := alter.Table
	var beforeFieldName string
	var alterLines []string
	var fieldCount int = 0

	// 比对字段 - Two-phase comparison strategy:
	// Phase 1: Compare text from SHOW CREATE TABLE first
	// Phase 2: Only if text differs, use INFORMATION_SCHEMA for detailed comparison
	useStructuredComparison := len(sourceMyS.FieldInfos) > 0 && len(destMyS.FieldInfos) > 0

	if useStructuredComparison {
		log.Printf("[Debug] Using two-phase field comparison for table %s", table)
		// Use two-phase comparison
		for el := sourceMyS.Fields.Front(); el != nil; el = el.Next() {
			fieldName := el.Key.(string)
			if sc.Config.IsIgnoreField(table, fieldName) {
				log.Printf("ignore column %s.%s", table, fieldName)
				continue
			}
			var alterSQL string

			if destValue, has := destMyS.Fields.Get(fieldName); has {
				// Field exists in destination
				// Phase 1: Compare text from SHOW CREATE TABLE directly
				if el.Value == destValue {
					// Text definitions are identical, skip this field
					log.Println("[Debug] check column.alter ", fmt.Sprintf("%s.%s", table, fieldName), "not change (text identical)")
					beforeFieldName = fieldName
					fieldCount++
					continue
				}

				// Phase 2: Text differs, use structured comparison to determine if change is needed
				sourceFieldInfo := sourceMyS.FieldInfos[fieldName]
				destFieldInfo := destMyS.FieldInfos[fieldName]

				if sourceFieldInfo != nil && destFieldInfo != nil {
					if sourceFieldInfo.Equals(destFieldInfo) {
						// Structured info shows they're semantically equal despite text difference
						log.Printf("[Debug] field %s.%s: text differs but semantically equal, skipping", table, fieldName)
						log.Printf("[Debug] source text: %s", el.Value)
						log.Printf("[Debug] dest text: %s", destValue)
						beforeFieldName = fieldName
						fieldCount++
						continue
					}
					// Fields are genuinely different
					alterSQL = fmt.Sprintf("CHANGE `%s` %s", fieldName, sourceFieldInfo.String())
					log.Printf("[Debug] field %s.%s: confirmed difference via structured comparison", table, fieldName)
					log.Printf("[Debug] source: %+v", sourceFieldInfo)
					log.Printf("[Debug] dest: %+v", destFieldInfo)
				} else {
					// No structured info, use text-based CHANGE
					alterSQL = fmt.Sprintf("CHANGE `%s` %s", fieldName, el.Value)
					log.Printf("[Debug] field %s.%s: text differs, using text-based change", table, fieldName)
				}
				beforeFieldName = fieldName
			} else {
				// Field doesn't exist in destination, ADD it
				if len(beforeFieldName) == 0 {
					if fieldCount == 0 {
						alterSQL = "ADD " + el.Value.(string) + " FIRST"
					} else {
						alterSQL = "ADD " + el.Value.(string)
					}
				} else {
					alterSQL = fmt.Sprintf("ADD %s AFTER `%s`", el.Value.(string), beforeFieldName)
				}
				beforeFieldName = fieldName
			}

			if len(alterSQL) != 0 {
				log.Println("[Debug] check column.alter ", fmt.Sprintf("%s.%s", table, fieldName), "alterSQL=", alterSQL)
				alterLines = append(alterLines, alterSQL)
			} else {
				log.Println("[Debug] check column.alter ", fmt.Sprintf("%s.%s", table, fieldName), "not change")
			}
			fieldCount++
		}
	} else {
		log.Printf("[Debug] Using legacy text-based field comparison for table %s", table)
		// Use legacy text-based comparison
		for el := sourceMyS.Fields.Front(); el != nil; el = el.Next() {
			if sc.Config.IsIgnoreField(table, el.Key.(string)) {
				log.Printf("ignore column %s.%s", table, el.Key.(string))
				continue
			}
			var alterSQL string
			if destDt, has := destMyS.Fields.Get(el.Key); has {
				if el.Value != destDt {
					alterSQL = fmt.Sprintf("CHANGE `%s` %s", el.Key, el.Value)
				}
				beforeFieldName = el.Key.(string)
			} else {
				if len(beforeFieldName) == 0 {
					if fieldCount == 0 {
						alterSQL = "ADD " + el.Value.(string) + " FIRST"
					} else {
						alterSQL = "ADD " + el.Value.(string)
					}
				} else {
					alterSQL = fmt.Sprintf("ADD %s AFTER `%s`", el.Value.(string), beforeFieldName)
				}
				beforeFieldName = el.Key.(string)
			}

			if len(alterSQL) != 0 {
				log.Println("[Debug] check column.alter ", fmt.Sprintf("%s.%s", table, el.Key.(string)), "alterSQL=", alterSQL)
				alterLines = append(alterLines, alterSQL)
			} else {
				log.Println("[Debug] check column.alter ", fmt.Sprintf("%s.%s", table, el.Key.(string)), "not change")
			}
			fieldCount++
		}
	}

	// 源库已经删除的字段
	if sc.Config.Drop {
		for _, name := range destMyS.Fields.Keys() {
			if sc.Config.IsIgnoreField(table, name.(string)) {
				log.Printf("ignore column %s.%s", table, name)
				continue
			}
			if _, has := sourceMyS.Fields.Get(name); !has {
				alterSQL := fmt.Sprintf("drop `%s`", name)
				alterLines = append(alterLines, alterSQL)
				log.Println("[Debug] check column.drop ", fmt.Sprintf("%s.%s", table, name), "alterSQL=", alterSQL)
			} else {
				log.Println("[Debug] check column.drop ", fmt.Sprintf("%s.%s", table, name), "not change")
			}
		}
	}

	// 多余的字段暂不删除

	// 比对索引
	for indexName, idx := range sourceMyS.IndexAll {
		if sc.Config.IsIgnoreIndex(table, indexName) {
			log.Printf("ignore index %s.%s", table, indexName)
			continue
		}
		dIdx, has := destMyS.IndexAll[indexName]
		log.Println("[Debug] indexName---->[", fmt.Sprintf("%s.%s", table, indexName),
			"] dest_has:", has, "\ndest_idx:", dIdx, "\nsource_idx:", idx)
		var alterSQLs []string
		if has {
			if idx.SQL != dIdx.SQL {
				alterSQLs = append(alterSQLs, idx.alterAddSQL(true)...)
			}
		} else {
			alterSQLs = append(alterSQLs, idx.alterAddSQL(false)...)
		}
		if len(alterSQLs) > 0 {
			alterLines = append(alterLines, alterSQLs...)
			log.Println("[Debug] check index.alter ", fmt.Sprintf("%s.%s", table, indexName), "alterSQL=", alterSQLs)
		} else {
			log.Println("[Debug] check index.alter ", fmt.Sprintf("%s.%s", table, indexName), "not change")
		}
	}

	// drop index
	if sc.Config.Drop {
		for indexName, dIdx := range destMyS.IndexAll {
			if sc.Config.IsIgnoreIndex(table, indexName) {
				log.Printf("ignore index %s.%s", table, indexName)
				continue
			}
			var dropSQL string
			if _, has := sourceMyS.IndexAll[indexName]; !has {
				dropSQL = dIdx.alterDropSQL()
			}

			if len(dropSQL) != 0 {
				alterLines = append(alterLines, dropSQL)
				log.Println("[Debug] check index.drop ", fmt.Sprintf("%s.%s", table, indexName), "alterSQL=", dropSQL)
			} else {
				log.Println("[Debug] check index.drop ", fmt.Sprintf("%s.%s", table, indexName), " not change")
			}
		}
	}

	// 比对外键
	for foreignName, idx := range sourceMyS.ForeignAll {
		if sc.Config.IsIgnoreForeignKey(table, foreignName) {
			log.Printf("ignore foreignName %s.%s", table, foreignName)
			continue
		}
		dIdx, has := destMyS.ForeignAll[foreignName]
		log.Println("[Debug] foreignName---->[", fmt.Sprintf("%s.%s", table, foreignName),
			"] dest_has:", has, "\ndest_idx:", dIdx, "\nsource_idx:", idx)
		var alterSQLs []string
		if has {
			if idx.SQL != dIdx.SQL {
				alterSQLs = append(alterSQLs, idx.alterAddSQL(true)...)
			}
		} else {
			alterSQLs = append(alterSQLs, idx.alterAddSQL(false)...)
		}
		if len(alterSQLs) > 0 {
			alterLines = append(alterLines, alterSQLs...)
			log.Println("[Debug] check foreignKey.alter ", fmt.Sprintf("%s.%s", table, foreignName), "alterSQL=", alterSQLs)
		} else {
			log.Println("[Debug] check foreignKey.alter ", fmt.Sprintf("%s.%s", table, foreignName), "not change")
		}
	}

	// drop 外键
	if sc.Config.Drop {
		for foreignName, dIdx := range destMyS.ForeignAll {
			if sc.Config.IsIgnoreForeignKey(table, foreignName) {
				log.Printf("ignore foreignName %s.%s", table, foreignName)
				continue
			}
			var dropSQL string
			if _, has := sourceMyS.ForeignAll[foreignName]; !has {
				log.Println("[Debug] foreignName --->[", fmt.Sprintf("%s.%s", table, foreignName), "]", "didx:", dIdx)
				dropSQL = dIdx.alterDropSQL()
			}
			if len(dropSQL) != 0 {
				alterLines = append(alterLines, dropSQL)
				log.Println("[Debug] check foreignKey.drop ", fmt.Sprintf("%s.%s", table, foreignName), "alterSQL=", dropSQL)
			} else {
				log.Println("[Debug] check foreignKey.drop ", fmt.Sprintf("%s.%s", table, foreignName), "not change")
			}
		}
	}

	return alterLines
}

// SyncSQL4Dest sync schema change
func (sc *SchemaSync) SyncSQL4Dest(sqlStr string, sqls []string) error {
	sqlStr = strings.TrimSpace(sqlStr)
	log.Print("Exec_SQL:\n>>>>>>\n", color.GreenString(sqlStr), "\n<<<<<<<<\n\n")
	if len(sqlStr) == 0 {
		log.Println("sql_is_empty, skip")
		return nil
	}
	t := newMyTimer()
	ret, err := sc.DestDb.Query(sqlStr)

	defer func() {
		if ret != nil {
			err := ret.Close()
			if err != nil {
				log.Println("close ret error:", errString(err))
				return
			}
		}
	}()

	// how to enable allowMultiQueries?
	if err != nil && len(sqls) > 1 {
		log.Println("Exec_mut_query failed, err=", errString(err), ", now try exec SQLs foreach")
		tx, errTx := sc.DestDb.sqlDB.Begin()
		if errTx != nil {
			log.Println("db.Begin failed", errString(err))
			return errTx
		}
		for _, sql := range sqls {
			ret, err = tx.Query(sql)
			log.Println("query_one:[", sql, "]", errString(err))
			if err != nil {
				break
			}
		}
		if err == nil {
			err = tx.Commit()
		} else {
			_ = tx.Rollback()
		}
	}
	t.stop()
	if err != nil {
		log.Println("EXEC_SQL_FAILED:", errString(err))
		return err
	}
	log.Println("EXEC_SQL_SUCCESS, used:", t.usedSecond())
	cl, err := ret.Columns()
	log.Println("EXEC_SQL_RET:", cl, err)
	return err
}
