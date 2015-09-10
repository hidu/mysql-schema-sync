package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html"
	"io/ioutil"
	"log"
	"net/smtp"
	"os"
	"regexp"
	"strings"
	"time"
)

var configPath = flag.String("conf", "./config.json", "json config file path")
var sync = flag.Bool("sync", false, "sync shcema change to dest db")
var drop = flag.Bool("drop", true, "drop fields and index")

var source = flag.String("source", "", "mysql dsn source,eg: test@(10.10.0.1:3306)/test\n\twhen it is not empty ignore [-conf] param")
var dest = flag.String("dest", "", "mysql dsn dest,eg test@(127.0.0.1:3306)/imis")
var tables = flag.String("tables", "", "table names to check\n\teg : product_base,order_*")

func init() {
	log.SetFlags(log.Lshortfile | log.Ldate)
	df := flag.Usage
	flag.Usage = func() {
		df()
		fmt.Fprintln(os.Stderr, "")
		fmt.Fprintln(os.Stderr, "mysql schema sync tools 0.1")
		fmt.Fprintln(os.Stderr, "https://github.com/hidu/tool/\n")
	}
}

var cfg *Config

func main() {
	flag.Parse()
	if *source == "" {
		cfg = LoadConfig(*configPath)
	} else {
		cfg = new(Config)
		cfg.SourceDSN = *source
		cfg.DestDSN = *dest
	}
	if cfg.Tables == nil {
		cfg.Tables = []string{}
	}
	if *tables != "" {
		_ts := strings.Split(*tables, ",")
		for _, _name := range _ts {
			_name = strings.TrimSpace(_name)
			if _name != "" {
				cfg.Tables = append(cfg.Tables, _name)
			}
		}
	}
	defer (func() {
		if err := recover(); err != nil {
			log.Println(err)
			sendMailFail(fmt.Sprintf("%s", err))
			log.Fatalln("exit")
		}
	})()

	cfg.check()
	checkSchema()
}

func sendMailFail(errStr string) {
	if cfg.Email == nil {
		log.Println("email conf is empty,skip send mail")
		return
	}
	_host, _ := os.Hostname()
	title := "[mysql-schema-sync][" + _host + "]failed"
	body := "error:<font color=red>" + errStr + "</font><br/>"
	body += "host:" + _host + "<br/>"
	body += "config-file:" + *configPath + "<br/>"
	body += "dest_dsn:" + cfg.DestDSN + "<br/>"
	pwd, _ := os.Getwd()
	body += "pwd:" + pwd + "<br/>"
	cfg.Email.SendMail(title, body)
}

func checkSchema() {
	statics := newStatics()
	defer (func() {
		statics.timer.stop()
		statics.sendMailNotice(cfg)
	})()

	sc := NewSchemaSync(cfg)
	newTables := sc.SourceDb.GetTableNames()

	for index, table := range newTables {
		log.Printf("Index : %d Table : %s\n", index, table)
		if !cfg.ChechMatchTables(table) {
			log.Println("Table:", table, "skip")
			continue
		}

		sd := sc.GetAlterDataByTable(table)

		st := statics.newTableStatics(table, sd)

		if sd.Type != ALTER_TYPE_NO {
			fmt.Println(sd)
			fmt.Println("")
		} else {
			log.Println("table:", table, "not change,", sd)
		}

		if *sync && sd.Type != ALTER_TYPE_NO {
			st.alterRet = sc.SyncSql4Dest(sd.SQL)
		}

		st.timer.stop()
	}
}

type SchemaSync struct {
	Config   *Config
	SourceDb *MyDb
	DestDb   *MyDb
}

func NewSchemaSync(config *Config) *SchemaSync {
	s := new(SchemaSync)
	s.Config = config
	s.SourceDb = NewMyDb(config.SourceDSN)
	s.DestDb = NewMyDb(config.DestDSN)
	return s
}

func (sc *SchemaSync) GetNewTableNames() []string {
	sourceTables := sc.SourceDb.GetTableNames()
	destTables := sc.DestDb.GetTableNames()

	newTables := make([]string, 0)

	for _, name := range sourceTables {
		if !In_StringSlice(name, destTables) {
			newTables = append(newTables, name)
		}
	}
	return newTables
}

