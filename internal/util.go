package internal

import (
	"bytes"
	"encoding/json"
	"html"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/xanygo/anygo/cli/xcolor"
)

// Version 版本号，格式：更新日期(8位).更新次数(累加)
const Version = "20251021.4"

// AppURL site
const AppURL = "https://github.com/hidu/mysql-schema-sync/"

const timeFormatStd string = "2006-01-02 15:04:05"

// loadJsonFile load json
func loadJSONFile(jsonPath string, val any) error {
	bs, err := os.ReadFile(jsonPath)
	if err != nil {
		return err
	}
	lines := strings.Split(string(bs), "\n")
	var bf bytes.Buffer
	for _, line := range lines {
		lineNew := strings.TrimSpace(line)
		if (len(lineNew) > 0 && lineNew[0] == '#') || (len(lineNew) > 1 && lineNew[0:2] == "//") {
			continue
		}
		bf.WriteString(lineNew)
	}
	return json.Unmarshal(bf.Bytes(), &val)
}

func inStringSlice(str string, strSli []string) bool {
	for _, v := range strSli {
		if str == v {
			return true
		}
	}
	return false
}

func simpleMatch(patternStr string, str string, msg ...string) bool {
	str = strings.TrimSpace(str)
	patternStr = strings.TrimSpace(patternStr)
	if patternStr == str {
		log.Println("simple_match:suc,equal", msg, "patternStr:", patternStr, "str:", str)
		return true
	}
	pattern := "^" + strings.ReplaceAll(patternStr, "*", `.*`) + "$"
	match, err := regexp.MatchString(pattern, str)
	if err != nil {
		log.Println("simple_match:error", msg, "patternStr:", patternStr, "pattern:", pattern, "str:", str, "err:", err)
	}
	// if match {
	// log.Println("simple_match:suc", msg, "patternStr:", patternStr, "pattern:", pattern, "str:", str)
	// }
	return match
}

func htmlPre(str string) string {
	return "<pre>" + html.EscapeString(str) + "</pre>"
}

func dsnShort(dsn string) string {
	i := strings.Index(dsn, "@")
	if i < 1 {
		return dsn
	}
	return dsn[i+1:]
}

func errString(err error) string {
	if err == nil {
		return xcolor.YellowString("<nil>")
	}
	return xcolor.RedString("%s", err.Error())
}

// normalizeIntegerType removes display width from integer types for MySQL 8.0.19+ compatibility.
// MySQL 8.0.19+ deprecated display width for integer types (TINYINT, SMALLINT, MEDIUMINT, INT, BIGINT).
// This function normalizes types like "int(11)" to "int" while preserving modifiers like "unsigned" and "zerofill".
//
// Examples:
//   - "int(11)" -> "int"
//   - "int(11) unsigned" -> "int unsigned"
//   - "bigint(20)" -> "bigint"
//   - "tinyint(1)" -> "tinyint"
//   - "varchar(255)" -> "varchar(255)" (unchanged, not an integer type)
func normalizeIntegerType(columnType string) string {
	// Pattern matches: (tinyint|smallint|mediumint|int|bigint) followed by optional (digits)
	// Captures the type name and everything after the display width
	re := regexp.MustCompile(`(?i)^(tinyint|smallint|mediumint|int|bigint)\(\d+\)(\s+.+)?$`)

	matches := re.FindStringSubmatch(columnType)
	if len(matches) > 0 {
		// matches[1] is the type name (e.g., "int")
		// matches[2] is the modifiers (e.g., " unsigned", " zerofill"), may be empty
		if len(matches) > 2 && matches[2] != "" {
			return matches[1] + matches[2] // e.g., "int unsigned"
		}
		return matches[1] // e.g., "int"
	}

	// Not an integer type with display width, return as-is
	return columnType
}
