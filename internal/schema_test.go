// Copyright(C) 2022 github.com/fsgo  All Rights Reserved.
// Author: hidu <duv123@gmail.com>
// Date: 2022/9/25

package internal

import (
	"os"
	"testing"

	"github.com/xanygo/anygo/ds/xmap"
	"github.com/xanygo/anygo/xt"
)

func testLoadFile(name string) string {
	bf, err := os.ReadFile(name)
	if err != nil {
		panic("read " + name + " failed:" + err.Error())
	}
	return string(bf)
}

func TestParseSchema(t *testing.T) {
	type args struct {
		schema string
	}
	tests := []struct {
		name string
		args args
		want *MySchema
	}{
		{
			name: "case 1",
			args: args{
				schema: testLoadFile("testdata/user/user_0.sql"),
			},
			want: &MySchema{
				Fields: (func() xmap.Ordered[string, string] {
					m := xmap.Ordered[string, string]{}
					// 不会检查 value
					m.Set("id", "`id` bigint unsigned NOT NULL AUTO_INCREMENT,")
					m.Set("email", "`email` varchar(1000) NOT NULL DEFAULT '',")
					m.Set("register_time", "`register_time` timestamp NOT NULL,")
					m.Set("password", "`password` varchar(1000) NOT NULL DEFAULT '',")
					m.Set("status", "`status` tinyint unsigned NOT NULL DEFAULT '0',")
					return m
				})(),
				IndexAll: map[string]*DbIndex{
					"PRIMARY KEY": {
						Name:      "PRIMARY KEY",
						SQL:       "PRIMARY KEY (`id`)",
						IndexType: indexTypePrimary,
					},
				},
			},
		},
		{
			name: "case 2",
			args: args{
				schema: testLoadFile("testdata/user/user_4.sql"),
			},
			want: &MySchema{
				Fields: (func() xmap.Ordered[string, string] {
					m := xmap.Ordered[string, string]{}
					// 不会检查 value
					m.Set("id", "\"id\" bigint unsigned NOT NULL AUTO_INCREMENT,")
					m.Set("email", "\"email\" varchar(1000) NOT NULL DEFAULT \"\",")
					m.Set("register_time", "\"register_time\" timestamp NOT NULL,")
					m.Set("password", "\"password\" varchar(1000) NOT NULL DEFAULT \"\",")
					m.Set("status", "\"status\" tinyint unsigned NOT NULL DEFAULT \"0\",")
					return m
				})(),
				IndexAll: map[string]*DbIndex{
					"PRIMARY KEY": {
						Name:      "PRIMARY KEY",
						SQL:       "PRIMARY KEY (\"id\")",
						IndexType: indexTypePrimary,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParseSchema(tt.args.schema)
			gs := got.String()
			ws := tt.want.String()
			xt.Equal(t, ws, gs)
		})
	}
}
