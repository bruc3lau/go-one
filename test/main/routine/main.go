package main

import (
	"fmt"
	"runtime"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("Number: %d\n", i)
	}
}

func printLetters() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("Letter: %c\n", i)
	}
}

func main() {
	// 获取当前的CPU核心数
	numCPU := runtime.NumCPU()
	fmt.Printf("本机的 CPU 核心数：%d\n", numCPU)

	// 获取当前GOMAXPROCS的设置
	fmt.Printf("当前GOMAXPROCS的值：%d\n", runtime.GOMAXPROCS(0))

	// 设置使用的线程数为2
	runtime.GOMAXPROCS(2)
	fmt.Printf("设置后GOMAXPROCS的值：%d\n", runtime.GOMAXPROCS(0))

	// 使用协程执行任务
	go printNumbers()
	go printLetters()

	// 等待足够的时间让协程执行完
	time.Sleep(1 * time.Second)
	fmt.Println("Main function finished")
}
