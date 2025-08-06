package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println()

	func(data int) {
		fmt.Println("data:", data)
	}(999)

	f := func(i int) int {
		return i * 2
	}(100)
	fmt.Println(f)

	f1 := func(i int) float64 {
		return math.Sqrt(float64(i))
	}
	fmt.Println(f1(88))

}
