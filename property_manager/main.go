package main

import (
	"log"
	"property_manager/api"
	"property_manager/manager"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	_, err := manager.GetPropertyManagerSingleton()
	if err != nil {
		log.Fatalln(err.Error())
	}

	// start api
	api.StartApi()
}
