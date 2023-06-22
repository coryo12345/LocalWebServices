package main

import (
	"queue/api"
	"queue/config"
	"queue/manager"

	"github.com/joho/godotenv"
)

func main() {
	// not capturing error - it's fine if .env doesn't exist
	godotenv.Load(".env")

	manager := manager.NewQueueManager()

	// config queues from file
	config.InitConfig(manager)

	qapi := api.NewResourceAPI(manager)
	qapi.StartApi()
}
