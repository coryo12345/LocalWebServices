package api

import (
	"encoding/json"
	"file_storage/manager"
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	DEFAULT_API_PORT = "8081"
	ENV_API_PORT     = "API_PORT"
)

type allSingletons struct {
	fileManager *manager.FileManager
}

func StartApi(fileManager *manager.FileManager) {
	singletons := allSingletons{
		fileManager: fileManager,
	}

	http.HandleFunc("/dir", CORS(func(w http.ResponseWriter, r *http.Request) {
		handleDirRequest(w, r, singletons)
	}))
	http.HandleFunc("/file", CORS(func(w http.ResponseWriter, r *http.Request) {
		handleFileRequest(w, r, singletons)
	}))

	apiPort := os.Getenv(ENV_API_PORT)
	if apiPort == "" {
		apiPort = DEFAULT_API_PORT
	}
	address := fmt.Sprintf(":%s", apiPort)

	log.Printf("Starting File Storage API on %s\n", address)
	log.Fatal(http.ListenAndServe(address, nil))
}

func handleDirRequest(w http.ResponseWriter, r *http.Request, singletons allSingletons) {
	log.Printf("Received request: %s /dir\n", r.Method)
	switch r.Method {
	case http.MethodGet:
		ListFilesHandler(w, r, singletons)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func ListFilesHandler(w http.ResponseWriter, r *http.Request, singletons allSingletons) {
	// TODO get path from query param
	path := ""

	files, err := singletons.fileManager.ListAllFiles(path)
	if err != nil {
		http.Error(w, "Unable to list files", http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(files)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Write(data)
}

func handleFileRequest(w http.ResponseWriter, r *http.Request, singletons allSingletons) {
	log.Printf("Received request: %s /file\n", r.Method)
	switch r.Method {
	case http.MethodGet:
		DownloadFileHandler(w, r, singletons)
	case http.MethodPut:
		UploadFileHandler(w, r, singletons)
	case http.MethodDelete:
		DeleteFileHandler(w, r, singletons)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func DownloadFileHandler(w http.ResponseWriter, r *http.Request, singletons allSingletons) {}

func UploadFileHandler(w http.ResponseWriter, r *http.Request, singletons allSingletons) {}

func DeleteFileHandler(w http.ResponseWriter, r *http.Request, singletons allSingletons) {}
