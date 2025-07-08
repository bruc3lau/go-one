package main

import "fmt"

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var db = make(map[int]User)
var db2 = make(map[int]*User)

func main() {

	//// 添加用户
	//db[1] = User{ID: 1, Name: "Alice"}
	//db[2] = User{ID: 2, Name: "Bob"}
	//
	//// 查询用户
	//for id, user := range db {
	//	println("User ID:", id, "Name:", user.Name)
	//}
	//
	//// 更新用户
	//db[1] = User{ID: 1, Name: "Alice Smith"}
	//
	//// 删除用户
	//delete(db, 2)
	//
	//// 打印剩余用户
	//for id, user := range db {
	//	println("User ID:", id, "Name:", user.Name)
	//}

	db[1] = User{ID: 1, Name: "Alice"}

	u := db[1]
	//fmt.Printf("user: %p\n", &db[1])
	fmt.Printf("user: %p\n", &u)

	fmt.Println(db[1])

	user := db[1]
	fmt.Printf("user: %p\n", &user)

	user.Name = "Bob"
	fmt.Println(db[1]) // 输出: {1 Bob}

	db2[1] = &User{ID: 1, Name: "Alice"}
	u2 := db2[1]
	u2.Name = "Bruce"
	fmt.Println(*db2[1])

	if user, ok := db2[1]; ok {
		fmt.Println("id", user)
	} else {
		fmt.Println("not found")
	}

}
