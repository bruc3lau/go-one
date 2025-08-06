package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println()

	str1 := "hello world!"
	fmt.Println(str1)
	str2 := process(str1)
	fmt.Println(str2)

	str3 := strCase(str1, process)
	fmt.Println(str3)

	str4 := strCase2(str1, process)
	fmt.Println(str4)
}

func process(str string) string {
	result := ""
	for i, x := range str {
		//fmt.Println(i, x)
		if i%2 == 0 {
			result += strings.ToUpper(string(x))
		} else {
			result += strings.ToLower(string(x))
		}
	}
	return result
}

func strCase(str string, f func(string) string) string {
	return f(str)
}

type caseFunc func(string) string

func strCase2(str string, f caseFunc) string {
	return f(str)
}
