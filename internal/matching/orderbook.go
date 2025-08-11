// internal/matching/orderbook.go

package matching

import (
	"sort"
	"time"
)

// NewOrderBook 创建一个新的订单簿
func NewOrderBook() *OrderBook {
	return &OrderBook{
		Bids: make([]*Order, 0),
		Asks: make([]*Order, 0),
	}
}

// AddOrder 将订单添加到订单簿
func (ob *OrderBook) AddOrder(order *Order) {
	if order.Side == Buy {
		ob.Bids = append(ob.Bids, order)
		// 买单按价格降序排序
		sort.Slice(ob.Bids, func(i, j int) bool {
			if ob.Bids[i].Price == ob.Bids[j].Price {
				return ob.Bids[i].Timestamp < ob.Bids[j].Timestamp
			}
			return ob.Bids[i].Price > ob.Bids[j].Price
		})
	} else {
		ob.Asks = append(ob.Asks, order)
		// 卖单按价格升序排序
		sort.Slice(ob.Asks, func(i, j int) bool {
			if ob.Asks[i].Price == ob.Asks[j].Price {
				return ob.Asks[i].Timestamp < ob.Asks[j].Timestamp
			}
			return ob.Asks[i].Price < ob.Asks[j].Price
		})
	}
}

// Match 撮合订单并返回成交记录
func (ob *OrderBook) Match(takerOrder *Order) []Trade {
	trades := make([]Trade, 0)

	if takerOrder.Side == Buy {
		// 新订单是买单，与卖单簿 (Asks) 撮合
		for len(ob.Asks) > 0 && takerOrder.Amount > 0 {
			makerOrder := ob.Asks[0] // 最优卖单
			if takerOrder.Price < makerOrder.Price {
				break // 买方出价低于卖方要价，无法撮合
			}

			tradeAmount := min(takerOrder.Amount, makerOrder.Amount)
			trades = append(trades, Trade{
				TakerOrderID: takerOrder.ID,
				MakerOrderID: makerOrder.ID,
				Price:        makerOrder.Price,
				Amount:       tradeAmount,
				Timestamp:    time.Now().UnixNano(),
			})

			takerOrder.Amount -= tradeAmount
			makerOrder.Amount -= tradeAmount

			if makerOrder.Amount == 0 {
				// 移除已完全成交的 maker order
				ob.Asks = ob.Asks[1:]
			}
		}
	} else {
		// 新订单是卖单，与买单簿 (Bids) 撮合
		for len(ob.Bids) > 0 && takerOrder.Amount > 0 {
			makerOrder := ob.Bids[0] // 最优买单
			if takerOrder.Price > makerOrder.Price {
				break // 卖方要价高于买方出价，无法撮合
			}

			tradeAmount := min(takerOrder.Amount, makerOrder.Amount)
			trades = append(trades, Trade{
				TakerOrderID: takerOrder.ID,
				MakerOrderID: makerOrder.ID,
				Price:        makerOrder.Price,
				Amount:       tradeAmount,
				Timestamp:    time.Now().UnixNano(),
			})

			takerOrder.Amount -= tradeAmount
			makerOrder.Amount -= tradeAmount

			if makerOrder.Amount == 0 {
				// 移除已完全成交的 maker order
				ob.Bids = ob.Bids[1:]
			}
		}
	}

	// 如果新订单未完全成交，则将其加入订单簿
	if takerOrder.Amount > 0 {
		ob.AddOrder(takerOrder)
	}

	return trades
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
