// Copyright(C) 2022 github.com/fsgo  All Rights Reserved.
// Author: hidu <duv123@gmail.com>
// Date: 2022/9/25

package internal

import (
	"testing"

	"github.com/stretchr/testify/require"
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.field1.Equals(tt.field2)
			require.Equal(t, tt.equal, got, "Fields should be equal: %v", tt.equal)
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
				CharsetName:   stringPtr("utf8mb4"),
				CollationName: stringPtr("utf8mb4_general_ci"),
			},
			want: "`name` varchar(64) NOT NULL",
		},
		{
			name: "field with default value",
			field: &FieldInfo{
				ColumnName:     "status",
				ColumnType:     "tinyint",
				IsNullAble:     "NO",
				ColumnDefault:  stringPtr("0"),
				CharsetName:    nil,
				CollationName:  nil,
			},
			want: "`status` tinyint NOT NULL DEFAULT 0",
		},
		{
			name: "nullable field",
			field: &FieldInfo{
				ColumnName:    "description",
				ColumnType:    "text",
				IsNullAble:    "YES",
				CharsetName:   stringPtr("utf8mb4"),
				CollationName: stringPtr("utf8mb4_general_ci"),
			},
			want: "`description` text NULL",
		},
		{
			name: "timestamp with auto update",
			field: &FieldInfo{
				ColumnName:     "updated_at",
				ColumnType:     "timestamp",
				IsNullAble:     "NO",
				ColumnDefault:  stringPtr("CURRENT_TIMESTAMP"),
				Extra:          "on update CURRENT_TIMESTAMP",
				CharsetName:    nil,
				CollationName:  nil,
			},
			want: "`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.field.String()
			require.Equal(t, tt.want, got)
		})
	}
}

// Helper function to create string pointers
func stringPtr(s string) *string {
	return &s
}