package api

import (
	"errors"
	"log"
	"net/http"

	"queue/manager"
)

type APIError struct {
	err  error
	code int
}

func NewApiError(msg string, code int) *APIError {
	return &APIError{
		err:  errors.New(msg),
		code: code,
	}
}

type ResourceAPI struct {
	QueueApi   QueueAPI
	MessageApi MessageAPI
}

func NewResourceAPI(manager manager.IQueueManager) *ResourceAPI {
	api := &ResourceAPI{
		QueueApi:   NewQueueAPI(manager),
		MessageApi: NewMessageAPI(manager),
	}
	return api
}

func (q *ResourceAPI) StartApi() {
	// GET - get all queues
	// PUT - create queue
	// DELETE - delete queue
	http.HandleFunc("/queues", q.QueueHandler)

	// assuming: given a name (id)
	// NOTE: SQS does polling by sending a request every X seconds over a period of length Y, where Y > X.
	// GET - Poll for messages
	// DELETE - delete a message
	// PUT - publish a message
	http.HandleFunc("/message", q.MessageHandler)

	// TODO: move port into env var
	log.Printf("Starting Queue API on :8080\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (api *ResourceAPI) MessageHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request: %s /message\n", r.Method)
	var data []byte
	var err *APIError
	switch r.Method {
	case "GET":
		data, err = api.MessageApi.GetMessages()
	case "PUT":
		data, err = api.MessageApi.PublishMessage()
	case "DELETE":
		data, err = api.MessageApi.DeleteMessage()
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	if err != nil {
		http.Error(w, err.err.Error(), err.code)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func (api *ResourceAPI) QueueHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request: %s /queues\n", r.Method)
	var data []byte
	var err *APIError
	switch r.Method {
	case "GET":
		data, err = api.QueueApi.GetAllQueues()
	case "PUT":
		data, err = api.QueueApi.CreateQueue(r)
	case "DELETE":
		data, err = api.QueueApi.DeleteQueue(r)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	if err != nil {
		http.Error(w, err.err.Error(), err.code)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
