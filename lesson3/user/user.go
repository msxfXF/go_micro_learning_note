package main

import (
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/web"
	"github.com/gin-gonic/gin"
	"github.com/go-micro/plugins/v3/registry/consul"
)

func main() {

	csReg := consul.NewRegistry(registry.Addrs("20.124.120.120:8500"))
	r := gin.Default()

	v1 := r.Group("/v1")
	v1.Handle("GET", "/user", func(context *gin.Context) {
		context.JSON(200, gin.H{"users": GetUserList(5)})
	})

	server := web.NewService(
		web.Name("user"),
		web.Address(":8002"),
		web.Handler(r),
		web.Registry(csReg),
	)
	server.Run()
}
