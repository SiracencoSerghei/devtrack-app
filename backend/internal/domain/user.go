package domain

import (
	"context"
	"time"
)

type User struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"created_at"`
}

type Session struct {
	ID           string    `json:"id"`
	UserID       string    `json:"user_id"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresAt    time.Time `json:"expires_at"`
	CreatedAt    time.Time `json:"created_at"`
}

type UserRepository interface {
	Create(ctx context.Context, u User, password string) (User, error)
	GetByEmail(ctx context.Context, email string) (User, error)
	GetAll(ctx context.Context) ([]User, error)
	
	CreateSession(ctx context.Context, session Session) error
	DeleteSession(ctx context.Context, refreshToken string) error
}

type UserUseCase interface {
	SignUp(ctx context.Context, name, email, password string) (User, error)
	Login(ctx context.Context, email, password string) (string, string, User, error)
	RefreshToken(ctx context.Context, refreshToken string) (string, string, error)
	GetAll(ctx context.Context) ([]User, error)
}