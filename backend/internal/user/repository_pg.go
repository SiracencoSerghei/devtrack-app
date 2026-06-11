package user

import (
	"context"
	"errors"
	"time"

	"github.com/SiracencoSerghei/devtrack-app/backend/internal/auth"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresRepository struct {
	db *pgxpool.Pool
}

func NewPostgresRepository(db *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) Create(ctx context.Context, u User, password string) (User, error) {
	u.ID = uuid.NewString()
	u.CreatedAt = time.Now()

	hash, err := auth.HashPassword(password)
	if err != nil {
		return User{}, err
	}
	u.PasswordHash = hash

	query := `INSERT INTO users (id, name, email, password_hash, created_at) VALUES ($1, $2, $3, $4, $5)`
	_, err = r.db.Exec(ctx, query, u.ID, u.Name, u.Email, u.PasswordHash, u.CreatedAt)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return User{}, ErrEmailAlreadyExists
			}
		}
		return User{}, err
	}

	return u, nil
}

func (r *PostgresRepository) GetByEmail(ctx context.Context, email string) (User, error) {
	var u User
	// Utilizziamo ARRAY_AGG per estrarre tutti i ruoli associati in una stringa/array nativo Postgres
	query := `
	SELECT u.id, u.name, u.email, u.password_hash, u.created_at, COALESCE(array_agg(r.name) FILTER (WHERE r.name IS NOT NULL), '{}')
	FROM users u
	LEFT JOIN user_roles ur ON ur.user_id = u.id
	LEFT JOIN roles r ON r.id = ur.role_id
	WHERE u.email = $1
	GROUP BY u.id
	`
	err := r.db.QueryRow(ctx, query, email).Scan(
		&u.ID, &u.Name, &u.Email, &u.PasswordHash, &u.CreatedAt, &u.Roles,
	)
	if err != nil {
		return User{}, err
	}
	return u, nil
}

func (r *PostgresRepository) GetAll(ctx context.Context) ([]User, error) {
	query := `
	SELECT u.id, u.name, u.email, u.created_at, COALESCE(array_agg(r.name) FILTER (WHERE r.name IS NOT NULL), '{}')
	FROM users u
	LEFT JOIN user_roles ur ON ur.user_id = u.id
	LEFT JOIN roles r ON r.id = ur.role_id
	GROUP BY u.id
	ORDER BY u.name
	`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt, &u.Roles)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}
func (r *PostgresRepository) Count(ctx context.Context) (int, error) {
	var count int
	query := `SELECT COUNT(*) FROM users`
	err := r.db.QueryRow(ctx, query).Scan(&count)
	return count, err
}