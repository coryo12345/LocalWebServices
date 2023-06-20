package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	manager "property_manager/manager"
)

const (
	DEFAULT_API_PORT        = "8081"
	ENV_API_PORT            = "API_PORT"
	INTERNAL_SERVER_ERR_MSG = "Something went wrong"
)

func StartApi() {
	http.HandleFunc("/", handleRequest)

	apiPort := os.Getenv(ENV_API_PORT)
	if apiPort == "" {
		apiPort = DEFAULT_API_PORT
	}
	address := fmt.Sprintf(":%s", apiPort)

	log.Printf("Starting Property Manager API on %s\n", address)
	log.Fatal(http.ListenAndServe(address, nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request: %s /\n", r.Method)
	switch r.Method {
	case http.MethodGet:
		GetPropertyHandler(w, r)
	case http.MethodPost:
		SetPropertyHandler(w, r)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}

}

func GetPropertyHandler(w http.ResponseWriter, r *http.Request) {
	man, err := manager.GetPropertyManagerSingleton()
	if err != nil {
		http.Error(w, INTERNAL_SERVER_ERR_MSG, http.StatusInternalServerError)
		return
	}

	queryValues := r.URL.Query()
	key := queryValues.Get("key")
	if key == "" {
		http.Error(w, "key query parameter must be provided", http.StatusBadRequest)
		return
	}

	value, err := man.GetProperty(key)
	if err != nil {
		http.Error(w, INTERNAL_SERVER_ERR_MSG, http.StatusInternalServerError)
		return
	}

	w.Write([]byte(value))
}

func SetPropertyHandler(w http.ResponseWriter, r *http.Request) {
	man, err := manager.GetPropertyManagerSingleton()
	if err != nil {
		http.Error(w, INTERNAL_SERVER_ERR_MSG, http.StatusInternalServerError)
		return
	}

	queryValues := r.URL.Query()
	key := queryValues.Get("key")
	if key == "" {
		http.Error(w, "key query parameter must be provided", http.StatusBadRequest)
		return
	}
	value := queryValues.Get("value")

	previousValue, err := man.SetProperty(key, value)
	if err != nil {
		http.Error(w, INTERNAL_SERVER_ERR_MSG, http.StatusInternalServerError)
		return
	}

	w.Write([]byte(previousValue))
}
