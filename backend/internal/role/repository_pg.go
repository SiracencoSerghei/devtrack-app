package role

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresRepository struct {
	db *pgxpool.Pool
}

func NewPostgresRepository(db *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) GetByName(ctx context.Context, name string) (string, error) {
	var id string
	err := r.db.QueryRow(ctx, `SELECT id FROM roles WHERE name=$1`, name).Scan(&id)
	return id, err
}

func (r *PostgresRepository) AssignRole(ctx context.Context, userID, roleID string) error {
	_, err := r.db.Exec(ctx,
		`INSERT INTO user_roles (user_id, role_id)
		 VALUES ($1, $2)
		 ON CONFLICT DO NOTHING`,
		userID, roleID,
	)
	return err
}

func (r *PostgresRepository) GetByUserID(ctx context.Context, userID string) ([]string, error) {
	query := `
	SELECT r.name 
	FROM roles r
	JOIN user_roles ur ON ur.role_id = r.id
	WHERE ur.user_id = $1
	`
	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []string
	for rows.Next() {
		var roleName string
		if err := rows.Scan(&roleName); err != nil {
			return nil, err
		}
		roles = append(roles, roleName)
	}
	return roles, nil
}