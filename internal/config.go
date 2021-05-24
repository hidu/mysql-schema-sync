package internal

import (
	"encoding/json"
	"log"
	"os"
)

// Config  config struct
type Config struct {
	// SourceDSN 同步的源头
	SourceDSN string `json:"source"`

	// DestDSN 将被同步
	DestDSN string `json:"dest"`

	// AlterIgnore 忽略配置， eg:   "tb1*":{"column":["aaa","a*"],"index":["aa"],"foreign":[]}
	AlterIgnore map[string]*AlterIgnoreTable `json:"alter_ignore"`

	// Tables 同步表的白名单，若为空，则同步全库
	Tables []string `json:"tables"`

	// TablesIGNORE 不同步的表
	TablesIGNORE []string `json:"tables_ignore"`

	// Email 完成同步后发送同步信息的邮件账号信息
	Email      *EmailStruct `json:"email"`
	ConfigPath string

	// Sync 是否真正的执行同步操作
	Sync bool

	// Drop 若目标数据库表比源头多了字段、索引，是否删除
	Drop bool
	
	// SingleSchemaChange 生成sql ddl语言每条命令只会进行单个修改操作
	SingleSchemaChange bool `json:"single_schema_change"`
}

func (cfg *Config) String() string {
	ds, _ := json.MarshalIndent(cfg, "  ", "  ")
	return string(ds)
}

// AlterIgnoreTable table's ignore info
type AlterIgnoreTable struct {
	Column []string `json:"column"`
	Index  []string `json:"index"`

	// 外键
	ForeignKey []string `json:"foreign"`
}

// IsIgnoreField isIgnore
func (cfg *Config) IsIgnoreField(table string, name string) bool {
	for tableName, dit := range cfg.AlterIgnore {
		if simpleMatch(tableName, table, "IsIgnoreField_table") {
			for _, col := range dit.Column {
				if simpleMatch(col, name, "IsIgnoreField_colum") {
					return true
				}
			}
		}
	}
	return false
}

// CheckMatchTables check table is match
func (cfg *Config) CheckMatchTables(name string) bool {
	// 若没有指定表，则意味对全库进行同步
	if len(cfg.Tables) == 0 {
		return true
	}
	for _, tableName := range cfg.Tables {
		if simpleMatch(tableName, name, "CheckMatchTables") {
			return true
		}
	}
	return false
}

// CheckMatchIgnoreTables check table_Ignore is match
func (cfg *Config) CheckMatchIgnoreTables(name string) bool {
	if len(cfg.TablesIGNORE) == 0 {
		return false
	}
	for _, tableName := range cfg.TablesIGNORE {
		if simpleMatch(tableName, name, "CheckMatchTables") {
			return true
		}
	}
	return false
}

// Check check config
func (cfg *Config) Check() {
	if cfg.SourceDSN == "" {
		log.Fatal("source DSN is empty")
	}
	if cfg.DestDSN == "" {
		log.Fatal("dest DSN is empty")
	}
	// log.Println("config:\n", cfg)
}

// IsIgnoreIndex is index ignore
func (cfg *Config) IsIgnoreIndex(table string, name string) bool {
	for tableName, dit := range cfg.AlterIgnore {
		if simpleMatch(tableName, table, "IsIgnoreIndex_table") {
			for _, index := range dit.Index {
				if simpleMatch(index, name) {
					return true
				}
			}
		}
	}
	return false
}

// IsIgnoreForeignKey 检查外键是否忽略掉
func (cfg *Config) IsIgnoreForeignKey(table string, name string) bool {
	for tableName, dit := range cfg.AlterIgnore {
		if simpleMatch(tableName, table, "IsIgnoreForeignKey_table") {
			for _, foreignName := range dit.ForeignKey {
				if simpleMatch(foreignName, name) {
					return true
				}
			}
		}
	}
	return false
}

// SendMailFail send fail mail
func (cfg *Config) SendMailFail(errStr string) {
	if cfg.Email == nil {
		log.Println("email conf is empty,skip send mail")
		return
	}
	_host, _ := os.Hostname()
	title := "[mysql-schema-sync][" + _host + "]failed"
	body := "error:<font color=red>" + errStr + "</font><br/>"
	body += "host:" + _host + "<br/>"
	body += "config-file:" + cfg.ConfigPath + "<br/>"
	body += "dest_dsn:" + cfg.DestDSN + "<br/>"
	pwd, _ := os.Getwd()
	body += "pwd:" + pwd + "<br/>"
	cfg.Email.SendMail(title, body)
}

// LoadConfig load config file
func LoadConfig(confPath string) *Config {
	var cfg *Config
	err := loadJSONFile(confPath, &cfg)
	if err != nil {
		log.Fatalln("load json conf:", confPath, "failed:", err)
	}
	cfg.ConfigPath = confPath
	return cfg
}
