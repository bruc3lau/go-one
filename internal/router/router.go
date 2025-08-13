package router

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// a new middleware to count HTTP requests.
var httpRequestsTotal = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total number of HTTP requests",
	},
	[]string{"method", "path"},
)

func init() {
	prometheus.MustRegister(httpRequestsTotal)
}

func prometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)
		httpRequestsTotal.With(prometheus.Labels{"method": c.Request.Method, "path": c.Request.URL.Path}).Inc()
		fmt.Println(duration)
	}
}

// 模拟一个共享变量
var counter int

func init() {
	// 根据环境变量设置 gin 模式
	mode := os.Getenv("GIN_MODE")
	if mode == "" {
		mode = os.Getenv("GO_ENV") // 如果没有设置 GIN_MODE，则查看 GO_ENV
	}

	switch mode {
	case "production", "prod":
		gin.SetMode(gin.ReleaseMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.DebugMode)
		//gin.SetMode(gin.ReleaseMode)
	}
}

func SetupRouter() *gin.Engine {
	// 使用 gin.New() 替代 gin.Default()
	r := gin.New()

	// 手动添加中间件
	r.Use(gin.Logger())   // 添加日志中间件
	r.Use(gin.Recovery()) // 添加恢复中间件
	r.Use(prometheusMiddleware())

	// 注册 prometheus 路由
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

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

	// 测试请求并发处理
	r.GET("/concurrent", func(c *gin.Context) {
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
	r.GET("/request-scope", func(c *gin.Context) {
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

	r.GET("/getUser", func(c *gin.Context) {
		user := &UserInfo{}
		if err := c.ShouldBind(&user); err == nil {
			c.JSON(http.StatusOK, user)
		} else {
			c.JSON(http.StatusOK, gin.H{
				//"code":    500,
				"message": err.Error(),
			})
		}
	})

	return r
}

type UserInfo struct {
	ID       int64  `json:"id" form:"id"`
	UserName string `json:"user_name" form:"user_name"`
}
