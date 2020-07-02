package grpc

import (

	"context"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	tracer "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"

	"go-common/app/service/store/api"
	"go-common/app/service/store/config"
	"go-common/app/service/store/internal/service"
	"go-common/library/net/rpc"
	"sync"
)

type grpcServer struct {
	svr *service.Service
}

func (g *grpcServer) Hello(ctx context.Context, req *api.Request, rsp *api.Response) error {
	//time.Sleep(time.Second*3)
	rsp.Greeting = "Hello " + req.Name
	return nil
}

// New new a grpc server.
func New(svc *service.Service, waiter *sync.WaitGroup) (err error) {

	t := opentracing.GlobalTracer()
	service := micro.NewService(
		micro.Name("go.micro.srv.greeterbo"),
		micro.Version("latest"),
		micro.Metadata(map[string]string{
			"type": "helloworld",
		}),
		// 服务注册
		micro.Registry(etcdv3.NewRegistry(func(options *registry.Options) {
			options.Addrs = []string{config.GetEtcd()}
		})),

		//挂载链路
		micro.WrapHandler(tracer.NewHandlerWrapper(t)),
		micro.WrapCall(tracer.NewCallWrapper(t)),

		micro.WrapHandler(rpc.LogHandler()),
	)

	err = api.RegisterGreeterHandler(service.Server(), &grpcServer{svr:svc})

	waiter.Add(1)
	go func() {
		defer waiter.Done()
		if err := service.Run(); err != nil {
			logger.Errorf("service run err:%v", err)
		}
	}()
	return
}


