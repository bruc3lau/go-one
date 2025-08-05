package main

import "fmt"

func main() {
	fmt.Println()
	a := 10
	b := 20.2
	fmt.Println((a + int(b)) / 2)
	fmt.Println((float64(a) + b) / 2)

	_ = true
	//int(flag)

	c := int('a')
	fmt.Println(c)

	result := string(666)
	fmt.Println(result)
}
