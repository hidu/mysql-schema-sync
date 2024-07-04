package internal

import (
	"fmt"
	"log"
	"strings"
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
	s.SourceDb = NewMyDb(config.SourceDSN, "source")
	s.DestDb = NewMyDb(config.DestDSN, "dest")
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

// 合并源数据库和目标数据库的表名
func (sc *SchemaSync) GetTableNames() []string {
	sourceTables := sc.SourceDb.GetTableNames()
	destTables := sc.DestDb.GetTableNames()
	var tables []string
	tables = append(tables, destTables...)
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
	alter.SchemaDiff = newSchemaDiff(table, RemoveTableSchemaConfig(sSchema), RemoveTableSchemaConfig(dSchema))

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
	// 比对字段
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
	log.Print("Exec_SQL_START:\n>>>>>>\n", sqlStr, "\n<<<<<<<<\n\n")
	sqlStr = strings.TrimSpace(sqlStr)
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
				log.Println("close ret error:", err)
				return
			}
		}
	}()

	// how to enable allowMultiQueries?
	if err != nil && len(sqls) > 1 {
		log.Println("exec_mut_query failed, err=", err, ",now exec SQLs foreach")
		tx, errTx := sc.DestDb.Db.Begin()
		if errTx == nil {
			for _, sql := range sqls {
				ret, err = tx.Query(sql)
				log.Println("query_one:[", sql, "]", err)
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
	}
	t.stop()
	if err != nil {
		log.Println("EXEC_SQL_FAILED:", err)
		return err
	}
	log.Println("EXEC_SQL_SUCCESS, used:", t.usedSecond())
	cl, err := ret.Columns()
	log.Println("EXEC_SQL_RET:", cl, err)
	return err
}

// CheckSchemaDiff 执行最终的 diff
func CheckSchemaDiff(cfg *Config) {
	scs := newStatics(cfg)
	defer func() {
		scs.timer.stop()
		scs.sendMailNotice(cfg)
	}()

	sc := NewSchemaSync(cfg)
	newTables := sc.GetTableNames()
	// log.Println("source db table total:", len(newTables))

	changedTables := make(map[string][]*TableAlterData)

	for _, table := range newTables {
		// log.Printf("Index : %d Table : %s\n", index, table)
		if !cfg.CheckMatchTables(table) {
			// log.Println("Table:", table, "skip")
			continue
		}

		if cfg.CheckMatchIgnoreTables(table) {
			log.Println("Table:", table, "skipped by ignore")
			continue
		}

		sd := sc.getAlterDataByTable(table, cfg)

		if sd.Type == alterTypeNo {
			log.Println("table:", table, "not change,", sd)
			continue
		}

		if sd.Type == alterTypeDropTable {
			log.Println("skipped table", table, ",only exists in dest's db")
			continue
		}

		fmt.Println(sd)
		fmt.Println("")
		relationTables := sd.SchemaDiff.RelationTables()
		// fmt.Println("relationTables:",table,relationTables)

		// 将所有有外键关联的单独放
		groupKey := "multi"
		if len(relationTables) == 0 {
			groupKey = "single_" + table
		}
		if _, has := changedTables[groupKey]; !has {
			changedTables[groupKey] = make([]*TableAlterData, 0)
		}
		changedTables[groupKey] = append(changedTables[groupKey], sd)
	}

	log.Println("[Debug] changedTables:", changedTables)

	var countSuccess int
	var countFailed int
	canRunTypePref := "single"

	// 先执行单个表的
runSync:
	for typeName, sds := range changedTables {
		if !strings.HasPrefix(typeName, canRunTypePref) {
			continue
		}
		log.Println("runSyncType:", typeName)
		var sqls []string
		var sts []*tableStatics
		for _, sd := range sds {
			for index := range sd.SQL {
				sql := strings.TrimRight(sd.SQL[index], ";")
				sqls = append(sqls, sql)

				st := scs.newTableStatics(sd.Table, sd, index)
				sts = append(sts, st)
			}
		}

		sql := strings.Join(sqls, ";\n") + ";"
		var ret error

		if sc.Config.Sync {
			ret = sc.SyncSQL4Dest(sql, sqls)
			if ret == nil {
				countSuccess++
			} else {
				countFailed++
			}
		}
		for _, st := range sts {
			st.alterRet = ret
			st.schemaAfter = sc.DestDb.GetTableSchema(st.table)
			st.timer.stop()
		}
	} // end for

	// 最后再执行多个表的 alter
	if canRunTypePref == "single" {
		canRunTypePref = "multi"
		goto runSync
	}

	if sc.Config.Sync {
		log.Println("execute_all_sql_done, success_total:", countSuccess, "failed_total:", countFailed)
	}
}
