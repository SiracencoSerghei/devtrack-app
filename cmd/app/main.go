package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/SiracencoSerghei/devtrack-app/internal/health"
	"github.com/SiracencoSerghei/devtrack-app/internal/router"
	"github.com/SiracencoSerghei/devtrack-app/internal/user"
	"github.com/SiracencoSerghei/devtrack-app/pkg/httpserver"
)

func main() {
	// Репозиторій і сервіс
	repo := user.NewRepository()
	svc := user.NewService(repo)

	// Хендлери
	uHandler := user.NewHandler(svc)
	hHandler := health.NewHandler()

	// Роутер chi
	r := router.New()
	router.RegisterRootRoute(r)
	router.RegisterUserRoutes(r, uHandler)
	router.RegisterHealthRoutes(r, hHandler)

	// HTTP сервер
	server := httpserver.New(":8080", r)

	// Канал для graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	errChan := make(chan error, 1)
	go func() {
		errChan <- server.Start()
	}()

	select {
	case err := <-errChan:
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("server error: %v", err)
		}
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