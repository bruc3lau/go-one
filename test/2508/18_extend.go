package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Gender string
}

func (p Person) live() {
	fmt.Println("person live", p)
}

type Man struct {
	Person
	Weight float64
}

func (p Man) smoke() {
	fmt.Println("man smoke", p)
}

type Woman struct {
}

func (p Woman) dress() {
	fmt.Println("woman dress", p)
}

func main() {
	man := Man{
		Person: Person{
			"Adam",
			18,
			"male",
		}}
	man.live()
	man.smoke()
	fmt.Println(man.Gender)
	fmt.Println(man.Weight)
}
