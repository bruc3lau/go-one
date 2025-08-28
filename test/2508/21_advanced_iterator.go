package main

import (
	"fmt"
	"iter"
	"maps"
	"slices"
)

// 高级迭代器示例

// 1. 带过滤的迭代器
func Filter[T any](seq iter.Seq[T], predicate func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range seq {
			if predicate(v) {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// 2. 映射变换迭代器
func Map[T, U any](seq iter.Seq[T], mapper func(T) U) iter.Seq[U] {
	return func(yield func(U) bool) {
		for v := range seq {
			if !yield(mapper(v)) {
				return
			}
		}
	}
}

// 3. 取前N个元素的迭代器
func Take[T any](seq iter.Seq[T], n int) iter.Seq[T] {
	return func(yield func(T) bool) {
		count := 0
		for v := range seq {
			if count >= n {
				return
			}
			if !yield(v) {
				return
			}
			count++
		}
	}
}

// 4. 无限序列迭代器
func Infinite(start int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := start; ; i++ {
			if !yield(i) {
				return
			}
		}
	}
}

// 5. 自定义集合类型
type Set[T comparable] map[T]struct{}

func NewSet[T comparable](items ...T) Set[T] {
	s := make(Set[T])
	for _, item := range items {
		s[item] = struct{}{}
	}
	return s
}

// 为 Set 实现迭代器
func (s Set[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for item := range s {
			if !yield(item) {
				return
			}
		}
	}
}

// 6. 并发安全的迭代器（示例）
type SafeSlice[T any] struct {
	items []T
}

func (s *SafeSlice[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		// 创建副本以避免并发问题
		items := make([]T, len(s.items))
		copy(items, s.items)

		for _, item := range items {
			if !yield(item) {
				return
			}
		}
	}
}

func main() {
	fmt.Println("=== Go 1.23 高级迭代器示例 ===")

	// 1. 链式操作：过滤 + 映射 + 限制
	fmt.Println("\n1. 链式操作示例:")
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// 过滤偶数 -> 平方 -> 取前3个
	result := Take(
		Map(
			Filter(slices.Values(numbers), func(x int) bool { return x%2 == 0 }),
			func(x int) int { return x * x },
		),
		3,
	)

	fmt.Print("偶数的平方(前3个): ")
	for v := range result {
		fmt.Printf("%d ", v)
	}
	fmt.Println()

	// 2. 无限序列的使用
	fmt.Println("\n2. 无限序列示例:")
	fmt.Print("从100开始的前5个数: ")
	for num := range Take(Infinite(100), 5) {
		fmt.Printf("%d ", num)
	}
	fmt.Println()

	// 3. 自定义集合迭代
	fmt.Println("\n3. 自定义Set迭代:")
	set := NewSet("apple", "banana", "cherry", "apple") // 重复的会被去除
	fmt.Print("Set内容: ")
	for item := range set.All() {
		fmt.Printf("%s ", item)
	}
	fmt.Println()

	// 4. 标准库的新迭代器函数
	fmt.Println("\n4. 标准库新功能:")

	// maps 包的新迭代器
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println("maps.All():")
	for k, v := range maps.All(m) {
		fmt.Printf("  %s: %d\n", k, v)
	}

	fmt.Println("maps.Keys():")
	for k := range maps.Keys(m) {
		fmt.Printf("  %s", k)
	}
	fmt.Println()

	fmt.Println("maps.Values():")
	for v := range maps.Values(m) {
		fmt.Printf("  %d", v)
	}
	fmt.Println()

	// 5. 迭代器的提前退出
	fmt.Println("\n5. 迭代器提前退出:")
	fmt.Print("查找第一个大于5的数: ")
	for num := range slices.Values([]int{1, 3, 4, 6, 8, 9}) {
		if num > 5 {
			fmt.Printf("找到: %d\n", num)
			break // 可以正常使用 break
		}
	}

	// 6. 错误处理模式
	fmt.Println("\n6. 带错误处理的迭代器模式:")
	processWithError := func() iter.Seq2[int, error] {
		return func(yield func(int, error) bool) {
			for i := 0; i < 5; i++ {
				var err error
				if i == 3 {
					err = fmt.Errorf("模拟错误在位置 %d", i)
				}
				if !yield(i, err) {
					return
				}
			}
		}
	}

	for value, err := range processWithError() {
		if err != nil {
			fmt.Printf("错误: %v\n", err)
			break
		}
		fmt.Printf("处理值: %d\n", value)
	}
}