func (sc *SchemaSync) GetAlterDataByTable(table string) *TableAlterData {
	alter := new(TableAlterData)
	alter.Table = table
	alter.Type = ALTER_TYPE_NO
	sschema := sc.SourceDb.GetTableSchema(table)
	dschema := sc.DestDb.GetTableSchema(table)
	if sschema == dschema {
		return alter
	}
	if sschema == "" {
		alter.Type = ALTER_TYPE_DROP
		alter.SQL = fmt.Sprintf("drop table `%s`;", table)
		return alter
	}
	if dschema == "" {
		alter.Type = ALTER_TYPE_CREATE
		alter.SQL = sschema + ";"
		return alter
	}

	diff := sc.GetSchemaDiff(table, sschema, dschema)
	if diff != "" {
		alter.Type = ALTER_TYPE_ALTER
		alter.SQL = fmt.Sprintf("ALTER TABLE `%s`\n%s;", table, diff)
	}

	return alter
}

func (sc *SchemaSync) GetSchemaDiff(table string, sourceSchema string, destSchema string) string {
	sourceMyS := ParseSchema(sourceSchema)
	destMyS := ParseSchema(destSchema)
	alterLines := make([]string, 0)
	//比对字段
	for name, dt := range sourceMyS.Fields {
		if cfg.IsIgnoreField(table, name) {
			log.Printf("ignore field %s.%s", table, name)
			continue
		}
		var alterSql string
		if destDt, has := destMyS.Fields[name]; has {
			if dt != destDt {
				alterSql = fmt.Sprintf("CHANGE `%s` %s", name, dt)
			}
		} else {
			alterSql = "ADD " + dt
		}
		if alterSql != "" {
			alterLines = append(alterLines, alterSql)
		}
	}

	//源库已经删除的字段
	if *drop {
		for name := range destMyS.Fields {
			if cfg.IsIgnoreField(table, name) {
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
		if cfg.IsIgnoreIndex(table, indexName) {
			log.Printf("ignore index %s.%s", table, indexName)
			continue
		}
		dIdx, has := destMyS.IndexAll[indexName]
		alterSql := ""
		if has {
			if idx.SQL != dIdx.SQL {
				alterSql = idx.AlterAddSql(true)
			}
		} else {
			alterSql = idx.AlterAddSql(false)
		}
		if alterSql != "" {
			alterLines = append(alterLines, alterSql)
		}
	}

	//drop index
	if *drop {
		for indexName, dIdx := range destMyS.IndexAll {
			if cfg.IsIgnoreIndex(table, indexName) {
				log.Printf("ignore index %s.%s", table, indexName)
				continue
			}

			if _, has := sourceMyS.IndexAll[indexName]; !has {
				if dropSql := dIdx.AlterDropSql(); dropSql != "" {
					alterLines = append(alterLines, dropSql)
				}
			}
		}
	}
	return strings.Join(alterLines, ",\n")
}

func (sc *SchemaSync) SyncSql4Dest(sqlStr string) error {
	log.Println("Exec_SQL_START:", sqlStr)
	sqlStr = strings.TrimSpace(sqlStr)
	if sqlStr == "" {
		log.Println("sql_is_empty,skip")
		return nil
	}
	ret, err := sc.DestDb.Db.Query(sqlStr)
	log.Println("Exec_SQL_DONE,err:", err)
	cl, err := ret.Columns()
	log.Println("ret:", cl, err)
	return err
}

type MySchema struct {
	Fields   map[string]string
	IndexAll map[string]*DbIndex
}

type IndexType string

const (
	IndexType_PRIMARY IndexType = "PRIMARY"
	IndexType_Index             = "index"
)

type DbIndex struct {
	IndexType IndexType
	Name      string
	SQL       string
}

func (idx *DbIndex) AlterAddSql(drop bool) string {
	alterSql := []string{}
	if drop {
		dropSql := idx.AlterDropSql()
		if dropSql != "" {
			alterSql = append(alterSql, dropSql)
		}
	}

	switch idx.IndexType {
	case IndexType_PRIMARY:
		alterSql = append(alterSql, "ADD "+idx.SQL)
	case IndexType_Index:
		alterSql = append(alterSql, fmt.Sprintf("ADD %s", idx.SQL))
	default:
		log.Fatalln("unknow indexType", idx.IndexType)
	}
	return strings.Join(alterSql, ",")
}

func (idx *DbIndex) AlterDropSql() string {
	switch idx.IndexType {
	case IndexType_PRIMARY:
		return "DROP PRIMARY KEY"
	case IndexType_Index:
		return fmt.Sprintf("DROP INDEX `%s`", idx.Name)
	default:
		log.Fatalln("unknow indexType", idx.IndexType)
	}
	return ""
}

func (mys *MySchema) String() string {
	s := "Fields:\n"
	for name, v := range mys.Fields {
		s += fmt.Sprintf("  %15s : %s\n", name, v)
	}
	s += "Index:\n  "
	for name, idx := range mys.IndexAll {
		s += "    " + name + " : " + idx.SQL
	}
	return s
}

func (mys *MySchema) GetFieldNames() []string {
	names := make([]string, 0)
	for name := range mys.Fields {
		names = append(names, name)
	}
	return names
}

func ParseSchema(schema string) *MySchema {
	schema = strings.TrimSpace(schema)
	lines := strings.Split(schema, "\n")
	mys := &MySchema{
		Fields:   make(map[string]string),
		IndexAll: make(map[string]*DbIndex, 0),
	}

	for i := 1; i < len(lines)-1; i++ {
		line := strings.TrimSpace(lines[i])
		if line == "" {
			continue
		}
		line = strings.TrimRight(line, ",")
		if line[0] == '`' {
			index := strings.Index(line[1:], "`")
			name := line[1 : index+1]
			mys.Fields[name] = line
		} else {
			idx := ParseDbIndexLine(line)
			mys.IndexAll[idx.Name] = idx
		}
	}
	return mys

}

func ParseDbIndexLine(line string) *DbIndex {
	line = strings.TrimSpace(line)
	idx := &DbIndex{
		SQL: line,
	}
	if strings.HasPrefix(line, "PRIMARY") {
		idx.IndexType = IndexType_PRIMARY
		idx.Name = "PRIMARY KEY"
		return idx
	}

	if strings.HasPrefix(line, "UNIQUE") || strings.HasPrefix(line, "KEY") {
		arr := strings.Split(line, "`")
		idx.IndexType = IndexType_Index
		idx.Name = arr[1]
		return idx
	}
	log.Fatalln("db_index parse failed,unsupport,line:", line)
	return nil
}

type MyDb struct {
	Db *sql.DB
}

type ALTER_TYPE int

const (
	ALTER_TYPE_NO     ALTER_TYPE = 0
	ALTER_TYPE_CREATE            = 1
	ALTER_TYPE_DROP              = 2
	ALTER_TYPE_ALTER             = 3
)

func (at ALTER_TYPE) String() string {
	switch at {
	case ALTER_TYPE_NO:
		return "not_change"
	case ALTER_TYPE_CREATE:
		return "create"
	case ALTER_TYPE_DROP:
		return "drop"
	case ALTER_TYPE_ALTER:
		return "alter"
	default:
		return "unknow"
	}

}

type TableAlterData struct {
	Table string
	Type  ALTER_TYPE
	SQL   string
}

func (ta *TableAlterData) String() string {
	return fmt.Sprintf("-- Table : %s\n-- Type  : %s\n-- SQL   :\n%s", ta.Table, ta.Type, ta.SQL)
}

func NewMyDb(dsn string) *MyDb {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Sprintf("connect to db [%s] failed,", dsn, err))
	}
	return &MyDb{
		Db: db,
	}
}

