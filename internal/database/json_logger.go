package database

import (
	"context"
	"encoding/json"
	"fmt"
	"gorm.io/gorm/logger"
	"time"
)

// JSONLogger 自定义的 JSON 格式日志记录器
type JSONLogger struct {
	LogLevel      logger.LogLevel
	SlowThreshold time.Duration
}

// NewJSONLogger 创建一个新的 JSON 日志记录器
func NewJSONLogger(config logger.Config) *JSONLogger {
	return &JSONLogger{
		LogLevel:      config.LogLevel,
		SlowThreshold: config.SlowThreshold,
	}
}

// LogMode 设置日志级别
func (l *JSONLogger) LogMode(level logger.LogLevel) logger.Interface {
	newLogger := *l
	newLogger.LogLevel = level
	return &newLogger
}

// Info 记录信息级别日志
func (l *JSONLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Info {
		l.log("info", msg, nil, data...)
	}
}

// Warn 记录警告级别日志
func (l *JSONLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Warn {
		l.log("warn", msg, nil, data...)
	}
}

// Error 记录错误级别日志
func (l *JSONLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Error {
		l.log("error", msg, nil, data...)
	}
}

// Trace 记录 SQL 执行日志
func (l JSONLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if l.LogLevel <= logger.Silent {
		return
	}

	elapsed := time.Since(begin)
	sql, rows := fc()

	logData := map[string]interface{}{
		"sql":           sql,
		"rows_affected": rows,
		"elapsed_time":  elapsed.String(),
	}

	if err != nil {
		l.log("error", "SQL执行错误", err, logData)
		return
	}

	if l.SlowThreshold > 0 && elapsed > l.SlowThreshold {
		l.log("warn", "慢查询", nil, logData)
		return
	}

	l.log("trace", "SQL执行", nil, logData)
}

// log 通用日志记录方法
func (l *JSONLogger) log(level, msg string, err error, data ...interface{}) {
	logData := map[string]interface{}{
		"timestamp": time.Now().Format(time.RFC3339),
		"level":     level,
		"message":   msg,
	}

	if err != nil {
		logData["error"] = err.Error()
	}

	if len(data) > 0 {
		logData["data"] = data
	}

	jsonBytes, err := json.MarshalIndent(logData, "", "    ")
	if err != nil {
		fmt.Printf("错误：日志序列化失败：%v\n", err)
		return
	}

	fmt.Println(string(jsonBytes))
}
