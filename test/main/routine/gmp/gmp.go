package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func init() {
	// 设置最大处理器数量，这里设置为2个
	runtime.GOMAXPROCS(2)
}

func showGoroutineID() {
	// 这个函数会打印当前 Goroutine 的大致信息
	// 实际上Go没有提供直接获取goroutine ID的方法
	buf := make([]byte, 64)
	runtime.Stack(buf, false)
	fmt.Printf("Goroutine: %s\n", buf)
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	// 模拟一些工作负载
	for i := 0; i < 3; i++ {
		fmt.Printf("Worker %d: 正在执行任务 %d\n", id, i)
		showGoroutineID()
		// 模拟工作耗时
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	fmt.Printf("当前系统使用的处理器数量: %d\n", runtime.GOMAXPROCS(0))
	fmt.Printf("CPU核心数: %d\n", runtime.NumCPU())

	// 方式1：直接声明（零值初始化）
	var wg1 sync.WaitGroup

	// 方式2：通过new创建指针（不常用）
	wg2 := new(sync.WaitGroup)

	// 使用方式1的WaitGroup
	fmt.Println("\n使用方式1（零值）:")
	for i := 1; i <= 3; i++ {
		wg1.Add(1)
		go worker(i, &wg1)
	}
	wg1.Wait()

	// 使用方式2的WaitGroup
	fmt.Println("\n使用方式2（指针）:")
	for i := 4; i <= 5; i++ {
		wg2.Add(1)
		go worker(i, wg2)
	}
	wg2.Wait()

	// 打印运行时统计信息
	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)
	fmt.Printf("\n运行时统计:\n")
	fmt.Printf("Goroutine数量: %d\n", runtime.NumGoroutine())
	fmt.Printf("系统线程数量: %d\n", runtime.NumCPU())
}