func (mydb *MyDb) GetTableNames() []string {
	rs, err := mydb.Db.Query("show tables")
	if err != nil {
		panic("show tables failed:" + err.Error())
	}
	defer rs.Close()
	tables := []string{}
	for rs.Next() {
		var name string
		if err := rs.Scan(&name); err != nil {
			panic("show tables failed when scan," + err.Error())
		}
		tables = append(tables, name)
	}
	return tables
}
func (mydb *MyDb) GetTableSchema(name string) (schema string) {
	rs, err := mydb.Db.Query(fmt.Sprintf("show create table `%s`", name))
	if err != nil {
		log.Println(err)
		return
	}
	defer rs.Close()
	for rs.Next() {
		var name string
		if err := rs.Scan(&name, &schema); err != nil {
			panic(fmt.Sprintf("get table %s 's schema failed,%s", name, err))
		}
	}
	return
}

type Config struct {
	SourceDSN  string                      `json:"source"`
	DestDSN    string                      `json:"dest"`
	DropIgnore map[string]*DropIgnoreTable `json:"drop_ignore"`
	Tables     []string                    `json:"tables"`
	Email      *EmailStruct                `json:"email"`
}

func (cfg *Config) String() string {
	ds, _ := json.MarshalIndent(cfg, "  ", "  ")
	return string(ds)
}

