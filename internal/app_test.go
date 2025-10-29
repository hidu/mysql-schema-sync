//  Copyright(C) 2025 github.com/hidu  All Rights Reserved.
//  Author: hidu <duv123+git@gmail.com>
//  Date: 2025-10-29

package internal_test

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/xanygo/anygo/cli/xcolor"
	"github.com/xanygo/anygo/xt"

	"github.com/hidu/mysql-schema-sync/internal"
)

func TestWithDB(t *testing.T) {
	source := strings.TrimSpace(os.Getenv("MSS_Test_Source"))
	dest := strings.TrimSpace(os.Getenv("MSS_Test_Dest"))
	if source == "" || dest == "" {
		t.Logf("env.MSS_Test_Source=%q, env.MSS_Test_Dest=%q  Test Skipped", source, dest)
		t.SkipNow()
		return
	}
	getDBS := func(t *testing.T) (s *sql.DB, d *sql.DB) {
		t.Helper()
		sourceDB, err := testConnectDB(t, source)
		xt.NoError(t, err)
		testImportTables(t, sourceDB, "testdata/app/source_tables")

		destDB, err := testConnectDB(t, dest)
		xt.NoError(t, err)
		testImportTables(t, destDB, "testdata/app/dest_tables")
		return sourceDB, destDB
	}

	t.Run("case 1 no sync", func(t *testing.T) {
		sourceDB, destDB := getDBS(t)
		defer sourceDB.Close()
		defer destDB.Close()

		cfg := &internal.Config{
			SourceDSN: source,
			DestDSN:   dest,
		}
		cfg.Check()
		internal.Execute(cfg)
	})
	t.Run("case 2 sync", func(t *testing.T) {
		sourceDB, destDB := getDBS(t)
		defer sourceDB.Close()
		defer destDB.Close()

		cfg := &internal.Config{
			SourceDSN: source,
			DestDSN:   dest,
			Drop:      true,
			Sync:      true,
		}
		cfg.Check()
		internal.Execute(cfg)
		// todo check tables
	})
}

func testImportTables(t *testing.T, db *sql.DB, dir string) {
	files, err := filepath.Glob(filepath.Join(dir, "*.sql"))
	xt.NoError(t, err)
	xt.NotEmpty(t, files)
	var sqls []string
	for _, file := range files {
		content, err := os.ReadFile(file)
		xt.NoError(t, err)
		sqls = append(sqls, string(content))
	}
	err = testDBExec(t, db, sqls...)
	xt.NoError(t, err)
}

func testConnectDB(t *testing.T, dsn string) (*sql.DB, error) {
	cfg, err := mysql.ParseDSN(dsn)
	if err != nil {
		return nil, err
	}
	if cfg.DBName == "" {
		return nil, errors.New("empty DBName")
	}
	nc := cfg.Clone()
	nc.DBName = ""
	db, err := sql.Open("mysql", nc.FormatDSN())
	if err != nil {
		return nil, err
	}
	sqls := []string{
		fmt.Sprintf("DROP DATABASE IF EXISTS %s", cfg.DBName),
		fmt.Sprintf("CREATE DATABASE %s CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci", cfg.DBName),
	}
	if err = testDBExec(t, db, sqls...); err != nil {
		db.Close()
		return nil, err
	}
	db.Close()
	return sql.Open("mysql", dsn)
}

func testDBExec(t *testing.T, db *sql.DB, sqls ...string) error {
	t.Logf("testExec:%s", strings.Repeat("-", 60))
	for _, sql := range sqls {
		t.Logf("start exec: \n%s", xcolor.GreenString(sql))
		ret, err := db.Exec(sql)
		t.Logf("exec result: err: %v", err)
		if err != nil {
			return err
		}
		num, err := ret.RowsAffected()
		t.Logf("RowsAffected %d, err: %v", num, err)
		if err != nil {
			return err
		}
	}
	return nil
}
