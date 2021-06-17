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

```mysql
mysql -h127.0.0.1 -uroot -p123456 -P3306
```

##### 选择数据库

```mysql
USE dbname
```

##### 查看数据库列表

```mysql
SHOW DATABASES;
```

##### 查看数据内的表的列表

```mysql
SHOW TABLES;
```

##### 查看表列

```mysql
SHOW COLUMNS FROM tablename;
```

```mysql
DESC tablename;
```

##### 查看表的注释

```mysql
SHOW CREATE TABLE tablename;
```

##### 通配符

​	* 所有s

##### 只返回不同的值

```mysql
SELECT DISTINCT id from tablename;
```

DISTINCT 应用于所有的列前面，并不只是作用于特定的列之前,当不存在不同的值时，匹配所有

```mysql
SELECT DISTINCT id, name from tablename;
```

##### 限制结果

LIMIT 带一个值总是从第一行开始 带两个值可以从指定的行号的第一个位置开始

```mysql
SELECT id from tablename LIMIT 1;
SELECT id from tablename LIMIT 1, 1;
```

##### 空值

NULL 无值 与字段包含0、空字符或仅仅包含空格不同

检查具有NULL值的列

```mysql
SELECT id from tablename WHERE id IS NULL;
```

##### 排序

 ORDER BY 也可以通过非选择项排序

```mysql
SELECT id from tablename ORDER BY id, name;
```

##### 指定排序方向

DESC 只作用于其前面的列名

默认升序，DESC 降序, ASC 升序

```mysql
SELECT id from tablename ORDER BY id, name DESC;
```

id 降序，name 升序

```mysql
SELECT id from tablename ORDER BY id DESC, name LIMIT 1;
```

##### WHERE字句操作符

```mysql
<> 不等于
!= 不等于
BETWEEN a AND b
```

```mysql
SELECT id from tablename WHERE id BETWEEN 2 AND 10;
```

##### AND操作符

检索满足所有给定条件的行 

```mysql
SELECT id, name from tablename WHERE id = 2 AND name = 2;
```

##### OR操作符

检索匹配任一给定条件的行

```mysql
SELECT id, name from tablename WHERE id = 2 OR name = 3;
```

AND 和 OR 同时存在时，优先匹配AND 多操作符匹配时加（）

##### IN操作符

匹配指点关键值,功能与OR相当

```mysql
SELECT id, name from tablename WHERE id IN (2, 3);
```

##### NOT操作符

否定后跟条件的关键字 mysql中NOT可以对IN、BETWEEN、EXISTS取反

```mysql
SELECT id, name from tablename WHERE id NOT IN (2, 3);
```

##### 通配符

LIKE 谓词，利用通配符匹配而不是直接相等匹配进行比较

* % 任何字符出现任意次数,包含0个字符,不匹配NULL

  匹配以a开头

  ```mysql
  SELECT id, name from tablename WHERE id LIKE `a%`
  ```

  匹配含a

  ```mysql
  SELECT id, name from tablename WHERE id LIKE `%a%`
  ```

* _ 只匹配单个字符

  ```mysql
  SELECT id, name from tablename WHERE id LIKE `%a_`
  ```

##### 拼接字段

RTrim() 去除右侧多余的空格 LTrim() 去除左侧多余的空格

```mysql
SELECT Concat(RTrim(id), '(', name, ')') FROM tablename;
```

##### 别名

```mysql
SELECT id, name as newname from tablename WHERE id
```

##### COUNT

COUNT(*) 对表中行的数目进行计数，不管表列中包含的是空值(NULL)还是非空值

COUNT(column)对特定列中具有值的行进行计数，忽略NULL值

##### 分组

WITH ROLLUP 得到每个分组以及每个分组汇总级别（针对每个分组）的值

```mysql
SELECT id, count(name) as newname from tablename GROUP BY id WITH ROLLUP ORDER BY id
```

WHERE 过滤指定行不是组，HAVING 过滤分组

WHERE 在数据分组前进行过滤 HAVING 在数据分组后进行过滤

##### 数据处理函数

Upper 将文本转换为大写

```mysql
SELECT name, Upper(name) as newname from tablename WHERE id
```

##### 日期处理函数

Date_Format() 返回一个格式化的日期或时间串

##### SELECT 子句顺序

```mysql
SELECT
FROM
WHERE
GROUP BY
HAVING
ORDER BY
LIMIT
```

##### 子查询

处理顺序：从内向外

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

##### 联结效率比子查询高

```mysql
select id, name from a, c where a.id = c.id and c.name = "adc"
select id, name inner join a, c on a,id = c,id and c.name = "abc"
```

##### 内部联结

inner join on等值联结

##### 外部联结

left outer join on包含哪些在相关表中没有关联行的行 

left\right 代表返回其所有行的表

##### 笛卡尔积

返回查找表的相乘结果

##### 组合查询

union 联结两个或两个以上的select语句 自动去重

union all 返回所有行 不去重

不允许使用多条order by语句，只能在最后一条where语句之后使用一条order by语句

##### 插入数据

完整一行数据

```mysql
insert into tablename values();
```

降低插入语句的优先级

```mysql
insert LOW_PRIORITY into tablename values();
```

##### 插入检索出来的数据

insert select

##### 更新数据

```mysql
update tablename set xxxxx where xxx;
```

更新多个值时，当一个值报错不停止继续执行

```mysql
update ignore tablename set xxxxx where xxx;
```

