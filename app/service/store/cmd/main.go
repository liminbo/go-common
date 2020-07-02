package main

import (
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source/file"
	"github.com/micro/go-micro/v2/logger"
	"go-common/app/service/store/internal/di"
	"go-common/library/net/trace"
	"github.com/opentracing/opentracing-go"
	conf "go-common/app/service/store/config"
)

func main() {
	// 加载配置文件
	if err := config.Load(file.NewSource(file.WithPath("./config.toml"))); err != nil {
		panic(err)
	}

	// 链路追踪 start
	t, closer, err := trace.NewJaegerTracer("go.micro.srv.greeterbo", conf.GetJaeger())
	if err != nil {
		logger.Fatalf("opentracing tracer create error:%v", err)
	}
	defer closer.Close()
	opentracing.SetGlobalTracer(t)
	// 链路追踪 end
	
	logger.Info("store start")
	_, closeFunc, err := di.InitApp()
	if err != nil {
		panic(err)
	}
	di.Waiter.Wait()
	closeFunc()
}