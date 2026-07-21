package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/SiracencoSerghei/devtrack-app/backend/internal/shared/auth"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/shared/httpx"
)

func Auth(tm *auth.TokenManager) func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			header := r.Header.Get("Authorization")

			if header == "" {
				httpx.WriteJSON(
					w,
					http.StatusUnauthorized,
					httpx.ErrorResponse{
						Error: "missing authorization header",
					},
				)
				return
			}


			fields := strings.Fields(header)

			if len(fields) != 2 || fields[0] != "Bearer" {

				httpx.WriteJSON(
					w,
					http.StatusUnauthorized,
					httpx.ErrorResponse{
						Error: "invalid authorization format",
					},
				)

				return
			}


			claims, err := tm.ValidateToken(fields[1])

			if err != nil {

				httpx.WriteJSON(
					w,
					http.StatusUnauthorized,
					httpx.ErrorResponse{
						Error: "invalid or expired token",
					},
				)

				return
			}


			user := auth.CurrentUser{
				UserID: claims.UserID,
				Email:  claims.Email,
				Roles:  claims.Roles,
			}


			ctx := context.WithValue(
				r.Context(),
				userKey,
				user,
			)


			next.ServeHTTP(
				w,
				r.WithContext(ctx),
			)
		})
	}
}