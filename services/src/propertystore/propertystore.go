package propertystore

import (
	"encoding/json"
	"fmt"
	"localwebservices/common"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type PropertyStore struct {
	storeName       string
	storageFilePath string
	values          map[string]string
	mutex           *sync.Mutex
	queue           chan bool
}

func NewPropertyStore(storeName string) common.Service {
	return &PropertyStore{
		storeName: storeName,
		values:    make(map[string]string),
		mutex:     &sync.Mutex{},
		queue:     make(chan bool),
	}
}

func (s *PropertyStore) Start(urlPrefix string) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	filename := s.storeName + ".json"
	s.storageFilePath = filepath.Join(dir, "..", "data", "propertystore", filename)

	// create file if it doesn't exist
	_, err = os.Stat(s.storageFilePath)
	if err != nil {
		err = os.MkdirAll(filepath.Dir(s.storageFilePath), os.ModePerm)
		if err != nil {
			return err
		}

		_, err = os.Create(s.storageFilePath)
		if err != nil {
			return err
		}
	} else {
		// read starting values
		b, err := os.ReadFile(s.storageFilePath)
		if err != nil {
			return err
		}
		err = json.Unmarshal(b, &s.values)
		if err != nil {
			return err
		}
	}

	// start save queue
	go startSaveWorker(s)

	http.HandleFunc(fmt.Sprintf("GET %s/", urlPrefix), s.getProperty)
	http.HandleFunc(fmt.Sprintf("POST %s/", urlPrefix), s.setProperty)
	return nil
}

func (s *PropertyStore) Shutdown() {
	log.Printf("Shutting down property store '%s'\n", s.storeName)
	savePropertyFile(s.storageFilePath, s.values)
}

func (s *PropertyStore) getProperty(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	s.mutex.Lock()
	value := s.values[name]
	s.mutex.Unlock()
	w.Write([]byte(value))
}

type setPropertyRequestBody struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (s *PropertyStore) setProperty(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var body setPropertyRequestBody
	err := decoder.Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("request format incorrect"))
		return
	}

	s.mutex.Lock()
	s.values[body.Name] = body.Value
	s.mutex.Unlock()

	// queue a save
	s.queue <- true

	w.Write([]byte("updated property"))
}

func savePropertyFile(filepath string, value map[string]string) error {
	b, err := json.MarshalIndent(value, "", "\t")
	if err != nil {
		log.Printf("ERROR: failed to convert properties to JSON\n")
		return err
	}
	err = os.WriteFile(filepath, b, 0)
	if err != nil {
		log.Printf("ERROR: failed to save property file: %s\n", filepath)
		return err
	}
	return nil
}

func startSaveWorker(s *PropertyStore) {
	save := common.CreateThrottledFunc(func() { savePropertyFile(s.storageFilePath, s.values) }, time.Second*5)
	for range s.queue {
		save(false)
	}
}
