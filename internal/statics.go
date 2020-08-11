package internal

import (
	"fmt"
	"html"
	"log"
	"os"
)

type statics struct {
	timer  *myTimer
	tables []*tableStatics
	Config *Config
}

type tableStatics struct {
	timer       *myTimer
	db          string
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

func (s *statics) newTableStatics(db string, table string, sd *TableAlterData) *tableStatics {
	ts := &tableStatics{
		timer: newMyTimer(),
		db:    db,
		table: table,
		alter: sd,
	}
	if sd.Type != alterTypeNo {
		s.tables = append(s.tables, ts)
	}
	return ts
}

func (s *statics) toHTML() string {
	code := "<h2>Result</h2>\n"
	code += "<h3>Tables</h3>\n"
	code += `<table class='tb_1'>
		<thead>
			<tr>
			<th width="60px">no</th>
			<th>table</th>
			<th>alter result</th>
			<th>used</th>
			</tr>
		</thead><tbody>
		`
	for idx, tb := range s.tables {
		code += "<tr>"
		code += "<td>" + fmt.Sprintf("%d", idx+1) + "</td>\n"
		code += "<td><a href='#detail_" + tb.db + "." + tb.table + "'>" + tb.db + "." + tb.table + "</a></td>\n"
		code += "<td>"
		if s.Config.Sync {
			if tb.alterRet == nil {
				code += "<font color=green>success</font>"
			} else {
				code += "<font color=red>failed," + html.EscapeString(tb.alterRet.Error()) + "</font>"
			}
		} else {
			code += "not sync"
		}
		code += "</td>\n"
		code += "<td>" + tb.timer.usedSecond() + "</td>\n"
		code += "</tr>\n"
	}
	code += "</tbody></table>\n<h3>Sqls</h3>\n<pre>"
	for _, tb := range s.tables {
		code += "<a name='detail_" + tb.table + "'></a>"
		code += html.EscapeString(tb.alter.String()) + "\n\n"
	}
	code += "</pre>\n\n"

	code += "<h3>Detail</h3>\n"
	code += `<table class='tb_1'>
		<thead>
			<tr>
			<th width="40px">no</th>
			<th width="80px">table</th>
			<th>&nbsp;</th>
			<th>&nbsp;</th>
			</tr>
		</thead><tbody>
		`
	for idx, tb := range s.tables {
		code += "<tr>"
		code += "<th rowspan=2>" + fmt.Sprintf("%d", idx+1) + "</th>\n"
		code += "<td rowspan=2>" + tb.table + "<br/><br/>"
		if s.Config.Sync {
			if tb.alterRet == nil {
				code += "<font color=green>success</font>"
			} else {
				code += "<font color=red>failed," + tb.alterRet.Error() + "</font>"
			}
		} else {
			code += "no sync"
		}
		code += "</td>\n"
		code += "<td valign=top><b>source schema:</b><br/>" + htmlPre(tb.alter.SchemaDiff.Source.SchemaRaw) + "</td>\n"
		code += "<td valign=top><b>dest schema:</b><br/>" + htmlPre(tb.alter.SchemaDiff.Dest.SchemaRaw) + "</td>\n"
		code += "</tr>\n"

		code += "<tr>\n"
		code += "<td valign=top><b>alter:</b><br/>" + htmlPre(tb.alter.SQL) + "</td>\n"
		code += "<td valign=top><b>alter after:</b><br/>" + htmlPre(tb.schemaAfter) + "</td>\n"
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
	if cfg.Email == nil {
		log.Println("mail conf is not set,skip send mail")
		return
	}
	alterTotal := len(s.tables)
	if alterTotal < 1 {
		log.Println("no table change,skip send mail")
		return
	}

	title := ""
	body := ""

	if !s.Config.Sync {
		title += "[表结构变化预览] "
		body += "<font color=red> 以下为 SQL 预览，并未执行!</font>\n"
	} else {
		title += "[[表结构自动同步] "
	}

	title += fmt.Sprintf("%d", alterTotal) + " tables change [" + dsnSort(cfg.DestDSN) + "]"

	hostName, _ := os.Hostname()
	body += "<h2>Info</h2>\n<pre>"
	body += "  from : " + dsnSort(cfg.SourceDSN) + "\n"
	body += "    to : " + dsnSort(cfg.DestDSN) + "\n"
	body += " alter : " + fmt.Sprintf("%d", len(s.tables)) + " tables\n"
	body += "<font color=green>  sync : " + fmt.Sprintf("%t", s.Config.Sync) + "</font>\n"
	if s.Config.Sync {
		fn := s.alterFailedNum()
		body += "<font color=red>failed : " + fmt.Sprintf("%d", fn) + "</font>\n"
		if fn > 0 {
			title += " [failed=" + fmt.Sprintf("%d", fn) + "]"
		}
	}
	body += "\n"
	body += "  host : " + hostName + "\n"
	body += " start : " + s.timer.start.Format(timeFormatStd) + "\n"
	body += "   end : " + s.timer.end.Format(timeFormatStd) + "\n"
	body += "  used : " + s.timer.usedSecond() + "\n"

	body += "</pre>\n"
	body += s.toHTML()
	cfg.Email.SendMail(title, body)
}
