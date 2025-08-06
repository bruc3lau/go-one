package main

import (
	"fmt"
	"math"
)

type myFunc5 func(float65 float64) float64

func main() {
	arr := []float64{2, 4, 6, 8, 9}

	arr2 := FilterArr(arr, func(n float64) float64 {
		return math.Sqrt(n)
	})

	fmt.Println(arr2)
	arr3 := FilterArr(arr, func(n float64) float64 {
		return math.Pow(n, 2)
	})

	fmt.Println(arr3)

}

func FilterArr(arr []float64, f myFunc5) []float64 {
	var result []float64
	for _, x := range arr {
		result = append(result, f(x))
	}
	return result
}

func sqrt(n float64) {
	math.Sqrt(n)
}
