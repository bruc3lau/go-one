package main

import "fmt"

func main() {
	s := make([]int, 3, 8)
	fmt.Println(s, len(s), cap(s))
	s = append(s, 1, 2, 3, 4, 5, 6)
	fmt.Println(s, len(s), cap(s))

	//delete
	s2 := s[1:]
	fmt.Println(s2, len(s2), cap(s2))

	//delete
	s3 := s2[:len(s2)-1]
	fmt.Println(s3, len(s3), cap(s3))

	s4 := append(s3[:2], s3[3:]...)
	fmt.Println(s4, len(s4), cap(s4))

	//var s5 []int
	s5 := make([]int, len(s4), cap(s4)*2)
	copy(s5, s4)
	fmt.Println(s5, len(s5), cap(s5))

	s5[0] = 888
	fmt.Println(s4, len(s4), cap(s4))

}
