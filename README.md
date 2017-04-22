# mysql-schema-sync
mysql表结构自动同步工具  

用于将 `线上` 数据库结构<b>变化</b>同步到 `本地环境`!  
支持功能：  
1.  同步**新表**  
2.  同步**字段** 变动：新增、修改  
3.  同步**索引** 变动：新增、修改   
4.  支持**预览**（只对比不同步变动）  
5.  **邮件**通知变动结果    
6.  支持屏蔽更新**表、字段、索引、外键**  
7.  支持本地比线上额外多一些表、字段、索引、外键



### 安装
>go get -u github.com/hidu/mysql-schema-sync


### 配置
参考 默认配置文件  config.json 配置同步源、目的地址。  
修改邮件接收人  当运行失败或者有表结构变化的时候你可以收到邮件通知。  

默认情况不会对多出的**表、字段、索引、外键**删除。若需要删除**字段、索引、外键** 可以使用 <code>-drop</code> 参数。

配置示例(config.json):  
```javascript
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

#### json配置项说明
source: 数据库同步源  
dest:   待同步的数据库  
tables： 数组，配置需要同步的表，为空则是不限制，eg: ["goods","order_*"]  
alter_ignore： 忽略修改的配置，表名为tableName，可以配置 column 和 index，支持通配符 *  
email ： 同步完成后发送邮件通知信息  

### 运行
### 直接运行
```
mysql-schema-sync -conf mydb_conf.json -sync
```
 
### 预览并生成变更sql
```
mysql-schema-sync -conf mydb_conf.json 2>/dev/null >db_alter.sql
```
### 使用shell调度
```
bash check.sh
```

每个json文件配置一个目的数据库，check.sh脚本会依次运行每份配置。
log存储在当前的log目录中。

### 自动定时运行
添加crontab 任务

<code>
30 * * * *  cd /your/path/xxx/ && bash check.sh >/dev/null 2>&1 
</code>

### 参数说明
<code>
mysql-schema-sync [-conf] [-dest] [-source] [-sync] [-drop]
</code>

说明：
<pre><code>
# mysql-schema-sync -help  
  -conf string
        配置文件名称
  -dest string
        mysql 同步源,eg test@(127.0.0.1:3306)/test_0
  -drop
        是否对本地多出的字段和索引进行删除 默认否
  -source string
        待同步的数据库 eg: test@(10.10.0.1:3306)/test_1
        该项不为空时，忽略读入 -conf参数项
  -sync
        是否将修改同步到数据库中去，默认否
  -tables string
        待检查同步的数据库表，为空则是全部
        eg : product_base,order_*

</code>
</pre>


