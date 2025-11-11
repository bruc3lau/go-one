package main

import (
	"context"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	//case1()
	case2()
}

func case1() {
	wg.Add(1)
	go func() {
		for {

		}
		wg.Done()
	}()
	wg.Wait()
}

func case2() {
	ctx, cancel := context.WithCancel(context.Background())
	fmt.Println(ctx)
	if false {
		cancel()
	}

	select {
	case <-ctx.Done():
		fmt.Println("done")
	default:
		fmt.Println("default")
	}
}
