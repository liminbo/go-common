package http

import (
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/web"
	"go-common/app/service/store/internal/service"
	"net/http/pprof"
	"sync"
)

var svr *service.Service

// New new a bm server.
func New(s *service.Service, waiter *sync.WaitGroup) (err error) {
	service := web.NewService(
		web.Name("go.micro.web.greeter"),
	)

	service.HandleFunc("/", pprof.Index)

	if err := service.Init(); err != nil {
		logger.Errorf("service init err:%v", err)
	}

	waiter.Add(1)
	go func() {
		defer waiter.Done()
		if err := service.Run(); err != nil {
			logger.Errorf("service run err:%v", err)
		}
	}()

	svr = s
	return
}

