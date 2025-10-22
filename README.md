# mysql-schema-sync

MySQL Schema 自动同步工具  

用于将 `线上` 数据库 Schema **变化**同步到 `本地测试环境`!
只同步 Schema、不同步数据。

支持功能：

1. 同步**新表**
2. 同步**字段** 变动：新增、修改
3. 同步**索引** 变动：新增、修改
4. 同步**字段顺序**：支持调整字段在表中的顺序
5. 支持**预览**（只对比不同步变动）
6. **邮件**通知变动结果
7. 支持屏蔽更新**表、字段、索引、外键**
8. 支持本地比线上额外多一些表、字段、索引、外键
9. 在该项目的基础上修复了比对过程中遇到分区表会终止后续操作的问题，支持分区表，对于分区表，会同步除了分区以外的变更。
10. 支持每条 ddl 只会执行单个的修改，目的兼容tidb ddl问题 Unsupported multi schema change，通过single_schema_change字段控制，默认关闭。

## 安装

```bash
go install github.com/hidu/mysql-schema-sync@master
```

## 配置

参考 默认配置文件  config.json 配置同步源、目的地址。  
修改邮件接收人  当运行失败或者有表结构变化的时候你可以收到邮件通知。  

默认情况不会对多出的**表、字段、索引、外键**删除。若需要删除**字段、索引、外键** 可以使用 `-drop` 参数。

默认情况不会同步字段顺序差异。若需要同步字段顺序，可以使用 `-field-order` 参数（注意：此操作可能需要重建表，影响性能）。

配置示例(config.json):  

```
cp config.json mydb_conf.json
```

```
{
      //source：同步源
      "source":"test:test@(127.0.0.1:3306)/test_0",
      //dest：待同步的数据库
      "dest":"test:test@(127.0.0.1:3306)/test_1",
      //alter_ignore： 同步时忽略的字段和索引
      "alter_ignore":{
        "tb1*":{
            "column":["aaa","a*"],
            "index":["aa"],
            "foreign":[]
        }
      },
      //  tables: table to check schema,default is all.eg :["order_*","goods"]
      "tables":[],
      //  tables_ignore: table to ignore check schema,default is Null :["order_*","goods"]
      "tables_ignore": [],
      //有变动或者失败时，邮件接收人
      "email":{
          "send_mail":false,
         "smtp_host":"smtp.163.com:25",
         "from":"xxx@163.com",
         "password":"xxx",
         "to":"xxx@163.com"
      }
}
```

### JSON 配置项说明

source: 数据库同步源  
dest:   待同步的数据库  
tables： 数组，配置需要同步的表，为空则是不限制，eg: ["goods","order_*"]  
alter_ignore： 忽略修改的配置，表名为tableName，可以配置 column 和 index，支持通配符 *  
email ： 同步完成后发送邮件通知信息  
single_schema_change：是否每个ddl只执行单个修改

### 运行

### 直接运行

```shell
./mysql-schema-sync -conf mydb_conf.json -sync
```

### 预览并生成变更sql

```shell
./mysql-schema-sync -drop -conf mydb_conf.json 2>/dev/null >db_alter.sql

```

### 使用shell调度

```shell
bash check.sh
```

每个json文件配置一个目的数据库，check.sh脚本会依次运行每份配置。
log存储在当前的log目录中。

### 自动定时运行

添加crontab 任务

```shell
30 * * * *  cd /your/path/xxx/ && bash check.sh >/dev/null 2>&1
```

### 参数说明

```shell
mysql-schema-sync [-conf] [-dest] [-source] [-sync] [-drop] [-field-order]
```

说明：

```shell
mysql-schema-sync -help
```

```text
  -conf string
        配置文件名称
  -dest string
        待同步的数据库 eg: test@(10.10.0.1:3306)/test_1
        该项不为空时，忽略读入 -conf参数项
  -drop
        是否对本地多出的字段和索引进行删除 默认否
  -field-order
        是否同步字段顺序（可能需要重建表，影响性能）默认否
  -http
        启用web站点显示运行结果报告的地址，如 :8080,默认否
  -source string
        mysql 同步源,eg test@(127.0.0.1:3306)/test_0
  -sync
        是否将修改同步到数据库中去，默认否
  -tables string
        待检查同步的数据库表，为空则是全部
        eg : product_base,order_*
  -single_schema_change
        生成 SQL DDL 语言每条命令是否只会进行单个修改操作，默认否
```
