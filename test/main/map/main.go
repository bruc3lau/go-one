package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 1. 使用 var 声明 - nil map
	var nilMap map[string]int
	fmt.Printf("1. nilMap 是否为 nil: %v\n", nilMap == nil)
	fmt.Printf("   nilMap 的值: %v\n", nilMap)
	// 注意：向 nil map 写入会导致 panic
	// nilMap["test"] = 1  // 会导致 panic

	// 2. 使用 make 初始化
	makeMap := make(map[string]int)
	fmt.Printf("\n2. makeMap 是否为 nil: %v\n", makeMap == nil)
	makeMap["test"] = 1
	fmt.Printf("   makeMap 的值: %v\n", makeMap)

	// 3. 使用字面量初始化
	literalMap := map[string]int{
		"one": 1,
		"two": 2,
	}
	fmt.Printf("\n3. literalMap 是否为 nil: %v\n", literalMap == nil)
	fmt.Printf("   literalMap 的值: %v\n", literalMap)

	// 4. 演示 nil map 的读取操作
	value, exists := nilMap["any"]
	fmt.Printf("\n4. 从 nil map 读取：\n")
	fmt.Printf("   值: %v, 是否存在: %v\n", value, exists)
	// 注意：读取 nil map 是安全的，只是总是返回零值

	// 5. make 时指定容量
	sizedMap := make(map[string]int, 10) // 预分配10个元素的空间
	fmt.Printf("\n5. 预分配空间的 map 是否为 nil: %v\n", sizedMap == nil)
	fmt.Printf("   初始长度: %d\n", len(sizedMap))

	// 6. 演示普通 map 的并发问题
	//unsafeMap := make(map[int]int)
	//
	// 这样的并发操作是不安全的
	//go func() {
	//	for i := 0; i < 1000; i++ {
	//		unsafeMap[i] = i // 可能导致 panic: concurrent map writes
	//	}
	//}()
	//
	//go func() {
	//	for i := 0; i < 1000; i++ {
	//		_ = unsafeMap[i] // 并发读写可能导致 panic
	//	}
	//}()

	// 7. 使用互斥锁保护 map
	var mutex sync.RWMutex
	safeMap := make(map[int]int)

	// 写入 goroutine
	go func() {
		for i := 0; i < 1000; i++ {
			mutex.Lock()
			safeMap[i] = i
			mutex.Unlock()
		}
	}()

	// 读取 goroutine
	go func() {
		for i := 0; i < 1000; i++ {
			mutex.RLock()
			_ = safeMap[i]
			mutex.RUnlock()
		}
	}()

	// 8. 使用 sync.Map（Go 提供的并发安全的 map）
	var syncMap sync.Map

	// 写入 goroutine
	go func() {
		for i := 0; i < 1000; i++ {
			syncMap.Store(i, i)
		}
	}()

	// 读取 goroutine
	go func() {
		for i := 0; i < 1000; i++ {
			if val, ok := syncMap.Load(i); ok {
				_ = val
			}
		}
	}()

	// 等待一段时间让 goroutines 执行
	time.Sleep(time.Second)

	// 9. 展示 sync.Map 的正确使用方式
	fmt.Println("\n演示 sync.Map 的使用：")
	var concurrentMap sync.Map

	// 存储值
	concurrentMap.Store("key1", 100)
	concurrentMap.Store("key2", 200)

	// 读取值
	if value, ok := concurrentMap.Load("key1"); ok {
		fmt.Printf("key1 的值: %v\n", value)
	}

	// 存储并获取旧值
	if actual, loaded := concurrentMap.LoadOrStore("key1", 150); loaded {
		fmt.Printf("key1 已存在，值为: %v\n", actual)
	}

	// 删除值
	concurrentMap.Delete("key2")

	// 遍历所有键值对
	concurrentMap.Range(func(key, value interface{}) bool {
		fmt.Printf("键: %v, 值: %v\n", key, value)
		return true // 返回 false 将停止遍历
	})
}
