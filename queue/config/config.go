package config

import (
	"log"
	"queue/manager"
)

type Queue struct {
	Name               string
	Visibility_timeout int
	Order              string
	Messages           []string
}

type QueueConfig struct {
	Version float32
	Queues  []Queue
}

func InitConfig(qm *manager.QueueManager) {
	config := getConfig()
	if config == nil || config.Queues == nil {
		return
	}
	log.Printf("Initializing %d queues...\n", len(config.Queues))

	// Apply config to queue manager
	for _, q := range config.Queues {
		qm.AddQueueWithMessages(q.Name, q.Order, q.Visibility_timeout, q.Messages)
	}
}
