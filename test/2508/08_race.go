package main

import (
	"fmt"
	"sync"
	"time"
)

// go run --race test/2508/08_race.go
func main() {
	var x int
	var mu sync.RWMutex

	go func() {
		mu.Lock()
		defer mu.Unlock()
		x = 1 // 写
	}()

	go func() {
		mu.Lock()
		defer mu.Unlock()
		x = 2 // 写
	}()

	// 等待 goroutine 执行完成
	time.Sleep(100 * time.Millisecond)

	mu.RLock()
	fmt.Println("final value of x:", x)
	mu.RUnlock()
}
