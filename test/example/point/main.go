package main

import (
	"fmt"
	"reflect"
)

// Point 是一个可比较的结构体
type Point struct {
	X int
	Y int
}

// Path 包含一个切片，是不可比较的结构体
type Path struct {
	Name   string
	Points []Point
}

func main() {
	fmt.Println("--- 1. 比较可比较的结构体 (使用 ==) ---")
	p1 := Point{X: 10, Y: 20}
	p2 := Point{X: 10, Y: 20}

	// 错误的方式：比较地址
	fmt.Printf("p1 的地址: %p, p2 的地址: %p\n", &p1, &p2)
	fmt.Printf("比较地址的结果: %v\n", &p1 == &p2) // false

	// 正确的方式：比较内容
	fmt.Printf("比较内容的结果 (p1 == p2): %v\n\n", p1 == p2) // true

	fmt.Println("--- 2. 比较不可比较的结构体 (使用 reflect.DeepEqual) ---")
	path1 := Path{Name: "A", Points: []Point{{1, 1}, {2, 2}}}
	path2 := Path{Name: "A", Points: []Point{{1, 1}, {2, 2}}}

	// 直接使用 '==' 会导致编译错误:
	// invalid operation: path1 == path2 (struct containing []Point is not comparable)

	// 正确的方式：深度比较
	areEqual := reflect.DeepEqual(path1, path2)
	fmt.Printf("深度比较内容的结果 (reflect.DeepEqual): %v\n", areEqual) // true
}
