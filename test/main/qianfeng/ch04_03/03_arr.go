package main

import "fmt"

func main() {
	arr := [4]int{1, 2, 3, 4}
	fmt.Println(arr)
	//change(arr)
	change2(&arr)

	fmt.Println(arr)

}

func change(arr [4]int) {
	arr[0] = 88
	fmt.Println(arr)
}

func change2(arr *[4]int) {
	(*arr)[0] = 99
	fmt.Println(arr)
}
