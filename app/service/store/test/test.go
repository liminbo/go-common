package main

import(
	"context"
	"fmt"
	"go-common/app/service/store/api"
)

func main() {

	service,_ := api.NewService()
	rep,err := service.Hello(context.Background(), &api.Request{Name: "John"})
	if err != nil{
		fmt.Printf("err:%v", err)
	}
	fmt.Printf("vvvv:%v", rep)
}
