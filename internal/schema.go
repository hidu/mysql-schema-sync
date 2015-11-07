package internal

import (
	"fmt"
	"strings"
)

// MySchema table schema
type MySchema struct {
	Fields     map[string]string
	IndexAll   map[string]*DbIndex
	ForeignAll map[string]*DbIndex
}

func (mys *MySchema) String() string {
	s := "Fields:\n"
	fl := maxMapKeyLen(mys.Fields, 2)
	for name, v := range mys.Fields {
		s += fmt.Sprintf("  %"+fl+"s : %s\n", name, v)
	}

	s += "Index:\n"
	fl = maxMapKeyLen(mys.IndexAll, 2)
	for name, idx := range mys.IndexAll {
		s += fmt.Sprintf("  %"+fl+"s : %s\n", name, idx.SQL)
	}
	s += "ForeignKey:\n"
	fl = maxMapKeyLen(mys.ForeignAll, 2)
	for name, idx := range mys.ForeignAll {
		s += fmt.Sprintf("  %"+fl+"s : %s\n", name, idx.SQL)
	}
	return s
}

// GetFieldNames table names
func (mys *MySchema) GetFieldNames() []string {
	var names []string
	for name := range mys.Fields {
		names = append(names, name)
	}
	return names
}

// ParseSchema parse table's schema
func ParseSchema(schema string) *MySchema {
	schema = strings.TrimSpace(schema)
	lines := strings.Split(schema, "\n")
	mys := &MySchema{
		Fields:     make(map[string]string),
		IndexAll:   make(map[string]*DbIndex, 0),
		ForeignAll: make(map[string]*DbIndex, 0),
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
			idx := parseDbIndexLine(line)
			if idx == nil {
				continue
			}
			switch idx.indexType {
			case indexTypeForeignKey:
				mys.ForeignAll[idx.Name] = idx
			default:
				mys.IndexAll[idx.Name] = idx
			}
		}
	}
//	fmt.Println(schema)
//	fmt.Println(mys)
//	fmt.Println("-----")
	return mys

}
