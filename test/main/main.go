package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hello World")

	map1 := make(map[string]string)
	map1["key2"] = "value2"
	map1["key1"] = "value1"
	map1["key3"] = "value3"
	map1["key4"] = "value4"

	// 第一次遍历
	fmt.Println("第一次遍历:")
	for key, value := range map1 {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}

	// 第二次遍历
	fmt.Println("\n第二次遍历:")
	for key, value := range map1 {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}

	// 第三次遍历
	fmt.Println("\n第三次遍历:")
	for key, value := range map1 {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}

	fmt.Println("\n随机顺序遍历:")
	for key, value := range map1 {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}

	// 获取所有的 key
	keys := make([]string, 0, len(map1))
	for k := range map1 {
		keys = append(keys, k)
	}

	// 对 key 进行排序
	sort.Strings(keys)

	// 按照排序后的 key 遍历 map
	fmt.Println("\n有序遍历:")
	for _, k := range keys {
		fmt.Printf("Key: %v, Value: %v\n", k, map1[k])
	}
}
