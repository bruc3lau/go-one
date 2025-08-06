package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	pb "go-one/pkg/simplepubsub"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:50053"
)

func main() {
	// 建立到服务端的连接
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("连接失败: %v", err)
	}
	defer conn.Close()

	client := pb.NewPublisherClient(conn)

	// 生成一个客户端 ID，这里我们简单地使用主机名
	clientID, err := os.Hostname()
	if err != nil {
		log.Fatalf("获取主机名失败: %v", err)
	}
	clientID = fmt.Sprintf("%s-%d", clientID, os.Getpid()) // 加上进程ID确保唯一性

	// 调用 Subscribe 方法来订阅消息
	stream, err := client.Subscribe(context.Background(), &pb.ClientInfo{ClientId: clientID})
	if err != nil {
		log.Fatalf("订阅失败: %v", err)
	}

	log.Printf("客户端 %s 已成功订阅，等待接收消息...", clientID)

	// 在一个无限循环中接收来自服务端的消息
	for {
		message, err := stream.Recv()
		if err == io.EOF {
			// 如果流被服务端关闭，会收到 EOF 错误
			log.Println("服务端已关闭连接。")
			break
		}
		if err != nil {
			log.Fatalf("接收消息失败: %v", err)
		}

		// 打印收到的消息
		log.Printf("收到消息: %s", message.GetContent())
	}
}
