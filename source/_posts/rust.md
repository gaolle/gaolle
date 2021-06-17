---
title: rust
date: 2021-06-27 12:10:00
tags: rust
categories: 语言
---



```rust
    let strSlice: &str = "str"; // 字符串切片 //  &str是一个不可变的引用  
    // 通过存储对第一个元素的引用和长度
    println!("{}", strSlice);
    let mut str = String::from("str"); // 字符串  // 可变
```

```rust
    if let Some(value) = temp {
        println!("{}", value);
    }
```

```
mod front_of_house;使用分号而不是使用块{}告诉 Rust 从另一个与模块同名的文件中加载模块的内容
```

```
同一作用域中不能有可变引用和不可变引用
```

```
&不可变引用
& mut 可变引用
```

```
unwrap: 如果Result值是Ok变体，unwrap将返回Ok. 如果Result是Err变体，unwrap将为panic!我们调用宏
```

```
?：只允许使用返回Result或Option或其他类型的实现 std::ops::Try
```

```
生命周期不是确保类型具有我们想要的行为，而是确保引用在我们需要时有效
生命周期的主要目的是防止悬空引用
```

```
let str = String::from("string"); // 字符串 可变 堆
let str = "string"; // 字符串切片 堆栈 不可变
引用堆栈不可变的变量不受生命周期的限制，生命周期的限制作用于引用分配在堆上的变量
```

```
函数返回的引用的生命周期与传入引用的生命周期中较小的生命周期相同
从函数返回引用时，返回类型的生命周期参数需要与参数之一的生命周期参数匹配
每个引用都有一个生命周期
生命周期第三条规则：函数或方法参数的生命周期称为输入生命周期，返回值的生命周期称为输出生命周期。
```

```
如果有多个输入生命周期参数，但其中一个是&self或&mut self因为这是一种方法，则生命周期self 被分配给所有输出生命周期参数
```

