package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/SiracencoSerghei/devtrack-app/backend/internal/db"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/health"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/router"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/user"
)

func main() {
	log.Println("[START] Inizializzazione del server DevTrack...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pool, err := db.Connect(ctx)
	if err != nil {
		log.Fatalf("[FATAL] Impossibile connettersi al database: %v", err)
	}
	defer pool.Close()

	userRepo := user.NewPostgresRepository(pool)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)

	healthHandler := health.NewHandler()

	appRouter := router.New(userHandler, healthHandler)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      appRouter,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	log.Println("[READY] Server in ascolto sulla porta :8080 🚀")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("[FATAL] Errore durante l'esecuzione del server: %v", err)
	}
}