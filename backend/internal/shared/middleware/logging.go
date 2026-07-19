package middleware

import (
	"context"
	"log/slog"
	"net/http"
	"time"
)

type responseWriterInterceptor struct {
	http.ResponseWriter
	statusCode int
}

func (w *responseWriterInterceptor) WriteHeader(statusCode int) {
	if w.statusCode != 0 {
		return
	}

	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *responseWriterInterceptor) Write(data []byte) (int, error) {
	if w.statusCode == 0 {
		w.WriteHeader(http.StatusOK)
	}

	return w.ResponseWriter.Write(data)
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		reqID := r.Header.Get("X-Request-ID")
		if reqID == "" {
			reqID = generateRequestID()
		}

		w.Header().Set("X-Request-ID", reqID)

		ctx := context.WithValue(
			r.Context(),
			traceKey,
			reqID,
		)

		interceptor := &responseWriterInterceptor{
			ResponseWriter: w,
		}

		next.ServeHTTP(interceptor, r.WithContext(ctx))

		status := interceptor.statusCode
		if status == 0 {
			status = http.StatusOK
		}

		slog.Info(
			"http request processed",
			"request_id", reqID,
			"method", r.Method,
			"path", r.URL.Path,
			"status", status,
			"duration", time.Since(start),
			"ip", r.RemoteAddr,
		)
	})
}