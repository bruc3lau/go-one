package main

import (
	"go-one/internal/log"
	"go-one/internal/router"
	"net/http"
	_ "net/http/pprof"
)

func main() {

	go func() {
		http.ListenAndServe("0.0.0.0:6060", nil)
	}()

	// 初始化数据库连接
	//if err := database.InitDB(); err != nil {
	//	log.Fatalf("数据库初始化失败: %v", err)
	//}

	//初始化日志系统
	//log.InitLoggerWithExample()
	log.InitLogger()

	r := router.SetupRouter()
	err := r.Run(":8080")
	if err != nil {
		return
	} // 监听并在 0.0.0.0:8080 上启动服务
}
