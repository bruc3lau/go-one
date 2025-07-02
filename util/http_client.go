package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// HTTPClient 是一个HTTP客户端工具
type HTTPClient struct {
	client  *http.Client
	baseURL string
}

// NewHTTPClient 创建一个新的HTTP客户端
func NewHTTPClient(baseURL string, timeout time.Duration) *HTTPClient {
	return &HTTPClient{
		client: &http.Client{
			Timeout: timeout,
		},
		baseURL: baseURL,
	}
}

// Get 发送GET请求
func (c *HTTPClient) Get(path string, headers map[string]string) ([]byte, error) {
	url := c.baseURL + path
	log.Printf("发送GET请求到: %s\n", url)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}

	// 添加请求头
	for key, value := range headers {
		req.Header.Add(key, value)
	}

	start := time.Now()
	resp, err := c.client.Do(req)
	if err != nil {
		log.Printf("GET请求失败: %v\n", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("读取响应失败: %v\n", err)
		return nil, err
	}

	log.Printf("GET请求完成 - 状态码: %d, 耗时: %v\n", resp.StatusCode, time.Since(start))
	return body, nil
}

// Post 发送POST请求
func (c *HTTPClient) Post(path string, data interface{}, headers map[string]string) ([]byte, error) {
	url := c.baseURL + path
	log.Printf("发送POST请求到: %s\n", url)

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("JSON序列化失败: %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}

	// 设置默认Content-Type
	if _, exists := headers["Content-Type"]; !exists {
		req.Header.Set("Content-Type", "application/json")
	}

	// 添加其他请求头
	for key, value := range headers {
		req.Header.Add(key, value)
	}

	start := time.Now()
	resp, err := c.client.Do(req)
	if err != nil {
		log.Printf("POST请求失败: %v\n", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("读取响应失败: %v\n", err)
		return nil, err
	}

	log.Printf("POST请求完成 - 状态码: %d, 耗时: %v\n", resp.StatusCode, time.Since(start))
	return body, nil
}
