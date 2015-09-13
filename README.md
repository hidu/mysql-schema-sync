# mysql-schema-sync

###安装
>go get -u github.com/hidu/mysql-schema-sync

### 配置
参考 默认配置文件  config.json 配置同步源、目的地址。  
修改邮件接收人  当运行失败或者有表结构变化的时候你可以收到邮件通知。  

默认情况不会对多出的*表、字段、索引*删除。若需要删除*字段、索引* 可以使用 <code>-drop</code> 参数。

#### json配置项说明
source: 数据库同步源  
dest:   待同步的数据库  
tables： 数组，配置需要同步的表，为空则是不限制，eg: ["goods","order_*"]  
alter_ignore： 忽略修改的配置，表名为tableName，可以配置 column 和 index，支持通配符 *  
email ： 同步完成后发送邮件通知信息  

### 运行
<code>
bash check.sh
</code>

每个json文件配置一个目的数据库，check.sh脚本会依次运行每份配置。
log存储在当前的log目录中。

###自动定时运行
添加crontab 任务

<code>
30 * * * *  cd /your/path/xxx/ && bash check.sh >/dev/null 2>&1 
</code>

###不使用配置文件
<code>
mysql-schema-sync [-conf] [-dest] [-source] [-sync] [-drop]
</code>

说明：
<pre>
<code>
#mysql-schema-sync -help  
  -conf string
        json config file path (default "./config.json")
  -dest string
        mysql dsn dest,eg test@(127.0.0.1:3306)/test_0
  -drop
        drop fields and index (default true)
  -source string
        mysql dsn source,eg: test@(10.10.0.1:3306)/test_1
        when it is not empty ignore [-conf] param
  -sync
        sync shcema change to dest db
  -tables string
        table names to check
        eg : product_base,order_*

</code>
</pre>

