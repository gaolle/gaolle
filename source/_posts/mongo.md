---
title: mongo
date: 2020-10-27 01:15:50
tags: mongo
categories: 数据库
---

# MongoDB

https://docs.mongodb.com/manual/tutorial/insert-documents/

mongo 本质是按照key-value map 存储的 json

```go
"go.mongodb.org/mongo-driver/mongo" // mongo驱动
```

##### 存储ma np

使用inc拼接（需要定义protobuf存储格式）

```json
{
    "school": {
        "class": {
            "name": "name",
            "total": 10,
            "可变值": 1
        }
    }
}
```

根据map更新内容


```go
update.Inc("school.class.name, "name")
update.Inc("school.class."+strconv.Itoa(int(可变值)), 1)
```

##### 读取map

```go
valMap := map[string]interface{}{}
if err := db.FindOne(ctx, filter).Decode(valMap); err != nil {
  return err
}
```

**存在时更新，不存在时创建, 主键("_id")不能更新**

set : 更新值

push：向数组压入值（可重复）

addToSet：向数组压入值（不可重复）

upsert：不存在时是否更新

```shell
db.test.update({"_id": 2}, {$set: {item: "a"}, $push: {tags: "a"}}, {upsert:true})
```
