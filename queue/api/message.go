package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"queue/manager"

	"github.com/google/uuid"
)

type MessageAPI struct {
	Manager manager.IQueueManager
}

func NewMessageAPI(man manager.IQueueManager) MessageAPI {
	qapi := MessageAPI{
		Manager: man,
	}
	return qapi
}

// ============================================================
type MessageDataResonse struct {
	Content string    `json:"content"`
	Id      uuid.UUID `json:"id"`
	Created string    `json:"created_timestamp"`
}

// Takes queue_name query parameter
func (m *MessageAPI) GetMessages(r *http.Request) ([]MessageDataResonse, *APIError) {
	name := r.URL.Query().Get("queue_name")
	if name == "" {
		return nil, NewApiError("Must provide queue_name query parameter", http.StatusBadRequest)
	}

	queue := m.Manager.GetQueue(name)
	if queue == nil {
		return nil, NewApiError(fmt.Sprintf("queue %s not found", name), http.StatusNotFound)
	}
	messages := queue.Poll()

	data := make([]MessageDataResonse, len(messages))
	for i, message := range messages {
		data[i].Content = message.GetContent()
		data[i].Id = message.GetId()
		data[i].Created = message.GetTimestamp()
	}

	return data, nil
}

// ============================================================

type DeleteMessageRequestBody struct {
	QueueName string `json:"queue_name"`
	MessageID string `json:"message_id"`
}

type DeleteMessageResponse struct {
	Deleted bool `json:"deleted"`
}

func (m *MessageAPI) DeleteMessage(r *http.Request) (DeleteMessageResponse, *APIError) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return DeleteMessageResponse{}, NewApiError("Error reading request body", http.StatusInternalServerError)
	}

	req := DeleteMessageRequestBody{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		return DeleteMessageResponse{}, NewApiError("Incorrect request body", http.StatusBadRequest)
	}

	if req.QueueName == "" || req.MessageID == "" {
		return DeleteMessageResponse{}, NewApiError("queue_name and message_id are required", http.StatusBadRequest)
	}

	queue := m.Manager.GetQueue(req.QueueName)
	if queue == nil {
		return DeleteMessageResponse{}, NewApiError(fmt.Sprintf("queue %s not found", req.QueueName), http.StatusNotFound)
	}
	deleted := queue.DeleteMessage(req.MessageID)
	if !deleted {
		return DeleteMessageResponse{}, NewApiError(fmt.Sprintf("message not found in queue %s", req.QueueName), http.StatusNotFound)
	}

	resp := DeleteMessageResponse{
		Deleted: deleted,
	}

	return resp, nil
}

// ============================================================

type PublishMessageRequest struct {
	QueueName string `json:"queue_name"`
	Content   string `json:"content"`
}

type PublishMessageResponse struct {
	Id string `json:"id"`
}

// returns message uuid (string)
func (m *MessageAPI) PublishMessage(r *http.Request) (PublishMessageResponse, *APIError) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return PublishMessageResponse{}, NewApiError("Error reading request body", http.StatusInternalServerError)
	}

	messReq := PublishMessageRequest{
		QueueName: "",
		Content:   "",
	}
	err = json.Unmarshal(body, &messReq)
	if err != nil {
		return PublishMessageResponse{}, NewApiError("Incorrect request body", http.StatusBadRequest)
	}

	if messReq.Content == "" || messReq.QueueName == "" {
		return PublishMessageResponse{}, NewApiError("content and queue_name must not be blank", http.StatusBadRequest)
	}

	queue := m.Manager.GetQueue(messReq.QueueName)
	if queue == nil {
		return PublishMessageResponse{}, NewApiError("queue does not exist", http.StatusNotFound)
	}
	message := queue.AddMessage(messReq.Content)

	log.Printf("Added message to queue %s\n", messReq.QueueName)
	resp := PublishMessageResponse{
		Id: message.GetId().String(),
	}
	return resp, nil
}

// ============================================================
