package repository

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/SiracencoSerghei/devtrack-app/backend/internal/domain"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrEmailAlreadyExists = errors.New("email già registrata nel sistema")
	ErrInternalDatabase   = errors.New("errore interno del database")
	ErrSessionNotFound    = errors.New("sessione o refresh token non valido")
)

type PostgresRepository struct {
	db *pgxpool.Pool
}

func NewPostgresRepository(db *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) Create(ctx context.Context, u domain.User, passwordHash string) (domain.User, error) {
	u.ID = uuid.NewString()
	u.CreatedAt = time.Now()
	u.PasswordHash = passwordHash

	query := `INSERT INTO users (id, name, email, password_hash, created_at) VALUES ($1, $2, $3, $4, $5);`
	_, err := r.db.Exec(ctx, query, u.ID, u.Name, u.Email, u.PasswordHash, u.CreatedAt)
	
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return domain.User{}, ErrEmailAlreadyExists
		}
		log.Printf("[DATABASE ERROR] Failed to insert user: %v", err)
		return domain.User{}, ErrInternalDatabase
	}
	return u, nil
}

func (r *PostgresRepository) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	query := `SELECT id, name, email, password_hash, created_at FROM users WHERE email = $1;`
	var u domain.User
	err := r.db.QueryRow(ctx, query, email).Scan(&u.ID, &u.Name, &u.Email, &u.PasswordHash, &u.CreatedAt)
	if err != nil {
		return domain.User{}, err
	}
	return u, nil
}

func (r *PostgresRepository) GetAll(ctx context.Context) ([]domain.User, error) {
	query := `SELECT id, name, email, created_at FROM users ORDER BY name ASC;`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		log.Printf("[DATABASE ERROR] GetAll users failed: %v", err)
		return nil, ErrInternalDatabase
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var u domain.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt); err != nil {
			return nil, ErrInternalDatabase
		}
		users = append(users, u)
	}
	return users, nil
}

func (r *PostgresRepository) CreateSession(ctx context.Context, s domain.Session) error {
	s.ID = uuid.NewString()
	s.CreatedAt = time.Now()

	query := `INSERT INTO user_sessions (id, user_id, refresh_token, expires_at, created_at) VALUES ($1, $2, $3, $4, $5);`
	_, err := r.db.Exec(ctx, query, s.ID, s.UserID, s.RefreshToken, s.ExpiresAt, s.CreatedAt)
	if err != nil {
		log.Printf("[DATABASE ERROR] CreateSession failed: %v", err)
		return ErrInternalDatabase
	}
	return nil
}

func (r *PostgresRepository) DeleteSession(ctx context.Context, refreshToken string) error {
	query := `DELETE FROM user_sessions WHERE refresh_token = $1;`
	cmdTag, err := r.db.Exec(ctx, query, refreshToken)
	if err != nil {
		log.Printf("[DATABASE ERROR] DeleteSession failed: %v", err)
		return ErrInternalDatabase
	}
	if cmdTag.RowsAffected() == 0 {
		return ErrSessionNotFound
	}
	return nil
}