// Copyright(C) 2022 github.com/fsgo  All Rights Reserved.
// Author: hidu <duv123@gmail.com>
// Date: 2022/9/25

package internal

import (
	"testing"

	"github.com/xanygo/anygo/xt"
)

func TestFieldInfo_Equals(t *testing.T) {
	tests := []struct {
		name   string
		field1 *FieldInfo
		field2 *FieldInfo
		equal  bool
	}{
		{
			name: "identical fields",
			field1: &FieldInfo{
				ColumnName:    "name",
				ColumnType:    "varchar(64)",
				IsNullAble:    "NO",
				CharsetName:   stringPtr("utf8mb4"),
				CollationName: stringPtr("utf8mb4_general_ci"),
			},
			field2: &FieldInfo{
				ColumnName:    "name",
				ColumnType:    "varchar(64)",
				IsNullAble:    "NO",
				CharsetName:   stringPtr("utf8mb4"),
				CollationName: stringPtr("utf8mb4_general_ci"),
			},
			equal: true,
		},
		{
			name: "same field with and without explicit charset/collation",
			field1: &FieldInfo{
				ColumnName:    "name",
				ColumnType:    "varchar(64)",
				IsNullAble:    "NO",
				CharsetName:   nil,
				CollationName: nil,
			},
			field2: &FieldInfo{
				ColumnName:    "name",
				ColumnType:    "varchar(64)",
				IsNullAble:    "NO",
				CharsetName:   stringPtr("utf8mb4"),
				CollationName: stringPtr("utf8mb4_general_ci"),
			},
			equal: true,
		},
		{
			name: "same field with different charset",
			field1: &FieldInfo{
				ColumnName:    "name",
				ColumnType:    "varchar(64)",
				IsNullAble:    "NO",
				CharsetName:   nil,
				CollationName: nil,
			},
			field2: &FieldInfo{
				ColumnName:    "name",
				ColumnType:    "varchar(64)",
				IsNullAble:    "NO",
				CharsetName:   stringPtr("utf8mb4"),
				CollationName: stringPtr("utf8mb4_general_ci"),
			},
			equal: true,
		},
		{
			name: "different field type",
			field1: &FieldInfo{
				ColumnName:    "name",
				ColumnType:    "varchar(64)",
				IsNullAble:    "NO",
				CharsetName:   stringPtr("utf8mb4"),
				CollationName: stringPtr("utf8mb4_general_ci"),
			},
			field2: &FieldInfo{
				ColumnName:    "name",
				ColumnType:    "varchar(128)",
				IsNullAble:    "NO",
				CharsetName:   stringPtr("utf8mb4"),
				CollationName: stringPtr("utf8mb4_general_ci"),
			},
			equal: false,
		},
		{
			name: "different nullable",
			field1: &FieldInfo{
				ColumnName:    "name",
				ColumnType:    "varchar(64)",
				IsNullAble:    "NO",
				CharsetName:   stringPtr("utf8mb4"),
				CollationName: stringPtr("utf8mb4_general_ci"),
			},
			field2: &FieldInfo{
				ColumnName:    "name",
				ColumnType:    "varchar(64)",
				IsNullAble:    "YES",
				CharsetName:   stringPtr("utf8mb4"),
				CollationName: stringPtr("utf8mb4_general_ci"),
			},
			equal: false,
		},
		{
			name: "same field with default collation",
			field1: &FieldInfo{
				ColumnName:    "name",
				ColumnType:    "varchar(64)",
				IsNullAble:    "NO",
				CharsetName:   nil,
				CollationName: nil,
			},
			field2: &FieldInfo{
				ColumnName:    "name",
				ColumnType:    "varchar(64)",
				IsNullAble:    "NO",
				CharsetName:   stringPtr("utf8mb4"),
				CollationName: stringPtr("utf8mb4_unicode_ci"),
			},
			equal: true,
		},
		// Integer type display width tests (MySQL 5.7 vs 8.0 compatibility)
		{
			name: "int(11) vs int should be equal",
			field1: &FieldInfo{
				ColumnName: "id",
				ColumnType: "int(11)",
				DataType:   "int",
				IsNullAble: "NO",
			},
			field2: &FieldInfo{
				ColumnName: "id",
				ColumnType: "int",
				DataType:   "int",
				IsNullAble: "NO",
			},
			equal: true,
		},
		{
			name: "bigint(20) vs bigint should be equal",
			field1: &FieldInfo{
				ColumnName: "user_id",
				ColumnType: "bigint(20)",
				DataType:   "bigint",
				IsNullAble: "NO",
			},
			field2: &FieldInfo{
				ColumnName: "user_id",
				ColumnType: "bigint",
				DataType:   "bigint",
				IsNullAble: "NO",
			},
			equal: true,
		},
		{
			name: "tinyint(1) vs tinyint should be equal",
			field1: &FieldInfo{
				ColumnName: "is_active",
				ColumnType: "tinyint(1)",
				DataType:   "tinyint",
				IsNullAble: "NO",
			},
			field2: &FieldInfo{
				ColumnName: "is_active",
				ColumnType: "tinyint",
				DataType:   "tinyint",
				IsNullAble: "NO",
			},
			equal: true,
		},
		{
			name: "tinyint(4) vs tinyint should be equal",
			field1: &FieldInfo{
				ColumnName: "status",
				ColumnType: "tinyint(4)",
				DataType:   "tinyint",
				IsNullAble: "NO",
			},
			field2: &FieldInfo{
				ColumnName: "status",
				ColumnType: "tinyint",
				DataType:   "tinyint",
				IsNullAble: "NO",
			},
			equal: true,
		},
		{
			name: "int(11) unsigned vs int unsigned should be equal",
			field1: &FieldInfo{
				ColumnName: "count",
				ColumnType: "int(11) unsigned",
				DataType:   "int",
				IsNullAble: "NO",
			},
			field2: &FieldInfo{
				ColumnName: "count",
				ColumnType: "int unsigned",
				DataType:   "int",
				IsNullAble: "NO",
			},
			equal: true,
		},
		{
			name: "bigint(20) unsigned vs bigint unsigned should be equal",
			field1: &FieldInfo{
				ColumnName: "total",
				ColumnType: "bigint(20) unsigned",
				DataType:   "bigint",
				IsNullAble: "NO",
			},
			field2: &FieldInfo{
				ColumnName: "total",
				ColumnType: "bigint unsigned",
				DataType:   "bigint",
				IsNullAble: "NO",
			},
			equal: true,
		},
		{
			name: "int(10) zerofill vs int zerofill should be equal",
			field1: &FieldInfo{
				ColumnName: "order_id",
				ColumnType: "int(10) zerofill",
				DataType:   "int",
				IsNullAble: "NO",
			},
			field2: &FieldInfo{
				ColumnName: "order_id",
				ColumnType: "int zerofill",
				DataType:   "int",
				IsNullAble: "NO",
			},
			equal: true,
		},
		{
			name: "int(10) unsigned zerofill vs int unsigned zerofill should be equal",
			field1: &FieldInfo{
				ColumnName: "code",
				ColumnType: "int(10) unsigned zerofill",
				DataType:   "int",
				IsNullAble: "NO",
			},
			field2: &FieldInfo{
				ColumnName: "code",
				ColumnType: "int unsigned zerofill",
				DataType:   "int",
				IsNullAble: "NO",
			},
			equal: true,
		},
		{
			name: "int vs bigint should not be equal",
			field1: &FieldInfo{
				ColumnName: "value",
				ColumnType: "int",
				DataType:   "int",
				IsNullAble: "NO",
			},
			field2: &FieldInfo{
				ColumnName: "value",
				ColumnType: "bigint",
				DataType:   "bigint",
				IsNullAble: "NO",
			},
			equal: false,
		},
		{
			name: "int unsigned vs int should not be equal (unsigned modifier difference)",
			field1: &FieldInfo{
				ColumnName: "amount",
				ColumnType: "int unsigned",
				DataType:   "int",
				IsNullAble: "NO",
			},
			field2: &FieldInfo{
				ColumnName: "amount",
				ColumnType: "int",
				DataType:   "int",
				IsNullAble: "NO",
			},
			equal: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.field1.Equals(tt.field2)
			xt.Equal(t, tt.equal, got)
		})
	}
}

