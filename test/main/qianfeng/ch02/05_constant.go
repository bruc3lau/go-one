package main

import "fmt"

func main() {
	const (
		a = 10
		b
		c
	)
	fmt.Println(a, b, c)

	const i1 = iota
	const i2 = iota
	fmt.Println(i1, i2)

	const (
		c1 = iota
		c2 = iota
		c3
		c4
	)
	fmt.Println(c1, c2, c3, c4)

	const (
		d1 = 1
		d2 = iota
		d3
	)
	fmt.Println(d1, d2, d3)

	const (
		L = iota
		M = iota
		N = iota
	)
	fmt.Println(L, M, N)

	const (
		a1 = ""
		a2
		a3 = 666
		a4
	)
	fmt.Println(a1, a2, a3, a4)

	const (
		e1 = "qqq"
		e2 = iota
		e3
		e4 = iota
	)

	fmt.Println(e1, e2, e3)

	const (
		f1 = 1 << iota
		f2 = 3 << iota
		f3
		f4
	)
	fmt.Println(f1, f2, f3)
	const (
		g1 = 1 + iota
		g2 = 3 + iota
		g3
		g4
	)
	const (
		h1 = '一'
		h2
		h3 = iota
		h4
	)
	fmt.Println(h1, h2, h3, h4)
}
