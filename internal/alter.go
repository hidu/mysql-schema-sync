package internal

import (
	"fmt"
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
		return "unknow"
	}

}

// TableAlterData 表的变更情况
type TableAlterData struct {
	Table        string
	Type         alterType
	SQL          string
	SourceSchema string
	DestSchema   string
}

func (ta *TableAlterData) String() string {
	return fmt.Sprintf("-- Table : %s\n-- Type  : %s\n-- SQL   :\n%s", ta.Table, ta.Type, ta.SQL)
}
