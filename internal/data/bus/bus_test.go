package bus

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

//  go test -v go-one/internal/data/bus
//  go test -v ./internal/data/bus/...

type SingleParallelConsumer struct {
	count int
	wg    *sync.WaitGroup
}

var lock sync.Mutex

func (c *SingleParallelConsumer) HandleMessage(topic string, message interface{}, isParallel bool) bool {
	if isParallel {
		//fmt.Println("consumer callback ", c.count)
		lock.Lock()
		c.count++
		if c.count == 10000 {
			c.wg.Done()
		}
		lock.Unlock()
	}
	return false
}

type SingleSerialConsumer struct {
}

func (c *SingleSerialConsumer) HandleMessage(topic string, message interface{}, isParallel bool) bool {
	return false
}

func TestSingleProducerConsumerPerform(t *testing.T) {
	CreateTopic("TrackRawData")
	var wg sync.WaitGroup
	wg.Add(2)
	AddSerialTopicConsumer("TrackRawData", &SingleSerialConsumer{})
	RegisterParallelTopicConsumer("TrackRawData", &SingleParallelConsumer{
		0,
		&wg,
	})
	start := time.Now().UnixNano()
	//producer
	go func() {
		for i := 0; i < 10000; i++ {
			Produce("TrackRawData", "test")
			//fmt.Println("produce")
		}
		fmt.Println("produce wg done")
		wg.Done()
	}()
	wg.Wait()
	end := time.Now().UnixNano()
	duration := end - start
	durationInMil := duration / 1e6
	fmt.Println("the time consumed for 10000 produce/consume is ", durationInMil, " millisecond")
	if durationInMil > 1000 {
		t.Error("10000 produce/consume more than 1 second")

	}
}

type CorrectTwoParallelConsumer struct {
	name  string
	count uint64 // 使用原子操作支持的类型
	wg    *sync.WaitGroup
}

func (c *CorrectTwoParallelConsumer) HandleMessage(topic string, message interface{}, isParallel bool) bool {
	// 原子地将 count 加一，这是并发安全的
	atomic.AddUint64(&c.count, 1)
	c.wg.Done() // 每处理完一条消息，就通知 WaitGroup
	return false
}

func TestSingleProducerTwoConsumer(t *testing.T) {

	topic := "SingleProducerTwoConsumer"

	CreateTopic(topic)

	n := 1000

	var wg sync.WaitGroup // 使用 WaitGroup 进行可靠同步

	consumer1 := &CorrectTwoParallelConsumer{
		name: "consumer1",
		wg:   &wg,
	}

	consumer2 := &CorrectTwoParallelConsumer{
		name: "consumer2",
		wg:   &wg,
	}

	RegisterParallelTopicConsumer(topic, consumer1)

	RegisterParallelTopicConsumer(topic, consumer2)

	// 因为有 2 个消费者，每个都会处理 n 条消息，所以总共要等待 2*n 次 Done()

	wg.Add(n * 2)

	// 生产消息

	for i := 0; i < n; i++ {

		Produce(topic, "test")

	}

	// 等待所有消息都被两个消费者处理完毕
	wg.Wait()

	// 读取最终结果，此时是完全安全的
	count1 := atomic.LoadUint64(&consumer1.count)
	count2 := atomic.LoadUint64(&consumer2.count)

	fmt.Println("consumer 1 count ", count1)
	fmt.Println("consumer 2 count ", count2)

	if count1 != uint64(n) {
		t.Errorf("consumer1 count is %d, want %d", count1, n)
	}

	if count2 != uint64(n) {
		t.Errorf("consumer2 count is %d, want %d", count2, n)
	}
}

type AddAndRemoveSerialConsumer struct {
	isRemove bool
	t        *testing.T
}

func (c *AddAndRemoveSerialConsumer) HandleMessage(topic string, message interface{}, isParallel bool) bool {
	if c.isRemove {
		c.t.Error("remove not work")
	}
	return false
}

func TestAddAndRemoveSerialConsumer(t *testing.T) {
	topic := "AddAndRemove"
	CreateTopic(topic)
	consumer := &AddAndRemoveSerialConsumer{
		isRemove: false,
		t:        t,
	}
	AddSerialTopicConsumer(topic, consumer)
	Produce(topic, "test")
	if !RemoveSerialTopicConsumer(topic, consumer) {
		t.Error("Remove not return succeeds")
	}
	consumer.isRemove = true
	Produce(topic, "test")
}

type PointerAndStructMessageConsumer struct {
	count int
	wg    *sync.WaitGroup
}

func (c *PointerAndStructMessageConsumer) HandleMessage(topic string, message interface{}, isParallel bool) bool {
	lock.Lock()
	c.count++
	if c.count == 10000 {
		c.wg.Done()
	}
	lock.Unlock()
	return false
}

func TestPointerAndStructMessage(t *testing.T) {
	topic := "PointerAndStructMessage"
	CreateTopic(topic)
	var wg sync.WaitGroup
	wg.Add(2)
	consumer := &PointerAndStructMessageConsumer{
		count: 0,
		wg:    &wg,
	}
	// 定义一个本地结构体，避免导入 models 包，解除循环依赖
	type LocalTrack struct {
		TrackId  string
		ImageUrl string
		CarNo    string
		Features [][]float32
	}

	feature := make([]float32, 2048)
	for i := 0; i < 2048; i++ {
		feature[i] = 1.0
	}
	track := LocalTrack{
		TrackId:  "cam12-04-dddddddddd",
		ImageUrl: "http://12.23.45.67/",
		CarNo:    "京A12345",
		Features: [][]float32{feature},
	}
	RegisterParallelTopicConsumer(topic, consumer)
	begin := time.Now().UnixNano()
	for i := 0; i < 10000; i++ {
		Produce(topic, track)
	}
	wg.Done()
	wg.Wait()
	end := time.Now().UnixNano()
	fmt.Println("the time for 10000 struct message is ", (end-begin)/1e6, " ms")

	wg.Add(2)
	begin = time.Now().UnixNano()
	consumer.count = 0
	for i := 0; i < 10000; i++ {
		Produce(topic, &track)
	}
	wg.Done()
	wg.Wait()
	end = time.Now().UnixNano()
	fmt.Println("the time for 10000 pointer message is ", (end-begin)/1e6, " ms")
}

type SerialReturnConsumer1 struct {
}

func (c *SerialReturnConsumer1) HandleMessage(topic string, message interface{}, isParallel bool) bool {
	return true
}

type SerialReturnConsumer2 struct {
	t *testing.T
}

func (c *SerialReturnConsumer2) HandleMessage(topic string, message interface{}, isParallel bool) bool {
	c.t.Error("HandleMessage should not be called")
	return false
}

func TestSerialReturn(t *testing.T) {
	topic := "SerialReturn"
	CreateTopic(topic)
	AddSerialTopicConsumer(topic, &SerialReturnConsumer1{})
	AddSerialTopicConsumer(topic, &SerialReturnConsumer2{
		t: t,
	})
	Produce(topic, "hello")
	time.Sleep(20 * time.Millisecond)
}
