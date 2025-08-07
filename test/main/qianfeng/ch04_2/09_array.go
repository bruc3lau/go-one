package main

import "fmt"

func main() {
	a := [4]string{"hello", "world", "aaa", "bbb"}
	fmt.Println(a)

	//var ptrArr [4]*string
	//fmt.Println(ptrArr)

	var ptr [4]*string
	fmt.Printf("%T %v\n", ptr, ptr)

	for i := 0; i < 4; i++ {
		ptr[i] = &a[i]
	}
	fmt.Println("%T %v\n", ptr)
	fmt.Println(ptr[0])

	for i := 0; i < 4; i++ {
		fmt.Println(*ptr[i])
	}
}
