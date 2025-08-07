package main

import "fmt"

func main() {

	//s2 := make([]int, 2, 3)
	//fmt.Println(s2, len(s2), cap(s2))
	//
	//s2 = append(s2, 1, 2, 3, 4, 5)
	//
	//fmt.Println(s2, len(s2), cap(s2))

	var s []int
	//s = append(s, 1, 2, 3)
	fmt.Println(s, len(s), cap(s))

	for i := 1; i < 100; i++ {
		s = append(s, i)
		fmt.Println(i, len(s), cap(s))
	}

	s2 := s[:2]
	fmt.Println(s2, len(s2), cap(s2))
	s3 := s2[3:9]
	fmt.Println(s3, len(s3), cap(s3))

	s3[5] = 666
	fmt.Println(s)

	//var s1 []int
	//fmt.Println(s1, s1 == nil)
	//fmt.Printf("%T %v\n", s1, s1)
	//fmt.Println(s1, len(s1), cap(s1))

}
