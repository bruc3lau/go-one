package main

import (
	"fmt"
	"runtime"
	"sort"
	"unsafe"
)

func printMemStats(stage string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("\n=== %s ===\n", stage)
	fmt.Printf("Alloc = %v KB\n", m.Alloc/1024)
	fmt.Printf("TotalAlloc = %v KB\n", m.TotalAlloc/1024)
	fmt.Printf("HeapAlloc = %v KB\n", m.HeapAlloc/1024)
	fmt.Printf("HeapObjects = %v\n", m.HeapObjects)
}

// compareMaps 比较两个 map 是否相等
func compareMaps(m1, m2 map[string]string) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v1 := range m1 {
		if v2, ok := m2[k]; !ok || v1 != v2 {
			return false
		}
	}
	return true
}

func demonstrateMapPanic() {
	// 声明一个 nil map
	var nilMap map[string]int

	// 注册 panic 处理器
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("捕获到 panic：%v\n", r)
		}
	}()

	fmt.Println("尝试向 nil map 写入...")
	// 这行会触发 panic
	nilMap["key"] = 1
	// 这行代码永远不会执行
	fmt.Println("这行不会被打印，因为上面的代码会 panic")
}

func main() {
	fmt.Println("Hello World")

	// 打印初始内存状态
	printMemStats("初始状态")

	// 创建空map
	map1 := make(map[string]string)
	printMemStats("创建空map后")

	// 添加第一个元素
	map1["key1"] = "value1"
	fmt.Printf("添加一个元素后占用字节数: %d\n", unsafe.Sizeof(map1))

	// 添加第二个元素
	map1["key2"] = "value2"
	fmt.Printf("添加两个元素后占用字节数: %d\n", unsafe.Sizeof(map1))

	// 添加第三个元素
	map1["key3"] = "value3"
	fmt.Printf("添加三个元素后占用字节数: %d\n", unsafe.Sizeof(map1))

	// 添加第四个元素
	map1["key4"] = "value4"
	fmt.Printf("添加四个元素后占用字节数: %d\n", unsafe.Sizeof(map1))

	// 注意：unsafe.Sizeof 只返回 map 头部结构的大小，不包括实际存储的键值对数据
	// 要获取更准确的内存使用情况，我们可以打印每个键值对的大小
	fmt.Println("\n各个键值对的大小：")
	for k, v := range map1 {
		keySize := unsafe.Sizeof(k)
		valueSize := unsafe.Sizeof(v)
		fmt.Printf("键 '%s' 的大小: %d 字节, 值 '%s' 的大小: %d 字节\n",
			k, keySize, v, valueSize)
	}

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

	// 添加100个元素
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("key%d", i)
		value := fmt.Sprintf("value%d", i)
		map1[key] = value

		if i%20 == 0 { // 每添加20个元素打印一次
			printMemStats(fmt.Sprintf("添加了%d个元素后", i+1))
		}
	}

	// 最终状态
	printMemStats("最终状态")

	// 打印map的大小
	fmt.Printf("\nmap中的元素个数: %d\n", len(map1))

	// 创建一个 map
	map2 := make(map[string]string)

	// 打印 map 的大小（实际上是指针的大小）
	fmt.Printf("map1 的大小: %d 字节\n", unsafe.Sizeof(map1))
	fmt.Printf("map2 的大小: %d 字节\n", unsafe.Sizeof(map2))

	// 向 map1 添加1000个元素
	for i := 0; i < 1000; i++ {
		map1[fmt.Sprintf("key%d", i)] = fmt.Sprintf("value%d", i)
	}

	// 再次打印大小，可以看到即使添加了很多元素，map变量的大小（指针大小）并没有改变
	fmt.Printf("\n添加1000个元素后:\n")
	fmt.Printf("map1 的大小: %d 字节\n", unsafe.Sizeof(map1))
	fmt.Printf("空的 map2 的大小: %d 字节\n", unsafe.Sizeof(map2))

	// 打印两个 map 的地址
	fmt.Printf("\nmap1 的指针地址: %p\n", &map1)
	fmt.Printf("map2 的指针地址: %p\n", &map2)

	// 打印元素数量对比
	fmt.Printf("\nmap1 的元素数量: %d\n", len(map1))
	fmt.Printf("map2 的元素数量: %d\n", len(map2))

	fmt.Printf("&map1: %p\n", &map1)
	fmt.Printf("map1: %p\n", map1)
	//fmt.Printf("*&map1: %p\n", *&map1)

	// 创建第二个 map 来对比
	map3 := map1
	fmt.Printf("\n复制后：\n")
	fmt.Printf("&map3  : %p (map3变量的地址，和 map1 不同)\n", &map3)
	fmt.Printf("map3   : %p (指向同一个 runtime.hmap，和 map1 相同)\n", map3)

	//b := map3 == map1
	//fmt.Printf("map3 == map1: %v (两个 map 指向同一个内存地址)\n", map3 == map1)

	// 创建一个内容相同但是不同底层数据的 map
	map4 := make(map[string]string)
	map4["key1"] = "value1"
	map4["key2"] = "value2"

	fmt.Printf("map1: %p\n", map1)
	fmt.Printf("map3: %p (和 map1 相同，因为指向同一个底层数据)\n", map3)
	fmt.Printf("map4: %p (和 map1 不同，因为是独立的 map)\n", map4)

	// 比较 map
	fmt.Printf("\n比较结果：\n")
	//fmt.Printf("map1 和 map3 是否指向同一个地址: %v\n", map1 == map3) // 这行会导致编译错误
	fmt.Printf("map1 和 map3 的内容是否相同: %v\n", compareMaps(map1, map3))
	fmt.Printf("map1 和 map4 的内容是否相同: %v\n", compareMaps(map1, map4))

	// map 只能和 nil 比较
	var nilMap map[string]string
	fmt.Printf("map1 == nil: %v\n", map1 == nil)
	fmt.Printf("nilMap == nil: %v\n", nilMap == nil)

	b := &map1 == &map3
	fmt.Printf("map1 的地址: %p\n", &map1)
	fmt.Printf("map3 的地址: %p\n", &map3)
	fmt.Printf("map1 和 map3 的地址是否相同: %v\n", b)
	//fmt.Printf("map1 和 map3 是否相同: %v\n", *&map1 == *&map3)

	// 1. 使用 make 创建
	map5 := make(map[string]int)
	map5["one"] = 1

	// 2. 使用字面量创建并初始化
	map6 := map[string]int{
		"one": 1,
		"two": 2,
	}

	// 3. 声明一个 nil map（未初始化）
	var map7 map[string]int

	// 4. 创建空map
	map8 := map[string]int{}

	// 5. 嵌套 map
	nestedMap := map[string]map[string]int{
		"outer1": {
			"inner1": 1,
			"inner2": 2,
		},
		"outer2": {
			"inner3": 3,
			"inner4": 4,
		},
	}

	// 6. 使用自定义类型作为键或值
	type Person struct {
		Name string
		Age  int
	}

	// map 的值是结构体
	peopleMap := map[string]Person{
		"person1": {"Alice", 20},
		"person2": {"Bob", 25},
	}

	// map 的键是结构体（注意：结构体作为键必须是可比较的）
	type Key struct {
		ID   int
		Type string
	}

	structKeyMap := map[Key]string{
		Key{1, "user"}:  "value1",
		Key{2, "admin"}: "value2",
	}

	// 打印结果
	fmt.Println("1. make 创建的 map:", map5)
	fmt.Println("2. 字面量创建的 map:", map6)
	fmt.Println("3. nil map:", map7, map7 == nil)
	fmt.Println("4. 空 map:", map8, map8 == nil)
	fmt.Println("5. 嵌套 map:", nestedMap)
	fmt.Println("6. 结构体值的 map:", peopleMap)
	fmt.Println("7. 结构体键的 map:", structKeyMap)

	// 演示 nil map 和空 map 的区别
	fmt.Println("\nnil map vs 空 map:")
	fmt.Printf("nil map == nil: %v\n", map7 == nil)
	fmt.Printf("空 map == nil: %v\n", map8 == nil)

	// nil map 不能添加元素（会 panic）
	// map3["test"] = 1  // 这会导致 panic

	// 空 map 可以添加元素
	map8["test"] = 1
	fmt.Println("向空 map 添加元素:", map8)

	// 1. nil map（只声明，没有初始化）
	var nilMap2 map[string]int

	// 2. 空 map（已初始化，但没有元素）
	emptyMap := map[string]int{}

	// 3. 使用 make 创建的空 map
	makeMap := make(map[string]int)

	// 打印它们的地址和状态
	fmt.Printf("1. nilMap: %p, 是否为nil: %v\n", nilMap2, nilMap2 == nil)
	fmt.Printf("2. emptyMap: %p, 是否为nil: %v\n", emptyMap, emptyMap == nil)
	fmt.Printf("3. makeMap: %p, 是否为nil: %v\n", makeMap, makeMap == nil)

	// 尝试获取元素（都是安全的）
	v1, ok1 := nilMap2["key"]
	v2, ok2 := emptyMap["key"]
	fmt.Printf("\n从 nilMap 中读取：值=%v, 存在=%v\n", v1, ok1)
	fmt.Printf("从 emptyMap 中读取：值=%v, 存在=%v\n", v2, ok2)

	// 尝试添加元素
	fmt.Println("\n尝试添加元素：")

	////1. 向 nil map 添加元素会 panic
	//defer func() {
	//	if r := recover(); r != nil {
	//		fmt.Println("向 nil map 添加元素导致了 panic")
	//	}
	//}()

	// 这行是安全的
	emptyMap["key"] = 1
	fmt.Println("向 empty map 添加元素成功")

	// 这行会导致 panic
	//nilMap2["key"] = 1

	fmt.Println("\n演示 panic 处理：")
	demonstrateMapPanic()
	fmt.Println("程序正常继续执行") // 这行会被打印，因为 panic 被处理了
}
