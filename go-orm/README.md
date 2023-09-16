# ORM

## sqlite 的 使用
### Install sqlite
```shell
yum install sqlite-devel
```

### Sqlite Feature

1. SQlite 使用 c编写，轻量级、快速。
2. SQLite 可以直接嵌入到代码中，不需要像 MySQL、PostgreSQL 需要启动独立的服务才能使用。
3. SQLite 将数据存储在单一的磁盘文件中，使用起来非常方便。



### Sqlite Command

**创建数据库**
```shell
sqlite3 gee.db
```

**查看命令**

```shell
.help
```

**打开显示列名的开关**
```shell
.head on
```

**查看当前数据库中所有的表**
```shell
.table
```

**查看建表的 SQL 语句**
```shell
.schema table
```

**查询SQLite常用命令**、

https://geektutu.com/post/cheat-sheet-sqlite.html

## 实现一个简单的log库
> 为什么不用原生的log库。log标准库没有日志分级，不打印文件和行号，这就意味着我们很难找到哪个地方发生了错误

### 简易log库的特性
+ 支持日志分级 info Error Disabled
+ 不同层次日志显示使用不用的颜色区分
+ 显示打印日志代码对应的文件名和行号