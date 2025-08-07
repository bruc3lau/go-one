package main

import "fmt"

func main() {
	a := 10
	b := &a

	fmt.Printf("%T %v\n", b, b)
	fmt.Println(*b)

	*b++
	fmt.Println(a)
}
