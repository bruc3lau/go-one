package main

import "fmt"

func adder() func(n int) int {
	sum := 0
	res := func(n int) int {
		sum += n
		return sum
	}
	return res
}

func main() {
	res := adder()
	for i := 0; i < 10; i++ {
		r := res(i)
		fmt.Println(r)
	}
}
