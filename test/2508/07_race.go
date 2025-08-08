package main

import "fmt"

// go run --race test/2508/07_race.go
func main() {
	var x int = 100
	go func() { fmt.Println(x) }() // 读
	go func() { fmt.Println(x) }() // 读
}
