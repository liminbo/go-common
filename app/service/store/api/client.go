package api

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	"go-common/app/service/store/config"
)


func NewService() (srv GreeterService, err error) {
	service := micro.NewService(
		micro.Name("go.micro.srv.greeterbo"),
		micro.Version("latest"),
		micro.Registry(etcdv3.NewRegistry(func(options *registry.Options) {
			options.Addrs = []string{config.GetEtcd()}
		})),
	)
	srv = NewGreeterService("go.micro.srv.greeterbo", service.Client())
	return
}
