package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/panjf2000/ants/v2"
)

func main() {
	// 创建一个容量为 3 的池，设置过期时间为 10 秒
	pool, err := ants.NewPool(3, ants.WithExpiryDuration(10*time.Second))
	if err != nil {
		panic(err)
	}
	defer pool.Release()

	var wg sync.WaitGroup

	// 提交 5 个任务
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		taskID := i

		// 提交任务到池中
		err := pool.Submit(func() {
			defer wg.Done()
			// 模拟耗时操作
			time.Sleep(time.Millisecond * 100)
			result := taskID * 2
			fmt.Printf("任务 %d 完成，结果: %d\n", taskID, result)
		})

		if err != nil {
			fmt.Printf("提交任务 %d 失败: %v\n", taskID, err)
			wg.Done()
		} else {
			fmt.Printf("提交任务 %d 成功\n", taskID)
		}
	}

	// 等待所有任务完成
	wg.Wait()

	fmt.Printf("运行完成，池的统计信息：\n")
	fmt.Printf("正在运行的 goroutine 数量: %d\n", pool.Running())
	fmt.Printf("空闲的 worker 数量: %d\n", pool.Free())
	fmt.Printf("池的容量: %d\n", pool.Cap())
}
