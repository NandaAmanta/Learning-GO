package main

import (
	"api/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", handler.GetAllUser)
	r.GET("/query", handler.GetUsersByEmail)
	r.GET("/:id", handler.GetDetailUser)
	r.POST("/", handler.AddUser)
	r.Run()
}
