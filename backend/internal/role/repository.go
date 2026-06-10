package role

import "context"

type Repository interface {
	GetByName(ctx context.Context, name string) (string, error)
}