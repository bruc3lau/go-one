package main

import "fmt"

func main() {
	//arr := [4]int{2, 3, 4, 5}
	//fmt.Println(arr)
	//arr2 := arr
	//fmt.Println(arr2)
	//arr[1] = 666
	//fmt.Println(arr)
	//fmt.Println(arr2)

	arr := []int{2, 3, 4, 5}
	fmt.Println(arr)
	arr2 := arr
	fmt.Println(arr2)
	arr[1] = 666
	fmt.Println(arr)
	fmt.Println(arr2)

}
