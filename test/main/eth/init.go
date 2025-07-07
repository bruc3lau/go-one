package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

var client *ethclient.Client

// 在这个文件中集中处理所有初始化逻辑
func init() {
	// 1. 明确地先加载配置
	if err := loadAppConfig(); err != nil {
		log.Fatalf("初始化失败：无法加载配置: %v", err)
	}

	// 2. 再根据配置初始化客户端
	if err := setupEthClient(); err != nil {
		log.Fatalf("初始化失败：无法连接到以太坊: %v", err)
	}
}

func loadAppConfig() error {
	// ... 加载 AppConfig 的逻辑 ...
	var err error
	// 调用 LoadConfig 并将结果赋值给全局变量 AppConfig
	AppConfig, err = LoadConfig()
	if err != nil {
		// 如果配置加载失败，程序无法正常运行，
		// 使用 log.Fatalf 打印错误并立即退出程序。
		log.Fatalf("错误：无法加载配置: %v", err)
	}
	fmt.Println("配置加载成功:", AppConfig.InfuraProjectID)
	return nil
}

func setupEthClient() error {
	// ... 使用 AppConfig 初始化 client 的逻辑 ...
	fmt.Println("程序启动成功！")
	fmt.Printf("Infura Project ID: %s\n", AppConfig.InfuraProjectID)

	fmt.Println("连接到 Sepolia 测试网络...")
	// 使用 fmt.Sprintf 从配置中动态构建 Infura URL
	infuraURL := fmt.Sprintf("https://sepolia.infura.io/v3/%s", AppConfig.InfuraProjectID)
	cl, err := ethclient.Dial(infuraURL)
	if err != nil {
		log.Fatal("连接失败:", err)
	}
	fmt.Println("连接成功！")
	client = cl
	return nil
}
