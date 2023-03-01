package manager

import (
	"errors"
	"fmt"
)

//go:generate mockgen -destination=./mocks/manager_mock.go . IQueueManager
type IQueueManager interface {
	AddQueue(name string, order string, timeout int) (IQueue, error)
	AddQueueWithMessages(name string, order string, timeout int, rawMessages []string) (IQueue, error)
	GetQueues() map[string]IQueue
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

	messages := make([]*Message, len(rawMessages))
	queue := NewQueueWithMessages(name, timeout, order, messages)
	if q.queues[queue.Name] != nil {
		return nil, fmt.Errorf("queue with name %s already exists", queue.Name)
	}

	q.queues[queue.Name] = queue

	return queue, nil
}

func (q *QueueManager) GetQueues() map[string]IQueue {
	return q.queues
}

func (q *QueueManager) DeleteQueue(name string) error {
	if q.queues[name] == nil {
		return fmt.Errorf("Queue with name %s does not exist", name)
	}

	delete(q.queues, name)
	return nil
}
