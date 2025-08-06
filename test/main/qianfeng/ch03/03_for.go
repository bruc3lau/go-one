package main

import "fmt"

func main() {
	for i := 10; i < 20; i++ {
		fmt.Println(i)
	}

	i := 0
	for ; i < 10; i++ {
		fmt.Printf("%d\n", i)
	}

	i = 0
	for {
		i++
		if i > 10 {
			break
		}
		fmt.Printf("%d", i)
	}
}
