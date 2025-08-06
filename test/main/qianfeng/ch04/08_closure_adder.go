package main

import "fmt"

func main() {
	fmt.Println()

	//res := counter()
	//fmt.Println(res)
	//fmt.Printf("%T\n", res)
	//fmt.Printf("%T\n", res())
	//res()
	//res()
	//result := res()
	//fmt.Println(result)

	res := counter()
	fmt.Println(res)
	fmt.Println("res:", res())
	fmt.Println("res:", res())
	fmt.Println("res:", res())

	res2 := counter()
	fmt.Println(res)
	fmt.Println("res2:", res2())
	fmt.Println("res2:", res2())
	fmt.Println("res2:", res2())
}

func counter() func() int {
	sum := 0
	res := func() int {
		sum++
		return sum
	}
	fmt.Println("Counter inner: ", res)
	//fmt.Printf("Counter inner: %T\n", res)
	return res
}
