package main

import "fmt"

func main() {
	i := 1
	defer func() { fmt.Println("defer1:", i) }()
	defer func() { i++ }()
	defer func(j int) { fmt.Println("j:", j) }(i + i)
	fmt.Println(i)

	i = 99
	defer func() { fmt.Println("defer2:", i) }()

}
