package main

import "fmt"

func main() {
	fmt.Println("开始执行")

	// 第一个 defer（最后执行）
	defer func() {
		fmt.Println("defer 1 执行")
		if r := recover(); r != nil {
			fmt.Println("defer 1 捕获到 panic:", r)
		}
	}()

	// 第二个 defer（先执行）
	defer func() {
		fmt.Println("defer 2 执行")
		if r := recover(); r != nil {
			fmt.Println("defer 2 捕获到 panic:", r)
		}
	}()

	fmt.Println("即将发生 panic")
	a := 6
	b := 0
	fmt.Println(a / b) // 这里会触发 panic

	fmt.Println("这行不会执行") // 这行代码不会执行，因为上面发生了 panic
}
