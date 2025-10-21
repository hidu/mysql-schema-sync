package internal

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql" // mysql driver
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
	Extra                  string  `json:"extra"`
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
		if strings.ToUpper(*f.ColumnDefault) == "CURRENT_TIMESTAMP" {
			parts = append(parts, "DEFAULT CURRENT_TIMESTAMP")
		} else {
			parts = append(parts, fmt.Sprintf("DEFAULT %s", *f.ColumnDefault))
		}
	}

	// Extra
	if f.Extra != "" {
		parts = append(parts, strings.ToUpper(f.Extra))
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
		f.ColumnType != other.ColumnType ||
		f.Extra != other.Extra {
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

// MyDb db struct
type MyDb struct {
	Db     *sql.DB
	dbType string
}

// NewMyDb parse dsn
func NewMyDb(dsn string, dbType string) *MyDb {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Sprintf("connected to db [%s] failed,err=%s", dsn, err))
	}
	return &MyDb{
		Db:     db,
		dbType: dbType,
	}
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
		log.Println(err)
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

// GetTableFieldsFromInformationSchema retrieves detailed field information from INFORMATION_SCHEMA.COLUMNS
func (db *MyDb) GetTableFieldsFromInformationSchema(tableName string) (map[string]*FieldInfo, error) {
	// Check if database connection is available
	if db == nil || db.Db == nil {
		return nil, errors.New("database connection is nil")
	}

	// Extract database name from DSN or use current database
	dbName := db.getDatabaseName()
	if dbName == "" {
		return nil, errors.New("could not determine database name from DSN")
	}

	query := `
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
			EXTRA
		FROM INFORMATION_SCHEMA.COLUMNS
		WHERE TABLE_SCHEMA = ? AND TABLE_NAME = ?
		ORDER BY ORDINAL_POSITION`

	log.Println("[SQL]", "["+db.dbType+"]", query, "args:", dbName, tableName)

	rows, err := db.Query(query, dbName, tableName)
	if err != nil {
		return nil, fmt.Errorf("failed to query INFORMATION_SCHEMA.COLUMNS for table %s: %v", tableName, err)
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
			&field.Extra,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan field information for table %s: %v", tableName, err)
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
		return nil, fmt.Errorf("error iterating field information for table %s: %v", tableName, err)
	}

	if len(fields) == 0 {
		return nil, fmt.Errorf("no fields found for table %s in database %s", tableName, dbName)
	}

	return fields, nil
}

// getDatabaseName extracts database name from the current database connection
func (db *MyDb) getDatabaseName() string {
	if db == nil || db.Db == nil {
		log.Print("database connection is nil")
		return ""
	}
	var dbName string
	err := db.Db.QueryRow("SELECT DATABASE()").Scan(&dbName)
	if err != nil {
		log.Printf("failed to get current database name: %v", err)
		return ""
	}
	return dbName
}

// Query execute sql query
func (db *MyDb) Query(query string, args ...any) (*sql.Rows, error) {
	log.Println("[SQL]", "["+db.dbType+"]", query, args)
	return db.Db.Query(query, args...)
}
