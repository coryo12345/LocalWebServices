package manager

import (
	"math"
	"sync"
)

const (
	DefaultPollCount         = 10
	DefaultVisibilityTimeout = 30 * 1000
)

//go:generate mockgen -destination=./mocks/queue_mock.go . IQueue
type IQueue interface {
	AddMessage(content string) *Message
	Poll() []*Message
	PollByCount(count int) []*Message
	DeleteMessage(messageId string) bool
	GetName() string
	GetOrder() string
	GetVisibilityTimeout() int
}

type Queue struct {
	name              string
	order             string
	messages          []*Message
	visibilityTimeout int // ms to block messages
	mutex             sync.Mutex
}

func NewQueueRaw() *Queue {
	return NewQueue("", "fifo", DefaultVisibilityTimeout)
}

func NewQueue(name string, order string, timeout int) *Queue {
	messages := make([]*Message, 0)
	return NewQueueWithMessages(name, timeout, order, messages)
}

func NewQueueWithMessages(name string, timeout int, order string, messages []*Message) *Queue {
	return &Queue{
		name:              name,
		order:             order,
		messages:          messages,
		visibilityTimeout: timeout,
		mutex:             sync.Mutex{},
	}
}

func (q *Queue) GetName() string {
	return q.name
}

func (q *Queue) GetOrder() string {
	return q.order
}

func (q *Queue) GetVisibilityTimeout() int {
	return q.visibilityTimeout
}

func (q *Queue) AddMessage(content string) *Message {
	message := NewMessage(content)
	q.mutex.Lock()
	q.messages = append(q.messages, message)
	q.mutex.Unlock()
	return message
}

func (q *Queue) DeleteMessage(messageId string) bool {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	index := -1
	for i, message := range q.messages {
		if message.id.String() == messageId {
			index = i
			break
		}
	}

	if index == -1 {
		return false
	}

	q.messages = append(q.messages[:index], q.messages[index+1:]...)
	return true
}

func (q *Queue) Poll() []*Message {
	return q.PollByCount(DefaultPollCount)
}

func (q *Queue) PollByCount(count int) []*Message {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	if q.order == "filo" {
		return pollQueueFilo(q.messages, q.visibilityTimeout, count)
	} else {
		return pollQueueFifo(q.messages, q.visibilityTimeout, count)
	}
}

func pollQueueFifo(m []*Message, visibilityTimeout int, count int) []*Message {
	size := int(math.Min(float64(count), float64(len(m))))
	polledMessages := make([]*Message, 0, size)
	i := 0
	for len(polledMessages) < size && i < len(m) {
		message := m[i]
		if message.visible {
			message.Lock(visibilityTimeout)
			polledMessages = append(polledMessages, message)
		}
		i++
	}
	return polledMessages
}

func pollQueueFilo(m []*Message, visibilityTimeout int, count int) []*Message {
	size := int(math.Min(float64(count), float64(len(m))))
	polledMessages := make([]*Message, 0, size)
	i := len(m) - 1
	for len(polledMessages) < size && i >= 0 {
		message := m[i]
		if message.visible {
			message.Lock(visibilityTimeout)
			polledMessages = append(polledMessages, message)
		}
		i--
	}
	return polledMessages
}
