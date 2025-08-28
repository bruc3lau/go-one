package main

import (
	"fmt"
	"iter"
	"slices"
)

// 自定义迭代器：生成斐波那契数列
func fibonacci(n int) iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 0, 1
		for i := 0; i < n; i++ {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

// 自定义迭代器：键值对迭代
func enumerate[T any](slice []T) iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i, v := range slice {
			if !yield(i, v) {
				return
			}
		}
	}
}

// 树结构的迭代器
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) InOrder() iter.Seq[int] {
	return func(yield func(int) bool) {
		t.inOrderHelper(yield)
	}
}

func (t *TreeNode) inOrderHelper(yield func(int) bool) bool {
	if t == nil {
		return true
	}
	if !t.Left.inOrderHelper(yield) {
		return false
	}
	if !yield(t.Value) {
		return false
	}
	return t.Right.inOrderHelper(yield)
}

func main() {
	fmt.Println("=== Go 1.23 迭代器示例 ===")

	// 1. 使用自定义斐波那契迭代器
	fmt.Println("\n1. 斐波那契数列迭代:")
	for num := range fibonacci(10) {
		fmt.Printf("%d ", num)
	}
	fmt.Println()

	// 2. 使用键值对迭代器
	fmt.Println("\n2. 枚举迭代器:")
	fruits := []string{"apple", "banana", "cherry"}
	for i, fruit := range enumerate(fruits) {
		fmt.Printf("%d: %s\n", i, fruit)
	}

	// 3. 标准库的新迭代器函数
	fmt.Println("\n3. 标准库迭代器:")
	numbers := []int{1, 2, 3, 4, 5}

	// slices.All 返回索引和值的迭代器
	fmt.Println("slices.All:")
	for i, v := range slices.All(numbers) {
		fmt.Printf("  [%d]=%d\n", i, v)
	}

	// slices.Values 只返回值的迭代器
	fmt.Println("slices.Values:")
	for v := range slices.Values(numbers) {
		fmt.Printf("  %d", v)
	}
	fmt.Println()

	// 4. 树的中序遍历
	fmt.Println("\n4. 树的中序遍历:")
	tree := &TreeNode{
		Value: 4,
		Left: &TreeNode{
			Value: 2,
			Left:  &TreeNode{Value: 1},
			Right: &TreeNode{Value: 3},
		},
		Right: &TreeNode{
			Value: 6,
			Left:  &TreeNode{Value: 5},
			Right: &TreeNode{Value: 7},
		},
	}

	for value := range tree.InOrder() {
		fmt.Printf("%d ", value)
	}
	fmt.Println()

	// 5. 迭代器组合和链式操作
	fmt.Println("\n5. 迭代器组合:")
	evenFib := func() iter.Seq[int] {
		return func(yield func(int) bool) {
			for num := range fibonacci(20) {
				if num%2 == 0 {
					if !yield(num) {
						return
					}
				}
			}
		}
	}

	fmt.Println("偶数斐波那契数:")
	for num := range evenFib() {
		fmt.Printf("%d ", num)
	}
	fmt.Println()
}
