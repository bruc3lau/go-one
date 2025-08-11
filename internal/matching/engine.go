// internal/matching/engine.go

package matching

import (
	"log"
)

// MatchingEngine 撮合引擎
type MatchingEngine struct {
	OrderBook  *OrderBook
	OrdersChan chan *Order
	TradesChan chan []Trade
}

// NewMatchingEngine 创建一个新的撮合引擎
func NewMatchingEngine() *MatchingEngine {
	return &MatchingEngine{
		OrderBook:  NewOrderBook(),
		OrdersChan: make(chan *Order, 1000),  // 带缓冲的 Channel
		TradesChan: make(chan []Trade, 1000), // 带缓冲的 Channel
	}
}

// Start 启动撮合引擎的核心循环
func (me *MatchingEngine) Start() {
	log.Println("Matching engine started")
	for order := range me.OrdersChan {
		log.Printf("Received order: %+v", order)
		trades := me.OrderBook.Match(order)
		if len(trades) > 0 {
			log.Printf("Trades generated: %+v", trades)
			me.TradesChan <- trades
		}
	}
	log.Println("Matching engine channel closed, stopping.")
	close(me.TradesChan)
}

// AddOrder 提供给外部调用的方法，用于提交订单
func (me *MatchingEngine) AddOrder(order *Order) {
	me.OrdersChan <- order
}

// Stop 停止引擎 (用于测试)
func (me *MatchingEngine) Stop() {
	close(me.OrdersChan)
	log.Println("Matching engine stopped")
}
