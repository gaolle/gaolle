---
title: golang
date: 2020-10-27 01:15:50
tags: golang
categories: 语言
---

# Golang

##### 语法

不存在隐式类型转换，只能强制类型转换

函数或者变量：首字母大写外部可导入，小写包内使用

函数包依次导入，不一次性直接包含

##### 数组

```go
var arr = [...]int{1, 2, 3} // 可变数量
for _, val := range arr {
    // _ : 不读取， 
    // 遍历时不会越界
}
```



##### Map 

底层 哈希表 

冲突方法 链地址法（数组加链表）

实际存储：key、value分别存储

https://mp.weixin.qq.com/s/OJSxIXH87mjCkQn76eNQsQ

插入key对应的value写入对应的内存，而是返回了value应该插入的内存地址

把一个值 赋给一个地址

把一个地址 赋给一个地址

##### String

```go
// 字符串元素不可改，只读，支持切片
str := "abcd" // 定义之后不能修改，相当于常量 修改时必须赋值
str[0] = 'b' // Cannot assign to str[0]
```

##### Mod

```shell
go mod init modname
go build
```
