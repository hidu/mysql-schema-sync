// Copyright(C) 2022 github.com/hidu  All Rights Reserved.
// Author: hidu <duv123@gmail.com>
// Date: 2022/3/11

package internal

import (
	"testing"
)

func Test_fmtTableCreateSQL(t *testing.T) {
	type args struct {
		sql string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "del auto_incr",
			args: args{
				sql: `CREATE TABLE user (
				id bigint unsigned NOT NULL AUTO_INCREMENT,
				email varchar(1000) NOT NULL DEFAULT '',
				PRIMARY KEY (id)
			) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb3`,
			},
			want: `CREATE TABLE user (
				id bigint unsigned NOT NULL AUTO_INCREMENT,
				email varchar(1000) NOT NULL DEFAULT '',
				PRIMARY KEY (id)
			) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3`,
		},
		{
			name: "del auto_incr 2",
			args: args{
				sql: `CREATE TABLE user (
				id bigint unsigned NOT NULL AUTO_INCREMENT,
				email varchar(1000) NOT NULL DEFAULT '',
				PRIMARY KEY (id)
			) ENGINE=InnoDB AUTO_INCREMENT=4049116 DEFAULT CHARSET=utf8mb4`,
			},
			want: `CREATE TABLE user (
				id bigint unsigned NOT NULL AUTO_INCREMENT,
				email varchar(1000) NOT NULL DEFAULT '',
				PRIMARY KEY (id)
			) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fmtTableCreateSQL(tt.args.sql); got != tt.want {
				t.Errorf("fmtTableCreateSQL() = %v, want %v", got, tt.want)
			}
		})
	}
}
