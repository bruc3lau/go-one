package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// weiToEth 将 Wei 转换为 ETH
func weiToEth(wei *big.Int) *big.Float {
	f := new(big.Float)
	f.SetString(wei.String())
	ethValue := new(big.Float).Quo(f, big.NewFloat(1e18))
	return ethValue
}

func test1() {

	//fmt.Println("程序启动成功！")
	//fmt.Printf("Infura Project ID: %s\n", AppConfig.InfuraProjectID)
	//
	//fmt.Println("连接到 Sepolia 测试网络...")
	//// 使用 fmt.Sprintf 从配置中动态构建 Infura URL
	//infuraURL := fmt.Sprintf("https://sepolia.infura.io/v3/%s", AppConfig.InfuraProjectID)
	//client, err := ethclient.Dial(infuraURL)
	//if err != nil {
	//	log.Fatal("连接失败:", err)
	//}
	//fmt.Println("连接成功！")

	// 获取最新区块号
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal("获取区块号失败:", err)
	}
	fmt.Printf("\n1. 最新区块号: %d\n", header.Number.Uint64())

	// 获取账户余额
	//account := common.HexToAddress("0x742d35Cc6634C0532925a3b844Bc454e4438f44e")
	//account := common.HexToAddress("0x35E1f780604C869Af0a9f94399B842244a5A6676")

	account := common.HexToAddress("0xf08C0a03792Dd0043107046Cb799a382ca3b5ba5")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal("获取余额失败:", err)
	}
	ethBalance := weiToEth(balance)
	fmt.Printf("\n2. 账户信息:\n   地址: %s\n   余额: %s ETH (%s Wei)\n", account.Hex(), ethBalance.String(), balance.String())

	// 获取当前 gas 价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal("获取 Gas 价格失败:", err)
	}
	fmt.Printf("\n3. 当前 Gas 价格: %s Gwei (%s Wei)\n", new(big.Int).Div(gasPrice, big.NewInt(1e9)).String(), gasPrice.String())

	// 获取指定区块的信息
	blockNumber := big.NewInt(100)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal("获取区块信息失败:", err)
	}
	fmt.Printf("\n4. 区块信息:\n   区块号: %d\n   哈希值: %s\n   交易数量: %d\n",
		blockNumber,
		block.Hash().Hex(),
		len(block.Transactions()),
	)
}

func test2() {

	account := common.HexToAddress("0xf08C0a03792Dd0043107046Cb799a382ca3b5ba5")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal("获取余额失败:", err)
	}
	ethBalance := weiToEth(balance)
	fmt.Printf("\n2. 账户信息:\n   地址: %s\n   余额: %s ETH (%s Wei)\n", account.Hex(), ethBalance.String(), balance.String())

}

func QueryBlock() {
	// 这里可以添加查询区块的逻辑

}

func init() {
	//fmt.Println(666)
}
func main() {
	//common.BytesToAddress()
	//test1()
	test2()
}
