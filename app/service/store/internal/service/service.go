package service

import (
	"github.com/micro/go-micro/v2/logger"
	"go-common/app/service/store/internal/dao"
)

// Service service.
type Service struct {
	dao dao.Dao
}

// New new a service and return.
func New(d dao.Dao) (s *Service, cf func(), err error) {
	s = &Service{
		dao: d,
	}
	cf = s.Close
	return
}

// Close close the resource.
func (s *Service) Close() {
	logger.Info("close service")
}




