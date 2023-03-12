package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"queue/manager"
)

type QueueAPI struct {
	Manager manager.IQueueManager
}

func NewQueueAPI(man manager.IQueueManager) QueueAPI {
	qapi := QueueAPI{
		Manager: man,
	}
	return qapi
}

// ==========================================================
type GetQueueResponse struct {
	Name              string
	Order             string
	VisibilityTimeout int
}

func (q *QueueAPI) GetAllQueues() ([]GetQueueResponse, *APIError) {
	queueMap := q.Manager.GetQueues()

	queues := make([]GetQueueResponse, 0, len(queueMap))
	for _, queue := range queueMap {
		qr := GetQueueResponse{
			Name:              queue.GetName(),
			Order:             queue.GetOrder(),
			VisibilityTimeout: queue.GetVisibilityTimeout(),
		}
		queues = append(queues, qr)
	}

	return queues, nil
}

// ======================================================

type CreateQueueRequest struct {
	Name    string `json:"name"`
	Order   string `json:"order"`
	Timeout int    `json:"timeout"`
}

type CreateQueueResponse struct{}

func (q *QueueAPI) CreateQueue(r *http.Request) (CreateQueueResponse, *APIError) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return CreateQueueResponse{}, NewApiError("Error reading request body", http.StatusInternalServerError)
	}
	queueReq := CreateQueueRequest{
		Name:    "",
		Order:   "",
		Timeout: -1,
	}
	err = json.Unmarshal(body, &queueReq)
	if err != nil {
		return CreateQueueResponse{}, NewApiError("Incorrect request body", http.StatusBadRequest)
	}

	if queueReq.Name == "" {
		return CreateQueueResponse{}, NewApiError("Name must be at least 1 character", http.StatusBadRequest)
	}

	if queueReq.Timeout == -1 {
		queueReq.Timeout = manager.DefaultVisibilityTimeout
	}

	_, err = q.Manager.AddQueue(queueReq.Name, queueReq.Order, queueReq.Timeout)
	if err != nil {
		return CreateQueueResponse{}, NewApiError(err.Error(), http.StatusBadRequest)
	}

	log.Printf("Created queue %s\n", queueReq.Name)
	return CreateQueueResponse{}, nil
}

// ===========================================

type DeleteQueueRequest struct {
	Name    string `json:"name"`
	Order   string `json:"order"`
	Timeout int    `json:"timeout"`
}

// not really needed at this point
type DeleteQueueResponse struct{}

func (q *QueueAPI) DeleteQueue(r *http.Request) (DeleteQueueResponse, *APIError) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return DeleteQueueResponse{}, NewApiError("Error reading request body", http.StatusInternalServerError)
	}

	queueReq := DeleteQueueRequest{}
	err = json.Unmarshal(body, &queueReq)
	if err != nil {
		return DeleteQueueResponse{}, NewApiError("Incorrect request body", http.StatusBadRequest)
	}

	if queueReq.Name == "" {
		return DeleteQueueResponse{}, NewApiError("Name must be at least 1 character", http.StatusBadRequest)
	}

	err = q.Manager.DeleteQueue(queueReq.Name)
	if err != nil {
		return DeleteQueueResponse{}, NewApiError(err.Error(), http.StatusBadRequest)
	}

	log.Printf("Deleted queue %s\n", queueReq.Name)
	return DeleteQueueResponse{}, nil
}
