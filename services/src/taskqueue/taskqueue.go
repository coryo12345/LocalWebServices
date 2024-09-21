package taskqueue

import (
	"encoding/json"
	"fmt"
	"localwebservices/common"
	"log"
	"net/http"
	"sync"
)

type TaskQueue struct {
	queueName string
	values    []Task
	mutex     *sync.Mutex
}

type Task struct {
	Priority int    `json:"priority"`
	Value    string `json:"value"`
}

func NewTaskQueue(name string) common.Service {
	return &TaskQueue{
		queueName: name,
		values:    make([]Task, 0),
		mutex:     &sync.Mutex{},
	}
}

func (s *TaskQueue) Start(urlPrefix string) error {
	log.Printf("Starting task queue: %s\n", s.queueName)
	http.HandleFunc(fmt.Sprintf("GET %s/preview", urlPrefix), s.previewEvents)
	http.HandleFunc(fmt.Sprintf("GET %s/latest", urlPrefix), s.getEventByTime)
	http.HandleFunc(fmt.Sprintf("GET %s/priority", urlPrefix), s.getEventByPriority)
	http.HandleFunc(fmt.Sprintf("POST %s/add", urlPrefix), s.addEvent)
	return nil
}

func (s *TaskQueue) Shutdown() {
	log.Printf("Shutting down task queue '%s'\n", s.queueName)
}

func (s *TaskQueue) previewEvents(w http.ResponseWriter, r *http.Request) {
	// just a preview, don't need to lock this
	b, err := json.Marshal(s.values)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(b)
}

func (s *TaskQueue) getEventByPriority(w http.ResponseWriter, r *http.Request) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if len(s.values) == 0 {
		w.Write([]byte("no tasks"))
		return
	}

	highestIdx := len(s.values) - 1
	highest := s.values[highestIdx]
	for i := len(s.values) - 1; i >= 0; i-- {
		if s.values[i].Priority > highest.Priority {
			highestIdx = i
			highest = s.values[i]
		}
	}

	// remove task from queue
	if highestIdx == len(s.values)-1 {
		s.values = s.values[:highestIdx]
	} else {
		s.values = append(s.values[:highestIdx], s.values[highestIdx+1:]...)
	}

	b, err := json.Marshal(highest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(b)
}

func (s *TaskQueue) getEventByTime(w http.ResponseWriter, r *http.Request) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if len(s.values) == 0 {
		w.Write([]byte("no tasks"))
		return
	}

	task := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]

	b, err := json.Marshal(task)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(b)
}

func (s *TaskQueue) addEvent(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	body := Task{
		Priority: -1,
	}
	err := decoder.Decode(&body)
	if err != nil || body.Priority < 0 || body.Value == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("request format incorrect"))
		return
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.values = append(s.values, body)
	w.WriteHeader(http.StatusOK)
}
