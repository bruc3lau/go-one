package main

import (
	"fmt"
	"sync"
)

func main() {
	//chan01()
	//chan02()
	//chan03()
	//chan04()
	//chan05()
	chan06()
}

func chan01() {
	c := make(chan int)
	go func() { c <- 1 }()
	fmt.Println(<-c)
}

func chan02() {
	c := make(chan int)
	go func() {
		c <- 1
	}()
	//fmt.Println(<-c)
	select {
	case v := <-c:
		fmt.Println(v)
	default:
		fmt.Println("no value received")
	}
}
func chan03() {
	c := make(chan int, 1)
	go func() {
		c <- 1
	}()
	//fmt.Println(<-c)
	select {
	case v := <-c:
		fmt.Println(v)
	default:
		fmt.Println("no value received")
	}
}

func chan04() {
	c := make(chan int)
	c <- 1
	c <- 2
}

func chan05() {
	c := make(chan int, 1)
	select {
	case c <- 1:
		fmt.Println("sent 1")
	case c <- 2:
		fmt.Println("sent 2")
	default:
		fmt.Println("no value sent")
	}
}
func chan06() {
	c := make(chan int, 1)
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		//fmt.Println("enter for loop")
		go func() {
			select {
			case c <- i:
				fmt.Println("sent value", i)
			case c <- i:
				fmt.Println("sent value ", i)
			default:
				fmt.Println("no value sent")
			}
			wg.Done()
		}()
	}
	for i := 0; i < 1000; i++ {
		wg.Add(1)

		//fmt.Println("enter for loop")
		go func() {
			select {
			case v := <-c:
				fmt.Println("get value ", v)
			default:
				fmt.Println("no value get")
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
