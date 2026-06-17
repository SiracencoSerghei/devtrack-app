package httpserver

import (
	"context"
	"log/slog"
	"net/http"
	"time"
)

type Server struct {
	*http.Server
}

func New(addr string, handler http.Handler) *Server {
	return &Server{
		Server: &http.Server{
			Addr:              addr,
			Handler:           handler,
			ReadTimeout:       10 * time.Second,
			WriteTimeout:      10 * time.Second,
			ReadHeaderTimeout: 3 * time.Second,
			IdleTimeout:       120 * time.Second,
			MaxHeaderBytes:    1 << 20,
		},
	}
}

func (s *Server) Start() error {
	slog.Info("server starting", "addr", s.Addr)
	return s.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	slog.Info("server shutting down")
	return s.Shutdown(ctx)
}