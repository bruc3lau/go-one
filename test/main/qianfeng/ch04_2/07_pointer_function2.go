package main

import "fmt"

func main() {
	a, b := 100, 200
	//a, b = swap(a, b)

	swap1(&a, &b)
	fmt.Println(a, b)

}

func swap(x, y int) (int, int) {
	return y, x
}

func swap1(x, y *int) {
	*x, *y = *y, *x
}
