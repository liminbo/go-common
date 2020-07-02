package config

import (
	"github.com/micro/go-micro/v2/config"
	"go-common/library/database/es"
	"go-common/library/database/orm"
	"os"
)

func GetMySQL() (db *orm.Config, err error){
	db = new(orm.Config)
	err = config.Get("MySQL").Scan(db)
	return
}

func GetES() (e *es.Config, err error){
	e = new(es.Config)
	err = config.Get("ES").Scan(e)
	return
}

// 通过环境变量获取
func GetEtcd() (addr string){
	addr = os.Getenv("etcd")
	return
}

// 通过环境变量获取
func GetJaeger() (addr string){
	addr = os.Getenv("jaeger")
	return
}