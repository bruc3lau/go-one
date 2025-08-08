package bus

import (
	"sync"
)

// key is the topic name, string type
// value is  the queue of messages callbacks
// messages put into the parallelMQ would be consumed parallel by topic message consumer
var parallelMQ sync.Map

// key is the topic name, string type
// value is the queue of messages callbacks
// messages put into the seriallMQ would be consumed serial by topic message consumer
var serialMQ sync.Map

const (
	ChanSize             = 10000
	ProduceWaitInSeconds = 15
)

type MessageCallback interface {
	//return true if message totally consumed by the serial consumer, in other words, other serial consumers would not be called
	//return false by default
	HandleMessage(topic string, message interface{}, isParallel bool) bool
}

func init() {

}

// return true if succeeds
// return false if topic already exists
func CreateTopic(topic string) bool {
	ret := true
	if _, ok := parallelMQ.Load(topic); ok {
		ret = false
	} else {
		parallelMQ.Store(topic, make([]MessageCallback, 0))
		serialMQ.Store(topic, make([]MessageCallback, 0))
	}
	return ret
}

func HasTopic(topic string) bool {
	ret := false
	if _, ok := parallelMQ.Load(topic); ok {
		ret = true
	}
	return ret
}

// return false if topic not yet created
func RegisterParallelTopicConsumer(topic string, messageCallback MessageCallback) bool {
	ret := false
	if HasTopic(topic) {
		if value, ok := parallelMQ.Load(topic); ok {
			if callbacks, ok := value.([]MessageCallback); ok {
				ret = true
				callbacks = append(callbacks, messageCallback)
				parallelMQ.Store(topic, callbacks)
			}
		}

		//go func() {
		//	if value, ok := parallelMQ.Load(topic); ok {
		//		fmt.Println("consumer has parallelMQ")
		//		if messageChan, ok := value.(chan interface{}); ok {
		//			fmt.Println("consumer waiting")
		//			for{
		//				select {
		//					case msg := <- messageChan :
		//						messageCallback.HandleMessage(topic, msg, true)
		//				}
		//			}
		//		}
		//
		//	}
		//}()
	}
	return ret
}

// return false if topic not created yet
func AddSerialTopicConsumer(topic string, messageCallback MessageCallback) bool {
	ret := false
	if HasTopic(topic) {
		if value, ok := serialMQ.Load(topic); ok {
			if callbacks, ok := value.([]MessageCallback); ok {
				ret = true
				callbacks = append(callbacks, messageCallback)
				serialMQ.Store(topic, callbacks)
			}
		}
	}
	return ret
}

// return false if messageCallback not created yet added
func RemoveSerialTopicConsumer(topic string, messageCallback MessageCallback) bool {
	ret := false
	if HasTopic(topic) {
		if value, ok := serialMQ.Load(topic); ok {
			if callbacks, ok := value.([]MessageCallback); ok {
				for index, callback := range callbacks {
					if callback == messageCallback {
						newCallbacks := append(callbacks[:index], callbacks[index+1:]...)
						serialMQ.Store(topic, newCallbacks)
						ret = true
						break
					}
				}
			}
		}
	}
	return ret
}

// return false if topic not created yet
// call serial consumers one by one sync
// and then call parallel consumers async
func Produce(topic string, message interface{}) bool {
	ret := false
	if HasTopic(topic) {
		ret = true
		//call serial callback firstly
		if value, ok := serialMQ.Load(topic); ok {
			if callbacks, ok := value.([]MessageCallback); ok {
				for _, callback := range callbacks {
					if callback.HandleMessage(topic, message, false) {
						break
					}
				}
			}
		}

		//call parallel consumer
		if value, ok := parallelMQ.Load(topic); ok {
			if callbacks, ok := value.([]MessageCallback); ok {
				for _, callback := range callbacks {
					go callback.HandleMessage(topic, message, true)
				}
			}
		}
	}
	return ret
}
