package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"queue/manager"
)

const (
	ENV_API_PORT     = "API_PORT"
	DEFAULT_API_PORT = "8080"
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
	// TODO document these endpoint

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

	apiPort := os.Getenv(ENV_API_PORT)
	if apiPort == "" {
		apiPort = DEFAULT_API_PORT
	}
	address := fmt.Sprintf(":%s", apiPort)

	log.Printf("Starting Queue API on %s\n", address)
	log.Fatal(http.ListenAndServe(address, nil))
}

func (api *ResourceAPI) MessageHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request: %s /message\n", r.Method)
	var data interface{} // will be a pointer to some marshalable data
	var err *APIError
	switch r.Method {
	case "GET":
		data, err = api.MessageApi.GetMessages(r)
	case "PUT":
		data, err = api.MessageApi.PublishMessage(r)
	case "DELETE":
		data, err = api.MessageApi.DeleteMessage(r)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	if err != nil {
		http.Error(w, err.err.Error(), err.code)
		return
	}

	bytedata, err2 := marshalData(data)
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytedata)
}

func (api *ResourceAPI) QueueHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request: %s /queues\n", r.Method)
	var data interface{}
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

	bytedata, err2 := marshalData(data)
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytedata)
}

func marshalData(data interface{}) ([]byte, error) {
	var bytedata []byte
	if data != nil {
		var err error
		bytedata, err = json.Marshal(data)
		if err != nil {
			return nil, errors.New("internal server error")
		}
	} else {
		bytedata = make([]byte, 0)
	}
	return bytedata, nil
}
