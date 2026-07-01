package infrastructure

import (
	"context"
	"errors"
	"time"

	"github.com/SiracencoSerghei/devtrack-app/backend/internal/contexts/logistics/application"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/contexts/logistics/domain"
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

func (r *PostgresRepository) Create(ctx context.Context, d domain.Driver) (domain.Driver, error) {
	d.ID = uuid.NewString()
	d.UpdatedAt = time.Now()
	if d.Status == "" {
		d.Status = "AVAILABLE"
	}

	query := `
	INSERT INTO drivers (id, user_id, license_number, phone, status, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := r.db.Exec(ctx, query, d.ID, d.UserID, d.LicenseNumber, d.Phone, d.Status, d.UpdatedAt)
	if err != nil {
		var pgErr *pgconn.PgError
		// Перевірка на унікальність унікального індексу drivers_user_id_key
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return domain.Driver{}, application.ErrDriverAlreadyExists
		}
		return domain.Driver{}, err
	}

	return d, nil
}

func (r *PostgresRepository) GetByUserID(ctx context.Context, userID string) (domain.Driver, error) {
	var d domain.Driver
	query := `
	SELECT id, user_id, license_number, phone, status, updated_at
	FROM drivers
	WHERE user_id = $1
	`
	err := r.db.QueryRow(ctx, query, userID).Scan(
		&d.ID, &d.UserID, &d.LicenseNumber, &d.Phone, &d.Status, &d.UpdatedAt,
	)
	return d, err
}

func (r *PostgresRepository) UpdateStatus(ctx context.Context, id string, status string) error {
	query := `UPDATE drivers SET status = $1, updated_at = NOW() WHERE id = $2`
	_, err := r.db.Exec(ctx, query, status, id)
	return err
}