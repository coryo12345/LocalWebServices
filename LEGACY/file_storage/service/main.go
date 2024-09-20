package main

import (
	"file_storage/api"
	"file_storage/manager"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	ENV_STORAGE_DIR     = "STORAGE_DIRECTORY"
	DEFAULT_STORAGE_DIR = "./tmp"
)

func main() {
	godotenv.Load(".env")

	storageDir := os.Getenv(ENV_STORAGE_DIR)
	if storageDir == "" {
		storageDir = DEFAULT_STORAGE_DIR
	}

	fileManager, err := manager.NewFileManager(storageDir)
	if err != nil {
		log.Fatalf("unable to create filemanager: %s\n", err.Error())
	}

	// start api
	api.StartApi(fileManager)
}
