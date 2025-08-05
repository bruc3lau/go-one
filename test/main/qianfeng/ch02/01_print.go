package main

import "fmt"

var m = "a"

func main() {
	var a int
	//fmt.Println(a)
	var (
		b string
		c func() bool
		d float32
		e []float64
		f int32  = 100
		g [3]int //#############
		h [3]string
		i bool
		j = 200
		k = 200.32 //#############
	)
	//fmt.Println(b)
	//fmt.Println(c)
	fmt.Printf("%T %v\n", a, a)
	fmt.Printf("%T %q\n", b, b) //###########
	fmt.Printf("%T %v\n", c, c)
	fmt.Printf("%T %v\n", d, d)
	fmt.Printf("%T %v\n", e, e)
	fmt.Printf("%T %v\n", f, f)
	fmt.Printf("%T %v\n", g, g)
	fmt.Printf("%T %p\n", h, h)
	fmt.Printf("%T %v\n", i, i)
	fmt.Printf("%T %v\n", j, j)
	fmt.Printf("%T %v\n", k, k)

	l := "hello"
	//l2 := "world"
	fmt.Printf("%T %v\n", l, l)
	fmt.Printf("%T %v\n", m, m)

	//m := "b"
	//fmt.Printf("%T %v\n", m, m)

	m := "a"
	n := "b"
	//m, n := "c", "d"		x
	//m, n, q := "c", "d", "e"
	//m, n, _ := "c", "d", "e" x
	m, _, q := "c", "d", "e"
	fmt.Printf("%T %v\n", m, m)
	fmt.Printf("%T %v\n", n, n)
	fmt.Printf("%T %v\n", q, q)

	x := 1
	y := 2
	x, y = y, x // ###############
	fmt.Println(x, y)

	var abc uint8 = 255
	fmt.Println(abc)
	abc += 1
	fmt.Println(abc)

}
