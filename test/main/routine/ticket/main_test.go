package main

import (
	"sync"
	"testing"
)

var testTicket int
var testMx sync.Mutex
var testWg sync.WaitGroup

// 使用 defer 的版本
func sellTicketWithDefer() int {
	testMx.Lock()
	defer testMx.Unlock()
	testTicket++
	return testTicket
}

// 不使用 defer 的版本
func sellTicketWithoutDefer() int {
	testMx.Lock()
	testTicket++
	current := testTicket
	testMx.Unlock()
	return current
}

// 测试使用 defer 的性能
func BenchmarkSellTicketWithDefer(b *testing.B) {
	testTicket = 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		for j := 0; j < 10; j++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for k := 0; k < 10; k++ {
					sellTicketWithDefer()
				}
			}()
		}
		wg.Wait()
	}
}

// 测试不使用 defer 的性能
func BenchmarkSellTicketWithoutDefer(b *testing.B) {
	testTicket = 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		for j := 0; j < 10; j++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for k := 0; k < 10; k++ {
					sellTicketWithoutDefer()
				}
			}()
		}
		wg.Wait()
	}
}

// 使用 *testing.T 的单元测试
func TestSellTicket(t *testing.T) {
	// 重置票数
	testTicket = 0

	// 创建 WaitGroup 来等待所有 goroutine
	var wg sync.WaitGroup

	// 启动 10 个 goroutine，每个售票 10 次
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				sellTicketWithDefer()
			}
		}()
	}

	// 等待所有 goroutine 完成
	wg.Wait()

	// 验证最终票数是否正确（应该是 100 张）
	if testTicket != 100 {
		t.Errorf("期望售出 100 张票，实际售出 %d 张", testTicket)
	}
}
