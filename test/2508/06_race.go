package main

import (
	"fmt"
	"time"
)

// go run --race test/2508/06_race.go
func main() {
	var x int
	go func() { x = 100 }()        // 写
	go func() { fmt.Println(x) }() // 读
	time.Sleep(time.Millisecond)
	//fmt.Println(x)

}
