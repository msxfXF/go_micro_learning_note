package main

import (
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"github.com/go-micro/plugins/v3/registry/consul"
	"log"
	"micro/helloworld2/lesson8/ServiceImpl"
	userService "micro/helloworld2/lesson8/Services"
)

func main() {
	csRegistry := consul.NewRegistry(registry.Addrs("20.124.120.120:8500"))
	mService := micro.NewService(
		micro.Address(":8003"),
		micro.Name("Grpc"),
		micro.Registry(csRegistry),
	)
	mService.Init()
	err := userService.RegisterUserServiceHandler(mService.Server(), new(ServiceImpl.UserServiceImpl))
	if err != nil {
		log.Fatal(err)
	}
	err = mService.Run()
	if err != nil {
		log.Fatal(err)
	}

}
