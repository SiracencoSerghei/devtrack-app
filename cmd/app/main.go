package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/SiracencoSerghei/devtrack-app/internal/router"
	"github.com/SiracencoSerghei/devtrack-app/internal/user"
	"github.com/SiracencoSerghei/devtrack-app/internal/handler"
	"github.com/SiracencoSerghei/devtrack-app/pkg/httpserver"
)

func main() {

	// --- user domain setup ---
	repo := user.NewRepository()
	svc := user.NewService(repo)
	h := user.NewHandler(svc)

	// --- health handler ---
	healthHandler := handler.NewHealthHandler()

	// --- router ---
	r := router.New()
	router.RegisterRootRoute(r)        
	router.RegisterUserRoutes(r, h)
	router.RegisterHealthRoutes(r, healthHandler)

	// --- HTTP server ---
	server := httpserver.New(":8080", r)

	// --- graceful shutdown ---
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	errChan := make(chan error, 1)
	go func() { errChan <- server.Start() }()

	select {
	case err := <-errChan:
		log.Fatalf("server crashed: %v", err)
	case <-stop:
		log.Println("shutdown signal received")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Stop(ctx); err != nil {
		log.Printf("shutdown error: %v", err)
	}

	log.Println("server stopped")
}