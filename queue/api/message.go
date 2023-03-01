package api

import "queue/manager"

type MessageAPI struct {
	Manager manager.IQueueManager
}

func NewMessageAPI(man manager.IQueueManager) *QueueAPI {
	qapi := QueueAPI{
		Manager: man,
	}
	return &qapi
}

// TODO what is this returning?
// What does this take?
// default polling length?
func (m *MessageAPI) GetMessages() {}

// TODO what are the options this takes?
// bool to not allow deleting invis messages?
// What does this return?
func (m *MessageAPI) DeleteMessage() {}

// TODO what does this take?
// What does this return
func (m *MessageAPI) PublishMessage() {}
