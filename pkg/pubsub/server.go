package pubsub

import (
	"context"
	"log"
	"sync"
)

// Server implements the PubSubService.
// It manages subscribers and broadcasts messages.
type Server struct {
	UnimplementedPubSubServiceServer
	mu          sync.Mutex
	subscribers map[string][]chan *Message
}

// NewServer creates a new server.
func NewServer() *Server {
	return &Server{
		subscribers: make(map[string][]chan *Message),
	}
}

// Subscribe handles a new subscription request.
func (s *Server) Subscribe(req *SubscribeRequest, stream PubSubService_SubscribeServer) error {
	log.Printf("Received subscription request for topic: %s", req.Topic)

	ch := make(chan *Message, 10) // Buffered channel

	s.mu.Lock()
	s.subscribers[req.Topic] = append(s.subscribers[req.Topic], ch)
	s.mu.Unlock()

	log.Printf("Client subscribed to topic: %s. Total subscribers for this topic: %d", req.Topic, len(s.subscribers[req.Topic]))

	for {
		select {
		case <-stream.Context().Done():
			log.Printf("Client for topic %s has disconnected.", req.Topic)
			s.mu.Lock()
			newSubscribers := []chan *Message{}
			for _, subscriber := range s.subscribers[req.Topic] {
				if subscriber != ch {
					newSubscribers = append(newSubscribers, subscriber)
				}
			}
			s.subscribers[req.Topic] = newSubscribers
			s.mu.Unlock()
			return nil
		case msg := <-ch:
			if err := stream.Send(msg); err != nil {
				log.Printf("Failed to send message to subscriber for topic %s: %v", req.Topic, err)
				return err
			}
		}
	}
}

// Publish handles a new message publication.
func (s *Server) Publish(ctx context.Context, req *PublishRequest) (*PublishResponse, error) {
	log.Printf("Received publish request for topic: %s, content: %s", req.Topic, req.Content)

	s.mu.Lock()
	defer s.mu.Unlock()

	msg := &Message{
		Topic:   req.Topic,
		Content: req.Content,
	}

	if subscribers, found := s.subscribers[req.Topic]; found {
		log.Printf("Broadcasting to %d subscribers for topic: %s", len(subscribers), req.Topic)
		for _, ch := range subscribers {
			select {
			case ch <- msg:
			default:
				log.Printf("Subscriber channel for topic %s is full. Message dropped.", req.Topic)
			}
		}
	}

	return &PublishResponse{Success: true}, nil
}
