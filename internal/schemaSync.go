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

	alter.SourceSchema = sschema
	alter.DestSchema = dschema

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

	diff := sc.getSchemaDiff(table, sschema, dschema)
	if diff != "" {
		alter.Type = alterTypeAlter
		alter.SQL = fmt.Sprintf("ALTER TABLE `%s`\n%s;", table, diff)
	}

	return alter
}

func (sc *SchemaSync) getSchemaDiff(table string, sourceSchema string, destSchema string) string {
	sourceMyS := ParseSchema(sourceSchema)
	destMyS := ParseSchema(destSchema)
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
		fmt.Println("indexName---->", indexName, "has:", has, dIdx, idx)
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
		fmt.Println("alterSQL:", alterSQL)
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
	return strings.Join(alterLines, ",\n")
}

// SyncSQL4Dest sync schema change
func (sc *SchemaSync) SyncSQL4Dest(sqlStr string) error {
	log.Println("Exec_SQL_START:", sqlStr)
	sqlStr = strings.TrimSpace(sqlStr)
	if sqlStr == "" {
		log.Println("sql_is_empty,skip")
		return nil
	}
	ret, err := sc.DestDb.Query(sqlStr)
	log.Println("EXEC_SQL_DONE,err:", err)
	if err != nil {
		return err
	}
	cl, err := ret.Columns()
	log.Println("ret:", cl, err)
	return err
}

// CheckSchemaDiff do check diff
func CheckSchemaDiff(cfg *Config) {
	statics := newStatics(cfg)
	defer (func() {
		statics.timer.stop()
		statics.sendMailNotice(cfg)
	})()

	sc := NewSchemaSync(cfg)
	newTables := sc.SourceDb.GetTableNames()
	log.Println("source db table total:", len(newTables))
	for index, table := range newTables {
		log.Printf("Index : %d Table : %s\n", index, table)
		if !cfg.ChechMatchTables(table) {
			log.Println("Table:", table, "skip")
			continue
		}

		sd := sc.getAlterDataByTable(table)

		st := statics.newTableStatics(table, sd)

		if sd.Type != alterTypeNo {
			fmt.Println(sd)
			fmt.Println("")
		} else {
			log.Println("table:", table, "not change,", sd)
		}

		if sc.Config.Sync && sd.Type != alterTypeNo {
			st.alterRet = sc.SyncSQL4Dest(sd.SQL)
		}
		st.schemaAfter = sc.DestDb.GetTableSchema(table)

		st.timer.stop()
	}
}
