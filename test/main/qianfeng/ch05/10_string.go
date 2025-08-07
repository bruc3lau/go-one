package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "hello world 你好"
	fmt.Println(str, len(str))

	for i, x := range str {
		fmt.Printf("%d %c\n", i, x)
	}
	for i, x := range []byte(str) {
		fmt.Printf("%d %c\n", i, x)
	}
	count := 0
	for i, x := range []rune(str) {
		fmt.Printf("%d %c\n", i, x)
		count++
	}
	fmt.Println(count)
	fmt.Println(utf8.RuneCountInString(str))
}
