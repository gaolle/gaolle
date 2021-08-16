### protobuf 安装

* https://github.com/protocolbuffers/protobuf/releases

  下载protoc

  window版本解压之后将bin目录下的protoc.exe放在$GOPATH/bin

  linux 需要编译

  

  安装go语言protobuf插件

* ```shell
  go get -u github.com/golang/protobuf/protoc-gen-go
  ```

安装完之后确保$GOPATH/bin下存在protoc.exe和protoc-gen-go.exe

将$GOPATH/bin加入环境变量

https://geektutu.com/post/quick-go-protobuf.html

处理结构体的驼峰tag

go get github.com/favadi/protoc-go-inject-tag

protoc-go-inject-tag.exe位于$GOPATH/bin









### grpc-go

```shell
go get -u google.golang.org/grpc
```



#### Protobuf语法

message：定义一个请求或相应的消息格式，可以包含多种类型字段

```protobuf
// proto编译之后，Go文件中，每个消息类型对应一个结构体
message StudentRequest {
    // 每个字段的数值型标识符唯一
    string name = 1;
    int32 age = 2;
    int32 height = 3;
    // message可嵌套
    // HelloReply hello = 4;
}
```

repeated 代表可重复，相当于数组

```protobuf
repeated name = 1;
```

服务定义格式

```protobuf
// service serviceName {
//	rpc method (requestMessage) returns (responseMessage) {}
// }
service StudentService {
    // rpc 远程方法调用
    rpc Student (StudentRequest) returns (StudentResponse) {}
}
```



##### unparsable Go source: 20:4: illegal UTF-8 encoding

https://www.jianshu.com/p/62939f187785

