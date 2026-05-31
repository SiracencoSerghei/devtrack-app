package delivery

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/SiracencoSerghei/devtrack-app/backend/internal/user/usecase"
)

type HTTPHandler struct {
	useCase *usecase.UserUseCase
}

func NewHTTPHandler(uc *usecase.UserUseCase) *HTTPHandler {
	return &HTTPHandler{useCase: uc}
}

func (h *HTTPHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var input usecase.SignUpInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		h.respondWithError(w, http.StatusBadRequest, "Payload JSON non valido")
		return
	}

	user, err := h.useCase.SignUp(r.Context(), input)
	if err != nil {
		if errors.Is(err, usecase.ErrValidationFailed) {
			h.respondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		h.respondWithError(w, http.StatusConflict, err.Error())
		return
	}

	h.respondWithJSON(w, http.StatusCreated, user)
}

func (h *HTTPHandler) Login(w http.ResponseWriter, r *http.Request) {
	var input usecase.LoginInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		h.respondWithError(w, http.StatusBadRequest, "Payload JSON non valido")
		return
	}

	accessToken, refreshToken, user, err := h.useCase.Login(r.Context(), input)
	if err != nil {
		if errors.Is(err, usecase.ErrInvalidAuth) {
			h.respondWithError(w, http.StatusUnauthorized, "Inidirizzo email o password errati")
			return
		}
		h.respondWithError(w, http.StatusInternalServerError, "Errore interno durante il login")
		return
	}

	h.respondWithJSON(w, http.StatusOK, map[string]interface{}{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"user":          user,
	})
}

func (h *HTTPHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := h.useCase.GetAll(r.Context())
	if err != nil {
		h.respondWithError(w, http.StatusInternalServerError, "Impossibile recuperare gli utenti")
		return
	}
	h.respondWithJSON(w, http.StatusOK, users)
}

func (h *HTTPHandler) respondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func (h *HTTPHandler) respondWithError(w http.ResponseWriter, status int, message string) {
	h.respondWithJSON(w, status, map[string]string{"error": message})
}