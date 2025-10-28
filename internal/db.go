package internal

import (
	"database/sql"
	"fmt"
	"log"
	"slices"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql" // mysql driver
	"github.com/xanygo/anygo/cli/xcolor"
)

// FieldInfo represents detailed field information from INFORMATION_SCHEMA.COLUMNS
type FieldInfo struct {
	ColumnName             string  `json:"column_name"`
	OrdinalPosition        int     `json:"ordinal_position"`
	ColumnDefault          *string `json:"column_default"`
	IsNullAble             string  `json:"is_nullable"`
	DataType               string  `json:"data_type"`
	CharacterMaximumLength *int    `json:"character_maximum_length"`
	NumericPrecision       *int    `json:"numeric_precision"`
	NumericScale           *int    `json:"numeric_scale"`
	CharsetName            *string `json:"character_set_name"`
	CollationName          *string `json:"collation_name"`
	ColumnType             string  `json:"column_type"`
	ColumnComment          string  `json:"column_comment"`
	Extra                  string  `json:"extra"`
}

// needsQuotedDefault returns true if the field type requires quoted default values
func (f *FieldInfo) needsQuotedDefault() bool {
	// String types that need quoted default values
	stringTypes := []string{
		"char", "varchar", "binary", "varbinary",
		"tinyblob", "blob", "mediumblob", "longblob",
		"tinytext", "text", "mediumtext", "longtext",
		"enum", "set", "json",
	}

	dataType := strings.ToLower(f.DataType)
	return slices.Contains(stringTypes, dataType)
}

// String returns the full column definition as used in CREATE TABLE
func (f *FieldInfo) String() string {
	var parts []string

	// Column name and type
	parts = append(parts, fmt.Sprintf("`%s` %s", f.ColumnName, f.ColumnType))

	// NULL/NOT NULL
	if strings.ToUpper(f.IsNullAble) == "NO" {
		parts = append(parts, "NOT NULL")
	} else {
		parts = append(parts, "NULL")
	}

	// Default value
	if f.ColumnDefault != nil {
		defaultValue := *f.ColumnDefault
		upperDefault := strings.ToUpper(defaultValue)

		// Special keywords that don't need quotes
		if upperDefault == "CURRENT_TIMESTAMP" || upperDefault == "NULL" {
			parts = append(parts, fmt.Sprintf("DEFAULT %s", upperDefault))
		} else if f.needsQuotedDefault() {
			// String types need quotes
			parts = append(parts, fmt.Sprintf("DEFAULT '%s'", defaultValue))
		} else {
			// Numeric types don't need quotes
			parts = append(parts, fmt.Sprintf("DEFAULT %s", defaultValue))
		}
	}

	// Extra
	if f.Extra != "" {
		parts = append(parts, strings.ToUpper(f.Extra))
	}

	// Comment
	if f.ColumnComment != "" {
		// Escape single quotes in comment by doubling them
		escapedComment := strings.ReplaceAll(f.ColumnComment, "'", "''")
		parts = append(parts, fmt.Sprintf("COMMENT '%s'", escapedComment))
	}

	return strings.Join(parts, " ")
}

// Equals compares two FieldInfo instances for semantic equality
func (f *FieldInfo) Equals(other *FieldInfo) bool {
	if f == nil || other == nil {
		return f == other
	}

	// Compare basic properties
	if f.ColumnName != other.ColumnName ||
		f.IsNullAble != other.IsNullAble ||
		f.DataType != other.DataType ||
		f.ColumnComment != other.ColumnComment ||
		f.Extra != other.Extra {
		return false
	}

	// Compare ColumnType with normalization for integer display width
	// MySQL 8.0.19+ removed display width for integer types (int(11) -> int)
	normalizedSourceType := normalizeIntegerType(f.ColumnType)
	normalizedDestType := normalizeIntegerType(other.ColumnType)
	if normalizedSourceType != normalizedDestType {
		return false
	}

	// Compare default values
	if (f.ColumnDefault == nil && other.ColumnDefault != nil) ||
		(f.ColumnDefault != nil && other.ColumnDefault == nil) {
		return false
	}
	if f.ColumnDefault != nil && other.ColumnDefault != nil {
		if *f.ColumnDefault != *other.ColumnDefault {
			return false
		}
	}

	// Compare character set and collation (handle NULL values gracefully)
	// For charset and collation, we consider them equal if:
	// 1. Both are NULL, or
	// 2. One is NULL and the other uses the default/collation, or
	// 3. Both are set and equal
	if !f.charsetEquals(other) || !f.collationEquals(other) {
		return false
	}

	return true
}

