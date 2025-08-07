package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4}
	change2(&arr)
	fmt.Println(arr)
}

func change2(arr *[]int) {
	//*arr = append(*arr, 10)
	(*arr)[0] = 88
}
