package main

import "fmt"

func main() {
	c := func() func() int {
		i := 10
		return func() int {
			i++
			return i
		}
	}()

	fmt.Println(c)
	fmt.Println(c())
}
