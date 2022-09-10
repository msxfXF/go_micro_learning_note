package main

import (
	"fmt"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/selector"
	"github.com/go-micro/plugins/v3/registry/consul"
	"log"
)

func main() {
	reg := consul.NewRegistry(registry.Addrs("20.124.120.120:8500"))
	myService, err := reg.GetService("user")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("GetService", myService)

	//node, err := selector.Random(myService)()
	next := selector.Random(myService)
	node, err := next()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v %v %v \n", node.Address, node.Id, node.Metadata)
}
