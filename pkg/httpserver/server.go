package httpserver

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/SiracencoSerghei/devtrack-app/internal/handler"
	"github.com/SiracencoSerghei/devtrack-app/internal/middleware"
	"github.com/SiracencoSerghei/devtrack-app/internal/service"
)

type Server struct {
	httpServer *http.Server
}

func New() *Server {
	userService := service.NewUserService()
	userHandler := handler.NewUserHandler(userService)

	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	mux.Handle("/user", middleware.Timeout(middleware.Logging(http.HandlerFunc(userHandler.GetUser))))

	server := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       120 * time.Second,
		ReadHeaderTimeout: 3 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}

	return &Server{httpServer: server}
}

func (s *Server) Start() error {
	log.Printf("Starting server on %s\n", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	log.Println("Shutting down server...")
	return s.httpServer.Shutdown(ctx)
}