package manager

import (
	"errors"
	"fmt"
	"regexp"
)

//go:generate mockgen -destination=./mocks/manager_mock.go . IQueueManager
type IQueueManager interface {
	AddQueue(name string, order string, timeout int) (IQueue, error)
	AddQueueWithMessages(name string, order string, timeout int, rawMessages []string) (IQueue, error)
	GetQueues() map[string]IQueue
	GetQueue(name string) IQueue
	DeleteQueue(name string) error
}

type QueueManager struct {
	queues map[string]IQueue
}

func NewQueueManager() *QueueManager {
	manager := &QueueManager{}
	manager.queues = make(map[string]IQueue)
	return manager
}

func (q *QueueManager) AddQueue(name string, order string, timeout int) (IQueue, error) {
	messages := make([]string, 0)
	return q.AddQueueWithMessages(name, order, timeout, messages)
}

func (q *QueueManager) AddQueueWithMessages(name string, order string, timeout int, rawMessages []string) (IQueue, error) {
	if order != "fifo" && order != "filo" {
		return nil, errors.New("order should be one of: ['fifo', 'filo'] ")
	}

	matched, err := regexp.Match(`^[0-9a-zA-Z_\-]+$`, []byte(name))
	if err != nil || !matched {
		return nil, errors.New("queue name must be alphanumeric with dashes or underscores")
	}

	messages := make([]*Message, len(rawMessages))
	queue := NewQueueWithMessages(name, timeout, order, messages)
	if q.queues[queue.name] != nil {
		return nil, fmt.Errorf("queue with name %s already exists", queue.name)
	}

	q.queues[queue.name] = queue

	return queue, nil
}

func (q *QueueManager) GetQueues() map[string]IQueue {
	return q.queues
}

func (q *QueueManager) GetQueue(name string) IQueue {
	return q.queues[name]
}

func (q *QueueManager) DeleteQueue(name string) error {
	if q.queues[name] == nil {
		return fmt.Errorf("Queue with name %s does not exist", name)
	}

	delete(q.queues, name)
	return nil
}
