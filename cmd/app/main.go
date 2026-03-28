package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"
	"net/http"

	"github.com/SiracencoSerghei/devtrack-app/pkg/httpserver"


	// "log"
	// "net/http"

	// "github.com/SiracencoSerghei/devtrack-app/internal/handler"
)

// func main() {
// 	mux := http.NewServeMux()

// 	mux.HandleFunc("/health", handler.HealthHandler)

// 	// server := &http.Server{
// 	// 	Addr: ":8080",
// 	// 	Handler: mux,
// 	// }

// 	// basic security configurations for the server anti-Slowloris attacks and other common issues:
// 	server := &http.Server{
// 	    Addr: ":8080",
// 	    Handler: mux,

// 	    ReadTimeout:       5 * time.Second,
// 	    ReadHeaderTimeout: 2 * time.Second,
// 	    WriteTimeout:      10 * time.Second,
// 	    IdleTimeout:       120 * time.Second,
// 	    MaxHeaderBytes:    1 << 20, // 1 MB

// 	}



// 	log.Println("Starting server on :8080")

// 	err:= server.ListenAndServe()
// 	if err != nil {
// 		log.Fatalf("Server failed to start: %v", err)
// 	}
// }

func main() {
	server := httpserver.New()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	errChan := make(chan error, 1)

	go func() {
		if err := server.Start(); err != nil && err != http.ErrServerClosed {
			errChan <- err
		}
		close(errChan)
	}()

	select {
	case <-stop:
		log.Println("Received OS signal, shutting down...")
	case err := <-errChan:
		if err != nil {
			log.Fatalf("Server crashed: %v", err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Stop(ctx); err != nil {
		log.Printf("Error during shutdown: %v", err)
	} else {
		log.Println("Server stopped gracefully")
	}
}
