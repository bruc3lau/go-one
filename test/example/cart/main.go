package main

import "fmt"

// Cart 是一个商品列表，本质是 slice
type Cart []string

// Add 方法使用【值接收者】
// 它向购物车切片中添加一个商品
func (c Cart) Add(item string) {
	// 这里的 append 会修改 c 指向的底层数组
	// 因为容量足够，所以没有发生重新分配
	c = append(c, item)
	fmt.Printf("  (在 Add 方法内部) c 的地址: %p, 内容: %v\n", &c, c)
}

// AddAndReallocate 方法同样使用【值接收者】
// 但这次我们会故意让 append 触发底层数组的重新分配
func (c Cart) AddAndReallocate(item string) Cart {
	// 当 append 发现容量不足时，会创建一个全新的、更大的底层数组，
	// 并返回一个指向新数组的新切片头。
	c = append(c, item)
	fmt.Printf("  (在 AddAndReallocate 方法内部) c 的地址: %p, 内容: %v\n", &c, c)
	return c // 必须返回这个新的切片，否则修改会丢失
}

func main() {
	// 创建一个容量为 5 的购物车
	myCart := make(Cart, 2, 5)
	myCart[0] = "apple"
	myCart[1] = "banana"

	fmt.Printf("原始 myCart 的地址: %p, 内容: %v\n", &myCart, myCart)
	fmt.Println("--- 调用 Add 方法 (值接收者，容量足够) ---")

	myCart.Add("orange") // 调用值接收者方法

	// 结果：修改成功！因为底层数组是共享的，且没有发生重分配。
	fmt.Printf("调用后 myCart 的内容: %v  <-- 修改成功!\n\n", myCart)

	fmt.Println("--- 调用 AddAndReallocate 方法 (值接收者，容量不足) ---")
	myCart = myCart.AddAndReallocate("grape") // 必须接收返回值
	myCart = myCart.AddAndReallocate("mango")
	myCart = myCart.AddAndReallocate("pear") // 这次调用会超出容量5，触发重分配

	// 结果：如果你不接收返回值，最后一个 "pear" 的添加就会丢失。
	// 因为方法内部的 c 得到了一个指向新数组的新切片，而外部的 myCart 仍然指向旧的。
	fmt.Printf("调用后 myCart 的内容: %v\n", myCart)
}
