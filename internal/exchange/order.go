package exchange

import (
	"errors"
	"sync"
)

type OrderType int

const (
	Buy OrderType = iota
	Sell
)

type Order struct {
	ID     string
	User   string
	Type   OrderType
	Amount float64
	Price  float64
	Filled bool
}

type OrderBook struct {
	BuyOrders  []*Order
	SellOrders []*Order
	mu         sync.Mutex
}

func NewOrderBook() *OrderBook {
	return &OrderBook{
		BuyOrders:  make([]*Order, 0),
		SellOrders: make([]*Order, 0),
	}
}

func (ob *OrderBook) PlaceOrder(order *Order) error {
	ob.mu.Lock()
	defer ob.mu.Unlock()

	switch order.Type {
	case Buy:
		ob.BuyOrders = append(ob.BuyOrders, order)
	case Sell:
		ob.SellOrders = append(ob.SellOrders, order)
	default:
		return errors.New("invalid order type")
	}
	return nil
}

// MatchOrders 执行简单的撮合逻辑
func (ob *OrderBook) MatchOrders() []string {
	ob.mu.Lock()
	defer ob.mu.Unlock()
	trades := []string{}
	for i := 0; i < len(ob.BuyOrders); i++ {
		buy := ob.BuyOrders[i]
		if buy.Filled {
			continue
		}
		for j := 0; j < len(ob.SellOrders); j++ {
			sell := ob.SellOrders[j]
			if sell.Filled {
				continue
			}
			if buy.Price >= sell.Price && buy.Amount == sell.Amount {
				buy.Filled = true
				sell.Filled = true
				trades = append(trades, buy.ID+"<->"+sell.ID)
				break
			}
		}
	}
	return trades
}
