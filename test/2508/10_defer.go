package main

import (
	"fmt"
)

func main() {
	defer func() {
		fmt.Println("enter panic1")
		if err := recover(); err != nil {
			fmt.Println("catch panic1")
		}
	}()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("catch panic2")
		}
	}()
	defer func() { fmt.Println("first defer") }()
	defer func() { fmt.Println("second defer") }()
	fmt.Println("hello")

	//open, err := os.Open("10_defer.go")
	//if err != nil {
	//	panic(err)
	//}
	//defer func(open *os.File) {
	//	err := open.Close()
	//	if err != nil {
	//		panic(err)
	//	}
	//}(open)
	//open.read
	//file, err := os.ReadFile("test/2508/10_defer.go")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(string(file))

	panic("panic")
}
