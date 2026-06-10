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

	if err := ensureSchema(ctx, pool); err != nil {
		return nil, fmt.Errorf("failed to run migrations: %w", err)
	}

	return pool, nil
}

func ensureSchema(ctx context.Context, pool *pgxpool.Pool) error {
	// 1. Tabella Utenti
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

	// 2. Tabella Ruoli
	roleQuery := `
	CREATE TABLE IF NOT EXISTS roles (
		id UUID PRIMARY KEY,
		name VARCHAR(50) UNIQUE NOT NULL
	);`
	if _, err := pool.Exec(ctx, roleQuery); err != nil {
		return fmt.Errorf("failed to create roles table: %w", err)
	}

	// Inserimento ruolo di default per il driver
	_, _ = pool.Exec(ctx, "INSERT INTO roles (id, name) VALUES ('6a2f72ec-5536-407b-bc83-49d799299446', 'driver') ON CONFLICT DO NOTHING")

	// 3. Tabella Relazione Utenti-Ruoli
	userRolesQuery := `
	CREATE TABLE IF NOT EXISTS user_roles (
		user_id UUID REFERENCES users(id) ON DELETE CASCADE,
		role_id UUID REFERENCES roles(id) ON DELETE CASCADE,
		PRIMARY KEY (user_id, role_id)
	);`
	if _, err := pool.Exec(ctx, userRolesQuery); err != nil {
		return fmt.Errorf("failed to create user_roles table: %w", err)
	}

	// 4. Tabella Sessioni
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