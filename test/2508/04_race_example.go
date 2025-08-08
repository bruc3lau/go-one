package main

import (
	"fmt"
	"sync"
)

// 一个全局变量，我们将并发地修改它
var counter int

func main() {
	var wg sync.WaitGroup
	// 启动 100 个 goroutine 同时增加 counter 的值
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// 这里没有加锁，直接对 counter 进行读和写
			counter++
		}()
	}

	wg.Wait()
	fmt.Printf("Final counter value: %d\n", counter)
}

//  go run -race race_example.go
//  go run -race test/2508/04_race_example.go
