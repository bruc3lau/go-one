package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, GoOne!",
		})
	})

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}
