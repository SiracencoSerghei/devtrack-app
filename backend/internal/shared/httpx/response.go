package httpx

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

// Інтерфейс для безпечних доменних помилок
type APIError interface {
	error
	APIError() (statusCode int, displayMessage string)
}

func WriteJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

func WriteError(w http.ResponseWriter, err error) {
	var apiErr APIError
	
	// Розумний розбір помилки, навіть якщо вона загорнута через fmt.Errorf("...: %w", err)
	if errors.As(err, &apiErr) {
		status, msg := apiErr.APIError()
		WriteJSON(w, status, ErrorResponse{Error: msg})
		return
	}

	// Якщо це сира помилка бази чи системи (дефолтний 500)
	slog.Error("internal server error occured", "err", err)
	WriteJSON(w, http.StatusInternalServerError, ErrorResponse{Error: "Internal server error. Please try again later."})
}