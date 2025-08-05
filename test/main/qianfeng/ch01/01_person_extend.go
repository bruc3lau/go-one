package main

import "fmt"

type Person struct {
}

func (p *Person) run() {
	fmt.Println("Person run")
}

type Teacher struct {
	Person
}

func (t *Teacher) run2() {
	fmt.Println("Teacher run")
}

func main() {
	//fmt.Println(666)
	t := Teacher{}
	t.run()
}
