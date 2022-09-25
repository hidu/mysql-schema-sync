package internal

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/elliotchance/orderedmap"
)

// MySchema table schema
type MySchema struct {
	Fields     *orderedmap.OrderedMap
	IndexAll   map[string]*DbIndex
	ForeignAll map[string]*DbIndex
	SchemaRaw  string
}

func (mys *MySchema) String() string {
	var buf bytes.Buffer
	buf.WriteString("Fields:\n")
	fl := maxMapKeyLen(mys.Fields, 2)
	for name, v := range mys.Fields.Keys() {
		buf.WriteString(fmt.Sprintf("  %"+fl+"s : %s\n", name, v))
	}

	buf.WriteString("Index:\n")
	fl = maxMapKeyLen(mys.IndexAll, 2)
	for name, idx := range mys.IndexAll {
		buf.WriteString(fmt.Sprintf("  %"+fl+"s : %s\n", name, idx.SQL))
	}
	buf.WriteString("ForeignKey:\n")
	fl = maxMapKeyLen(mys.ForeignAll, 2)
	for name, idx := range mys.ForeignAll {
		buf.WriteString(fmt.Sprintf("  %"+fl+"s : %s\n", name, idx.SQL))
	}
	return buf.String()
}

// GetFieldNames table names
func (mys *MySchema) GetFieldNames() []string {
	var names []string
	for _, name := range mys.Fields.Keys() {
		names = append(names, name.(string))
	}
	return names
}

func (mys *MySchema) RelationTables() []string {
	tbs := make(map[string]int)
	for _, idx := range mys.ForeignAll {
		for _, tb := range idx.RelationTables {
			tbs[tb] = 1
		}
	}
	var tables []string
	for tb := range tbs {
		tables = append(tables, tb)
	}
	return tables
}

// ParseSchema parse table's schema
func ParseSchema(schema string) *MySchema {
	schema = strings.TrimSpace(schema)
	lines := strings.Split(schema, "\n")
	mys := &MySchema{
		SchemaRaw:  schema,
		Fields:     orderedmap.NewOrderedMap(),
		IndexAll:   make(map[string]*DbIndex),
		ForeignAll: make(map[string]*DbIndex),
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
			mys.Fields.Set(name, line)
			// log.Println("$$$$$$$$$$$$$$$$111111 = ", name)
			// log.Printf("%+v\n", line)
			// log.Println("$$$$$$$$$$$$$$$$222222")

		} else {
			idx := parseDbIndexLine(line)
			if idx == nil {
				continue
			}
			switch idx.IndexType {
			case indexTypeForeignKey:
				mys.ForeignAll[idx.Name] = idx
			default:
				mys.IndexAll[idx.Name] = idx
			}
		}
	}
	// fmt.Println(schema)
	// fmt.Println(mys)
	// fmt.Println("-----")
	return mys

}

type SchemaDiff struct {
	Source *MySchema
	Dest   *MySchema
	Table  string
}

func newSchemaDiff(table, source, dest string) *SchemaDiff {
	return &SchemaDiff{
		Table:  table,
		Source: ParseSchema(source),
		Dest:   ParseSchema(dest),
	}
}

func (sdiff *SchemaDiff) RelationTables() []string {
	return sdiff.Source.RelationTables()
}
