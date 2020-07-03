package service

import (
	"context"
	"github.com/micro/go-micro/v2/logger"
	storeApi "go-common/app/service/store/api"
	"go-common/app/service/test/internal/dao"
	"go-common/app/service/test/internal/microservice"
)

// Service service.
type Service struct {
	dao dao.Dao
	storeRpc storeApi.GreeterService
}

// New new a service and return.
func New(d dao.Dao) (s *Service, cf func(), err error) {
	storeRpc,err := storeApi.NewClient(microservice.MicroService)
	if err != nil{
		logger.Warnf("client error:%v", err)
	}
	s = &Service{
		dao: d,
		storeRpc:storeRpc,
	}
	cf = s.Close
	return
}

// Close close the resource.
func (s *Service) Close() {
	logger.Info("close service")
}

func (s *Service) CallClient() (str string, err error){
	clientRep,err := s.storeRpc.Hello(context.Background(), &storeApi.Request{Name: "John"})
	if err != nil{
		logger.Warnf("err:%v", err)
		return
	}
	str = clientRep.Greeting
	return
}




