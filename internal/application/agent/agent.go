package agent

import (
	"log"
)

func StartAgent() {
	client := NewClient("http://localhost:8080")
	if client == nil {
		log.Fatal("Failed to create client")
	}

	worker := NewWorker(client)
	if worker == nil {
		log.Fatal("Failed to create worker")
	}

	go worker.Start()
	log.Println("Agent started")
}
