package main

import(
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	"go-common/app/service/store/api"
	"go-common/app/service/store/config"
)

func main() {
	service := micro.NewService(
		micro.Registry(etcdv3.NewRegistry(func(options *registry.Options) {
			options.Addrs = []string{config.GetEtcd()}
		})),
	)
	srv := api.NewGreeterService(api.AppID, service.Client())
	rep,err := srv.Hello(context.Background(), &api.Request{Name: "John"})
	if err != nil{
		fmt.Printf("err:%v", err)
	}
	fmt.Printf("vvvv:%v", rep)
}
