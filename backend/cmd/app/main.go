package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/SiracencoSerghei/devtrack-app/backend/internal/bootstrap"
)

func main() {
	log.Println("[START] Inizializzazione del server DevTrack...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Avvia il bootstrap dell'intera applicazione
	app, err := bootstrap.Initialize(ctx)
	if err != nil {
		log.Fatalf("[FATAL] Errore durante l'inizializzazione: %v", err)
	}
	defer app.Close()

	log.Println("[READY] Server in ascolto sulla porta :8080 🚀")
	if err := app.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("[FATAL] Errore durante l'esecuzione del server: %v", err)
	}
}