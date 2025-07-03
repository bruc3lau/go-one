package main

import (
	"fmt"
	"time"
)

func main() {
	timeout := time.Duration(5 * time.Second)
	fmt.Printf("设置超时时间为: %v\n", timeout)

	now := time.Now()
	fmt.Printf("当前时间: %s\n", now)
	//format
	fmt.Printf("当前时间格式化为: %s\n", now.Format("2006-01-02 15:04:05"))
	fmt.Printf("当前时间格式化为: %s\n", now.Format("2006/01/02 15:04:05"))
	//get timestamp
	timestamp := now.Unix()
	fmt.Printf("当前时间的时间戳: %d\n", timestamp)
}
