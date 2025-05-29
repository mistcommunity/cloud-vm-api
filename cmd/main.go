package main

import (
	"log"
	"net/http"

	"github.com/mistcommunity/cloud-vm-api/internal/api"
)

func main() {
	router := api.NewRouter()
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
