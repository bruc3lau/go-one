package main

import "fmt"

func main() {
	result1 := jiecheng1(10)
	fmt.Println(result1)
	fmt.Println(jiecheng2(10))
}

func jiecheng1(n int) (result int) {
	result = 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return
}

func jiecheng2(n int) int {
	if n == 0 {
		return 1
	}
	return n * jiecheng2(n-1)
}
