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

func (sc *SchemaSync) getAlterDataByTable(table string) *TableAlterData {
	alter := new(TableAlterData)
	alter.Table = table
	alter.Type = alterTypeNo

	sschema := sc.SourceDb.GetTableSchema(table)
	dschema := sc.DestDb.GetTableSchema(table)

	alter.SchemaDiff = newSchemaDiff(table, sschema, dschema)

	if sschema == dschema {
		return alter
	}
	if sschema == "" {
		alter.Type = alterTypeDrop
		alter.SQL = fmt.Sprintf("drop table `%s`;", table)
		return alter
	}
	if dschema == "" {
		alter.Type = alterTypeCreate
		alter.SQL = sschema + ";"
		return alter
	}

	diff := sc.getSchemaDiff(alter)
	if diff != "" {
		alter.Type = alterTypeAlter
		alter.SQL = fmt.Sprintf("ALTER TABLE `%s`\n%s;", table, diff)
	}

	return alter
}

func (sc *SchemaSync) getSchemaDiff(alter *TableAlterData) string {
	sourceMyS := alter.SchemaDiff.Source
	destMyS := alter.SchemaDiff.Dest
	table := alter.Table

	var alterLines []string
	//比对字段
	for name, dt := range sourceMyS.Fields {
		if sc.Config.IsIgnoreField(table, name) {
			log.Printf("ignore field %s.%s", table, name)
			continue
		}
		var alterSQL string
		if destDt, has := destMyS.Fields[name]; has {
			if dt != destDt {
				alterSQL = fmt.Sprintf("CHANGE `%s` %s", name, dt)
			}
		} else {
			alterSQL = "ADD " + dt
		}
		if alterSQL != "" {
			alterLines = append(alterLines, alterSQL)
		}
	}

	//源库已经删除的字段
	if sc.Config.Drop {
		for name := range destMyS.Fields {
			if sc.Config.IsIgnoreField(table, name) {
				log.Printf("ignore field %s.%s", table, name)
				continue
			}
			if _, has := sourceMyS.Fields[name]; !has {
				alterLines = append(alterLines, fmt.Sprintf("drop `%s`", name))
			}
		}
	}

	//多余的字段暂不删除

	//比对索引
	for indexName, idx := range sourceMyS.IndexAll {
		if sc.Config.IsIgnoreIndex(table, indexName) {
			log.Printf("ignore index %s.%s", table, indexName)
			continue
		}
		dIdx, has := destMyS.IndexAll[indexName]
		log.Println("trace indexName---->[", indexName, "] has:", has, "\ndest_idx:", dIdx, "\nsource_idx:", idx)
		alterSQL := ""
		if has {
			if idx.SQL != dIdx.SQL {
				alterSQL = idx.alterAddSQL(true)
			}
		} else {
			alterSQL = idx.alterAddSQL(false)
		}
		if alterSQL != "" {
			alterLines = append(alterLines, alterSQL)
		}
		//		fmt.Println("alterSQL:", alterSQL)
	}

	//drop index
	if sc.Config.Drop {
		for indexName, dIdx := range destMyS.IndexAll {
			if sc.Config.IsIgnoreIndex(table, indexName) {
				log.Printf("ignore index %s.%s", table, indexName)
				continue
			}

			if _, has := sourceMyS.IndexAll[indexName]; !has {
				if dropSQL := dIdx.alterDropSQL(); dropSQL != "" {
					alterLines = append(alterLines, dropSQL)
				}
			}
		}
	}

	//比对外键
	for foreignName, idx := range sourceMyS.ForeignAll {
		if sc.Config.IsIgnoreForeignKey(table, foreignName) {
			log.Printf("ignore foreignName %s.%s", table, foreignName)
			continue
		}
		dIdx, has := destMyS.ForeignAll[foreignName]
		log.Println("trace foreignName---->[", foreignName, "] has:", has, "\ndest_idx:", dIdx, "\nsource_idx:", idx)
		alterSQL := ""
		if has {
			if idx.SQL != dIdx.SQL {
				alterSQL = idx.alterAddSQL(true)
			}
		} else {
			alterSQL = idx.alterAddSQL(false)
		}
		if alterSQL != "" {
			alterLines = append(alterLines, alterSQL)
		}
		//		fmt.Println("alterSQL:", alterSQL)
	}

	//drop 外键
	if sc.Config.Drop {
		for foreignName, dIdx := range destMyS.ForeignAll {
			if sc.Config.IsIgnoreForeignKey(table, foreignName) {
				log.Printf("ignore foreignName %s.%s", table, foreignName)
				continue
			}

			if _, has := sourceMyS.ForeignAll[foreignName]; !has {
				log.Println("trace foreignName --->[", foreignName, "]", "didx:", dIdx)
				if dropSQL := dIdx.alterDropSQL(); dropSQL != "" {
					alterLines = append(alterLines, dropSQL)
				}
			}
		}
	}

	return strings.Join(alterLines, ",\n")
}

