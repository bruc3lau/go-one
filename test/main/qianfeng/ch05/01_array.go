package main

import "fmt"

func main() {
	array1()
}

var arr []int
var arr2 [2]int
var arr3 = [...]int{1, 2, 3}

func array1() {
	fmt.Println(arr, len(arr), cap(arr))
	//fmt.Println(arr[0])

	fmt.Println(arr2, len(arr2), cap(arr2))
	fmt.Println(arr2[0])

	fmt.Println(arr3, len(arr3), cap(arr3))
	fmt.Println(arr3)

}
