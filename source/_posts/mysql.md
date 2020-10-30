---
title: mysql
date: 2020-10-27 01:15:50
tags: mysql
categories: 数据库
---

# MySQL

不区分大小写

##### 数据库

保存有组织的数据的容器（通常是一个文件或者一组文件）

##### 数据库软件

DBMS（数据库管理系统）

##### 数据库和数据库软件的关系

数据库是通过DBMS创建和操纵的容器

##### 表

某种特定类型数据的结构化清单

##### 主键

一列（或一组列），其值能够唯一区分表中的每个行

条件

* 任意两行都不具有相同的主键值
* 每个行都必须具有一个主键值（主键列不允许null值）

##### SQL

结构化查询语言，一种专门用来与数据库通信的语言

##### 连接

```shell
mysql -h127.0.0.1 -uroot -p123456 -P3306
```

##### 选择数据库

```shell
USE dbname
```

##### 查看数据库列表

```shell
SHOW DATABASES;
```

##### 查看数据内的表的列表

```shell
SHOW TABLES;
```

##### 查看表列

```shell
SHOW COLUMNS FROM tablename;
```

```shell
DESC tablename;
```

##### 查看表的注释

```shell
SHOW CREATE TABLE tablename;
```

##### 通配符

​	* 所有

##### 只返回不同的值

```shell
SELECT DISTINCT id from tablename;
```

DISTINCT 应用于所有的列前面，并不只是作用于特定的列之前,当不存在不同的值时，匹配所有

```shell
SELECT DISTINCT id, name from tablename;
```

##### 限制结果

LIMIT 带一个值总是从第一行开始 带两个值可以从指定的行号的第一个位置开始

```shell
SELECT id from tablename LIMIT 1;
SELECT id from tablename LIMIT 1, 1;
```

##### 空值

NULL 无值 与字段包含0、空字符或仅仅包含空格不同

检查具有NULL值的列

```shell
SELECT id from tablename WHERE id IS NULL;
```

##### 排序

 ORDER BY 也可以通过非选择项排序

```shell
SELECT id from tablename ORDER BY id, name;
```

##### 指定排序方向

DESC 只作用于其前面的列名

默认升序，DESC 降序, ASC 升序

```shell
SELECT id from tablename ORDER BY id, name DESC;
```

id 降序，name 升序

```shell
SELECT id from tablename ORDER BY id DESC, name LIMIT 1;
```

##### WHERE字句操作符

```shell
<> 不等于
!= 不等于
BETWEEN a AND b
```

```shell
SELECT id from tablename WHERE id BETWEEN 2 AND 10;
```

##### AND操作符

检索满足所有给定条件的行 

```shell
SELECT id, name from tablename WHERE id = 2 AND name = 2;
```

##### OR操作符

检索匹配任一给定条件的行

```shell
SELECT id, name from tablename WHERE id = 2 OR name = 3;
```

AND 和 OR 同时存在时，优先匹配AND 多操作符匹配时加（）

##### IN操作符

匹配指点关键值,功能与OR相当

```shell
SELECT id, name from tablename WHERE id IN (2, 3);
```

##### NOT操作符

否定后跟条件的关键字 mysql中NOT可以对IN、BETWEEN、EXISTS取反

```shell
SELECT id, name from tablename WHERE id NOT IN (2, 3);
```

##### 通配符

LIKE 谓词，利用通配符匹配而不是直接相等匹配进行比较

* % 任何字符出现任意次数,包含0个字符,不匹配NULL

  匹配以a开头

  ```shell
  SELECT id, name from tablename WHERE id LIKE `a%`
  ```

  匹配含a

  ```shell
  SELECT id, name from tablename WHERE id LIKE `%a%`
  ```

* _ 只匹配单个字符

  ```shell
  SELECT id, name from tablename WHERE id LIKE `%a_`
  ```

##### 拼接字段

RTrim() 去除右侧多余的空格 LTrim() 去除左侧多余的空格

```shell
SELECT Concat(RTrim(id), '(', name, ')') FROM tablename;
```

##### 别名

```shell
SELECT id, name as newname from tablename WHERE id
```

##### COUNT

COUNT(*) 对表中行的数目进行计数，不管表列中包含的是空值(NULL)还是非空值

COUNT(column)对特定列中具有值的行进行计数，忽略NULL值





```go
_ "github.com/go-sql-driver/mysql" // mysql驱动
```

##### 默认值NULL判断

```go
val := sql.NullFloat64 
// 读取值操作
if sum.Valid {
	// 使用sum.Valid 进行判断是否为空值
}
```

##### 数据处理函数

Upper 将文本转换为大写

```shell
SELECT name, Upper(name) as newname from tablename WHERE id
```

##### 日期处理函数

Date_Format() 返回一个格式化的日期或时间串



```

  1 #!/bin/bash
  2  
  3 session="backstage"
  4 if tmux has-session -t $session; then
  5     tmux kill-session -t $session
  6 fi
  7  
  8 tmux new -s $session -d
  9 tmux send -t $session "./backstage" Enter 
```

