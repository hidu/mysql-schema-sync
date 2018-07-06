package internal

import (
	"fmt"
)

//DestDb struct
type DestDb struct {
	Config *Config
	Db     *MyDb
	DbName string
}

//NewDestDb 初始化
func NewDestDb(config *Config) *DestDb {
	s := new(DestDb)
	s.Config = config

	newDsn, DbName := newDsnDbName(s.Config.DestDSN)
	newDsn = fmt.Sprintf("%s%s", newDsn, ")/")

	s.DbName = DbName
	s.Db = NewMyDb(newDsn, "dest")
	return s
}

// CheckDestDb 检查目的数据库是否存在
func (sc *DestDb) CheckDestDb() {
	sc.Db.GetDestDbName(sc.DbName)
	println("...Create database...")
}

// Check 检查数据库
func Check(config *Config) {
	dest := NewDestDb(config)
	dest.CheckDestDb()
}

/*
func main() {
	dsn := "root:Datahub828!@tcp(116.62.116.178:3306)/"
	stat := "CREATE DATABASE IF NOT EXISTS jumping;"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db.Exec(stat)
	defer db.Close()
}
*/
