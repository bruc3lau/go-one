package main

import (
	"fmt"
)

func main() {
	fmt.Println(666)

	//a1 := new(chan int)
	//*a1 = make(chan int)
	//go func() {
	//	*a1 <- 1
	//}()
	//go func() {
	//	//get <- *a1
	//	fmt.Println("get ", <-*a1)
	//}()
	//<-*a1

	//a2 := make(chan int)
	//go func() {
	//	a2 <- 1
	//}()
	//<-a2

	//a3 := make(chan int, 1)
	//go func() {
	//	time.Sleep(2 * time.Second)
	//	a3 <- 1
	//	fmt.Println("发送成功")
	//}()
	//get3 := <-a3
	//fmt.Println("get:", get3)

	//a4 := make(chan int, 1)
	//a4 <- 1
	//fmt.Println("a4", <-a4)

	a5 := make(chan bool)
	go func() {
		fmt.Println("done")
		a5 <- true
	}()
	<-a5
}
