package pubsub

import (
	"sync"

	"github.com/yordanos-habtamu/GormWithGraphql/graph/model"
)

// PubSub handles topic-based in-memory pub-sub
type PubSub struct {
	mu        sync.RWMutex
	subscribers map[string][]chan *model.Message
}

// NewPubSub initializes the map
func NewPubSub() *PubSub {
	return &PubSub{
		subscribers: make(map[string][]chan *model.Message), 
	}
}

// Subscribe adds a subscriber to the topic
func (ps *PubSub) Subscribe(chatID string, ch chan *model.Message) error {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	ps.subscribers[chatID] = append(ps.subscribers[chatID], ch)
	return nil
}

// Publish sends the message to all subscribers of the topic
func (ps *PubSub) Publish(chatID string, msg *model.Message) {
	ps.mu.RLock()
	defer ps.mu.RUnlock()

	for _, ch := range ps.subscribers[chatID] {
		ch <- msg
	}
}
