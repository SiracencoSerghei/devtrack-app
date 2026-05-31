package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/SiracencoSerghei/devtrack-app/backend/internal/auth"
)

type contextKey string
const UserIDKey contextKey = "userID"

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Accesso negato: Token mancante", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Accesso negato: Formato token non valido", http.StatusUnauthorized)
			return
		}

		tokenString := parts[1]
		claims, err := auth.ValidateToken(tokenString) // Assicurati che nel tuo pacchetto auth ci sia un metodo per validare e ritornare i claims
		if err != nil {
			http.Error(w, "Accesso negato: Token scaduto o non valido", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, claims.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}