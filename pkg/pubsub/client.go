package pubsub

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"google.golang.org/grpc"
)

// Client represents a gRPC client for the PubSub service.
type Client struct {
	conn   *grpc.ClientConn
	client PubSubServiceClient
}

// NewClient creates a new gRPC client.
func NewClient(addr string, opts ...grpc.DialOption) (*Client, error) {
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		return nil, err
	}

	return &Client{
		conn:   conn,
		client: NewPubSubServiceClient(conn),
	}, nil
}

// Close closes the client connection.
func (c *Client) Close() {
	c.conn.Close()
}

// Subscribe subscribes to a topic and calls the handler function for each received message.
func (c *Client) Subscribe(ctx context.Context, topic string, handler func(msg *Message)) error {
	log.Printf("Subscribing to topic: %s", topic)

	stream, err := c.client.Subscribe(ctx, &SubscribeRequest{Topic: topic})
	if err != nil {
		return fmt.Errorf("failed to subscribe: %w", err)
	}

	for {
		msg, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("failed to receive message: %w", err)
		}
		handler(msg)
	}
	return nil
}

// Publish sends a message to a topic.
func (c *Client) Publish(ctx context.Context, topic, content string) error {
	_, err := c.client.Publish(ctx, &PublishRequest{Topic: topic, Content: content})
	if err != nil {
		return fmt.Errorf("failed to publish: %w", err)
	}
	return nil
}

// RunCLI starts a command-line interface for the client.
func (c *Client) RunCLI(ctx context.Context, topic string) {
	go func() {
		handler := func(msg *Message) {
			log.Printf("Received message on topic '%s': %s", msg.Topic, msg.Content)
		}
		if err := c.Subscribe(ctx, topic, handler); err != nil {
			log.Fatalf("Subscription failed: %v", err)
		}
	}()

	log.Println("Enter a message to publish (or 'exit' to quit):")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if strings.ToLower(text) == "exit" {
			break
		}

		if err := c.Publish(ctx, topic, text); err != nil {
			log.Printf("Failed to publish: %v", err)
		}
	}
}
