package domain

import "context"

type Repository interface {
	Create(ctx context.Context, u User, password string) (User, error)
	GetByEmail(ctx context.Context, email string) (User, error)
	GetAll(ctx context.Context) ([]User, error)
	Count(ctx context.Context) (int, error)
}

type RoleRepository interface {
	AssignRole(ctx context.Context, userID, roleID string) error
	GetByName(ctx context.Context, name string) (string, error)
	GetByUserID(ctx context.Context, userID string) ([]string, error)
}