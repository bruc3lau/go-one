package log

import (
	"log/slog"
	"os"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"
)

func InitLoggerWithExample() {
	// 1. 配置 lumberjack
	// 这是所有魔法发生的地方
	logRoller := &lumberjack.Logger{
		Filename:   "./logs/app.log", // 日志文件的位置
		MaxSize:    100,              // 每个日志文件的最大大小，单位是 MB
		MaxBackups: 3,                // 保留的旧日志文件的最大数量
		MaxAge:     28,               // 旧日志文件保留的最大天数
		Compress:   true,             // 是否压缩/归档旧日志
	}

	// 创建 handler: 一个用于文件(JSON)，一个用于控制台(Text)
	fileHandler := slog.NewJSONHandler(logRoller, nil)
	consoleHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})

	// 使用 MultiHandler 组合它们
	multiHandler := NewMultiHandler(fileHandler, consoleHandler)
	logger := slog.New(multiHandler)
	slog.SetDefault(logger)

	// 3. 将 logger 设置为全局默认 logger
	slog.SetDefault(logger)

	slog.Info("日志系统初始化成功", "log_file", logRoller.Filename)

	// 4. 模拟持续写入日志
	// 在真实应用中，这些日志会来自你的各个业务模块
	for i := 0; i < 200; i++ {
		slog.Info(
			"这是一条业务日志记录",
			"iteration", i+1,
			"user_id", "user-54321",
		)
		time.Sleep(5 * time.Millisecond)
	}

	slog.Warn("一个潜在的问题发生了，需要关注")
	slog.Error("一个严重的错误发生了", "error", "database connection lost")
	slog.Info("程序运行结束")
}
