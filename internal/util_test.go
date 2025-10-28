package internal

import (
	"testing"
)

func TestNormalizeIntegerType(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		// Basic integer types with display width
		{
			name:     "int with display width",
			input:    "int(11)",
			expected: "int",
		},
		{
			name:     "bigint with display width",
			input:    "bigint(20)",
			expected: "bigint",
		},
		{
			name:     "tinyint with display width",
			input:    "tinyint(1)",
			expected: "tinyint",
		},
		{
			name:     "tinyint(4) with display width",
			input:    "tinyint(4)",
			expected: "tinyint",
		},
		{
			name:     "smallint with display width",
			input:    "smallint(5)",
			expected: "smallint",
		},
		{
			name:     "mediumint with display width",
			input:    "mediumint(8)",
			expected: "mediumint",
		},

		// Integer types with unsigned modifier
		{
			name:     "int(11) unsigned",
			input:    "int(11) unsigned",
			expected: "int unsigned",
		},
		{
			name:     "bigint(20) unsigned",
			input:    "bigint(20) unsigned",
			expected: "bigint unsigned",
		},
		{
			name:     "tinyint(1) unsigned",
			input:    "tinyint(1) unsigned",
			expected: "tinyint unsigned",
		},

		// Integer types with zerofill modifier
		{
			name:     "int(11) zerofill",
			input:    "int(11) zerofill",
			expected: "int zerofill",
		},
		{
			name:     "int(10) unsigned zerofill",
			input:    "int(10) unsigned zerofill",
			expected: "int unsigned zerofill",
		},

		// Integer types without display width (already normalized)
		{
			name:     "int without display width",
			input:    "int",
			expected: "int",
		},
		{
			name:     "bigint without display width",
			input:    "bigint",
			expected: "bigint",
		},
		{
			name:     "int unsigned without display width",
			input:    "int unsigned",
			expected: "int unsigned",
		},

		// Non-integer types (should not be affected)
		{
			name:     "varchar with length",
			input:    "varchar(255)",
			expected: "varchar(255)",
		},
		{
			name:     "char with length",
			input:    "char(10)",
			expected: "char(10)",
		},
		{
			name:     "decimal with precision",
			input:    "decimal(10,2)",
			expected: "decimal(10,2)",
		},
		{
			name:     "text type",
			input:    "text",
			expected: "text",
		},
		{
			name:     "timestamp",
			input:    "timestamp",
			expected: "timestamp",
		},

		// Case insensitive matching
		{
			name:     "INT(11) uppercase",
			input:    "INT(11)",
			expected: "INT",
		},
		{
			name:     "BIGINT(20) UNSIGNED uppercase",
			input:    "BIGINT(20) UNSIGNED",
			expected: "BIGINT UNSIGNED",
		},
		{
			name:     "TinyInt(1) mixed case",
			input:    "TinyInt(1)",
			expected: "TinyInt",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := normalizeIntegerType(tt.input)
			if result != tt.expected {
				t.Errorf("normalizeIntegerType(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}
