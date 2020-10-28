---
title: mysql
date: 2020-10-27 01:15:50
tags: mysql
categories: 数据库
---

# MySQL

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
