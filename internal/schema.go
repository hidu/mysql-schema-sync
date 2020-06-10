package internal

import (
	"bytes"
	"fmt"
	"log"
	"strings"
)

// MySchema table schema
type MySchema struct {
	SchemaRaw  string
	Fields     map[string]string
	IndexAll   map[string]*DbIndex
	ForeignAll map[string]*DbIndex
}

func (mys *MySchema) String() string {
	var buf bytes.Buffer
	buf.WriteString("Fields:\n")
	fl := maxMapKeyLen(mys.Fields, 2)
	for name, v := range mys.Fields {
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
	for name := range mys.Fields {
		names = append(names, name)
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
		Fields:     make(map[string]string),
		IndexAll:   make(map[string]*DbIndex, 0),
		ForeignAll: make(map[string]*DbIndex, 0),
	}

	iTabBodyEnd := len(lines) - 1
	for i := 1; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if strings.Index(line, ") ENGINE=") != 0 {
			continue
		} else {
			iTabBodyEnd = i
			break
		}
	}

	if iTabBodyEnd != len(lines)-1 {
		// not accurate still - but reduced the parsing work
		// - would fatal definitely if continue to parse this ddl
		// and given a bit clear fatal infomation about why failed/exit
		log.Fatalln("some ddl (e.g 'partition') not supported:", strings.Join(lines[(iTabBodyEnd+1):], "\n"))
	}

	for i := 1; i < iTabBodyEnd; i++ {
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
	Table  string
	Source *MySchema
	Dest   *MySchema
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
