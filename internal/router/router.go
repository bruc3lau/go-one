package router

import (
	"fmt"
	"go-one/internal/log"
	"io"
	"log/slog"
	"os"
	"strings"
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

// logWriter is a bridge between gin's log output and slog.
type logWriter struct {
	logger *slog.Logger
}

// Write implements io.Writer.
func (w *logWriter) Write(p []byte) (n int, err error) {
	// Trim trailing newline because slog adds its own.
	message := strings.TrimRight(string(p), "\n")
	w.logger.Info(message)
	return len(p), nil
}

func SetupRouter() *gin.Engine {
	// 初始化 slog logger
	logger, _ := log.InitLogger()

	// 使用 gin.New() 替代 gin.Default()
	r := gin.New()

	// 创建一个 writer，将 gin 的日志写入 slog
	writer := &logWriter{logger: logger}

	// 手动添加中间件
	// 使用 LoggerWithConfig 并将输出重定向到 slog
	r.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		Output:    io.MultiWriter(writer, os.Stdout), // 同时输出到 slog 和标准输出
		SkipPaths: []string{"/metrics"},              // 可选：跳过特定路径的日志
	}))
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
