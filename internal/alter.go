package internal

import (
	"fmt"
	"strings"
)

type alterType int

const (
	alterTypeNo     alterType = 0
	alterTypeCreate           = 1
	alterTypeDrop             = 2
	alterTypeAlter            = 3
)

func (at alterType) String() string {
	switch at {
	case alterTypeNo:
		return "not_change"
	case alterTypeCreate:
		return "create"
	case alterTypeDrop:
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
	SQL        string
	SchemaDiff *SchemaDiff
}

func (ta *TableAlterData) String() string {
	relationTables := ta.SchemaDiff.RelationTables()
	fmtStr := `
-- Table : %s
-- Type  : %s
-- RelationTables : %s
-- SQL   : 
%s
`
	return fmt.Sprintf(fmtStr, ta.Table, ta.Type, strings.Join(relationTables, ","), ta.SQL)
}
