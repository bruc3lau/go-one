package main

import "fmt"

var a = 666
var b = 7

func main() {
	//a := 1
	//fmt.Println(a, b)
	//
	//b, c := 2, 3
	//
	//fmt.Println(b, c)

	a, b, c := 10, 20, 30

	fmt.Println(a, b, c)

	c, d := sum(a, b)
	fmt.Println(c, d)
}

//	func sum(a, b int) int {
//		a++
//		b += 2
//		c := a + b
//		return c
//	}

//func sum(a, b int) (c int) {
//	fmt.Println("c:", c)
//	a++
//	b += 2
//	c = a + b
//	return
//}

func sum(a, b int) (c, d int) {
	fmt.Println("c:", c)
	a++
	b += 2
	c = a + b

	d = a * b
	return
}
