# Demo

## 项目简介

###安装etcdv3服务
```
1、下载最新版本：https://github.com/etcd-io/etcd/releases/tag/v3.4.9
2、启动etcd ./etcd --listen-client-urls http://0.0.0.0:2379 --advertise-client-urls http://0.0.0.0:2379 --listen-peer-urls http://0.0.0.0:2380
```

###安装jaeger
```
1、docker pull jaegertracing/all-in-one
2、docker run -d --name jaeger -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 -p 5775:5775/udp -p 6831:6831/udp -p 6832:6832/udp -p 5778:5778 -p 16686:16686 -p 14268:14268 -p 9411:9411 jaegertracing/all-in-one:latest
```

###配置环境变量
```
go micro默认是以下地址，如果不是这样的可以设置环境变量
1、etcd=127.0.0.1:6379
2、jaeger=127.0.0.1:6831
```

###修改配置文件里数据库配置
```
配置文件目录：/cmd/config.toml
```

###编译
```
cd /cmd && go build main.go
```

###启动
```
./main
```

###测试
```
cd /store/test && go build test.go
启动 ./test
```
