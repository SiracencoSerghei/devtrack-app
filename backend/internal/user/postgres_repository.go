// backend/internal/user/postgres_repository.go
package user

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/auth"
)

type PostgresRepository struct {
	db *pgxpool.Pool
}

func NewPostgresRepository(db *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) Create(ctx context.Context, u User, password string) (User, error) {
	u.ID = uuid.NewString()
	
	hashedPassword, err := auth.HashPassword(password)
	if err != nil {
		return User{}, fmt.Errorf("failed to hash password: %w", err)
	}
	u.PasswordHash = hashedPassword

	query := `INSERT INTO users (id, name, email, password_hash) VALUES ($1, $2, $3, $4);`
	_, err = r.db.Exec(ctx, query, u.ID, u.Name, u.Email, u.PasswordHash)
	if err != nil {
		return User{}, fmt.Errorf("failed to insert user: %w", err)
	}

	return u, nil
}

func (r *PostgresRepository) GetByEmail(ctx context.Context, email string) (User, error) {
	query := `SELECT id, name, email, password_hash FROM users WHERE email = $1;`
	
	var u User
	err := r.db.QueryRow(ctx, query, email).Scan(&u.ID, &u.Name, &u.Email, &u.PasswordHash)
	if err != nil {
		return User{}, err
	}
	return u, nil
}

func (r *PostgresRepository) GetAll(ctx context.Context) ([]User, error) {
	query := `SELECT id, name, email FROM users ORDER BY name ASC;`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}