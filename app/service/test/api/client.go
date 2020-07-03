package api

import (
	"github.com/micro/go-micro/v2"
)

const(
	AppID = "go.micro.srv.demo2"
	WebAppID = "go.micro.web.demo2"
)


func NewClient(service micro.Service) (srv TestDemoService, err error) {
	srv = NewTestDemoService(AppID, service.Client())
	return
}
