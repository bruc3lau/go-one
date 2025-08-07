package main

import "fmt"

func main() {
	//slice1()
	//array1()
	struct1()
}

func slice1() {
	arr := []int{1, 2, 3}
	fmt.Printf("%T %v %p\n", arr, arr, &arr)
	change1(arr)
	//sliceChange2(&arr)
	fmt.Println(arr)

}
func change1(arr []int) {
	arr[1] = 250
	fmt.Printf("%T %v %p\n", arr, arr, &arr)
}

func sliceChange2(arr *[]int) {
	(*arr)[1] = 123
	fmt.Printf("%T %v %p\n", arr, arr, arr)
}

func array1() {
	arr := [4]int{1, 2, 3, 4}
	fmt.Printf("%T %v %p\n", arr, arr, &arr)
	arrayChange1(arr)
	//arrayChange2(&arr)
	fmt.Printf("%T %v %p\n", arr, arr, &arr)
}

func arrayChange1(arr [4]int) {
	arr[1] = 666
	fmt.Printf("%T %v %p\n", arr, arr, arr)
}
func arrayChange2(arr *[4]int) {
	arr[1] = 222
	fmt.Printf("%T %v %p\n", arr, arr, arr)
}

type Point struct {
	x, y float64
}

func struct1() {
	p := Point{1, 2}
	fmt.Printf("%T %v %p\n", p, p, &p)
	//structChange1(p)
	structChange2(&p)
	fmt.Printf("%T %v %p\n", p, p, &p)
}

func structChange1(p Point) {
	p.x = 3
	fmt.Printf("%T %v %p\n", p, p, &p)
}

func structChange2(p *Point) {
	p.x = 4
	fmt.Printf("%T %v %p\n", p, p, p)
}