type DropIgnoreTable struct {
	Column []string `json:"column"`
	Index  []string `json:"index"`
}

func (cfg *Config) IsIgnoreField(table string, name string) bool {
	if dit, has := cfg.DropIgnore[table]; has {
		for _, col := range dit.Column {
			if col == name {
				return true
			}
		}
	}
	return false
}

func (cfg *Config) ChechMatchTables(name string) bool {
	if len(cfg.Tables) == 0 {
		return true
	}
	for _, tableName := range cfg.Tables {
		tableName = strings.TrimSpace(tableName)
		if tableName == name {
			return true
		}
		tableNameRegStr := "^" + strings.Replace(tableName, "*", `.*`, -1) + "$"
		match, err := regexp.MatchString(tableNameRegStr, name)
		if err != nil {
			log.Println("ChechMatchTables failed,tableName=", tableName, ",err:", err)
		} else {
			log.Println("check tables,", name, " match ", tableName, ",reg:", tableNameRegStr)
		}
		return match
	}
	return false
}

func (cfg *Config) check() {
	if cfg.SourceDSN == "" {
		log.Fatal("source dns is empty")
	}
	if cfg.DestDSN == "" {
		log.Fatal("dest dns is empty")
	}
	log.Println("config:\n", cfg)
}

func (cfg *Config) IsIgnoreIndex(table string, name string) bool {
	if dit, has := cfg.DropIgnore[table]; has {
		for _, index := range dit.Index {
			if index == name {
				return true
			}
		}
	}
	return false
}

type EmailStruct struct {
	SendMailAble bool   `json:"send_mail"`
	SmtpHost     string `json:"smtp_host"`
	From         string `json:"from"`
	Password     string `json:"password"`
	To           string `json:"to"`
}

func (m *EmailStruct) SendMail(title string, body string) {
	if !m.SendMailAble {
		log.Println("disbale send email")
		return
	}
	if m.SmtpHost == "" || m.From == "" || m.To == "" {
		log.Println("smtp_host,from,to is empty")
		return
	}
	addr_info := strings.Split(m.SmtpHost, ":")
	if len(addr_info) != 2 {
		log.Println("smtp_host wrong,eg: host_name:25")
		return
	}
	auth := smtp.PlainAuth("", m.From, m.Password, addr_info[0])
	_sendTo := strings.Split(m.To, ";")
	var sendTo []string
	for _, _to := range _sendTo {
		_to = strings.TrimSpace(_to)
		if _to != "" {
			sendTo = append(sendTo, _to)
		}
	}
	msgBody := fmt.Sprintf("To: %s\r\nContent-Type: text/html;charset=utf-8\r\nSubject: %s\r\n\r\n%s", strings.Join(sendTo, ";"), title, body)
	err := smtp.SendMail(m.SmtpHost, auth, m.From, sendTo, []byte(msgBody))
	if err == nil {
		log.Println("send mail success")
	} else {
		log.Println("send mail failed,err:", err)
	}
}

func LoadConfig(confPath string) *Config {
	bd, err := ioutil.ReadFile(confPath)
	if err != nil {
		log.Fatalln(err)
	}
	var cfg *Config
	err = json.Unmarshal(bd, &cfg)
	if err != nil {
		log.Fatalln(err)
	}
	return cfg
}

func In_StringSlice(str string, strSli []string) bool {
	for _, v := range strSli {
		if str == v {
			return true
		}
	}
	return false
}

type myTimer struct {
	start time.Time
	end   time.Time
}

func newMyTimer() *myTimer {
	return &myTimer{
		start: time.Now(),
	}
}

