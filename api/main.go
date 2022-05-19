package main

import (
	"api/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", handler.Hello)
	r.GET("/query", handler.Query)
	r.GET("/:id", handler.Path)
	r.Run()
}
