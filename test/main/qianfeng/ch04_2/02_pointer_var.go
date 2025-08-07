package main

import "fmt"

func main() {

	//var c *int
	//fmt.Printf("%x", c)

	a := 100
	var ip *int

	ip = &a
	fmt.Println(ip)

	fmt.Printf("%T %v\n", a, a)
	fmt.Printf("%T %v\n", *ip, *ip)
	fmt.Printf("%T %v\n", &a, &a)
	fmt.Printf("%T %v\n", ip, ip)
	fmt.Printf("%T %v\n", *&a, *&a)

	fmt.Println(ip, &ip, *ip, *(&ip), &(*ip))
}
