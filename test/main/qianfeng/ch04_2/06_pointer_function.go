package main

import "fmt"

func main() {
	a := 10

	change(&a)

	fmt.Printf("a=%d\n", a)
}

func change(n *int) {
	*n = 20
}
