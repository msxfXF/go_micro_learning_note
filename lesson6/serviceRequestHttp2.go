package main

import (
	"context"
	"github.com/asim/go-micro/v3/client"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/selector"
	"github.com/go-micro/plugins/v3/client/http"
	"github.com/go-micro/plugins/v3/registry/consul"
	"log"
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
		if err != nil {
			log.Println(err)
		}
		log.Println(res["data"])
		time.Sleep(time.Second)
	}
}

func CallAPI(s selector.Selector) (map[string]interface{}, error) {

	myClient := http.NewClient(client.Selector(s), client.ContentType("application/json"))
	req := myClient.NewRequest("user", "/v1/user", map[string]string{})
	rsp := map[string]interface{}{}
	err := myClient.Call(context.Background(), req, &rsp)
	if err != nil {
		return nil, err
	}
	return rsp, err
}
