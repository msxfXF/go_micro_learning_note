package main

import (
	"context"
	"github.com/asim/go-micro/v3/client"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/selector"
	"github.com/go-micro/plugins/v3/client/http"
	"github.com/go-micro/plugins/v3/registry/consul"
	"log"
	"micro/helloworld2/lesson7/Models"
	"time"
)

func main() {
	reg := consul.NewRegistry(registry.Addrs("20.124.120.120:8500"))
	for {
		mySelector := selector.NewSelector(
			selector.Registry(reg),
			selector.SetStrategy(selector.RoundRobin),
		)
		res, err := CallAPI(mySelector)
		if err != nil && res != nil {
			log.Println(err)
		}
		log.Println(res.Data)
		time.Sleep(time.Second)
	}
}

func CallAPI(s selector.Selector) (*Models.UserResp, error) {

	myClient := http.NewClient(client.Selector(s), client.ContentType("application/json"))
	req := myClient.NewRequest("user", "/v1/user", Models.UserReq{Size: 1})
	rsp := Models.UserResp{}
	err := myClient.Call(context.Background(), req, &rsp)
	if err != nil {
		return nil, err
	}
	return &rsp, err
}
