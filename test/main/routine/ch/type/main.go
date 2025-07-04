package main

import (
	"fmt"
	"time"
)

// 1. 用于数据传递
func dataTransfer() {
	ch := make(chan string)
	go func() {
		// 发送数据
		ch <- "hello from goroutine"
	}()
	// 接收数据
	msg := <-ch
	fmt.Println("数据传递示例:", msg)
}

// 2. 用于同步控制
func syncControl() {
	done := make(chan bool)
	go func() {
		fmt.Println("goroutine 开始工作")
		time.Sleep(time.Second) // 模拟工作负载
		fmt.Println("goroutine 完成工作")
		done <- true
	}()
	<-done // 等待 goroutine 完成
	fmt.Println("主程序继续执行")
}

// 3. 用于控制并发数量
func workerPool() {
	const maxWorkers = 3
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	// 启动固定数量的 worker
	for i := 0; i < maxWorkers; i++ {
		go func(id int) {
			for job := range jobs {
				fmt.Printf("worker %d 处理任务 %d\n", id, job)
				results <- job * 2
			}
		}(i)
	}

	// 发送任务
	for i := 1; i <= 5; i++ {
		jobs <- i
	}
	close(jobs)

	// 收集结果
	for i := 1; i <= 5; i++ {
		<-results
	}
}

func main() {
	fmt.Println("=== 1. 数据传递示例 ===")
	dataTransfer()

	fmt.Println("\n=== 2. 同步控制示例 ===")
	syncControl()

	fmt.Println("\n=== 3. 并发控制示例 ===")
	workerPool()
}
