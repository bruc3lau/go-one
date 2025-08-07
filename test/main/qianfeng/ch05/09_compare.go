package main

import "fmt"

func main() {
	arr := []int{2, 4, 3, 46, 23, 12, 1}

	for i := 0; i < len(arr); i++ {
		//fmt.Println(arr[i])
		for j := 0; j < i; j++ {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	fmt.Println(arr)
}
