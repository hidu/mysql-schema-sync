package internal

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

// Config  config struct
type Config struct {
	// AlterIgnore 忽略配置， eg:   "tb1*":{"column":["aaa","a*"],"index":["aa"],"foreign":[]}
	AlterIgnore map[string]*AlterIgnoreTable `json:"alter_ignore"`

	// Email 完成同步后发送同步信息的邮件账号信息
	Email *EmailStruct `json:"email"`

	// SourceDSN 同步的源头
	SourceDSN string `json:"source"`

	// DestDSN 将被同步
	DestDSN string `json:"dest"`

	ConfigPath string

	// Tables 同步表的白名单，若为空，则同步全库
	Tables []string `json:"tables"`

	// TablesIgnore 不同步的表
	TablesIgnore []string `json:"tables_ignore"`

	// Sync 是否真正的执行同步操作
	Sync bool

	// Drop 若目标数据库表比源头多了字段、索引，是否删除
	Drop bool

	// FieldOrder 是否同步字段顺序（需要重建表，可能影响性能）
	FieldOrder bool

	// HTTPAddress 生成站点报告的地址，如 :8080
	HTTPAddress string

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

func (cfg *Config) SetTables(tables []string) {
	for _, name := range tables {
		name = strings.TrimSpace(name)
		if len(name) > 0 {
			cfg.Tables = append(cfg.Tables, name)
		}
	}
}

// SetTablesIgnore 设置忽略
func (cfg *Config) SetTablesIgnore(tables []string) {
	for _, name := range tables {
		name = strings.TrimSpace(name)
		if len(name) > 0 {
			cfg.TablesIgnore = append(cfg.TablesIgnore, name)
		}
	}
}

// CheckMatchIgnoreTables check table_Ignore is match
func (cfg *Config) CheckMatchIgnoreTables(name string) bool {
	if len(cfg.TablesIgnore) == 0 {
		return false
	}
	for _, tableName := range cfg.TablesIgnore {
		if simpleMatch(tableName, name, "CheckMatchTables") {
			return true
		}
	}
	return false
}

// Check check config
func (cfg *Config) Check() {
	if len(cfg.SourceDSN) == 0 {
		log.Fatal("source DSN is empty")
	}
	if len(cfg.DestDSN) == 0 {
		log.Fatal("dest DSN is empty")
	}
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
