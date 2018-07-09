package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

// Config  config struct
type Config struct {
	SourceDSN   string                       `json:"source"`
	DestDSN     string                       `json:"dest"`
	AlterIgnore map[string]*AlterIgnoreTable `json:"alter_ignore"`
	Tables      []string                     `json:"tables"`
	Email       *EmailStruct                 `json:"email"`
	ConfigPath  string
	Sync        bool
	Drop        bool
	CreateDb    bool
}

func (cfg *Config) String() string {
	ds, _ := json.MarshalIndent(cfg, "  ", "  ")
	return string(ds)
}

// AlterIgnoreTable table's ignore info
type AlterIgnoreTable struct {
	Column     []string `json:"column"`
	Index      []string `json:"index"`
	ForeignKey []string `json:"foreign"` //外键
}

// IsIgnoreField isIgnore
func (cfg *Config) IsIgnoreField(table string, name string) bool {
	for tname, dit := range cfg.AlterIgnore {
		if simpleMatch(tname, table, "IsIgnoreField_table") {
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

// Check check config
func (cfg *Config) Check(dbName string) {
	if cfg.SourceDSN == "" {
		log.Fatal("source dns is empty")
	}
	if cfg.DestDSN == "" {
		log.Fatal("dest dns is empty")
	}

	if dbName != "" {
		cfg.SourceDSN = cfg.AppendDb(cfg.SourceDSN, dbName)
		cfg.DestDSN = cfg.AppendDb(cfg.DestDSN, dbName)
	}
	//	log.Println("config:\n", cfg)
}

// AppendDb dbname which will ignore the db name in the configure json file
func (cfg *Config) AppendDb(dsn, dbName string) string {
	prefix, dbname := splitDsn(dsn)
	if strings.ToLower(dbname) != strings.ToLower(dbName) {
		return fmt.Sprintf("%s)/%s", prefix, dbName)
	} else {
		return dsn
	}
}

// IsIgnoreIndex is index ignore
func (cfg *Config) IsIgnoreIndex(table string, name string) bool {
	for tname, dit := range cfg.AlterIgnore {
		if simpleMatch(tname, table, "IsIgnoreIndex_table") {
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
	for tname, dit := range cfg.AlterIgnore {
		if simpleMatch(tname, table, "IsIgnoreForeignKey_table") {
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
	//	if *mailTo != "" {
	//		cfg.Email.To = *mailTo
	//	}
	return cfg
}