// charsetEquals checks if character sets are semantically equal
func (f *FieldInfo) charsetEquals(other *FieldInfo) bool {
	// Both NULL
	if f.CharsetName == nil && other.CharsetName == nil {
		return true
	}

	// One NULL, one not NULL
	if (f.CharsetName == nil) != (other.CharsetName == nil) {
		// If one is NULL, check if the other is the default charset
		if f.CharsetName != nil {
			return *f.CharsetName == "utf8mb4" || *f.CharsetName == "utf8" || *f.CharsetName == "latin1"
		}
		return *other.CharsetName == "utf8mb4" || *other.CharsetName == "utf8" || *other.CharsetName == "latin1"
	}

	// Both not NULL, compare values
	return *f.CharsetName == *other.CharsetName
}

// collationEquals checks if collations are semantically equal
func (f *FieldInfo) collationEquals(other *FieldInfo) bool {
	// Both NULL
	if f.CollationName == nil && other.CollationName == nil {
		return true
	}

	// One NULL, one not NULL
	if (f.CollationName == nil) != (other.CollationName == nil) {
		// If one is NULL, check if the other is the default collation
		if f.CollationName != nil {
			return *f.CollationName == "utf8mb4_general_ci" ||
				*f.CollationName == "utf8mb4_unicode_ci" ||
				*f.CollationName == "utf8_general_ci" ||
				*f.CollationName == "latin1_swedish_ci"
		}
		return *other.CollationName == "utf8mb4_general_ci" ||
			*other.CollationName == "utf8mb4_unicode_ci" ||
			*other.CollationName == "utf8_general_ci" ||
			*other.CollationName == "latin1_swedish_ci"
	}

	// Both not NULL, compare values
	return *f.CollationName == *other.CollationName
}

type dbType string

const (
	dbTypeSource = "source"
	dbTypeDest   = "dest"
)

// MyDb db struct
type MyDb struct {
	sqlDB  *sql.DB
	dbType dbType
	dbName string // 数据库名称
}

// NewMyDb parse dsn
func NewMyDb(dsn string, dbType dbType) *MyDb {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Sprintf("connected to db [%s] failed,err=%s", dsn, err))
	}
	dbName, err := getDatabaseName(db)
	if err != nil {
		panic(fmt.Sprintf("get database name failed,err=%s", err))
	}
	return &MyDb{
		sqlDB:  db,
		dbType: dbType,
		dbName: dbName,
	}
}

// getDatabaseName extracts database name from the current database connection
func getDatabaseName(db *sql.DB) (string, error) {
	var dbName string
	const query = "SELECT DATABASE()"
	err := db.QueryRow(query).Scan(&dbName)
	if err != nil {
		log.Printf("QueryRow %q, Result=%q, Err=%v", query, dbName, err)
	}
	return dbName, err
}

// GetTableNames table names
func (db *MyDb) GetTableNames() []string {
	rs, err := db.Query("show table status")
	if err != nil {
		panic("show tables failed:" + err.Error())
	}
	defer rs.Close()

	var tables []string
	columns, _ := rs.Columns()
	for rs.Next() {
		var values = make([]any, len(columns))
		valuePtrs := make([]any, len(columns))
		for i := range columns {
			valuePtrs[i] = &values[i]
		}
		if err := rs.Scan(valuePtrs...); err != nil {
			panic("show tables failed when scan," + err.Error())
		}
		var valObj = make(map[string]any)
		for i, col := range columns {
			var v any
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			valObj[col] = v
		}
		if valObj["Engine"] != nil {
			tables = append(tables, valObj["Name"].(string))
		}
	}
	return tables
}

