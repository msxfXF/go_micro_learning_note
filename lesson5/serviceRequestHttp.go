package main

import (
	"bytes"
	"fmt"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/selector"
	"github.com/go-micro/plugins/v3/registry/consul"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	reg := consul.NewRegistry(registry.Addrs("20.124.120.120:8500"))
	myService, err := reg.GetService("user")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("GetService", myService)

	for {
		next := selector.RoundRobin(myService)
		node, err := next()
		if err != nil {
			log.Fatal(err)
		}
		res, err := CallHttpAPI(node.Address, "/v1/user", "GET", nil)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v %v %v %s\n", node.Address, node.Id, node.Metadata, res)
		time.Sleep(time.Second)
	}
}

func CallHttpAPI(address string, path string, method string, body []byte) (string, error) {
	req, err := http.NewRequest(method, "http://"+address+path, bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	res, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(res), err
}
