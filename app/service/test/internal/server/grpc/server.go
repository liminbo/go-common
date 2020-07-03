package grpc

import (
	"context"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"go-common/app/service/test/api"
	"go-common/app/service/test/internal/service"
	"sync"
)

type grpcServer struct {
	svr *service.Service
}

func (g *grpcServer) TestHello(ctx context.Context, req *api.TestRequest, rsp *api.TestResponse) (err error) {
	if rsp.Greeting,err = g.svr.CallClient(); err != nil{
		return
	}
	rsp.Greeting = rsp.Greeting + "test--demo"
	return nil
}

// New new a grpc server.
func New(microService micro.Service ,svc *service.Service, waiter *sync.WaitGroup) (err error) {
	err = api.RegisterTestDemoHandler(microService.Server(), &grpcServer{svr:svc})

	waiter.Add(1)
	go func() {
		defer waiter.Done()
		if err := microService.Run(); err != nil {
			logger.Errorf("service run err:%v", err)
		}
	}()
	return
}


