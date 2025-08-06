package main

import (
	"fmt"
	"log"
	"net"
	"time"

	pb "go-one/pkg/simplepubsub"

	"google.golang.org/grpc"
)

const (
	port = ":50053"
)

// server 结构体实现了我们定义的 PublisherServer 接口
type server struct {
	pb.UnimplementedPublisherServer
}

// Subscribe 是我们实现的核心方法
func (s *server) Subscribe(info *pb.ClientInfo, stream pb.Publisher_SubscribeServer) error {
	log.Printf("客户端 %s 已连接并订阅。", info.ClientId)

	// 使用一个 for 循环和 time.Ticker 来每秒发送一次消息
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-stream.Context().Done():
			// 如果客户端断开连接，这个 context 会被取消
			log.Printf("客户端 %s 已断开连接。", info.ClientId)
			return nil // 正常退出
		case t := <-ticker.C:
			// 构建要发送的消息
			message := &pb.Message{
				Content: fmt.Sprintf("来自客户端 %s 的消息，时间戳: %d", info.ClientId, t.UnixMilli()),
			}

			// 通过流将消息发送给客户端
			if err := stream.Send(message); err != nil {
				log.Printf("发送消息失败: %v", err)
				return err // 发送失败，也退出
			}
			log.Printf("已向客户端 %s 发送消息。", info.ClientId)
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("监听端口失败: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterPublisherServer(s, &server{})

	log.Printf("服务端正在监听 %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("启动服务失败: %v", err)
	}
}
