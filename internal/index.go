package internal

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

// DbIndex db index
type DbIndex struct {
	indexType     indexType
	Name          string
	SQL           string
	RelationTbles []string //相关联的表
}

type indexType string

const (
	indexTypePrimary    indexType = "PRIMARY"
	indexTypeIndex                = "INDEX"
	indexTypeForeignKey           = "FOREIGN KEY"
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
	case indexTypeIndex, indexTypeForeignKey:
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
	case indexTypeForeignKey:
		return fmt.Sprintf("DROP FOREIGN KEY `%s`", idx.Name)
	default:
		log.Fatalln("unknow indexType", idx.indexType)
	}
	return ""
}

func (idx *DbIndex) addRelationTable(table string) {
	table = strings.TrimSpace(table)
	if table != "" {
		idx.RelationTbles = append(idx.RelationTbles, table)
	}
}

//匹配索引字段
var indexReg = regexp.MustCompile(`^([A-Z]+\s)?KEY\s`)

//匹配外键
var foreignKeyReg = regexp.MustCompile("^CONSTRAINT `(.+)` FOREIGN KEY.+ REFERENCES `(.+)` ")

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

	//CONSTRAINT `busi_table_ibfk_1` FOREIGN KEY (`repo_id`) REFERENCES `repo_table` (`repo_id`)
	foreignMatches := foreignKeyReg.FindStringSubmatch(line)
	if len(foreignMatches) > 0 {
		idx.indexType = indexTypeForeignKey
		idx.Name = foreignMatches[1]
		idx.addRelationTable(foreignMatches[2])
		return idx
	}

	log.Fatalln("db_index parse failed,unsupport,line:", line)
	return nil
}
