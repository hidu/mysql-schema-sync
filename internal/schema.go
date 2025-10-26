package internal

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/xanygo/anygo/ds/xmap"
)

// MySchema table schema
type MySchema struct {
	Fields     xmap.Ordered[string, string] // Legacy: field name -> field definition string
	FieldInfos map[string]*FieldInfo        // New: structured field information
	IndexAll   map[string]*DbIndex
	ForeignAll map[string]*DbIndex
	SchemaRaw  string
}

func (mys *MySchema) String() string {
	if mys.Fields.Len() == 0 {
		return "nil"
	}
	var buf bytes.Buffer
	buf.WriteString("Fields:\n")
	for name, v := range mys.Fields.Keys() {
		buf.WriteString(fmt.Sprintf(" %v : %s\n", name, v))
	}

	if len(mys.FieldInfos) > 0 {
		buf.WriteString("FieldInfos:\n")
		for name, info := range mys.FieldInfos {
			buf.WriteString(fmt.Sprintf(" %s : %s (charset: %v, collation: %v)\n",
				name, info.String(), info.CharsetName, info.CollationName))
		}
	}

	buf.WriteString("Index:\n")
	for name, idx := range mys.IndexAll {
		buf.WriteString(fmt.Sprintf(" %s : %s\n", name, idx.SQL))
	}
	buf.WriteString("ForeignKey:\n")
	for name, idx := range mys.ForeignAll {
		buf.WriteString(fmt.Sprintf("  %s : %s\n", name, idx.SQL))
	}
	return buf.String()
}

// GetFieldNames table names
func (mys *MySchema) GetFieldNames() []string {
	return mys.Fields.Keys()
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
		FieldInfos: make(map[string]*FieldInfo),
		IndexAll:   make(map[string]*DbIndex),
		ForeignAll: make(map[string]*DbIndex),
	}

	for i := 1; i < len(lines)-1; i++ {
		line := strings.TrimSpace(lines[i])
		if len(line) == 0 {
			continue
		}

		line = strings.TrimRight(line, ",")
		switch line[0] {
		case '`':
			index := strings.Index(line[1:], "`")
			name := line[1 : index+1]
			mys.Fields.Set(name, line)

		case '"':
			index := strings.Index(line[1:], "\"")
			name := line[1 : index+1]
			mys.Fields.Set(name, line)

		default:
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

// NewSchemaWithFieldInfos creates a MySchema with structured field information
func NewSchemaWithFieldInfos(schema string, fieldInfos map[string]*FieldInfo) *MySchema {
	mys := ParseSchema(schema)
	if mys != nil {
		mys.FieldInfos = fieldInfos
	}
	return mys
}

// NewSchemaDiffWithFieldInfos creates a SchemaDiff with structured field information
func NewSchemaDiffWithFieldInfos(table, sourceSchema, destSchema string, sourceFields, destFields map[string]*FieldInfo) *SchemaDiff {
	return &SchemaDiff{
		Table:  table,
		Source: NewSchemaWithFieldInfos(sourceSchema, sourceFields),
		Dest:   NewSchemaWithFieldInfos(destSchema, destFields),
	}
}

func (sdiff *SchemaDiff) RelationTables() []string {
	return sdiff.Source.RelationTables()
}
