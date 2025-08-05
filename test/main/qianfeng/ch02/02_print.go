package main

import "fmt"

func main() {
	a := 100

	var b byte = 100
	var c rune = 100

	fmt.Printf("%T %v\n", a, a)
	fmt.Printf("%T %v\n", b, b)
	fmt.Printf("%T %v\n", c, c)

	var d byte = 'a'
	var e rune = 'a'
	fmt.Printf("%T %v\n", d, d)
	fmt.Printf("%T %v\n", e, e)

	var f rune = '一'
	fmt.Printf("%T %v\n", f, f)

	var g string = `aaa
bbb
 	ccc`
	fmt.Println(g)

}
