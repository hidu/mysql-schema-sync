package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"reflect"
	"regexp"
	"strings"
)

// Version version
const Version = "0.3"

// AppURL site
const AppURL = "https://github.com/hidu/mysql-schema-sync/"

const timeFormatStd string = "2006-01-02 15:04:05"

// loadJsonFile load json
func loadJSONFile(jsonPath string, val interface{}) error {
	bs, err := ioutil.ReadFile(jsonPath)
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
	pattern := "^" + strings.Replace(patternStr, "*", `.*`, -1) + "$"
	match, err := regexp.MatchString(pattern, str)
	if err != nil {
		log.Println("simple_match:error", msg, "patternStr:", patternStr, "pattern:", pattern, "str:", str, "err:", err)
	}
	if match {
		//log.Println("simple_match:suc", msg, "patternStr:", patternStr, "pattern:", pattern, "str:", str)
	}
	return match
}

func htmlPre(str string) string {
	return "<pre>" + html.EscapeString(str) + "</pre>"
}

func splitDsn(dsn string) (prefix, dbName string) {
	i := strings.Split(dsn, ")/")
	if len(i) < 1 {
		panic(fmt.Sprintf("split dsn %s failed.", dsn))
	}
	return i[0], i[len(i)-1]
}

func newDsnDbName(dsn string) (newDsn, dbName string) {
	prefix, dbName := splitDsn(dsn)
	return prefix, dbName
}

func dsnSort(dsn string) string {
	i := strings.Index(dsn, "@")
	if i < 1 {
		return dsn
	}
	return dsn[i+1:]
}

func maxMapKeyLen(data interface{}, ext int) string {
	l := 0

	for _, k := range reflect.ValueOf(data).MapKeys() {
		if k.Len() > l {
			l = k.Len()
		}
	}
	return fmt.Sprintf("%d", l+ext)
}
