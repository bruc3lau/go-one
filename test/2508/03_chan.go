package main

import "fmt"

func main() {
	//chan1()
	//chan2()
	//chan3()
	//chan4()
	chan5()
}

func chan1() {
	c := make(chan int, 1) // 创建一个带缓冲区的通道
	c <- 1                 // 向通道发送数据
	c <- 2                 // 向通道发送数据
	data := <-c            // 从通道接收数据
	fmt.Println(data)      // 打印接收到的数据
}

func chan2() {
	c := make(chan int) // 创建一个无缓冲区的通道
	go func() {
		c <- 1 // 在 goroutine 中向通道发送数据
	}()
	data := <-c       // 从通道接收数据
	fmt.Println(data) // 打印接收到的数据
}

func chan3() {
	c := make(chan int) // 创建一个无缓冲区的通道
	go func() {
		for i := 0; i < 100; i++ {
			c <- i // 在 goroutine 中向通道发送数据
		}
	}()
	data := <-c       // 从通道接收数据
	fmt.Println(data) // 打印接收到的数据
}

func chan4() {
	c := make(chan int, 10) // 创建一个带缓冲区的通道
	for i := 0; i < 10; i++ {
		c <- i // 向通道发送数据
	}
	close(c)              // 关闭通道
	for data := range c { // 从通道接收数据，直到通道被关闭
		fmt.Println(data) // 打印接收到的数据
	}
}

// fatal error: all goroutines are asleep - deadlock!
func chan5() {
	var c chan int
	fmt.Println(c)
	c <- 1
}
