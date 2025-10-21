package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strings"
)

// DbIndex db index
type DbIndex struct {
	IndexType indexType
	Name      string
	SQL       string

	// 相关联的表
	RelationTables []string
}

type indexType string

const (
	indexTypePrimary    indexType = "PRIMARY"
	indexTypeIndex      indexType = "INDEX"
	indexTypeForeignKey indexType = "FOREIGN KEY"
	checkConstraint     indexType = "CHECK"
)

func (idx *DbIndex) alterAddSQL(drop bool) []string {
	var alterSQL []string
	if drop {
		dropSQL := idx.alterDropSQL()
		if len(dropSQL) != 0 {
			alterSQL = append(alterSQL, dropSQL)
		}
	}

	switch idx.IndexType {
	case indexTypePrimary:
		alterSQL = append(alterSQL, "ADD "+idx.SQL)
	case indexTypeIndex, indexTypeForeignKey:
		alterSQL = append(alterSQL, fmt.Sprintf("ADD %s", idx.SQL))
	case checkConstraint:
		alterSQL = append(alterSQL, fmt.Sprintf("ADD %s", idx.SQL))
	default:
		log.Fatalln("unknown indexType", idx.IndexType)
	}
	return alterSQL
}

func (idx *DbIndex) String() string {
	bs, _ := json.MarshalIndent(idx, "  ", " ")
	return string(bs)
}

func (idx *DbIndex) alterDropSQL() string {
	switch idx.IndexType {
	case indexTypePrimary:
		return "DROP PRIMARY KEY"
	case indexTypeIndex:
		return fmt.Sprintf("DROP INDEX `%s`", idx.Name)
	case indexTypeForeignKey:
		return fmt.Sprintf("DROP FOREIGN KEY `%s`", idx.Name)
	case checkConstraint:
		return fmt.Sprintf("DROP CHECK `%s`", idx.Name)
	default:
		log.Fatalln("unknown indexType", idx.IndexType)
	}
	return ""
}

func (idx *DbIndex) addRelationTable(table string) {
	table = strings.TrimSpace(table)
	if len(table) != 0 {
		idx.RelationTables = append(idx.RelationTables, table)
	}
}

// 匹配索引字段
var indexReg = regexp.MustCompile(`^([A-Z]+\s)?KEY\s`)

// 匹配外键
var foreignKeyReg = regexp.MustCompile("^CONSTRAINT `(.+)` FOREIGN KEY.+ REFERENCES `(.+)` ")

// Check约束
var checkConstraintReg = regexp.MustCompile("^CONSTRAINT `([^`]+)` CHECK \\(\\((.+)\\)\\)")

func parseDbIndexLine(line string) *DbIndex {
	line = strings.TrimSpace(line)
	idx := &DbIndex{
		SQL:            line,
		RelationTables: []string{},
	}
	if strings.HasPrefix(line, "PRIMARY") {
		idx.IndexType = indexTypePrimary
		idx.Name = "PRIMARY KEY"
		return idx
	}

	// UNIQUE KEY `idx_a` (`a`) USING HASH COMMENT '注释',
	// FULLTEXT KEY `c` (`c`)
	// PRIMARY KEY (`d`)
	// KEY `idx_e` (`e`),
	if indexReg.MatchString(line) {
		arr := strings.Split(line, "`")
		idx.IndexType = indexTypeIndex
		idx.Name = arr[1]
		return idx
	}

	// CONSTRAINT `busi_table_ibfk_1` FOREIGN KEY (`repo_id`) REFERENCES `repo_table` (`repo_id`)
	foreignMatches := foreignKeyReg.FindStringSubmatch(line)
	if len(foreignMatches) > 0 {
		idx.IndexType = indexTypeForeignKey
		idx.Name = foreignMatches[1]
		idx.addRelationTable(foreignMatches[2])
		return idx
	}

	// CONSTRAINT `chk_xx_1` CHECK ((`x` >= 0 and `y` <= 100))
	checkMatches := checkConstraintReg.FindStringSubmatch(line)
	if len(checkMatches) > 0 {
		idx.IndexType = checkConstraint
		idx.Name = checkMatches[1]
		return idx
	}

	log.Fatalln("db_index parse failed, unsupported, line:", line)
	return nil
}
