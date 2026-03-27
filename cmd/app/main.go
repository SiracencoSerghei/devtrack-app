package main

import (
	"log"
	"net/http"

	"github.com/SiracencoSerghei/devtrack-app/internal/handler"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", handler.HealthHandler)

	server := &http.Server{
		Addr: ":8080",
		Handler: mux,
	}

	log.Println("Starting server on :8080")

	err:= server.ListenAndServe()
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}