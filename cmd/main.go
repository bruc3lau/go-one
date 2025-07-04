package main

import (
	"go-one/internal/router"
)

func main() {
	// 初始化数据库连接
	//if err := database.InitDB(); err != nil {
	//	log.Fatalf("数据库初始化失败: %v", err)
	//}

	r := router.SetupRouter()
	r.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}
