// Copyright(C) 2022 github.com/fsgo  All Rights Reserved.
// Author: hidu <duv123@gmail.com>
// Date: 2022/9/25

package internal

import (
	"testing"

	"github.com/xanygo/anygo/xt"
)

func TestSchemaSync_getAlterDataBySchema(t *testing.T) {
	type args struct {
		table   string
		sSchema string
		dSchema string
		cfg     *Config
	}
	tests := []struct {
		name string
		sc   *SchemaSync
		args args
		want string
	}{
		{
			name: "user 0-1",
			args: args{
				table:   "user",
				sSchema: testLoadFile("testdata/user/user_0.sql"),
				dSchema: testLoadFile("testdata/user/user_1.sql"),
				cfg:     &Config{},
			},
			sc: &SchemaSync{
				Config: &Config{},
			},
			want: testLoadFile("testdata/user/result_1.sql"),
		},
		{
			name: "user 0-1 ssc",
			args: args{
				table:   "user",
				sSchema: testLoadFile("testdata/user/user_0.sql"),
				dSchema: testLoadFile("testdata/user/user_1.sql"),
				cfg: &Config{
					SingleSchemaChange: true,
				},
			},
			sc: &SchemaSync{
				Config: &Config{},
			},
			want: testLoadFile("testdata/user/result_2.sql"),
		},
		{
			name: "user 0-1 ssc",
			args: args{
				table:   "user",
				sSchema: testLoadFile("testdata/user/user_0.sql"),
				dSchema: testLoadFile("testdata/user/user_1.sql"),
				cfg: &Config{
					SingleSchemaChange: true,
				},
			},
			sc: &SchemaSync{
				Config: &Config{},
			},
			want: testLoadFile("testdata/user/result_2.sql"),
		},
		{
			name: "user 1-0 ssc",
			args: args{
				table:   "user",
				sSchema: testLoadFile("testdata/user/user_1.sql"),
				dSchema: testLoadFile("testdata/user/user_0.sql"),
				cfg: &Config{
					SingleSchemaChange: true,
				},
			},
			sc: &SchemaSync{
				Config: &Config{},
			},
			want: testLoadFile("testdata/user/result_3.sql"),
		},
		{
			name: "user 2-0 ssc",
			args: args{
				table:   "user",
				sSchema: testLoadFile("testdata/user/user_2.sql"),
				dSchema: testLoadFile("testdata/user/user_0.sql"),
				cfg:     &Config{},
			},
			sc: &SchemaSync{
				Config: &Config{},
			},
			want: testLoadFile("testdata/user/result_4.sql"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.sc.getAlterDataBySchema(tt.args.table, tt.args.sSchema, tt.args.dSchema, tt.args.cfg)
			t.Log("got alter:\n", got.String())
			xt.Equal(t, tt.want, got.String())
		})
	}
}
