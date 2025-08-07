package main

import (
	"fmt"
	"sync"
)

var sm sync.Map

func main() {
	//unsafeMapExampleError()
	//unsafeMapExample()
	syncMapExample()
}

// fatal error: concurrent map read and map write
func unsafeMapExampleError() {
	unsafeMap := make(map[int]int)

	// 使用 WaitGroup 等待所有 goroutine 完成
	var wg sync.WaitGroup
	wg.Add(2) // 添加两个 goroutine

	// Goroutine 1: 写入 map
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			unsafeMap[i] = i * i // 写入操作
		}
	}()

	// Goroutine 2: 读取 map
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			// 在另一个 goroutine 正在写入的同时进行读取
			_ = unsafeMap[i] // 读取操作

		}
	}()
	// 等待两个 goroutine 结束
	wg.Wait()
	fmt.Println("Done (you will likely not see this message)")
}

func unsafeMapExample() {
	unsafeMap := make(map[int]int)
	var mu sync.RWMutex

	for i := 0; i < 1000; i++ {
		mu.Lock()
		unsafeMap[i] = i * i // 写入操作
		mu.Unlock()
	}

	for i := 0; i < 1000; i++ {
		mu.RLock()
		fmt.Println(unsafeMap[i]) // 读取操作
		mu.RUnlock()
	}
	fmt.Println("Done (you will likely not see this message)")
}

func syncMapExample() {
	var sm sync.Map

	// 使用 WaitGroup 等待所有 goroutine 完成
	var wg sync.WaitGroup
	wg.Add(2) // 添加两个 goroutine

	// Goroutine 1: 写入 map
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			sm.Store(i, i*i)
		}
	}()

	// Goroutine 2: 读取 map
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			// 在另一个 goroutine 正在写入的同时进行读取
			_, _ = sm.Load(i) // 读取操作

		}
	}()

	// 等待两个 goroutine 结束
	wg.Wait()
	fmt.Println("Done (you will likely not see this message)")
}
