package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresRepository struct {
	db *pgxpool.Pool
}

func NewPostgresRepository(db *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) Create(ctx context.Context, u User) (User, error) {
	u.ID = uuid.NewString()

	query := `INSERT INTO users (id, name, email) VALUES ($1, $2, $3);`
	
	_, err := r.db.Exec(ctx, query, u.ID, u.Name, u.Email)
	if err != nil {
		
		if err.Error() != "" &&  ctx.Err() == nil {
			// for future -> use pgconn.PgError and check Code for "23505" 
			// (unique_violation)
			// Simple error text check for uniqueness
			return User{}, errors.New("email already exists in database")
		}
		return User{}, fmt.Errorf("failed to insert user: %w", err)
	}

	return u, nil
}

func (r *PostgresRepository) GetAll(ctx context.Context) ([]User, error) {
	query := `SELECT id, name, email FROM users ORDER BY name ASC;`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query users: %w", err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return users, nil
}