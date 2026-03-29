package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/SiracencoSerghei/devtrack-app/internal/handler"
	"github.com/SiracencoSerghei/devtrack-app/internal/router"
	"github.com/SiracencoSerghei/devtrack-app/internal/service"
	"github.com/SiracencoSerghei/devtrack-app/pkg/httpserver"
)

func main() {

	userService := service.NewUserService()

	userHandler := handler.NewUserHandler(userService)
	healthHandler := handler.NewHealthHandler()

	r := router.New()

	router.RegisterHealthRoutes(r, healthHandler)
	router.RegisterUserRoutes(r, userHandler)

	server := httpserver.New(":8080", r)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	errChan := make(chan error, 1)

	go func() {
		errChan <- server.Start()
	}()

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