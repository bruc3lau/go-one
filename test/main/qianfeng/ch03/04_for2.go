package main

import "fmt"

func main() {
	//sum := 0
	//for i := 1; i <= 100; i++ {
	//	sum += i
	//}
	//fmt.Println(sum)
	//

	//sum := 0
	//for i := 1; i <= 100; i++ {
	//	if i%3 == 0 {
	//		sum += i
	//	}
	//}
	//fmt.Println(sum)

	//n := 0
	//for l := 32.0; l > 4; l = l - 1.5 {
	//	fmt.Println(l)
	//	n++
	//}
	//fmt.Println(n)

	//str := "hello world 一你好" //汉字三字节
	//for i, v := range str {
	//	//fmt.Println(i, v)
	//	fmt.Printf("%d %T %v %c", i, v, v, v)
	//	fmt.Print(" " + string(v) + "\n")
	//}
	//
	//fmt.Println(string('一' + 1))

	arr := []int{1, 2, 3, 4}
	for i, v := range arr {
		fmt.Println(i, v)
	}
}
