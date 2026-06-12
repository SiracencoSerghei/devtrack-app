package role

import "context"

type Repository interface {
	GetByName(ctx context.Context, name string) (string, error)
	AssignRole(ctx context.Context, userID, roleID string) error
	GetByUserID(ctx context.Context, userID string) ([]string, error)
}