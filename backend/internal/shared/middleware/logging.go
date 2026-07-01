package middleware

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"log/slog"
	"net/http"
	"time"
)

type responseWriterInterceptor struct {
	http.ResponseWriter
	statusCode int
}

func (w *responseWriterInterceptor) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func generateGlobalRequestID() string {
	bytes := make([]byte, 16)
	_, _ = rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Гарантуємо наявність Request ID на самому вході в систему
		reqID := r.Header.Get("X-Request-ID")
		if reqID == "" {
			reqID = generateGlobalRequestID()
		}
		w.Header().Set("X-Request-ID", reqID)

		// Насичуємо контекст trace-ключем БЕЗПЕЧНО для використання у всіх наступних шарах
		ctx := context.WithValue(r.Context(), traceKey, reqID)

		interceptor := &responseWriterInterceptor{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		// Передаємо оновлений контекст далі по ланцюжку
		next.ServeHTTP(interceptor, r.WithContext(ctx))

		slog.Info("http request processed",
			"request_id", reqID,
			"method", r.Method,
			"path", r.URL.Path,
			"status", interceptor.statusCode,
			"duration", time.Since(start).String(),
			"ip", r.RemoteAddr,
		)
	})
}