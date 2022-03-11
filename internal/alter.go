package internal

import (
	"fmt"
	"regexp"
	"strings"
)

type alterType int

const (
	alterTypeNo alterType = iota
	alterTypeCreate
	alterTypeDropTable
	alterTypeAlter
)

func (at alterType) String() string {
	switch at {
	case alterTypeNo:
		return "not_change"
	case alterTypeCreate:
		return "create"
	case alterTypeDropTable:
		return "drop"
	case alterTypeAlter:
		return "alter"
	default:
		return "unknown"
	}

}

// TableAlterData 表的变更情况
type TableAlterData struct {
	Table      string
	Type       alterType
	Comment    string
	SQL        []string
	SchemaDiff *SchemaDiff
}

func (ta *TableAlterData) String() string {
	relationTables := ta.SchemaDiff.RelationTables()
	sqlTpl := `
-- Table : %s
-- Type  : %s
-- RelationTables : %s
-- Comment: %s
-- SQL   : 
%s
`
	return fmt.Sprintf(sqlTpl, ta.Table, ta.Type, strings.Join(relationTables, ","), ta.Comment, strings.Join(ta.SQL, "\n"))
}

var autoIncrReg = regexp.MustCompile(`\sAUTO_INCREMENT=[1-9]\d*\s`)

func fmtTableCreateSQL(sql string) string {
	return autoIncrReg.ReplaceAllString(sql, " ")
}