// GetTableSchema table schema
func (db *MyDb) GetTableSchema(name string) (schema string) {
	rs, err := db.Query(fmt.Sprintf("show create table `%s`", name))
	if err != nil {
		return
	}
	defer rs.Close()
	for rs.Next() {
		var vname string
		if err := rs.Scan(&vname, &schema); err != nil {
			panic(fmt.Sprintf("get table %s 's schema failed, %s", name, err))
		}
	}
	return
}

// TableFieldsFromInformationSchema retrieves detailed field information from INFORMATION_SCHEMA.COLUMNS
func (db *MyDb) TableFieldsFromInformationSchema(tableName string) (map[string]*FieldInfo, error) {
	const query = `
		SELECT
			COLUMN_NAME,
			ORDINAL_POSITION,
			COLUMN_DEFAULT,
			IS_NULLABLE,
			DATA_TYPE,
			CHARACTER_MAXIMUM_LENGTH,
			NUMERIC_PRECISION,
			NUMERIC_SCALE,
			CHARACTER_SET_NAME,
			COLLATION_NAME,
			COLUMN_TYPE,
			COLUMN_COMMENT,
			EXTRA
		FROM INFORMATION_SCHEMA.COLUMNS
		WHERE TABLE_SCHEMA = ? AND TABLE_NAME = ?
		ORDER BY ORDINAL_POSITION`

	rows, err := db.Query(query, db.dbName, tableName)
	if err != nil {
		return nil, fmt.Errorf("failed to query INFORMATION_SCHEMA.COLUMNS for table %q: %v", tableName, err)
	}
	defer rows.Close()

	fields := make(map[string]*FieldInfo)

	for rows.Next() {
		field := &FieldInfo{}
		var charMaxLen, numericPrecision, numericScale sql.NullInt64
		var charset, collation, columnDefault sql.NullString

		err := rows.Scan(
			&field.ColumnName,
			&field.OrdinalPosition,
			&columnDefault,
			&field.IsNullAble,
			&field.DataType,
			&charMaxLen,
			&numericPrecision,
			&numericScale,
			&charset,
			&collation,
			&field.ColumnType,
			&field.ColumnComment,
			&field.Extra,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan field information for table %q: %v", tableName, err)
		}

		// Handle nullable fields
		if columnDefault.Valid {
			field.ColumnDefault = &columnDefault.String
		}
		if charMaxLen.Valid {
			val := int(charMaxLen.Int64)
			field.CharacterMaximumLength = &val
		}
		if numericPrecision.Valid {
			val := int(numericPrecision.Int64)
			field.NumericPrecision = &val
		}
		if numericScale.Valid {
			val := int(numericScale.Int64)
			field.NumericScale = &val
		}
		if charset.Valid {
			field.CharsetName = &charset.String
		}
		if collation.Valid {
			field.CollationName = &collation.String
		}

		fields[field.ColumnName] = field
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating field information for table %q: %v", tableName, err)
	}

	if len(fields) == 0 {
		return nil, fmt.Errorf("no fields found for table %q in database %q", tableName, db.dbName)
	}

	return fields, nil
}

// Query execute sql query
func (db *MyDb) Query(query string, args ...any) (rows *sql.Rows, err error) {
	txt := fmt.Sprintf("[%-6s: %s] [Query] Start SQL=%s Args=%s\n",
		db.dbType,
		db.dbName,
		xcolor.GreenString("%s", strings.TrimSpace(query)),
		xcolor.GreenString("%v", args),
	)
	log.Output(2, txt)
	start := time.Now()
	defer func() {
		cost := time.Since(start)
		txt = fmt.Sprintf("[%-6s: %s] [Query] End   Cost=%s Err=%s\n", db.dbType, db.dbName, cost.String(), errString(err))
		log.Output(3, txt)
	}()
	return db.sqlDB.Query(query, args...)
}
