package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"go-one/pkg/pubsub"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50052, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pubsubServer := pubsub.NewServer()
	pubsub.RegisterPubSubServiceServer(grpcServer, pubsubServer)

	log.Printf("Pub/Sub server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
