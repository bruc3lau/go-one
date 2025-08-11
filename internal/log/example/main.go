// file: cmd/complexlogger/main.go

package main

import (
	"go-one/internal/log" // 导入我们刚刚创建的包
	"log/slog"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	// 目标1的 Writer：文件轮转
	logRoller := &lumberjack.Logger{
		Filename: "./logs/app-multi.log",
		MaxSize:  1, // 为了演示，设置为 1MB
		Compress: true,
	}

	// 目标1的 Handler：写入文件的 JSON Handler
	fileHandler := slog.NewJSONHandler(logRoller, nil)

	// 目标2的 Handler：写入控制台的 Text Handler
	consoleHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	})

	// 使用我们的 MultiHandler 将上述两个 Handler 组合起来
	multiHandler := log.NewMultiHandler(fileHandler, consoleHandler)

	// 创建最终的 logger
	logger := slog.New(multiHandler)

	// 将其设置为全局 logger
	slog.SetDefault(logger)

	// --- 现在，每一次日志调用都会同时走向两个目标 ---
	slog.Info("系统启动", "version", "1.0.5")
	slog.Debug("这是一个调试信息")
	slog.Warn("CPU 占用率过高", "current_usage", "85%")
	slog.Error("无法连接到数据库", "db_host", "10.0.5.12")
}
