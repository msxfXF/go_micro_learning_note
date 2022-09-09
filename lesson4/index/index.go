package main

import (
	"github.com/asim/go-micro/v3/web"
	"github.com/gin-gonic/gin"
)

func main() {

	//consul.NewRegistry(registry.Addrs("20.124.120.120:8500"))
	r := gin.Default()
	r.Handle("GET", "/", func(context *gin.Context) {
		context.JSON(200, gin.H{"page": "index"})
	})

	server := web.NewService(
		web.Name("index"),
		web.Address(":8001"),
		web.Handler(r),
	)
	server.Run()
}
