package main

import "fmt"

func main() {
	fmt.Println()

	res := addSum(1, 2, 3)
	fmt.Println(res)
	fmt.Println(addSum())

	arr := []int{2, 4, 6, 8}
	fmt.Println(addSum(arr...))
}

func addSum(args ...int) (sum int) {
	fmt.Printf("%T\n", args)
	for _, x := range args {
		//fmt.Println(x)
		sum += x
	}
	return
}
