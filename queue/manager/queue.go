package manager

import (
	"math"
	"sync"
	"time"
)

const (
	DefaultPollCount         = 10
	DefaultVisibilityTimeout = 30 * 1000
)

//go:generate mockgen -destination=./mocks/queue_mock.go . IQueue
type IQueue interface {
	AddMessage(content string)
	Poll() *[]Message
	PollByCount(count int) *[]Message
}

type Queue struct {
	Name              string
	Persistent        bool
	Order             string
	Messages          []*Message
	VisibilityTimeout int // ms to block messages
	Mutex             sync.Mutex
}

func NewQueueRaw() *Queue {
	return NewQueue("", false, "fifo", DefaultVisibilityTimeout)
}

func NewQueue(name string, persistent bool, order string, timeout int) *Queue {
	messages := make([]*Message, 0)
	return NewQueueWithMessages(name, timeout, order, messages)
}

func NewQueueWithMessages(name string, timeout int, order string, messages []*Message) *Queue {
	return &Queue{
		Name:              name,
		Order:             order,
		Messages:          messages,
		VisibilityTimeout: int(time.Duration(timeout) * time.Millisecond),
		Mutex:             sync.Mutex{},
	}
}

func (q *Queue) AddMessage(content string) {
	message := NewMessage(content)
	q.Mutex.Lock()
	q.Messages = append(q.Messages, message)
	q.Mutex.Unlock()
}

func (q *Queue) Poll() *[]Message {
	return q.PollByCount(DefaultPollCount)
}

func (q *Queue) PollByCount(count int) *[]Message {
	q.Mutex.Lock()
	defer q.Mutex.Unlock()
	if q.Order == "filo" {
		return pollQueueFilo(q.Messages, q.VisibilityTimeout, count)
	} else {
		return pollQueueFifo(q.Messages, q.VisibilityTimeout, count)
	}
}

func pollQueueFifo(m []*Message, visibilityTimeout int, count int) *[]Message {
	size := int(math.Min(float64(count), float64(len(m))))
	polledMessages := make([]Message, 0, size)
	i := 0
	for len(polledMessages) < size && i < len(m) {
		message := m[i]
		if message.visible {
			message.Lock(visibilityTimeout)
			polledMessages = append(polledMessages, *message)
		}
		i++
	}
	return &polledMessages
}

func pollQueueFilo(m []*Message, visibilityTimeout int, count int) *[]Message {
	size := int(math.Min(float64(count), float64(len(m))))
	polledMessages := make([]Message, 0, size)
	i := len(m) - 1
	for len(polledMessages) < size && i >= 0 {
		message := m[i]
		if message.visible {
			message.Lock(visibilityTimeout)
			polledMessages = append(polledMessages, *message)
		}
		i--
	}
	return &polledMessages
}
