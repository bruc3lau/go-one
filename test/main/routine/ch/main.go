package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// SubmitStrategy 定义任务提交策略
type SubmitStrategy int

const (
	// Block 策略：队列满时阻塞等待
	Block SubmitStrategy = iota
	// Discard 策略：队列满时丢弃任务
	Discard
	// DiscardOldest 策略：队列满时丢弃最旧的任务
	DiscardOldest
)

// Task 代表一个任务
type Task struct {
	ID      int
	Handler func() interface{}
}

// Pool 代表一个工作池
type Pool struct {
	workers   int
	taskQueue chan Task
	results   chan interface{}
	wg        sync.WaitGroup
	strategy  SubmitStrategy
}

// NewPool 创建一个新的工作池
func NewPool(workers, queueSize int, strategy SubmitStrategy) *Pool {
	return &Pool{
		workers:   workers,
		taskQueue: make(chan Task, queueSize),
		results:   make(chan interface{}, queueSize),
		strategy:  strategy,
	}
}

// Start 启动工作池
func (p *Pool) Start() {
	for i := 0; i < p.workers; i++ {
		go func(workerID int) {
			for task := range p.taskQueue {
				result := task.Handler()
				fmt.Printf("Worker %d 完成任务 %d, 结果: %v\n", workerID, task.ID, result)
				p.results <- result
				p.wg.Done()
			}
		}(i)
	}
}

var ErrQueueFull = errors.New("任务队列已满")

// Submit 提交任务到工作池，返回错误表示是否提交成功
func (p *Pool) Submit(task Task) error {
	switch p.strategy {
	case Block:
		// 阻塞策略：一直等待直到可以提交
		p.wg.Add(1)
		p.taskQueue <- task
		return nil

	case Discard:
		// 丢弃策略：队列满时直接返回错误
		select {
		case p.taskQueue <- task:
			p.wg.Add(1)
			return nil
		default:
			return ErrQueueFull
		}

	case DiscardOldest:
		// 丢弃最旧任务策略：如果队列满，先移除一个旧任务再添加新任务
		select {
		case p.taskQueue <- task:
			p.wg.Add(1)
			return nil
		default:
			// 尝试移除一个旧任务
			select {
			case <-p.taskQueue:
				// 成功移除旧任务，提交新任务
				p.taskQueue <- task
				return nil
			default:
				return ErrQueueFull
			}
		}
	}

	return nil
}

// Wait 等待所有任务完成
func (p *Pool) Wait() {
	p.wg.Wait()
	close(p.taskQueue)
	close(p.results)
}

func main() {
	fmt.Println("\n=== 3. 并发控制示例 ===")
	// 创建一个有 3 个 worker 的工作池
	pool := NewPool(3, 10, DiscardOldest)
	pool.Start()

	// 提交 5 个任务
	for i := 1; i <= 5; i++ {
		taskID := i
		err := pool.Submit(Task{
			ID: taskID,
			Handler: func() interface{} {
				// 模拟耗时操作
				time.Sleep(time.Millisecond * 100)
				return taskID * 2
			},
		})
		if err != nil {
			fmt.Printf("提交任务 %d 失败: %v\n", taskID, err)
		} else {
			fmt.Printf("提交任务 %d 成功\n", taskID)
		}
	}

	// 等待所有任务完成
	pool.Wait()

	fmt.Println("所有任务已完成")
}
