package main

import (
	"fmt"
	"log" // 推荐使用 log 包来处理启动错误
	"os"
)

// Config 存储应用配置
type Config struct {
	InfuraProjectID string
}

// AppConfig 是一个全局变量，用于存储加载后的配置。
// 程序的其他部分可以直接访问这个变量。
var AppConfig *Config

// init 函数会在 main 函数执行之前自动运行。
func init() {
	var err error
	// 调用 LoadConfig 并将结果赋值给全局变量 AppConfig
	AppConfig, err = LoadConfig()
	if err != nil {
		// 如果配置加载失败，程序无法正常运行，
		// 使用 log.Fatalf 打印错误并立即退出程序。
		log.Fatalf("错误：无法加载配置: %v", err)
	}
	fmt.Println("配置加载成功:", AppConfig)
}

// LoadConfig 从环境变量加载配置
func LoadConfig() (*Config, error) {
	projectID := os.Getenv("INFURA_PROJECT_ID")
	if projectID == "" {
		return nil, fmt.Errorf("环境变量 INFURA_PROJECT_ID 未设置")
	}

	return &Config{
		InfuraProjectID: projectID,
	}, nil
}
