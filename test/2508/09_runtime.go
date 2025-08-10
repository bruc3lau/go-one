package main

import "fmt"

type small struct {
	Arr [1 << 10]int
}

func main() {
	s := small{}
	fmt.Println(&s)
}
