package user

import "context"

type Repository interface {
	Create(ctx context.Context, u User, password string) (User, error)
	GetByEmail(ctx context.Context, email string) (User, error)
	GetAll(ctx context.Context) ([]User, error)
}