func (mt *myTimer) stop() {
	mt.end = time.Now()
}
func (mt *myTimer) usedSecond() string {
	return fmt.Sprintf("%f s", mt.end.Sub(mt.start).Seconds())
}

type statics struct {
	timer  *myTimer
	tables []*tableStatics
}

func newStatics() *statics {
	return &statics{
		timer:  newMyTimer(),
		tables: make([]*tableStatics, 0),
	}
}

func (s *statics) newTableStatics(table string, sd *TableAlterData) *tableStatics {
	ts := &tableStatics{
		timer: newMyTimer(),
		table: table,
		alter: sd,
	}
	if sd.Type != ALTER_TYPE_NO {
		s.tables = append(s.tables, ts)
	}
	return ts
}

func (s *statics) toHtml() string {
	code := "<h2>Detail</h2>"
	code += "<h3>Tables</h3>"
	code += `<table>
		<thead>
			<tr>
			<th>no</th>
			<th>table</th>
			<th>alter result</th>
			<th>used</th>
			</tr>
		</thead><tbody>`
	for idx, tb := range s.tables {
		code += "<tr>"
		code += "<td>" + fmt.Sprintf("%d", idx+1) + "</td>\n"
		code += "<td><a href='#detail_" + tb.table + "'>" + tb.table + "</a></td>\n"
		code += "<td>"
		if *sync {
			if tb.alterRet == nil {
				code += "<font color=green>success</font>"
			} else {
				code += "<font color=red>failed," + tb.alterRet.Error() + "</font>"
			}
		} else {
			code += "no sync"
		}
		code += "</td>\n"
		code += "<td>" + tb.timer.usedSecond() + "</td>\n"
		code += "</tr>\n"
	}
	code += "</tbody></table>\n<h3>Sqls</h3><pre>"
	for _, tb := range s.tables {
		code += "<a name='detail_" + tb.table + "'></a>"
		code += html.EscapeString(tb.alter.String()) + "\n\n"
	}
	code += "</pre>"

	return code
}

func (s *statics) alterFailedNum() int {
	n := 0
	for _, tb := range s.tables {
		if tb.alterRet != nil {
			n++
		}
	}
	return n
}

const timeFormatStd string = "2006-01-02 15:04:05"

func (s *statics) sendMailNotice(cfg *Config) {
	if cfg.Email == nil {
		log.Println("mail conf is not set,skip send mail")
		return
	}
	alterTotal := len(s.tables)
	if alterTotal < 1 {
		log.Println("no table change,skip send mail")
		return
	}
	title := "[mysql_schema_sync] " + fmt.Sprintf("%d", alterTotal) + " tables change [" + dsnSort(cfg.DestDSN) + "]"
	body := ""

	if !*sync {
		title += "[preview]"
		body += "<font color=red>this is preview,all sql never execute!</font>"
	}

	host_name, _ := os.Hostname()
	body += "<h2>Info</h2><pre>"
	body += "  from : " + dsnSort(cfg.SourceDSN) + "\n"
	body += "    to : " + dsnSort(cfg.DestDSN) + "\n"
	body += " alter : " + fmt.Sprintf("%d", len(s.tables)) + " tables\n"
	body += "<font color=green>  sync : " + fmt.Sprintf("%t", *sync) + "</font>\n"
	if *sync {
		fn := s.alterFailedNum()
		body += "<font color=red>failed : " + fmt.Sprintf("%d", fn) + "</font>\n"
		if fn > 0 {
			title += " [failed=" + fmt.Sprintf("%d", fn) + "]"
		}
	}
	body += "\n"
	body += "  host : " + host_name + "\n"
	body += " start : " + s.timer.start.Format(timeFormatStd) + "\n"
	body += "   end : " + s.timer.end.Format(timeFormatStd) + "\n"
	body += "  used : " + s.timer.usedSecond() + "\n"

	body += "</pre>"
	body += s.toHtml()
	cfg.Email.SendMail(title, body)
}

func dsnSort(dsn string) string {
	i := strings.Index(dsn, "@")
	if i < 1 {
		return dsn
	}
	return dsn[i+1:]
}

type tableStatics struct {
	timer    *myTimer
	table    string
	alter    *TableAlterData
	alterRet error
}
