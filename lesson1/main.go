package main

import (
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/web"
	"github.com/gin-gonic/gin"
	"github.com/go-micro/plugins/v3/registry/consul"
)

func main() {

	//consul.NewRegistry(registry.Addrs("20.124.120.120:8500"))
	csReg := consul.NewRegistry(registry.Addrs("20.124.120.120:8500"))
	r := gin.Default()
	r.Handle("GET", "/user", func(context *gin.Context) {
		context.String(200, "user api")
	})
	r.Handle("GET", "/admin", func(context *gin.Context) {
		context.String(200, "admin api")
	})
	server := web.NewService(
		web.Name("hello"),
		web.Address(":8001"),
		web.Handler(r),
		web.Registry(csReg),
	)
	server.Run()
}
