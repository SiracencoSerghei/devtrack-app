package service

import (
	"context"
	"time"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) GetUser(ctx context.Context) (string, error) {

	select {

	case <-time.After(2 * time.Second):
		return "user data", nil

	case <-ctx.Done():
		return "", ctx.Err()

	}
}