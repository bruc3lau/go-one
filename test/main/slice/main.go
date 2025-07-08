package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("slice", slice, len(slice), cap(slice))

	slice2 := slice[:1]
	fmt.Println("slice2", slice2, len(slice2), cap(slice2))

	slice3 := append(slice2, 100)
	fmt.Println("slice3", slice3, len(slice3), cap(slice3))

	fmt.Println("slice")

	fmt.Println(append([]int(nil), slice...))

	ints := []int{}
	fmt.Println("ints", ints)

	//a := 1
	//n := a.(interface{})

	// 1. int 的零值是 0，不是 nil
	var myInt int
	fmt.Printf("int的零值是: %d\n", myInt)
	// 下面这行代码会编译失败: invalid operation: myInt == nil
	// if myInt == nil { ... }

	// 2. 接口(interface)的零值是 nil
	var myInterface interface{}
	fmt.Printf("空接口的零值是 nil 吗? %v\n", myInterface == nil) // 输出: true

	// 3. 将 int 赋值给接口 (正确的转换方式)
	a := 1
	myInterface = a // 直接赋值

	// 4. 赋值后，接口就不再是 nil 了
	// 因为它现在持有一个类型(int)和一个值(1)
	fmt.Printf("赋值后的接口是 nil 吗? %v\n", myInterface == nil) // 输出: false
	fmt.Printf("接口持有的值: %v, 接口持有的类型: %T\n", myInterface, myInterface)

	i := myInterface.(int)
	fmt.Println("i", i)

	var myInterface2 interface{}
	myInterface2 = "hello world"

	// 尝试断言为 int，这会失败
	value, ok := myInterface2.(int)
	if ok {
		fmt.Println("断言成功，值是:", value)
	} else {
		fmt.Println("断言失败！接口里装的不是 int，而是", myInterface2)
	}
	// 输出: 断言失败！接口里装的不是 int，而是 hello world

}