为了删除某个列的值，可设置它为null

```mysql
update tablename set xxxxx=NULL where xxx;
```

##### 删除

delete 删除整行不是删除列

```mysql
delete from tablename where xxx;
```

删除表中的内容不是表

##### 更快的删除

```mysql
truncate table
```

删除原来的表并重新创建一个表，而不是逐行删除表中的数据

##### 创建表

```mysql
create table tablename (
    .
    .
    .
primary key (主键)
)
```

##### 数据库引擎

InnoDB是一个可靠的事务处理引擎，它不支持全文本搜索

MyISAM是一个性能极高的引擎，它支持全文搜索，但不支持事务处理

MEMORY在功能上等同于MyISAM,但由于数据存储在内存中，速度很快（特别适合于临时表）

##### 更新表

增加列

```mysql
alter table tablename add columnname type 
```

删除列

```mysql
alter table tablename drop column columnname
```

##### 删除表

```mysql
drop table tablename
```

##### 重命名表

```mysql
rename table tablename to tablenewname
```

##### 变量名

所有MySQL变量都必须以@开始

##### 游标

MySQL游标只能用于存储过程（和函数）

##### 触发器

MySQL响应delete、insert、update时自动执行的一条MySQL语句（或位于begin和end语句之间的一组语句）

触发器名必须在每个表中唯一，但不是在每个数据库中唯一。这表示同一个数据库中的两个表可具有相同名字的触发器

只有表才支持触发器，视图和临时表不支持

##### 事务

维护数据库的完整性，保证成批的MySQL操作要么完全执行，要么完全不执行

事务：一组SQL语句

回退：撤销指定SQL语句的过程

提交：将未存储的SQL语句结果写入数据库表

保留点：事务处理中设置的临时占位符，可以对它发布回退（与回退整个事务处理不同）



### MySQL：关系数据库管理系统（RDBMS）(将数据保存在不同的表中)

#### 特点：

* 数据以表格的形式出现
* 每行为各种记录名称
* 每列为记录名称所对应的数据域
* 许多的行和列组成一张表单
* 若干的表单组成database

#### RDBMS：

* 数据库：关联表的集合
* 数据表：数据的矩阵
* 冗余：存储两倍数据，让系统速度更快
* 主键：唯一，一个数据表只能包含一个主键
* 外键：关联两个表
* 复合键：将多个列作为一个索引键
* 索引：快速索引特定信息。对数据库表中一列或多列的值进行排序的一种结构
* 参照完整性：保证数据的一致性

* 创建数据库

  ```mysql
  mysqladmin -u root -p create users
  ```

* 删除数据库

  ```mysql
  mysqladmin -u root -p drop users
  ```

* 选择数据库

  ```mysql
  use users
  ```

* 创建数据表

  ```mysql
  CREATE TABLE users_tbl
  ```

* 删除数据表

  ```mysql
  DROP TABLE users_tbl
  ```

* 插入数据

  ```mysql
  INSERT INTO users_tbl
  (id, name)
  VALUES
  (1, "A")
  ```

* 查询数据库

  ```mysql
  SELECT * from users_tbl
  ```

### docker安装MySql

* 拉取镜像

  ```shell
  docker pull mysql:latest
  ```

* 查看本地镜像

  ```shell
  docker images
  ```

* 运行一个容器

  ```shell
  docker run  --name Master -p 3306:3306 --restart=always  -e MYSQL_ROOT_PASSWORD=123456 -d mariadb:10.1
  ```

* 进入容器

  ```shell
  docker exec -it Master bash
  ```

* 登录mysql

  ```shell
  mysql -u root -p
  ```

  
#### 主键（primarily key）

能够唯一标志表中某一行的属性，一个表只能有一个主键

#### 外键（foreign key）

建立和加强两个表之间的链接，约束两个表之间数据的一致性。表的外键为另一个表的主键

创建父表A（只有一个）

```mysql
CREATE TABLE A  
(A_ID int NOT NULL, 
A_NAME VARCHAR(100) NOT NULL,
PRIMARY KEY (A_ID)；
```

创建子表B（一个或者多个）

```mysql
CREATE TABLE B 
( B_ID int NOT NULL, 
B_NAME VARCHAR(100) NOT NULL, 
A_ID int NOT NULL, 
PRIMARY KEY (B_ID)); 
```

A_ID为主键，B_ID为外键

```mysql
ALTER TABLE B ADD FOREIGN KEY (A_ID) REFERENCES A(A_ID);
```

A添加数据

| A_ID | A_NAME |
| ---- | ------ |
| 1    | A      |
| 2    | B      |
| 3    | C      |

B添加数据

| B_ID | B_NAME | A_ID |
| ---- | ------ | ---- |
| 1    | C      | 1    |
| 2    | C      | 2    |

向B中添加数据父表中A_ID不存在的值时强制不执行

```mysql
INSERT INTO B (B_ID, B_NAME, A_ID) VALUES (3, "C", 4);
ERROR 1452 (23000): Cannot add or update a child row: a foreign key constraint fails (`users`.`B`, CONSTRAINT `B_ibfk_1` FOREIGN KEY (`A_ID`) REFERENCES `A` (`A_ID`))
```

#### 索引

可以用来快速定位具有特征值的数据

#### 唯一索引

索引列的值必须唯一，允许空值

创建唯一索引

```mysql
CREATE UNIQUE INDEX A_INDEX ON A(A_NAME)
```



