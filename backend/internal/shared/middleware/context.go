package middleware

import (
	"context"
	"crypto/rand"
	"encoding/hex"

	"github.com/SiracencoSerghei/devtrack-app/backend/internal/shared/auth"
)

type contextKey string

const (
	userKey  contextKey = "current_user"
	traceKey contextKey = "trace_id"
)

func generateRequestID() string {

	bytes := make([]byte, 16)

	_, err := rand.Read(bytes)
	if err != nil {
		return "unknown"
	}

	return hex.EncodeToString(bytes)
}


func GetTraceID(ctx context.Context) (string, bool) {

	value := ctx.Value(traceKey)

	id, ok := value.(string)

	return id, ok
}


func GetUser(ctx context.Context) (auth.CurrentUser, bool) {

	value := ctx.Value(userKey)

	user, ok := value.(auth.CurrentUser)

	return user, ok
}