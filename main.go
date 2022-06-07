package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	var router = gin.Default()
	var addr = ":8080"

	v1 := router.Group("/api/v1")

	v1.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})
	router.GET("/get", func(c *gin.Context) {
		c.String(200, "get")
	})
	v1.POST("/post", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "post")
	})
	log.Fatal(router.Run(addr))
}
