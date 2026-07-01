package middleware

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"strings"

	"github.com/SiracencoSerghei/devtrack-app/backend/internal/shared/auth"
)

// Унікальний тип-структура для ключів контексту
type ctxUserKey struct{}
type ctxTraceKey struct{}

var (
	userKey  = ctxUserKey{}
	traceKey = ctxTraceKey{}
)

type AuthClaims struct {
	UserID string
	Email  string
	Roles  []string
}

func generateRequestID() string {
	bytes := make([]byte, 16)
	_, _ = rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

func Auth(tm *auth.TokenManager) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")
			if header == "" {
				http.Error(w, "missing authorization header", http.StatusUnauthorized)
				return
			}

			parts := strings.SplitN(header, " ", 2)
			if len(parts) != 2 || parts[0] != "Bearer" {
				http.Error(w, "invalid authorization format", http.StatusUnauthorized)
				return
			}

			claims, err := tm.ValidateToken(parts[1])
			if err != nil {
				http.Error(w, "invalid or expired token", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), userKey, AuthClaims{
				UserID: claims.UserID,
				Email:  claims.Email,
				Roles:  claims.Roles,
			})

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetUser(ctx context.Context) (AuthClaims, bool) {
	val := ctx.Value(userKey)
	if val == nil {
		return AuthClaims{}, false
	}

	claims, ok := val.(AuthClaims)
	return claims, ok
}