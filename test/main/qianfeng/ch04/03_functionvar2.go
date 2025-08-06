package main

import "fmt"

type myFunc func(int) bool

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(arr)

	arr1 := Filter(arr, isOdd)
	fmt.Println(arr1)

	arr2 := Filter(arr, isEven)
	fmt.Println(arr2)
}

func Filter(arr []int, f myFunc) []int {
	var result []int
	for _, x := range arr {
		if f(x) {
			result = append(result, x)
		}
	}
	return result
}

func isEven(x int) bool {
	if x%2 == 0 {
		return true
	}
	return false
}

func isOdd(x int) bool {
	if x%2 != 0 {
		return true
	}
	return false
}
