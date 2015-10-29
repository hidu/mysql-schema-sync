package internal

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

// DbIndex db index
type DbIndex struct {
	indexType indexType
	Name      string
	SQL       string
}

type indexType string

const (
	indexTypePrimary indexType = "PRIMARY"
	indexTypeIndex             = "index"
)

func (idx *DbIndex) alterAddSQL(drop bool) string {
	alterSQL := []string{}
	if drop {
		dropSQL := idx.alterDropSQL()
		if dropSQL != "" {
			alterSQL = append(alterSQL, dropSQL)
		}
	}

	switch idx.indexType {
	case indexTypePrimary:
		alterSQL = append(alterSQL, "ADD "+idx.SQL)
	case indexTypeIndex:
		alterSQL = append(alterSQL, fmt.Sprintf("ADD %s", idx.SQL))
	default:
		log.Fatalln("unknow indexType", idx.indexType)
	}
	return strings.Join(alterSQL, ",\n")
}

func (idx *DbIndex) alterDropSQL() string {
	switch idx.indexType {
	case indexTypePrimary:
		return "DROP PRIMARY KEY"
	case indexTypeIndex:
		return fmt.Sprintf("DROP INDEX `%s`", idx.Name)
	default:
		log.Fatalln("unknow indexType", idx.indexType)
	}
	return ""
}

var indexReg = regexp.MustCompile(`^([A-Z]+\s)?KEY\s`)

func parseDbIndexLine(line string) *DbIndex {
	line = strings.TrimSpace(line)
	idx := &DbIndex{
		SQL: line,
	}
	if strings.HasPrefix(line, "PRIMARY") {
		idx.indexType = indexTypePrimary
		idx.Name = "PRIMARY KEY"
		return idx
	}

	//  UNIQUE KEY `idx_a` (`a`) USING HASH COMMENT '注释',
	//  FULLTEXT KEY `c` (`c`)
	//  PRIMARY KEY (`d`)
	//  KEY `idx_e` (`e`),
	if indexReg.MatchString(line) {
		arr := strings.Split(line, "`")
		idx.indexType = indexTypeIndex
		idx.Name = arr[1]
		return idx
	}
	log.Fatalln("db_index parse failed,unsupport,line:", line)
	return nil
}
