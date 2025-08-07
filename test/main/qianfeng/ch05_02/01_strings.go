package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "hello world,你好！"

	containsAny := strings.Contains(str, "ll")
	fmt.Println(containsAny)

	fmt.Println(strings.ContainsAny(str, "eo"))

	fmt.Println(strings.ContainsRune(str, '你'))
	fmt.Println(strings.Fields("a b c"))
	fmt.Println(strings.Split("a b c", " "))

	fmt.Println(strconv.Atoi("123.2"))

	fmt.Println(strconv.ParseInt("01100001", 2, 64))
	fmt.Println(strconv.ParseInt("4e00", 16, 64))

	fmt.Println(strconv.FormatInt(100, 16))
}
