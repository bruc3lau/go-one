package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	//client, err := ethclient.Dial("https://cloudflare-eth.com")
	//client, err := ethclient.Dial("http://127.0.0.1:8545")
	//if err != nil {
	//	log.Fatal(err)
	//}

	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}

	//account := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
	//account := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	//balance, err := client.BalanceAt(context.Background(), account, nil)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(balance, "WEI") // 25893860171173005034
	//
	////blockNumber := big.NewInt(5532993)
	//blockNumber := big.NewInt(0)
	//balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(balanceAt, "WEI") // 25729324269165216042
	//
	//fbalance := new(big.Float)
	//fbalance.SetString(balanceAt.String())
	////ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	//ethValue := new(big.Float).Quo(fbalance, big.NewFloat(1e18))
	//fmt.Println(ethValue, "ETH") // 25.729324269165216041
	//
	//pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	//fmt.Println(pendingBalance, "WEI") // 25729324269165216042

	account := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	balance, err := client.BalanceAt(context.Background(), account, big.NewInt(0))
	fmt.Println(balance)

	floatBalance := new(big.Float)
	floatBalance.SetString(balance.String())

	//ethValue := new(big.Float).Quo(floatBalance, big.NewFloat(math.Pow10(18)))
	n := new(big.Float)
	ethValue := n.Quo(floatBalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(ethValue)
	fmt.Println(ethValue.Text('g', 10))

}
