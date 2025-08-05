package main

import "fmt"

type point struct {
	x, y int
}

func main() {
	str := "hello"
	fmt.Printf("%T %v\n", str, str)

	p := point{1, 2}
	fmt.Printf("%T %v\n", p, p)

	var r rune = '一'
	fmt.Printf("%T %v\n", r, r)

	fmt.Printf("%T %v\n", true, true)

	fmt.Printf("%T %08d\n", 1, 1)

	fmt.Printf("%T %b\n", 10, 10)
	s1 := fmt.Sprintf("%b", 10)
	fmt.Println(s1)
	fmt.Printf("%T %x\n", 12, 12)
	fmt.Printf("%T %X\n", 12, 12)
	fmt.Printf("%T %U\n", 'a', 'a')
	fmt.Printf("%U\n", '一')

	fmt.Printf("%.2f\n", 12.34567)

	fmt.Printf("%s\n", "hello world")
	fmt.Printf("%q\n", "hello world")

	arr := []byte{'a', 98, 99, 65}
	fmt.Printf("%T %s\n", arr, arr)
	arr2 := []rune{'a', 98, '一', 65}

	//%s 只能输出字符串或，字节数组[]byte ，汉字不在字节数组
	fmt.Printf("%T %s\n", arr2, arr2)
	fmt.Printf("%T %x\n", arr, arr)
}
