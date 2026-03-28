package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/SiracencoSerghei/devtrack-app/internal/handler"
	"github.com/SiracencoSerghei/devtrack-app/internal/router"
	"github.com/SiracencoSerghei/devtrack-app/internal/service"
	"github.com/SiracencoSerghei/devtrack-app/pkg/httpserver"
)

func main() {

	userService := service.NewUserService()
	userHandler := handler.NewUserHandler(userService)

	r := router.New(userHandler)

	server := httpserver.New(":8080", r)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	errChan := make(chan error, 1)

	go func() {
		if err := server.Start(); err != nil {
			errChan <- err
		}
	}()

	select {

	case <-stop:
		log.Println("Shutdown signal received")

	case err := <-errChan:
		log.Fatalf("Server crashed: %v", err)

	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Stop(ctx); err != nil {
		log.Printf("Shutdown error: %v", err)
	}

	log.Println("Server stopped")
}