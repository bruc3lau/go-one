package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Upgrader 持有 WebSocket 的配置，比如读写缓冲区大小。
// 我们这里使用默认设置，但需要提供一个 CheckOrigin 函数。
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// CheckOrigin 会检查请求的来源。在生产环境中，你应该在这里进行域名校验。
	// 为了简单起见，我们允许任何来源的连接。
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// handleConnections 是处理 WebSocket 请求的 HTTP 处理器。
func handleConnections(w http.ResponseWriter, r *http.Request) {
	// 将 HTTP 连接升级为 WebSocket 连接。
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	// 确保在函数返回时关闭连接。
	defer ws.Close()

	log.Println("Client Connected")

	// 创建一个无限循环来监听新消息。
	for {
		// ReadMessage 从客户端读取一条消息。
		// messageType 是消息类型 (文本, 二进制等)。
		// p 是消息内容 (一个字节切片)。
		messageType, p, err := ws.ReadMessage()
		if err != nil {
			log.Println(err)
			break // 如果有错误（比如客户端断开），就退出循环。
		}

		// 打印接收到的消息。
		log.Printf("Received msg: %s", p)

		// WriteMessage 将消息写回给客户端（实现回声效果）。
		if err := ws.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			break
		}
	}
}

func main() {
	// 配置 HTTP 路由，将 "/ws" 路径的请求交给 handleConnections 处理。
	http.HandleFunc("/ws", handleConnections)

	// 启动 HTTP 服务器，监听在 8080 端口。
	log.Println("http server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
