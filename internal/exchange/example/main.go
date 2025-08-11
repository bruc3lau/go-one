// cmd/exchange/main.go

package main

import (
	"log"
	"sync"
	"time"

	"go-one/internal/matching"
)

func main() {
	var wg sync.WaitGroup

	// 1. 创建撮合引擎

	engine := matching.NewMatchingEngine()

	wg.Add(2) // 等待两个后台 goroutine

	// 2. 启动撮合引擎 Goroutine
	go func() {
		defer wg.Done()
		engine.Start()
	}()

	// 3. 启动一个 Goroutine 来处理成交结果
	go func() {
		defer wg.Done()
		for trades := range engine.TradesChan {
			log.Printf(">>>>>> Processed Trades: %+v\n", trades)
		}
		log.Println("Trade processing finished.")
	}()

	// 4. 模拟提交一系列订单
	// 挂一个卖单: 卖 10 个，价格 100
	sellOrder1 := &matching.Order{ID: 1, UserID: 101, Side: matching.Sell, Price: 100, Amount: 10, Timestamp: time.Now().UnixNano()}
	engine.AddOrder(sellOrder1)
	log.Printf("Submitted order: %+v", sellOrder1)

	// 挂一个买单: 买 5 个，价格 99

	buyOrder1 := &matching.Order{ID: 2, UserID: 102, Side: matching.Buy, Price: 99, Amount: 5, Timestamp: time.Now().UnixNano()}
	engine.AddOrder(buyOrder1)
	log.Printf("Submitted order: %+v", buyOrder1)

	// 挂一个新的卖单: 卖 8 个，价格 101
	sellOrder2 := &matching.Order{ID: 3, UserID: 103, Side: matching.Sell, Price: 101, Amount: 8, Timestamp: time.Now().UnixNano()}
	engine.AddOrder(sellOrder2)
	log.Printf("Submitted order: %+v", sellOrder2)

	// 提交一个吃单的买单，价格 102，数量 20，应该会吃掉所有卖单
	log.Println(">>> Submitting a large buy order to match existing asks...")
	buyOrder2 := &matching.Order{ID: 4, UserID: 104, Side: matching.Buy, Price: 102, Amount: 20, Timestamp: time.Now().UnixNano()}
	engine.AddOrder(buyOrder2)
	log.Printf("Submitted order: %+v", buyOrder2)

	// 给一点时间让最后一个订单完成撮合
	time.Sleep(200 * time.Millisecond)
	logCurrentOrderBook(engine.OrderBook)

	// 5. 所有订单提交完毕，关闭订单通道，这是优雅关闭的信号
	log.Println("All orders submitted. Closing orders channel.")
	close(engine.OrdersChan)

	// 6. 等待所有 goroutine 执行完毕
	log.Println("Waiting for goroutines to finish...")
	wg.Wait()
	log.Println("All goroutines finished. Exiting.")
}

func logCurrentOrderBook(ob *matching.OrderBook) {
	log.Println("--- Current Order Book ---")
	log.Printf("Asks (%d):", len(ob.Asks))
	for _, ask := range ob.Asks {
		log.Printf("  %+v", *ask)
	}
	log.Printf("Bids (%d):", len(ob.Bids))
	for _, bid := range ob.Bids {
		log.Printf("  %+v", *bid)
	}
	log.Println("--------------------------")
}
