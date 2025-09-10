package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

// go build -gcflags="-m" 02_reflect.go
func main() {
	var a any = &struct {
		Name string
		Age  int
	}{"Alice", 30}

	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(a).Elem())
	fmt.Println(reflect.TypeOf(a).Elem().Kind())
	fmt.Println(reflect.TypeOf(a).Elem().Name())

	//var b any = &Person{"Alice", 30}
	//fmt.Println(reflect.TypeOf(b))
	//fmt.Println(reflect.TypeOf(b).Elem())
	//fmt.Println(reflect.TypeOf(b).Elem().Kind())
	//fmt.Println(reflect.TypeOf(b).Elem().Name())
	//fmt.Println(reflect.TypeOf(b).Elem().PkgPath())
	//fmt.Println(reflect.TypeOf(b).Elem().FieldByName("Name"))

	// 在堆上分配 Person 对象
	b := Person{
		Name: "Alice",
		Age:  30,
	}
	modify(b)
	fmt.Println("print 39", b)

	(&b).Name = "Charlie"
	fmt.Println("print 42", b)

	//b := &Person{
	//var b any = p

	// 现在可以安全地使用反射
	t := reflect.TypeOf(b)
	if t.Kind() == reflect.Ptr {
		field, ok := t.Elem().FieldByName("Name")
		if ok {
			fmt.Printf("Field: %+v\n", field)
		}
	}

}

//func modify(person *Person) {
//	person.Name = "Bob"
//}

var global interface{} // 全局变量

func modify(v interface{}) {
	person, ok := v.(*Person)
	if ok {
		fmt.Println("modify ok")
		person.Name = "Bob"
	}
	global = v // 将v赋值给全局变量，v的生命周期超出了modify函数
}

// 1.结构体语法糖 a.x =(&a).x
// 2.逃逸分析
