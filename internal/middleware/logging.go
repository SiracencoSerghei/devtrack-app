package middleware

import (
    "log"
    "net"
    "net/http"
    "time"

    chimiddleware "github.com/go-chi/chi/v5/middleware"
)

type statusRecorder struct {
    http.ResponseWriter
    status int
}

func (r *statusRecorder) WriteHeader(status int) {
    r.status = status
    r.ResponseWriter.WriteHeader(status)
}

func Logging(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()

        rec := &statusRecorder{
            ResponseWriter: w,
            status:         http.StatusOK,
        }

        next.ServeHTTP(rec, r)

        ip, _, _ := net.SplitHostPort(r.RemoteAddr)
        requestID := chimiddleware.GetReqID(r.Context())

        log.Printf(
            "rid=%s method=%s path=%s status=%d duration=%s ip=%s ua=%s",
            requestID,
            r.Method,
            r.URL.Path,
            rec.status,
            time.Since(start),
            ip,
            r.UserAgent(),
        )
    })
}