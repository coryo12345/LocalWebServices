package api

import (
	"encoding/json"
	"fmt"
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

func (q *QueueAPI) GetAllQueues() ([]byte, *APIError) {
	queueMap := q.Manager.GetQueues()
	keys := make([]string, 0, len(queueMap))
	for k := range queueMap {
		keys = append(keys, k)
	}

	jsonData, err := json.Marshal(keys)
	if err != nil {
		return nil, NewApiError("Unable to retrieve queues", http.StatusInternalServerError)
	}
	return jsonData, nil
}

type QueueRequest struct {
	Name    string `json:"name"`
	Order   string `json:"order"`
	Timeout int    `json:"timeout"`
}

func (q *QueueAPI) CreateQueue(r *http.Request) ([]byte, *APIError) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, NewApiError("Error reading request body", http.StatusInternalServerError)
	}
	queueReq := QueueRequest{
		Name:    "",
		Order:   "",
		Timeout: -1,
	}
	err = json.Unmarshal(body, &queueReq)
	if err != nil {
		return nil, NewApiError("Incorrect request body", http.StatusBadRequest)
	}

	if queueReq.Name == "" {
		return nil, NewApiError("Name must be at least 1 character", http.StatusBadRequest)
	}

	if queueReq.Timeout == -1 {
		queueReq.Timeout = manager.DefaultVisibilityTimeout
	}

	_, err = q.Manager.AddQueue(queueReq.Name, queueReq.Order, queueReq.Timeout)
	if err != nil {
		return nil, NewApiError(err.Error(), http.StatusBadRequest)
	}

	log.Printf("Created queue %s\n", queueReq.Name)
	return []byte(fmt.Sprintf("created queue %s", queueReq.Name)), nil
}

func (q *QueueAPI) DeleteQueue(r *http.Request) ([]byte, *APIError) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, NewApiError("Error reading request body", http.StatusInternalServerError)
	}

	queueReq := QueueRequest{}
	err = json.Unmarshal(body, &queueReq)
	if err != nil {
		return nil, NewApiError("Incorrect request body", http.StatusBadRequest)
	}

	if queueReq.Name == "" {
		return nil, NewApiError("Name must be at least 1 character", http.StatusBadRequest)
	}

	err = q.Manager.DeleteQueue(queueReq.Name)
	if err != nil {
		return nil, NewApiError(err.Error(), http.StatusBadRequest)
	}

	log.Printf("Deleted queue %s\n", queueReq.Name)
	return []byte(fmt.Sprintf("deleted queue %s", queueReq.Name)), nil
}
