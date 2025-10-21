package internal

import (
	"flag"
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type statics struct {
	timer  *myTimer
	Config *Config
	tables []*tableStatics
}

type tableStatics struct {
	timer       *myTimer
	table       string
	alter       *TableAlterData
	alterRet    error
	schemaAfter string
}

func newStatics(cfg *Config) *statics {
	return &statics{
		timer:  newMyTimer(),
		tables: make([]*tableStatics, 0),
		Config: cfg,
	}
}

func (s *statics) newTableStatics(table string, sd *TableAlterData, index int) *tableStatics {
	ts := &tableStatics{
		timer: newMyTimer(),
		table: table,
		alter: sd,
	}
	if sd.Type == alterTypeNo {
		return ts
	}
	if s.Config.SingleSchemaChange {
		sds := sd.Split()
		nts := &tableStatics{}
		*nts = *ts
		nts.alter = sds[index]
		s.tables = append(s.tables, nts)
	} else {
		s.tables = append(s.tables, ts)
	}
	return ts
}

func (s *statics) toHTML() string {
	code := "<h2>运行结果</h2>\n"
	code += "<h3>Tables</h3>\n"
	code += `<table class='tb_1'>
		<thead>
			<tr>
			<th width="60px">序号</th>
			<th>Table </th>
			<th>同步(alter) 结果</th>
			<th>耗时</th>
			</tr>
		</thead><tbody>
		`
	for idx, tb := range s.tables {
		code += "<tr>"
		code += "<td>" + strconv.Itoa(idx+1) + "</td>\n"
		code += "<td><a href='#detail_" + tb.table + "'>" + tb.table + "</a></td>\n"
		code += "<td>"
		if s.Config.Sync {
			if tb.alterRet == nil {
				code += "<font color=green>成功</font>"
			} else {
				code += "<font color=red>失败：" + html.EscapeString(tb.alterRet.Error()) + "</font>"
			}
		} else {
			code += "未同步"
		}
		code += "</td>\n"
		code += "<td>" + tb.timer.usedSecond() + "</td>\n"
		code += "</tr>\n"
	}
	code += "</tbody></table>\n<h3>SQLs</h3>\n<pre>"
	for _, tb := range s.tables {
		code += "<a name='detail_" + tb.table + "'></a>"
		code += html.EscapeString(tb.alter.String()) + "\n\n"
	}
	code += "</pre>\n\n"

	code += "<h3>详情</h3>\n"
	code += `<table class='tb_1'>
		<thead>
			<tr>
			<th width="40px">序号</th>
			<th width="80px">Table</th>
			<th>&nbsp;</th>
			<th>&nbsp;</th>
			</tr>
		</thead><tbody>
		`
	for idx, tb := range s.tables {
		code += "<tr>"
		code += "<th rowspan=2>" + strconv.Itoa(idx+1) + "</th>\n"
		code += "<td rowspan=2>" + tb.table + "<br/><br/>"
		if s.Config.Sync {
			if tb.alterRet == nil {
				code += "<font color=green>成功</font>"
			} else {
				code += "<font color=red>失败：" + tb.alterRet.Error() + "</font>"
			}
		} else {
			code += "未同步"
		}
		code += "</td>\n"
		code += "<td valign=top><b>数据源 Schema:</b><br/>"
		if len(tb.alter.SchemaDiff.Source.SchemaRaw) == 0 {
			code += "<font color=red>在源数据源不存在，在目标数据库存在</font>"
		} else {
			code += htmlPre(tb.alter.SchemaDiff.Source.SchemaRaw)
		}
		code += "</td>\n"

		code += "<td valign=top><b>目标 Schema:</b><br/>"
		if len(tb.alter.SchemaDiff.Dest.SchemaRaw) == 0 {
			code += "不存在"
		} else {
			code += htmlPre(tb.alter.SchemaDiff.Dest.SchemaRaw)
		}
		code += "</td>\n"
		code += "</tr>\n"

		code += "<tr>\n"
		code += "<td valign=top><b>请在目标库执行如下 SQL:</b><br/>"
		code += htmlPre(strings.Join(tb.alter.SQL, ","))
		code += "</td>\n"
		code += "<td valign=top>"
		if s.Config.Sync {
			code += "<b>执行后:</b><br/>" + htmlPre(tb.schemaAfter)
		}
		code += "</td>\n"
		code += "</tr>\n"
	}
	code += "</tbody></table>\n"
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

func (s *statics) sendMailNotice(cfg *Config) {
	alterTotal := len(s.tables)
	if alterTotal < 1 {
		writeHTMLResult("no table change")
		log.Println("no table change, skip send mail")
		return
	}
	title := "[mysql_schema_sync][" + dsnShort(cfg.DestDSN) + "]" + strconv.Itoa(alterTotal) + "张表发生变化"
	body := `
<style>
.tb_1,.tb_1 td,.tb_1 th{border: 1px solid;border-collapse: collapse;}
.tb_1 thead{ background-color: #e0e0e0;}
</style>`

	if !s.Config.Sync {
		title += "[preview]"
		body += "<font color=red>所有 SQL 均未执行!</font>\n"
	}

	hostName, _ := os.Hostname()
	body += "<h2>任务信息</h2>\n<pre>"
	body += " 数据源：" + dsnShort(cfg.SourceDSN) + "\n"
	body += "   目标：" + dsnShort(cfg.DestDSN) + "\n"
	body += " 有变化：" + strconv.Itoa(len(s.tables)) + " 张表/条语句\n"
	body += "<font color=green>是否同步：" + fmt.Sprintf("%t", s.Config.Sync) + "</font>\n"
	if s.Config.Sync {
		fn := s.alterFailedNum()
		body += "<font color=red>失败数 : " + strconv.Itoa(fn) + "</font>\n"
		if fn > 0 {
			title += " [失败-" + strconv.Itoa(fn) + "]"
		}
	}
	body += "\n"
	body += "  主机名： " + hostName + "\n"
	body += "开始时间： " + s.timer.start.Format(timeFormatStd) + "\n"
	body += "截止时间： " + s.timer.end.Format(timeFormatStd) + "\n"
	body += "运行耗时： " + s.timer.usedSecond() + "\n"

	body += "</pre>\n"
	body += s.toHTML()

	writeHTMLResult(body)
	if cfg.Email != nil {
		cfg.Email.SendMail(title, body)
	}
	if cfg.HTTPAddress != "" {
		startWebServer(cfg.HTTPAddress)
	}
}

func startWebServer(addr string) {
	fp := filepath.Join(os.TempDir(), "mysql-schema-sync_last.html")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		bf, err := os.ReadFile(fp)
		if err != nil {
			http.NotFoundHandler().ServeHTTP(w, r)
			return
		}
		_, _ = w.Write(bf)
	})
	log.Printf("http://%s", addr)
	log.Println("Press Ctrl-C to terminate the program")
	ser := &http.Server{
		Addr: addr,
	}
	log.Println(ser.ListenAndServe())
}

func writeHTMLResult(str string) {
	fp := filepath.Join(os.TempDir(), "mysql-schema-sync_last.html")
	if len(htmlResultPath) > 0 {
		fp = htmlResultPath
	}
	err := os.WriteFile(fp, []byte(str), 0666)
	log.Println("html result:", fp, err)
}

func init() {
	flag.StringVar(&htmlResultPath, "html", "", "html result file path")
}

var htmlResultPath string
