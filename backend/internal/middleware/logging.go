package middleware

import (
	"log"
	"net/http"
	"time"
)

// responseWriterInterceptor ci serve per catturare lo status code (es. 200, 400, 500)
type responseWriterInterceptor struct {
	http.ResponseWriter
	statusCode int
}

func (w *responseWriterInterceptor) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		
		interceptor := &responseWriterInterceptor{ResponseWriter: w, statusCode: http.StatusOK}
		
		next.ServeHTTP(interceptor, r)
		
		log.Printf(
			"[HTTP] %s %s | Status: %d | Duration: %v | IP: %s",
			r.Method,
			r.URL.Path,
			interceptor.statusCode,
			time.Since(start),
			r.RemoteAddr,
		)
	})
}