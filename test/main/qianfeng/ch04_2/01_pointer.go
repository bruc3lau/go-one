package main

import "fmt"

func main() {
	a := 10
	fmt.Printf("%x\n", &a)
	b := []int{1, 2, 3, 4}
	fmt.Printf("%x\n", &b)

	//&a++

	var c *int
	fmt.Printf("%x", c)
}
