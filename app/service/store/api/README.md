# Demo

## protobuf使用

###安装相关组件
```
1、安装protoc https://github.com/protocolbuffers/protobuf/releases
2、go get -u github.com/golang/protobuf/proto
3、go get -u github.com/golang/protobuf/protoc-gen-go
4、go get -u github.com/micro/micro/v2/cmd/protoc-gen-micro
5、生成的bin可执行文件要配置到系统的PATH里
```

###生成proto文件
```
protoc --proto_path=. --micro_out=. --go_out=.  greeter.proto
```
