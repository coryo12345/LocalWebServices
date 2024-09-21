package main

import (
	"localwebservices/common"
	"localwebservices/propertystore"
	"localwebservices/taskqueue"
	"log"
	"net/http"
	"os"
	"os/signal"
)

type Service interface {
	Start(urlPrefix string)
}

func main() {
	// create services
	services := make(map[string]common.Service)
	services["propStore"] = propertystore.NewPropertyStore("properties")
	services["taskQueue"] = taskqueue.NewTaskQueue("queue")

	// start services
	handleServiceError(services["propStore"].Start("/api/propertystore"))
	handleServiceError(services["taskQueue"].Start("/api/taskqueue"))

	// handle shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		for _, service := range services {
			service.Shutdown()
		}
		os.Exit(1)
	}()

	http.ListenAndServe(":3000", nil)
}

func handleServiceError(err error) {
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}
