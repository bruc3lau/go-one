//go:build ignore

package main

import (
	"fmt"
	_ "strconv"
	_ "sync"
	"time"

	"go-one/internal/data/bus"
)

// ParallelConsumer 是一个并行的消费者实现
type ParallelConsumer struct {
	Name string // 消费者名称
}

// HandleMessage 实现了 MessageCallback 接口
func (p *ParallelConsumer) HandleMessage(topic string, message interface{}, isParallel bool) bool {
	fmt.Printf("并行消费者 [%s] 接收到消息: 主题='%s', 消息='%v'\n", p.Name, topic, message)
	return false
}

// SerialConsumer 是一个串行的消费者实现
type SerialConsumer struct {
	Name          string // 消费者名称
	StopExecution bool   // 是否要中断后续消费者
}

// HandleMessage 实现了 MessageCallback 接口
func (s *SerialConsumer) HandleMessage(topic string, message interface{}, isParallel bool) bool {
	fmt.Printf("串行消费者 [%s] 接收到消息: 主题='%s', 消息='%v'\n", s.Name, topic, message)
	if s.StopExecution {
		fmt.Printf("串行消费者 [%s] 中断了消息链\n", s.Name)
		return true // 返回 true 来中断后续的串行消费者
	}
	return false
}

func main() {
	fmt.Println("--- 消息总线功能测试 ---")

	// 定义一个主题
	const testTopic = "TestTopic"

	// 1. 创建主题
	bus.CreateTopic(testTopic)
	fmt.Printf("主题 '%s' 创建成功\n", testTopic)

	// 2. 注册消费者
	// 注册两个并行消费者
	bus.RegisterParallelTopicConsumer(testTopic, &ParallelConsumer{Name: "P1"})
	bus.RegisterParallelTopicConsumer(testTopic, &ParallelConsumer{Name: "P2"})
	fmt.Println("注册了两个并行消费者 [P1, P2]")

	// 注册两个串行消费者
	s2 := &SerialConsumer{Name: "S2", StopExecution: true}
	bus.AddSerialTopicConsumer(testTopic, &SerialConsumer{Name: "S1", StopExecution: false})
	bus.AddSerialTopicConsumer(testTopic, s2)                                                // S2 将会中断执行
	bus.AddSerialTopicConsumer(testTopic, &SerialConsumer{Name: "S3", StopExecution: false}) // S3 将不会被执行
	fmt.Println("注册了三个串行消费者 [S1, S2, S3], S2会中断消息链")

	// 3. 生产消息
	fmt.Println("\n--- 开始生产第一条消息 ---")
	bus.Produce(testTopic, "这是第一条测试消息")

	// 等待一小段时间，让并行消费者有时间执行
	time.Sleep(100 * time.Millisecond)

	fmt.Println("\n--- 开始生产第二条消息 (S2 不再中断) ---")
	// 移除旧的 S2，添加一个新的 S2，它不会中断执行
	bus.RemoveSerialTopicConsumer(testTopic, s2) // 移除需要一个匹配的实例，这里为了演示，实际应用中需要传递同一个实例指针
	bus.AddSerialTopicConsumer(testTopic, &SerialConsumer{Name: "S2-New", StopExecution: false})
	bus.Produce(testTopic, "这是第二条测试消息")

	// 等待一小段时间
	time.Sleep(100 * time.Millisecond)

	fmt.Println("\n--- 测试完成 ---")
}
