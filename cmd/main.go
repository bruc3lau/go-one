package main

import (
	"go-one/internal/router"
)

func main() {
	r := router.SetupRouter()
	r.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}
