package manager

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	id        uuid.UUID
	content   string
	timestamp string
	visible   bool
}

func NewMessage(content string) *Message {
	message := Message{
		id:        uuid.New(),
		content:   content,
		timestamp: time.Now().String(),
		visible:   true,
	}
	return &message
}

func (m *Message) GetContent() string {
	return m.content
}

func (m *Message) GetTimestamp() string {
	return m.timestamp
}

func (m *Message) GetId() uuid.UUID {
	return m.id
}

func (m *Message) Lock(ms int) {
	m.visible = false
	go func() {
		time.Sleep(time.Millisecond * time.Duration(ms))
		// this can run even if the mutex is not acquired for the current queue
		m.visible = true
	}()
}
