package main

import (
	"fmt"
	"time"
)

func main() {
	var x int
	go func() { x = 1 }() // 写
	go func() { x = 2 }() //
	time.Sleep(time.Millisecond)
	fmt.Println(x)

	//go run --race test/2508/05_race.go
}
