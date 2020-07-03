package grpc

import (
	"context"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"go-common/app/service/store/api"
	"go-common/app/service/store/internal/service"
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
func New(microService micro.Service ,svc *service.Service, waiter *sync.WaitGroup) (err error) {
	err = api.RegisterGreeterHandler(microService.Server(), &grpcServer{svr:svc})

	waiter.Add(1)
	go func() {
		defer waiter.Done()
		if err := microService.Run(); err != nil {
			logger.Errorf("service run err:%v", err)
		}
	}()
	return
}


