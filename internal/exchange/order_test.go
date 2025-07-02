package exchange

import (
	"testing"
)

func TestOrderBook_PlaceAndMatchOrders(t *testing.T) {
	ob := NewOrderBook()
	buyOrder := &Order{ID: "1", User: "Alice", Type: Buy, Amount: 10, Price: 100}
	sellOrder := &Order{ID: "2", User: "Bob", Type: Sell, Amount: 10, Price: 90}

	err := ob.PlaceOrder(buyOrder)
	if err != nil {
		t.Fatalf("failed to place buy order: %v", err)
	}
	err = ob.PlaceOrder(sellOrder)
	if err != nil {
		t.Fatalf("failed to place sell order: %v", err)
	}

	trades := ob.MatchOrders()
	if len(trades) != 1 {
		t.Fatalf("expected 1 trade, got %d", len(trades))
	}
	if !buyOrder.Filled || !sellOrder.Filled {
		t.Fatalf("orders should be filled after match")
	}
	if trades[0] != "1<->2" {
		t.Fatalf("unexpected trade result: %v", trades[0])
	}
}

func TestOrderBook_NoMatch(t *testing.T) {
	ob := NewOrderBook()
	buyOrder := &Order{ID: "1", User: "Alice", Type: Buy, Amount: 10, Price: 80}
	sellOrder := &Order{ID: "2", User: "Bob", Type: Sell, Amount: 10, Price: 90}

	_ = ob.PlaceOrder(buyOrder)
	_ = ob.PlaceOrder(sellOrder)

	trades := ob.MatchOrders()
	if len(trades) != 0 {
		t.Fatalf("expected 0 trades, got %d", len(trades))
	}
	if buyOrder.Filled || sellOrder.Filled {
		t.Fatalf("orders should not be filled if no match")
	}
}
