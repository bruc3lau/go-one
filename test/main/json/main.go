package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonStr := "{\"name\":\"John\", \"age\":30, \"city\":\"New York\"}"
	fmt.Println(jsonStr)

	//解析json字符串为一个对象
	// 这里使用了一个空接口，表示可以接受任何类型的数据
	// 这种方式可以解析任意结构的 JSON 字符串
	// 但需要注意，这样解析后得到的对象是一个 map[string]interface{} 类型
	var interfaceObj interface{}
	_ = json.Unmarshal([]byte(jsonStr), &interfaceObj)

	fmt.Println(interfaceObj)

	s := "Jane"
	interfaceObj.(map[string]interface{})["name"] = s

	fmt.Println(interfaceObj)

	jsonStr2, _ := json.Marshal(interfaceObj)

	fmt.Println(string(jsonStr2))

}
