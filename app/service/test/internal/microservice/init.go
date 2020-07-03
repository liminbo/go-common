package microservice

import (
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source/file"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	conf "go-common/app/service/store/config"
	"go-common/app/service/test/api"
)

var MicroService micro.Service

func InitMicroService() micro.Service{
	MicroService = micro.NewService(
		micro.Name(api.AppID),
		micro.Version("latest"),
		micro.Metadata(map[string]string{
			"type": "helloworld",
		}),
		// 服务注册
		micro.Registry(etcdv3.NewRegistry(func(options *registry.Options) {
			options.Addrs = []string{conf.GetEtcd()}
		})),

		micro.Flags(
			&cli.StringFlag{
				Name:  "conf_path",
				Value: "./config.toml",
				Usage: "配置文件目录",
			},
		),
		micro.Action(func(ctx *cli.Context) (err error) {
			confPath := ctx.String("conf_path")
			logger.Infof("config:%v", confPath)
			// 加载配置文件
			if err = config.Load(file.NewSource(file.WithPath(confPath))); err != nil {
			}
			return err
		}),
	)
	MicroService.Init()
	return MicroService
}
