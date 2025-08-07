package main

import "fmt"

func main() {
	sum, avg := SumAndAverage(1, 2, 3, 4, 5)
	fmt.Println(sum, avg)
}

func SumAndAverage(args ...int) (sum int, avg float64) {
	for _, a := range args {
		sum += a
		avg += float64(a)
	}
	avg = avg / float64(len(args))
	return
}
