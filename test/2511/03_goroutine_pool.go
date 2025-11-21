package main

import (
	"fmt"
	"log"
)

type Request struct {
}

func main() {
	fmt.Println(666)

}

// 创建一个缓冲为 100 的 Channel 作为信号量
var limit = make(chan struct{}, 100)

func ProcessRequest(r Request) {
	// 1. 申请令牌（如果满了会阻塞，达到限流效果）
	limit <- struct{}{}

	go func() {
		defer func() {
			// 3. 释放令牌
			<-limit
			// 处理 panic，防止 crash
			if err := recover(); err != nil {
				log.Println("recover", err)
			}
		}()

		// 2. 执行业务逻辑
		doBusinessLogic(r)
	}()
}

func doBusinessLogic(r Request) {
	//do something
}
