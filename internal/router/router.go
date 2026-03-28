package router

import (
	"net/http"

	"github.com/SiracencoSerghei/devtrack-app/internal/handler"
	"github.com/SiracencoSerghei/devtrack-app/internal/middleware"
)

func New(userHandler *handler.UserHandler) http.Handler {

	mux := http.NewServeMux()

	mux.HandleFunc("/health", handler.HealthHandler)

	mux.Handle(
		"/user",
		middleware.Logging(http.HandlerFunc(userHandler.GetUser)),
	)

	return mux
}