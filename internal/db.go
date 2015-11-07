package internal

import (
	"database/sql"
	"fmt"
	//load mysql
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// MyDb db struct
type MyDb struct {
	Db     *sql.DB
	dbType string
}

// NewMyDb parse dsn
func NewMyDb(dsn string, dbType string) *MyDb {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Sprintf("connect to db [%s] failed,", dsn, err))
	}
	return &MyDb{
		Db:     db,
		dbType: dbType,
	}
}

// GetTableNames table names
func (mydb *MyDb) GetTableNames() []string {
	rs, err := mydb.Query("show table status")
	if err != nil {
		panic("show tables failed:" + err.Error())
	}
	defer rs.Close()
	tables := []string{}
	columns, _ := rs.Columns()
	for rs.Next() {
		var values = make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))
		for i := range columns {
			valuePtrs[i] = &values[i]
		}
		if err := rs.Scan(valuePtrs...); err != nil {
			panic("show tables failed when scan," + err.Error())
		}
		var valObj = make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			valObj[col] = v
		}
		if valObj["Engine"] != nil {
			tables = append(tables, valObj["Name"].(string))
		}
	}
	return tables
}

// GetTableSchema table schema
func (mydb *MyDb) GetTableSchema(name string) (schema string) {
	rs, err := mydb.Query(fmt.Sprintf("show create table `%s`", name))
	if err != nil {
		log.Println(err)
		return
	}
	defer rs.Close()
	for rs.Next() {
		var vname string
		if err := rs.Scan(&vname, &schema); err != nil {
			panic(fmt.Sprintf("get table %s 's schema failed,%s", name, err))
		}
	}
	return
}

// Query execute sql query
func (mydb *MyDb) Query(query string, args ...interface{}) (*sql.Rows, error) {
	log.Println("[SQL]", "["+mydb.dbType+"]", query, args)
	return mydb.Db.Query(query, args...)
}
