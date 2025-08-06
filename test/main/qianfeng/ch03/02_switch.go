package main

import "fmt"

func main() {
	var opt = 3
	switch opt {
	case 1:
		fmt.Println(444)
	case 3:
		fmt.Println(666)
	default:
		fmt.Println(555)

	}

	switch 1 {
	case 1:
		fmt.Println("OK")
	}

	switch {
	case 2 > 1, false:
		fmt.Println("GOOD")
		fallthrough
	default:
		fmt.Println("BAD")
	}
}
