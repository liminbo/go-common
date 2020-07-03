package api

import (
	"github.com/micro/go-micro/v2"
)

const(
	AppID = "go.micro.srv.demo"
	WebAppID = "go.micro.web.demo"
)


func NewClient(service micro.Service) (srv GreeterService, err error) {
	srv = NewGreeterService(AppID, service.Client())
	return
}
