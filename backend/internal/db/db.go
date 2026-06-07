package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(ctx context.Context) (*pgxpool.Pool, error) {

	host := os.Getenv("DB_HOST")
	if host == "" { host = "localhost" }
	
	user := os.Getenv("DB_USER")
	if user == "" { user = "postgres" }
	
	password := os.Getenv("DB_PASSWORD")
	if password == "" { password = "postgres" }
	
	dbname := os.Getenv("DB_NAME")
	if dbname == "" { dbname = "devtrack_db" }

	dsn := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable", user, password, host, dbname)

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("unable to ping database: %w", err)
	}
// Автоматично створюємо/оновлюємо таблиці (наш Clean-шар)
	if err := ensureSchema(ctx, pool); err != nil {
		return nil, fmt.Errorf("failed to run migrations: %w", err)
	}

	return pool, nil
}

func ensureSchema(ctx context.Context, pool *pgxpool.Pool) error {
	userQuery := `
	CREATE TABLE IF NOT EXISTS users (
		id UUID PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) UNIQUE NOT NULL,
		password_hash VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`
	
	if _, err := pool.Exec(ctx, userQuery); err != nil {
		return fmt.Errorf("failed to create users table: %w", err)
	}

	sessionQuery := `
	CREATE TABLE IF NOT EXISTS user_sessions (
		id UUID PRIMARY KEY,
		user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
		refresh_token VARCHAR(255) UNIQUE NOT NULL,
		expires_at TIMESTAMP NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	if _, err := pool.Exec(ctx, sessionQuery); err != nil {
		return fmt.Errorf("failed to create user_sessions table: %w", err)
	}

	return nil
}