func TestFieldInfo_String(t *testing.T) {
	tests := []struct {
		name  string
		field *FieldInfo
		want  string
	}{
		{
			name: "simple varchar field",
			field: &FieldInfo{
				ColumnName:    "name",
				ColumnType:    "varchar(64)",
				IsNullAble:    "NO",
				DataType:      "varchar",
				CharsetName:   stringPtr("utf8mb4"),
				CollationName: stringPtr("utf8mb4_general_ci"),
			},
			want: "`name` varchar(64) NOT NULL",
		},
		{
			name: "field with default value",
			field: &FieldInfo{
				ColumnName:    "status",
				ColumnType:    "tinyint",
				IsNullAble:    "NO",
				DataType:      "tinyint",
				ColumnDefault: stringPtr("0"),
				CharsetName:   nil,
				CollationName: nil,
			},
			want: "`status` tinyint NOT NULL DEFAULT 0",
		},
		{
			name: "varchar field with string default value",
			field: &FieldInfo{
				ColumnName:    "f_status",
				ColumnType:    "varchar(32)",
				IsNullAble:    "NO",
				DataType:      "varchar",
				ColumnDefault: stringPtr("queue"),
				CharsetName:   stringPtr("utf8mb3"),
				CollationName: stringPtr("utf8mb3_general_ci"),
			},
			want: "`f_status` varchar(32) NOT NULL DEFAULT 'queue'",
		},
		{
			name: "char field with string default value",
			field: &FieldInfo{
				ColumnName:    "type",
				ColumnType:    "char(10)",
				IsNullAble:    "NO",
				DataType:      "char",
				ColumnDefault: stringPtr("active"),
			},
			want: "`type` char(10) NOT NULL DEFAULT 'active'",
		},
		{
			name: "text field with string default value",
			field: &FieldInfo{
				ColumnName:    "description",
				ColumnType:    "text",
				IsNullAble:    "YES",
				DataType:      "text",
				ColumnDefault: stringPtr("default text"),
				CharsetName:   stringPtr("utf8mb4"),
				CollationName: stringPtr("utf8mb4_general_ci"),
			},
			want: "`description` text NULL DEFAULT 'default text'",
		},
		{
			name: "int field with numeric default value",
			field: &FieldInfo{
				ColumnName:    "count",
				ColumnType:    "int",
				IsNullAble:    "NO",
				DataType:      "int",
				ColumnDefault: stringPtr("100"),
			},
			want: "`count` int NOT NULL DEFAULT 100",
		},
		{
			name: "nullable field",
			field: &FieldInfo{
				ColumnName:    "description",
				ColumnType:    "text",
				IsNullAble:    "YES",
				DataType:      "text",
				CharsetName:   stringPtr("utf8mb4"),
				CollationName: stringPtr("utf8mb4_general_ci"),
			},
			want: "`description` text NULL",
		},
		{
			name: "timestamp with auto update",
			field: &FieldInfo{
				ColumnName:    "updated_at",
				ColumnType:    "timestamp",
				IsNullAble:    "NO",
				DataType:      "timestamp",
				ColumnDefault: stringPtr("CURRENT_TIMESTAMP"),
				Extra:         "on update CURRENT_TIMESTAMP",
				CharsetName:   nil,
				CollationName: nil,
			},
			want: "`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.field.String()
			xt.Equal(t, tt.want, got)
		})
	}
}

// Helper function to create string pointers
func stringPtr(s string) *string {
	return &s
}
