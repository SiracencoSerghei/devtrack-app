package user

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type RolePostgresRepository struct {
	db *pgxpool.Pool
}

func NewRolePostgresRepository(db *pgxpool.Pool) *RolePostgresRepository {
	return &RolePostgresRepository{db: db}
}

func (r *RolePostgresRepository) AssignRole(ctx context.Context, userID, roleID string) error {
	_, err := r.db.Exec(ctx,
		`INSERT INTO user_roles (user_id, role_id)
		 VALUES ($1,$2)
		 ON CONFLICT DO NOTHING`,
		userID, roleID,
	)
	return err
}