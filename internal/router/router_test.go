package router

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"
)

func TestRouter(t *testing.T) {
	// 设置为测试模式
	gin.SetMode(gin.TestMode)
	router := SetupRouter()

	t.Run("Test Ping", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/ping", nil)
		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("期望状态码 %d，实际得到 %d", http.StatusOK, w.Code)
		}

		var response map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &response)
		if err != nil {
			t.Fatalf("解析响应失败: %v", err)
		}

		if response["message"] != "pong" {
			t.Errorf("期望消息为 'pong'，实际得到 %s", response["message"])
		}
	})

	t.Run("Test Concurrent Requests", func(t *testing.T) {
		// 重置计数器
		counter = 0

		// 并发请求数
		concurrentRequests := 10
		var wg sync.WaitGroup
		responses := make([]map[string]interface{}, concurrentRequests)

		// 同时发起多个请求
		for i := 0; i < concurrentRequests; i++ {
			wg.Add(1)
			go func(index int) {
				defer wg.Done()
				w := httptest.NewRecorder()
				req, _ := http.NewRequest("GET", "/concurrent", nil)
				router.ServeHTTP(w, req)

				var response map[string]interface{}
				json.Unmarshal(w.Body.Bytes(), &response)
				responses[index] = response
			}(i)
		}
		wg.Wait()

		// 验证结果
		routineIDs := make(map[string]bool)
		for _, resp := range responses {
			// 检查 routineID 是否唯一
			routineID := resp["routineID"].(string)
			if routineIDs[routineID] {
				t.Errorf("发现重复的 routineID: %s", routineID)
			}
			routineIDs[routineID] = true

			// 验证计数器值是否在预期范围内
			count := int(resp["counter"].(float64))
			if count < 1 || count > concurrentRequests {
				t.Errorf("计数器值超出预期范围: %d", count)
			}
		}

		// 验证最终计数器值
		if counter != concurrentRequests {
			t.Errorf("期望最终计数器值为 %d，实际为 %d", concurrentRequests, counter)
		}
	})

	t.Run("Test Request Scope Isolation", func(t *testing.T) {
		// 连续发送两个请求，验证请求隔离性
		w1 := httptest.NewRecorder()
		req1, _ := http.NewRequest("GET", "/request-scope", nil)
		router.ServeHTTP(w1, req1)

		var response1 map[string]interface{}
		json.Unmarshal(w1.Body.Bytes(), &response1)

		// 等待一小段时间后发送第二个请求
		time.Sleep(time.Millisecond * 100)

		w2 := httptest.NewRecorder()
		req2, _ := http.NewRequest("GET", "/request-scope", nil)
		router.ServeHTTP(w2, req2)

		var response2 map[string]interface{}
		json.Unmarshal(w2.Body.Bytes(), &response2)

		// 验证两个请求的 routineID 不同
		if response1["routineID"] == response2["routineID"] {
			t.Error("不同请求的 routineID 不应该相同")
		}

		// 验证两个请求都没有获取到前一个请求的值
		if response1["exists"].(bool) || response2["exists"].(bool) {
			t.Error("请求不应该能够访问其他请求的上下文数据")
		}
	})
}

func TestRouterConcurrencyWithMutex(t *testing.T) {
	// 设置为测试模式
	gin.SetMode(gin.TestMode)
	router := SetupRouter()

	// 使用互斥锁保护计数器
	var mu sync.Mutex
	var safeCounter int

	// 添加一个使用互斥锁的路由
	router.GET("/concurrent-safe", func(c *gin.Context) {
		routineID := fmt.Sprintf("%p", c.Request)

		mu.Lock()
		safeCounter++
		currentCount := safeCounter
		mu.Unlock()

		time.Sleep(time.Second)

		c.JSON(http.StatusOK, gin.H{
			"routineID": routineID,
			"counter":   currentCount,
			"time":      time.Now().Format("15:04:05.000"),
		})
	})

	t.Run("Test Thread-Safe Counter", func(t *testing.T) {
		concurrentRequests := 10
		var wg sync.WaitGroup
		responses := make([]map[string]interface{}, concurrentRequests)

		for i := 0; i < concurrentRequests; i++ {
			wg.Add(1)
			go func(index int) {
				defer wg.Done()
				w := httptest.NewRecorder()
				req, _ := http.NewRequest("GET", "/concurrent-safe", nil)
				router.ServeHTTP(w, req)

				var response map[string]interface{}
				json.Unmarshal(w.Body.Bytes(), &response)
				responses[index] = response
			}(i)
		}
		wg.Wait()

		// 验证计数器值的唯一性
		counts := make(map[float64]bool)
		for _, resp := range responses {
			count := resp["counter"].(float64)
			if counts[count] {
				t.Errorf("发现重复的计数器值: %f", count)
			}
			counts[count] = true
		}

		// 验证最终计数器值
		if safeCounter != concurrentRequests {
			t.Errorf("期望最终计数器值为 %d，实际为 %d", concurrentRequests, safeCounter)
		}
	})
}
