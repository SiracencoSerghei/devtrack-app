package driver

import "context"

type Repository interface {
	Create(ctx context.Context, d Driver) (Driver, error)
	GetByUserID(ctx context.Context, userID string) (Driver, error)
	UpdateStatus(ctx context.Context, id string, status string) error
}