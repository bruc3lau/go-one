package main

import "fmt"

func main() {
	var ptr *int
	fmt.Printf("%T %v\n", ptr, ptr)

	if ptr == nil {
		fmt.Println("ptr is nil")
	}
}
