package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingRoute(t *testing.T) {
	// 获取路由
	router := setupRouter()

	// 创建一个测试请求记录器
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	// 断言状态码是200
	if w.Code != http.StatusOK {
		t.Errorf("期望状态码为200，但获得%v", w.Code)
	}

	// 检查返回的JSON数据
	var got map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
		t.Fatalf("无法解析响应JSON：%v", err)
	}

	// 验证返回的消息
	if got["message"] != "pong" {
		t.Errorf("期望消息为'pong'，但获得'%v'", got["message"])
	}
}

func TestHelloRoute(t *testing.T) {
	// 获取路由
	router := setupRouter()

	// 创建一个测试请求记录器
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/hello", nil)
	router.ServeHTTP(w, req)

	// 断言状态码是200
	if w.Code != http.StatusOK {
		t.Errorf("期望状态码为200，但获得%v", w.Code)
	}

	// 检查返回的JSON数据
	var got map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
		t.Fatalf("无法解析响应JSON：%v", err)
	}

	// 验证返回的消息
	if got["message"] != "Hello, GoOne!" {
		t.Errorf("期望消息为'Hello, GoOne!'，但获得'%v'", got["message"])
	}
}
