package router

import (
	"fmt"
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

	//test
	TestInit(r)

	//user
	UserInit(r)

	return r
}
