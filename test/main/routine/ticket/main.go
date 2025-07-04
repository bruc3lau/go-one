package main

import (
	"fmt"
	"sync"
)

var ticket int
var mx sync.Mutex
var wg sync.WaitGroup

func sellTicket() int {
	mx.Lock()
	defer mx.Unlock()
	ticket++
	return ticket
}

func main() {
	fmt.Println("初始票数:", ticket)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 10; i++ {
				current := sellTicket()
				fmt.Printf("售出第 %d 张票\n", current)
			}
		}()
	}

	wg.Wait()
	fmt.Println("最终售出票数:", ticket)
}
