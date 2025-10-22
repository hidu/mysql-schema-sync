package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/hidu/mysql-schema-sync/internal"
)

var configPath = flag.String("conf", "./mydb_conf.json", "json config file path")
var sync = flag.Bool("sync", false, "sync schema changes to dest's db\non default, only show difference")
var drop = flag.Bool("drop", false, "drop fields,index,foreign key only on dest's table")
var fieldOrder = flag.Bool("field-order", false, "sync field order (may require table rebuild, affecting performance)")
var httpAddress = flag.String("http", "", "HTTP service address, eg. :8080")

var source = flag.String("source", "", "sync from, eg: test@(10.10.0.1:3306)/my_online_db_name\nwhen it is not empty,[-conf] while ignore")
var dest = flag.String("dest", "", "sync to, eg: test@(127.0.0.1:3306)/my_local_db_name")
var tables = flag.String("tables", "", "tables to sync\neg : product_base,order_*")
var tablesIgnore = flag.String("tables_ignore", "", "tables ignore sync\neg : product_base,order_*")
var mailTo = flag.String("mail_to", "", "overwrite config's email.to")
var singleSchemaChange = flag.Bool("single_schema_change", false, "single schema changes ddl command a single schema change")

func init() {
	log.SetFlags(log.Lshortfile | log.Ldate)
	df := flag.Usage
	flag.Usage = func() {
		df()
		fmt.Fprintln(os.Stderr, "")
		fmt.Fprintln(os.Stderr, "mysql schema sync tools "+internal.Version)
		fmt.Fprint(os.Stderr, internal.AppURL+"\n\n")
	}
}

var cfg *internal.Config

func main() {
	flag.Parse()
	if len(*source) == 0 {
		cfg = internal.LoadConfig(*configPath)
	} else {
		cfg = new(internal.Config)
		cfg.SourceDSN = *source
		cfg.DestDSN = *dest
	}
	cfg.Sync = *sync
	cfg.Drop = *drop
	cfg.FieldOrder = *fieldOrder
	cfg.HTTPAddress = *httpAddress
	cfg.SingleSchemaChange = *singleSchemaChange

	if len(*mailTo) != 0 && cfg.Email != nil {
		cfg.Email.To = *mailTo
	}
	cfg.SetTables(strings.Split(*tables, ","))
	cfg.SetTablesIgnore(strings.Split(*tablesIgnore, ","))

	defer (func() {
		if re := recover(); re != nil {
			log.Println(re)
			bf := make([]byte, 4096)
			n := runtime.Stack(bf, false)
			cfg.SendMailFail(fmt.Sprintf("panic:%s\n trace=%s", re, bf[:n]))
			log.Fatalln("panic:", string(bf[:n]))
		}
	})()

	cfg.Check()
	internal.Execute(cfg)
}
