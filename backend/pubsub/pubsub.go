package pubsub

import (
	"sync"
)

// PubSub handles topic-based in-memory pub-sub
type PubSub struct {
	mu          sync.RWMutex
	subscribers map[string][]chan interface{}
}

// NewPubSub initializes the map
func NewPubSub() *PubSub {
	return &PubSub{
		subscribers: make(map[string][]chan interface{}),
	}
}

// Subscribe adds a subscriber to the topic
func (ps *PubSub) Subscribe(topic string, ch chan interface{}) error {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	ps.subscribers[topic] = append(ps.subscribers[topic], ch)
	return nil
}

// Publish sends the data to all subscribers of the topic
func (ps *PubSub) Publish(topic string, data interface{}) {
	ps.mu.RLock()
	defer ps.mu.RUnlock()

	for _, ch := range ps.subscribers[topic] {
		// Use a non-blocking send to prevent deadlocks
		select {
		case ch <- data:
		default:
		}
	}
}
