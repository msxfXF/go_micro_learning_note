package main

import (
	"flag"
	"fmt"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/web"
	"github.com/gin-gonic/gin"
	"github.com/go-micro/plugins/v3/registry/consul"
)

func main() {
	address := flag.String("address", ":8002", "ip:port")
	flag.Parse()
	fmt.Println(*address)
	csReg := consul.NewRegistry(registry.Addrs("20.124.120.120:8500"))
	r := gin.Default()

	v1 := r.Group("/v1")
	v1.Handle("POST", "/user", func(context *gin.Context) {
		context.JSON(200, gin.H{"data": GetUserList(5)})
	})

	server := web.NewService(
		web.Name("user"),
		web.Address(*address),
		web.Handler(r),
		web.Registry(csReg),
		web.Metadata(map[string]string{"protocol": "http"}),
	)
	server.Run()
}
