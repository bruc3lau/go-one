package util

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHTTPClient(t *testing.T) {
	// 创建测试服务器
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/test-get":
			// 测试GET请求
			if r.Method != "GET" {
				t.Errorf("期望请求方法为GET，实际获得：%s", r.Method)
			}
			if v := r.Header.Get("X-Test-Header"); v != "test-value" {
				t.Errorf("期望请求头X-Test-Header值为test-value，实际获得：%s", v)
			}

			response := map[string]string{"message": "get success"}
			json.NewEncoder(w).Encode(response)

		case "/test-post":
			// 测试POST请求
			if r.Method != "POST" {
				t.Errorf("期望请求方法为POST，实际获得：%s", r.Method)
			}
			if ct := r.Header.Get("Content-Type"); ct != "application/json" {
				t.Errorf("期望Content-Type为application/json，实际获得：%s", ct)
			}

			// 解析请求体
			var requestBody map[string]interface{}
			if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
				t.Errorf("解析请求体失败：%v", err)
			}

			// 验证请求体
			if data, ok := requestBody["data"].(string); !ok || data != "test data" {
				t.Errorf("期望请求体data字段为'test data'，实际获得：%v", requestBody["data"])
			}

			// 返回响应
			response := map[string]string{"message": "post success"}
			json.NewEncoder(w).Encode(response)
		}
	}))
	defer ts.Close()

	// 创建HTTP客户端
	client := NewHTTPClient(ts.URL, 5*time.Second)

	t.Run("Test GET Request", func(t *testing.T) {
		headers := map[string]string{
			"X-Test-Header": "test-value",
		}

		response, err := client.Get("/test-get", headers)
		if err != nil {
			t.Fatalf("GET请求失败：%v", err)
		}

		var result map[string]string
		if err = json.Unmarshal(response, &result); err != nil {
			t.Fatalf("解析响应失败：%v", err)
		}

		if msg := result["message"]; msg != "get success" {
			t.Errorf("期望响应message为'get success'，实际获得：%s", msg)
		}
	})

	t.Run("Test POST Request", func(t *testing.T) {
		data := map[string]interface{}{
			"data": "test data",
		}

		headers := map[string]string{
			"Content-Type": "application/json",
		}

		response, err := client.Post("/test-post", data, headers)
		if err != nil {
			t.Fatalf("POST请求失败：%v", err)
		}

		var result map[string]string
		if err = json.Unmarshal(response, &result); err != nil {
			t.Fatalf("解析响应失败：%v", err)
		}

		if msg := result["message"]; msg != "post success" {
			t.Errorf("期望响应message为'post success'，实际获得：%s", msg)
		}
	})
}
