# 数据库的基础操作

## 数据库操作

- 列出已有数据库：`show databases;`
- 显示用来创建指定数据库名的语句：`SHOW CREATE DATABASE 数据库名;`
- 切换到指定数据库：`use 数据库名;`
- 新增数据库：`create database 数据库名;`
- 删除数据库：`drop database 数据库名;`

## 数据表操作

- 列出已有的数据表：`show tables;`，在执行之前应已经切换到指定的数据库(`use database_name;`)。
- 显示指定数据表的列：`SHOW COLUMNS FROM 数据表名;`
- 显示用来创建指定数据表名的语句：`SHOW CREATE TABLE 数据表名;`
- 删除数据表：`drop table 数据表名;`
- 创建数据表：

```SQL
CREATE TABLE `table_name` (
    # [列名] [类型] [约束]
    `column_0` int NOT NULL,
    PRIMARY KEY (`column_name`),
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
```  

## 其他操作

- 显示MySQL服务器状态：`SHOW STATUS;`
- 显示授予用户的安全权限：`SHOW GRANTS;`

## 其他

- SQL关键词的大小写：SQL的关键词不区分大小写，比如`select`和`Select`及`SELECT`是一样的。为了便于阅读，推荐*关键词一律用大写形式*。
