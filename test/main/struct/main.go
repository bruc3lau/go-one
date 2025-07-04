package main

import (
	"fmt"
	"unsafe"
)

type Car struct {
	Id    string  `json:"id"`
	Brand string  `json:"brand"`
	Price float32 `json:"price"`
	Color string  `json:"color"`
}

func main() {
	car1 := &Car{Id: "京A88888"}
	fmt.Printf("car1: %#v\n", car1)

	model1 := &struct {
		Model string
	}{Model: "tesla"}
	fmt.Printf("model1: %#v\n", model1)

	fmt.Printf("car1: %v\n", unsafe.Sizeof(car1))
	fmt.Printf("car1.Id: %v\n", unsafe.Offsetof(car1.Id))
	fmt.Printf("car1.Brand: %v\n", unsafe.Offsetof(car1.Brand))
	fmt.Printf("car1.Price: %v\n", unsafe.Offsetof(car1.Price))

	var car2 Car
	fmt.Printf("car2: %v\n", car2)
	fmt.Printf("car2 size: %v\n", unsafe.Sizeof(car2))

	// 1. 判断结构体是否为零值
	emptycar := Car{}
	fmt.Printf("car2 是否为零值: %v\n", car2 == emptycar)

	// 2. 使用指针判断
	var car3 *Car
	fmt.Printf("car3(指针) 是否为 nil: %v\n", car3 == nil)

	// 3. 判断具体字段是否为零值
	fmt.Printf("car2.Id 是否为空: %v\n", car2.Id == "")
	fmt.Printf("car2.Price 是否为 0: %v\n", car2.Price == 0)

	changeCar(car2)
	fmt.Printf("car2: %v\n", car2)

	changeCarPointer(&car2)
	fmt.Printf("car2: %v\n", car2)
}

func changeCar(car Car) {
	fmt.Printf("修改前 car: %v\n", car)
	car.Id = "京B99999"
	fmt.Printf("修改后 car: %v\n", car)
}

func changeCarPointer(car *Car) {
	fmt.Printf("修改前 car: %v\n", car)
	car.Id = "京B99999"
	fmt.Printf("修改后 car: %v\n", car)
}
