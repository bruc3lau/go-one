package main

import (
	"context"
	"flag"
	"log"

	"go-one/pkg/pubsub"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr  = flag.String("addr", "localhost:50052", "the address to connect to")
	topic = flag.String("topic", "news", "the topic to subscribe and publish to")
)

func main() {
	flag.Parse()

	// The line below is now fixed because `grpc` is imported.
	client, err := pubsub.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	client.RunCLI(context.Background(), *topic)
}
