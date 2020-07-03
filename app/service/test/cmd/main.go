package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/opentracing/opentracing-go"
	conf "go-common/app/service/store/config"
	"go-common/app/service/test/api"
	"go-common/app/service/test/internal/di"
	"go-common/app/service/test/internal/microservice"
	"go-common/library/net/rpc"
	"go-common/library/net/trace"

	tracer "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
)

func main() {
	// 初始化 micro service
	microService := microservice.InitMicroService()

	// 链路追踪 start
	t, closer, err := trace.NewJaegerTracer(api.AppID, conf.GetJaeger())
	if err != nil {
		logger.Fatalf("opentracing tracer create error:%v", err)
	}
	defer closer.Close()
	opentracing.SetGlobalTracer(t)
	// 链路追踪 end

	microService.Init(
		//挂载链路
		micro.WrapHandler(tracer.NewHandlerWrapper(t)),
		micro.WrapCall(tracer.NewCallWrapper(t)),
		micro.WrapHandler(rpc.LogHandler()),
	)
	
	logger.Info("store start")
	_, closeFunc, err := di.InitApp()
	if err != nil {
		panic(err)
	}
	di.Waiter.Wait()
	closeFunc()
}