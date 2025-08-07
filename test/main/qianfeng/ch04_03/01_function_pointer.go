package main

import "fmt"

func main() {
	a := 10
	//changeStringVal(a)
	changeIntPointer(&a)

	fmt.Printf("%p,%v\n", &a, a)
}
func changeIntVal(a int) {
	fmt.Printf("changeStringVal %p %v\n", &a, a)
	a = 90
}
func changeIntPointer(a *int) {
	fmt.Printf("changeStringPointer %p, %v\n", a, *a)
	*a = 88
}
