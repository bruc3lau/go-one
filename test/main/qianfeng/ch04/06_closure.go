package main

import "fmt"

func main() {
	fmt.Println()

	result := 0
	for i := 0; i <= 10; i++ {
		//fmt.Printf("%d \t", i)
		//r := add(i)
		//fmt.Println(r)

		func(n int) {
			result += n
		}(i)

		fmt.Println(result)
	}
}

func add(n int) int {
	sum := 0
	sum += n
	return sum
}
