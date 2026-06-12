package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/SiracencoSerghei/devtrack-app/backend/internal/auth"
)

type contextKey string

const userKey contextKey = "user"

type AuthClaims struct {
	UserID string
	Email  string
	Roles  []string
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

			tokenStr := parts[1]

			claims, err := tm.ValidateToken(tokenStr)
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
	user, ok := ctx.Value(userKey).(AuthClaims)
	return user, ok
}