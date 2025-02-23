package main

import (
	"fmt"
	"sync"
	"time"
)

type pubsub struct {
	mu sync.RWMutex
    subscribers map[string][]chan string
}

func NewPubSub() *pubsub {
    return &pubsub{subscribers: make(map[string][]chan string)};
}

// subscribe add subriber to the topic
func (ps *pubsub) subcribe(topic string) <-chan string {
    ps.mu.Lock()
    defer ps.mu.Unlock()

    ch := make(chan string, 10);
    ps.subscribers[topic] = append(ps.subscribers[topic], ch);
    return ch
}

// Publish sends a message to all subscribers of a topic
func (ps *pubsub) publish(topic, message string) {
    ps.mu.RLock()
    defer ps.mu.RUnlock()

    for _, ch := range ps.subscribers[topic] {
        ch <- message
    }
}

// Unsubscribe closes the channel and removes it from subscribers
func (ps *pubsub) unsubscribe(topic string, sub <- chan string) {
    ps.mu.Lock()
    defer ps.mu.Unlock()

    subscribers := ps.subscribers[topic]

    for i, ch := range subscribers {
        if ch == sub {
            close(ch)
            ps.subscribers[topic] = append(subscribers[:i], subscribers[i+1:]...)
            break
        }
    }
}

func main() {
    ps := NewPubSub()

    // Subscriber 1
	sub1 := ps.subcribe("cricket")

    go func() {
		for msg := range sub1 {
			fmt.Println("Subscriber 1 received:", msg)
		}
	}()

    // Subscriber 2
	sub2 := ps.subcribe("cricket")
	go func() {
		for msg := range sub2 {
			fmt.Println("Subscriber 2 received:", msg)
		}
	}()

    // Publisher sending messages
	ps.publish("cricket", "King Kohli hits century!!")
	ps.publish("cricket", "India won the champions trophy!!!")

    time.Sleep(time.Second)

    ps.unsubscribe("cricket", sub1)

	// Publish more messages
	ps.publish("cricket", "Australia into semi-finals!")

	time.Sleep(time.Second)
}