package di

import (
	"go-common/app/service/store/internal/dao"
	"go-common/app/service/store/internal/microservice"
	"go-common/app/service/store/internal/server/grpc"
	"go-common/app/service/store/internal/server/http"
	"go-common/app/service/store/internal/service"
	"sync"
)

var Waiter sync.WaitGroup

func InitApp() (*App, func(), error) {

	db, closeDB, err := dao.NewDB()
	if err != nil {
		return nil, nil, err
	}

	daoDao, closeDao, err := dao.New(db)
	if err != nil {
		closeDB()
		return nil, nil, err
	}
	serviceService, closeService, err := service.New(daoDao)
	if err != nil {
		closeDao()
		closeDB()
		return nil, nil, err
	}
	err = http.New(serviceService, &Waiter)
	if err != nil {
		closeService()
		closeDao()
		closeDB()
		return nil, nil, err
	}
	err = grpc.New(microservice.MicroService, serviceService, &Waiter)
	if err != nil {
		closeService()
		closeDao()
		closeDB()
		return nil, nil, err
	}
	app, closeApp, err := NewApp()
	if err != nil {
		closeService()
		closeDao()
		closeDB()
		return nil, nil, err
	}
	return app, func() {
		closeApp()
		closeService()
		closeDao()
		closeDB()
	}, nil
}