// SyncSQL4Dest sync schema change
func (sc *SchemaSync) SyncSQL4Dest(sqlStr string) error {
	log.Println("Exec_SQL_START:\n>>>>>>\n", sqlStr, "\n<<<<<<<<\n")
	sqlStr = strings.TrimSpace(sqlStr)
	if sqlStr == "" {
		log.Println("sql_is_empty,skip")
		return nil
	}
	t := newMyTimer()
	ret, err := sc.DestDb.Query(sqlStr)
	t.stop()
	if err != nil {
		log.Println("EXEC_SQL_FAIELD", err)
		return err
	}
	log.Println("EXEC_SQL_SUCCESS,used:", t.usedSecond())
	cl, err := ret.Columns()
	log.Println("EXEC_SQL_RET:", cl, err)
	return err
}

// CheckSchemaDiff 执行最终的diff
func CheckSchemaDiff(cfg *Config) {
	statics := newStatics(cfg)
	defer (func() {
		statics.timer.stop()
		statics.sendMailNotice(cfg)
	})()

	sc := NewSchemaSync(cfg)
	newTables := sc.SourceDb.GetTableNames()
	log.Println("source db table total:", len(newTables))

	changedTables := make(map[string][]*TableAlterData)

	for index, table := range newTables {
		log.Printf("Index : %d Table : %s\n", index, table)
		if !cfg.ChechMatchTables(table) {
			log.Println("Table:", table, "skip")
			continue
		}

		sd := sc.getAlterDataByTable(table)

		if sd.Type != alterTypeNo {
			fmt.Println(sd)
			fmt.Println("")
			relationTables := sd.SchemaDiff.RelationTables()
			//			fmt.Println("relationTables:",table,relationTables)

			//将所有有外键关联的单独放
			groupKey := "multi"
			if len(relationTables) == 0 {
				groupKey = "single_" + table
			}
			if _, has := changedTables[groupKey]; !has {
				changedTables[groupKey] = make([]*TableAlterData, 0)
			}
			changedTables[groupKey] = append(changedTables[groupKey], sd)
		} else {
			log.Println("table:", table, "not change,", sd)
		}
	}

	log.Println("trace changedTables:", changedTables)

	countSuccess := 0
	countFailed := 0
	canRunTypePref := "single"
	//先执行单个表的
run_sync:
	for typeName, sds := range changedTables {
		if !strings.HasPrefix(typeName, canRunTypePref) {
			continue
		}
		log.Println("runSyncType:", typeName)
		var sqls []string
		var sts []*tableStatics
		for _, sd := range sds {
			sqls = append(sqls, sd.SQL)

			st := statics.newTableStatics(sd.Table, sd)
			sts = append(sts, st)
		}

		sql := strings.Join(sqls, ";\n")
		var ret error

		if sc.Config.Sync {
			ret = sc.SyncSQL4Dest(sql)
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

	} //end for

	//最后在执行多个表的alter
	if canRunTypePref == "single" {
		canRunTypePref = "multi"
		goto run_sync
	}

	if sc.Config.Sync {
		log.Println("execute_all_sql_done,success_total:", countSuccess, "failed_total:", countFailed)
	}

}
