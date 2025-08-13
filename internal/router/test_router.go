package router

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func TestInit(r *gin.Engine) {
	// 路由配置
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

	testRouter := r.Group("/test")
	{
		// 测试请求并发处理
		testRouter.GET("/concurrent", func(c *gin.Context) {
			// 获取当前 goroutine 的ID
			routineID := fmt.Sprintf("%p", c.Request)
			counter++ // 修改共享变量

			// 模拟耗时操作
			time.Sleep(time.Second)

			c.JSON(http.StatusOK, gin.H{
				"routineID": routineID,
				"counter":   counter,
				"time":      time.Now().Format("15:04:05.000"),
			})
		})

		// 测试请求隔离性
		testRouter.GET("/request-scope", func(c *gin.Context) {
			// 从上下文中获取值（如果存在）
			value, exists := c.Get("testKey")

			// 在当前请求上下文中设置值
			c.Set("testKey", time.Now().String())

			c.JSON(http.StatusOK, gin.H{
				"routineID":     fmt.Sprintf("%p", c.Request),
				"previousValue": value,
				"exists":        exists,
				"currentTime":   time.Now().Format("15:04:05.000"),
			})
		})
	}

}
