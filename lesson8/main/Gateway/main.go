package main

import (
	"context"
	"fmt"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"github.com/go-micro/plugins/v3/registry/consul"
	"log"
	userService "micro/helloworld2/lesson8/Services"
)

func main() {
	csRegistry := consul.NewRegistry(registry.Addrs("20.124.120.120:8500"))
	gService := micro.NewService()
	gService.Init(micro.Registry(csRegistry))
	uService := userService.NewUserService("Grpc", gService.Client())
	list, err := uService.GetUserList(context.Background(), &userService.GetUserListReq{Size: 5})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(list)
}
