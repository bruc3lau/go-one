package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func main() {
	//fmt.Println()
	s1 := Student{Name: "BOB", Age: 20}
	fmt.Println(s1)

	s2 := &Student{Name: "ADAM", Age: 18}
	fmt.Println(*s2)

	fmt.Printf("%T %v\n", s1, s1)

	fmt.Println(s1.Name, (&s1).Age)

}
