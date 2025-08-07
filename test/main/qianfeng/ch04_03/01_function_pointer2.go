package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "hello"
	//changeStringVal(a)
	changeStringPointer(&a)

	fmt.Printf("%p,%v\n", &a, a)
}
func changeStringVal(a string) {
	fmt.Printf("changeStringVal %p %v\n", &a, a)
	a = "bbb"
}
func changeStringPointer(a *string) {
	fmt.Printf("changeStringPointer %p, %v\n", a, *a)
	//*a = "world"
	*a = strings.ToUpper(*a)
}
