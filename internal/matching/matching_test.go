// internal/matching/matching_test.go

package matching

import (
	"testing"
	"time"
)

// assertFloatEqual 比较两个浮点数是否在误差范围内相等
func assertFloatEqual(t *testing.T, a, b, epsilon float64) {
	t.Helper()
	if (a-b) > epsilon || (b-a) > epsilon {
		t.Errorf("Expected %.2f, got %.2f", a, b)
	}
}

// TestSimpleMatch 测试一个简单的买单吃掉一个卖单
func TestSimpleMatch(t *testing.T) {
	ob := NewOrderBook()
	sellOrder := &Order{ID: 1, Side: Sell, Price: 100, Amount: 10, Timestamp: time.Now().UnixNano()}
	ob.AddOrder(sellOrder)

	buyOrder := &Order{ID: 2, Side: Buy, Price: 100, Amount: 5, Timestamp: time.Now().UnixNano()}
	trades := ob.Match(buyOrder)

	if len(trades) != 1 {
		t.Fatalf("Expected 1 trade, got %d", len(trades))
	}

	trade := trades[0]
	assertFloatEqual(t, trade.Price, 100, 1e-9)
	assertFloatEqual(t, trade.Amount, 5, 1e-9)
	assertFloatEqual(t, ob.Asks[0].Amount, 5, 1e-9) // 卖单剩余数量
	if len(ob.Bids) != 0 {
		t.Errorf("Expected 0 bids, got %d", len(ob.Bids))
	}
}

// TestPartialMatch 测试部分成交
func TestPartialMatch(t *testing.T) {
	ob := NewOrderBook()
	sellOrder := &Order{ID: 1, Side: Sell, Price: 100, Amount: 3, Timestamp: time.Now().UnixNano()}
	ob.AddOrder(sellOrder)

	buyOrder := &Order{ID: 2, Side: Buy, Price: 100, Amount: 5, Timestamp: time.Now().UnixNano()}
	trades := ob.Match(buyOrder)

	if len(trades) != 1 {
		t.Fatalf("Expected 1 trade, got %d", len(trades))
	}

	if len(ob.Asks) != 0 {
		t.Errorf("Expected 0 asks, got %d", len(ob.Asks))
	}

	if len(ob.Bids) != 1 {
		t.Fatalf("Expected 1 bid, got %d", len(ob.Bids))
	}

	assertFloatEqual(t, ob.Bids[0].Amount, 2, 1e-9) // 买单剩余数量
}

// TestMultiLevelMatch 测试吃掉多档订单
func TestMultiLevelMatch(t *testing.T) {
	ob := NewOrderBook()
	ob.AddOrder(&Order{ID: 1, Side: Sell, Price: 100, Amount: 5, Timestamp: time.Now().UnixNano()})
	ob.AddOrder(&Order{ID: 2, Side: Sell, Price: 101, Amount: 5, Timestamp: time.Now().UnixNano()})

	buyOrder := &Order{ID: 3, Side: Buy, Price: 101, Amount: 8, Timestamp: time.Now().UnixNano()}
	trades := ob.Match(buyOrder)

	if len(trades) != 2 {
		t.Fatalf("Expected 2 trades, got %d", len(trades))
	}

	if len(ob.Asks) != 1 {
		t.Fatalf("Expected 1 ask, got %d", len(ob.Asks))
	}
	assertFloatEqual(t, ob.Asks[0].Price, 101, 1e-9)
	assertFloatEqual(t, ob.Asks[0].Amount, 2, 1e-9) // 第二个卖单剩余数量

	if len(ob.Bids) != 0 {
		t.Errorf("Expected 0 bids, got %d", len(ob.Bids))
	}
}

// TestNoMatch 测试无法撮合的情况
func TestNoMatch(t *testing.T) {
	ob := NewOrderBook()
	ob.AddOrder(&Order{ID: 1, Side: Sell, Price: 101, Amount: 10, Timestamp: time.Now().UnixNano()})
	buyOrder := &Order{ID: 2, Side: Buy, Price: 100, Amount: 10, Timestamp: time.Now().UnixNano()}
	trades := ob.Match(buyOrder)

	if len(trades) != 0 {
		t.Fatalf("Expected 0 trades, got %d", len(trades))
	}

	if len(ob.Asks) != 1 {
		t.Errorf("Expected 1 ask")
	}
	if len(ob.Bids) != 1 {
		t.Errorf("Expected 1 bid")
	}
}
