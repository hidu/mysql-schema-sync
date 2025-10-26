// Copyright(C) 2022 github.com/fsgo  All Rights Reserved.
// Author: hidu <duv123@gmail.com>
// Date: 2022/9/25

package internal

import (
	"fmt"
	"testing"

	"github.com/xanygo/anygo/xt"
)

func TestFieldInfo_CharsetCollationComparison(t *testing.T) {
	// Test the exact scenario described in the issue
	sourceField := &FieldInfo{
		ColumnName:    "name",
		ColumnType:    "varchar(64)",
		IsNullAble:    "NO",
		CharsetName:   nil, // No explicit charset
		CollationName: nil, // No explicit collation
	}

	destField := &FieldInfo{
		ColumnName:    "name",
		ColumnType:    "varchar(64)",
		IsNullAble:    "NO",
		CharsetName:   stringPtr("utf8mb4"),            // Explicit charset
		CollationName: stringPtr("utf8mb4_general_ci"), // Explicit collation
	}

	// These should be considered equal
	xt.True(t, sourceField.Equals(destField))
	xt.True(t, destField.Equals(sourceField))
}

func TestFieldInfo_DifferentCharsetCollation(t *testing.T) {
	// Test fields with actually different charset/collation
	sourceField := &FieldInfo{
		ColumnName:    "name",
		ColumnType:    "varchar(64)",
		IsNullAble:    "NO",
		CharsetName:   stringPtr("latin1"),
		CollationName: stringPtr("latin1_swedish_ci"),
	}

	destField := &FieldInfo{
		ColumnName:    "name",
		ColumnType:    "varchar(64)",
		IsNullAble:    "NO",
		CharsetName:   stringPtr("utf8mb4"),
		CollationName: stringPtr("utf8mb4_general_ci"),
	}

	// These should be considered different
	xt.False(t, sourceField.Equals(destField))
	xt.False(t, destField.Equals(sourceField))
}

func TestFieldInfo_WithTimestamps(t *testing.T) {
	// Test the exact example from the issue: t_shedlock table
	sourceFields := map[string]*FieldInfo{
		"name": {
			ColumnName:    "name",
			ColumnType:    "varchar(64)",
			IsNullAble:    "NO",
			CharsetName:   nil,
			CollationName: nil,
		},
		"lock_until": {
			ColumnName:    "lock_until",
			ColumnType:    "timestamp(3)",
			IsNullAble:    "NO",
			ColumnDefault: stringPtr("CURRENT_TIMESTAMP(3)"),
			Extra:         "DEFAULT_GENERATED on update CURRENT_TIMESTAMP(3)",
			CharsetName:   nil,
			CollationName: nil,
		},
		"locked_at": {
			ColumnName:    "locked_at",
			ColumnType:    "timestamp(3)",
			IsNullAble:    "NO",
			ColumnDefault: stringPtr("CURRENT_TIMESTAMP(3)"),
			Extra:         "DEFAULT_GENERATED",
			CharsetName:   nil,
			CollationName: nil,
		},
		"locked_by": {
			ColumnName:    "locked_by",
			ColumnType:    "varchar(255)",
			IsNullAble:    "NO",
			CharsetName:   nil,
			CollationName: nil,
		},
	}

	destFields := map[string]*FieldInfo{
		"name": {
			ColumnName:    "name",
			ColumnType:    "varchar(64)",
			IsNullAble:    "NO",
			CharsetName:   stringPtr("utf8mb4"),
			CollationName: stringPtr("utf8mb4_general_ci"),
		},
		"lock_until": {
			ColumnName:    "lock_until",
			ColumnType:    "timestamp(3)",
			IsNullAble:    "NO",
			ColumnDefault: stringPtr("CURRENT_TIMESTAMP(3)"),
			Extra:         "DEFAULT_GENERATED on update CURRENT_TIMESTAMP(3)",
			CharsetName:   nil,
			CollationName: nil,
		},
		"locked_at": {
			ColumnName:    "locked_at",
			ColumnType:    "timestamp(3)",
			IsNullAble:    "NO",
			ColumnDefault: stringPtr("CURRENT_TIMESTAMP(3)"),
			Extra:         "DEFAULT_GENERATED",
			CharsetName:   nil,
			CollationName: nil,
		},
		"locked_by": {
			ColumnName:    "locked_by",
			ColumnType:    "varchar(255)",
			IsNullAble:    "NO",
			CharsetName:   stringPtr("utf8mb4"),
			CollationName: stringPtr("utf8mb4_general_ci"),
		},
	}

	// All fields should be considered equal
	for fieldName, sourceField := range sourceFields {
		t.Run(fmt.Sprintf("field_%s", fieldName), func(t *testing.T) {
			destField := destFields[fieldName]
			xt.True(t, sourceField.Equals(destField))
			xt.True(t, destField.Equals(sourceField))
		})
	}
}

func TestFieldInfo_DefaultCharsets(t *testing.T) {
	// Test that default charsets are handled correctly
	testCases := []struct {
		name          string
		charsetName   *string
		collationName *string
		shouldEqual   bool
	}{
		{
			name:          "both nil",
			charsetName:   nil,
			collationName: nil,
			shouldEqual:   true,
		},
		{
			name:          "charset nil, collation nil vs utf8mb4",
			charsetName:   nil,
			collationName: nil,
			shouldEqual:   true,
		},
		{
			name:          "charset nil vs utf8mb4, collation nil",
			charsetName:   stringPtr("utf8mb4"),
			collationName: nil,
			shouldEqual:   true,
		},
		{
			name:          "charset nil, collation nil vs general_ci",
			charsetName:   nil,
			collationName: stringPtr("utf8mb4_general_ci"),
			shouldEqual:   true,
		},
		{
			name:          "both utf8mb4 general_ci",
			charsetName:   stringPtr("utf8mb4"),
			collationName: stringPtr("utf8mb4_general_ci"),
			shouldEqual:   true,
		},
		{
			name:          "both utf8 general_ci",
			charsetName:   stringPtr("utf8"),
			collationName: stringPtr("utf8_general_ci"),
			shouldEqual:   true,
		},
		{
			name:          "both latin1 swedish_ci",
			charsetName:   stringPtr("latin1"),
			collationName: stringPtr("latin1_swedish_ci"),
			shouldEqual:   true,
		},
		{
			name:          "different charset: ascii vs utf8mb4",
			charsetName:   stringPtr("ascii"),
			collationName: stringPtr("ascii_general_ci"),
			shouldEqual:   false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			field1 := &FieldInfo{
				ColumnName:    "test_field",
				ColumnType:    "varchar(100)",
				IsNullAble:    "NO",
				CharsetName:   nil,
				CollationName: nil,
			}

			field2 := &FieldInfo{
				ColumnName:    "test_field",
				ColumnType:    "varchar(100)",
				IsNullAble:    "NO",
				CharsetName:   tc.charsetName,
				CollationName: tc.collationName,
			}

			if tc.shouldEqual {
				xt.True(t, field1.Equals(field2))
				xt.True(t, field2.Equals(field1))
			} else {
				xt.False(t, field1.Equals(field2))
				xt.False(t, field2.Equals(field1))
			}
		})
	}
}
