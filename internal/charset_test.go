// Copyright(C) 2022 github.com/fsgo  All Rights Reserved.
// Author: hidu <duv123@gmail.com>
// Date: 2022/9/25

package internal

import (
	"testing"

	"github.com/stretchr/testify/require"
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
		CharsetName:   stringPtr("utf8mb4"),      // Explicit charset
		CollationName: stringPtr("utf8mb4_general_ci"), // Explicit collation
	}

	// These should be considered equal
	require.True(t, sourceField.Equals(destField), "Fields with implicit and explicit charset/collation should be equal")
	require.True(t, destField.Equals(sourceField), "Fields with explicit and implicit charset/collation should be equal")
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
	require.False(t, sourceField.Equals(destField), "Fields with different charset should be different")
	require.False(t, destField.Equals(sourceField), "Fields with different charset should be different")
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
			ColumnName:     "lock_until",
			ColumnType:     "timestamp(3)",
			IsNullAble:     "NO",
			ColumnDefault:  stringPtr("CURRENT_TIMESTAMP(3)"),
			Extra:          "DEFAULT_GENERATED on update CURRENT_TIMESTAMP(3)",
			CharsetName:    nil,
			CollationName:  nil,
		},
		"locked_at": {
			ColumnName:     "locked_at",
			ColumnType:     "timestamp(3)",
			IsNullAble:     "NO",
			ColumnDefault:  stringPtr("CURRENT_TIMESTAMP(3)"),
			Extra:          "DEFAULT_GENERATED",
			CharsetName:    nil,
			CollationName:  nil,
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
			ColumnName:     "lock_until",
			ColumnType:     "timestamp(3)",
			IsNullAble:     "NO",
			ColumnDefault:  stringPtr("CURRENT_TIMESTAMP(3)"),
			Extra:          "DEFAULT_GENERATED on update CURRENT_TIMESTAMP(3)",
			CharsetName:    nil,
			CollationName:  nil,
		},
		"locked_at": {
			ColumnName:     "locked_at",
			ColumnType:     "timestamp(3)",
			IsNullAble:     "NO",
			ColumnDefault:  stringPtr("CURRENT_TIMESTAMP(3)"),
			Extra:          "DEFAULT_GENERATED",
			CharsetName:    nil,
			CollationName:  nil,
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
		destField := destFields[fieldName]
		require.True(t, sourceField.Equals(destField),
			"Field %s should be equal between source and dest", fieldName)
		require.True(t, destField.Equals(sourceField),
			"Field %s should be equal between dest and source", fieldName)
	}
}

func TestFieldInfo_DefaultCharsets(t *testing.T) {
	// Test that default charsets are handled correctly
	testCases := []struct {
		name         string
		charsetName  *string
		collationName *string
		shouldEqual  bool
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
				require.True(t, field1.Equals(field2), "Fields should be equal: %s", tc.name)
				require.True(t, field2.Equals(field1), "Fields should be equal (reverse): %s", tc.name)
			} else {
				require.False(t, field1.Equals(field2), "Fields should be different: %s", tc.name)
				require.False(t, field2.Equals(field1), "Fields should be different (reverse): %s", tc.name)
			}
		})
	}
}