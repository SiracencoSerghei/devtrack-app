package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/SiracencoSerghei/devtrack-app/backend/internal/db"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/health"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/router"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/user"
	"github.com/SiracencoSerghei/devtrack-app/backend/pkg/httpserver"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://postgres:password@localhost:5432/devtrack?sslmode=disable"
	}

	dbPool, err := db.NewPool(ctx, dbURL)
	if err != nil {
		log.Fatalf("database connection failed: %v", err)
	}
	defer dbPool.Close()
	log.Println("Connected to PostgreSQL successfully")

	repo := user.NewPostgresRepository(dbPool)
	svc := user.NewService(repo)

	userHandler := user.NewHandler(svc)
	healthHandler := health.NewHandler()

	r := router.New(userHandler, healthHandler)
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server := httpserver.New(":"+port, r)

	errChan := make(chan error, 1)
	go func() {
		errChan <- server.Start()
	}()

	log.Printf("Server is running on port %s", port)

	<-ctx.Done()
	log.Println("Shutdown signal received")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Stop(shutdownCtx); err != nil {
		log.Printf("Server shutdown error: %v", err)
	}

	log.Println("Server stopped cleanly")
}