//  Copyright(C) 2025 github.com/hidu  All Rights Reserved.
//  Author: hidu <duv123+git@gmail.com>
//  Date: 2025-10-21

package internal

import (
	"fmt"
	"log"
	"strings"

	"github.com/xanygo/anygo/cli/xcolor"
)

func Execute(cfg *Config) {
	scs := newStatics(cfg)
	defer func() {
		scs.timer.stop()
		scs.sendMailNotice(cfg)
	}()

	sc := NewSchemaSync(cfg)
	allTables := sc.AllDBTables()
	// log.Println("source db table total:", len(allTables))

	changedTables := make(map[string][]*TableAlterData)

	for _, table := range allTables {
		xcolor.Green("start checking table %q ...", table)
		if !cfg.CheckMatchTables(table) {
			xcolor.Cyan("table %q skipped by not match", table)
			continue
		}

		if cfg.CheckMatchIgnoreTables(table) {
			xcolor.Cyan("table %q skipped by ignore", table)
			continue
		}

		sd := sc.getAlterDataByTable(table, cfg)

		switch sd.Type {
		case alterTypeNo:
			xcolor.Yellow("table %q not changed", table)
			continue
		case alterTypeDropTable:
			xcolor.Yellow("table %q skipped, only exists in destination's database", table)
			continue
		default:
		}

		fmt.Printf("\n%s\n\n", sd)

		relationTables := sd.SchemaDiff.RelationTables()
		log.Printf("table %q RelationTables: %q", table, relationTables)

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
