package infrastructure

import (
	"context"
	"time"

	"github.com/SiracencoSerghei/devtrack-app/backend/internal/contexts/logistics/domain"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresRepository struct {
	db *pgxpool.Pool
}

func NewPostgresRepository(db *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) Create(ctx context.Context, o domain.Order) (domain.Order, error) {
	o.ID = uuid.NewString()
	o.CreatedAt = time.Now()
	if o.Status == "" {
		o.Status = "PENDING"
	}

	query := `
	INSERT INTO orders (id, customer_id, pickup_address, delivery_address, status, created_at)
	VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := r.db.Exec(ctx, query, o.ID, o.CustomerID, o.PickupAddress, o.DeliveryAddress, o.Status, o.CreatedAt)
	if err != nil {
		return domain.Order{}, err
	}

	return o, nil
}

func (r *PostgresRepository) GetByID(ctx context.Context, id string) (domain.Order, error) {
	var o domain.Order
	query := `
	SELECT id, customer_id, pickup_address, delivery_address, status, created_at
	FROM orders
	WHERE id = $1
	`
	err := r.db.QueryRow(ctx, query, id).Scan(
		&o.ID, &o.CustomerID, &o.PickupAddress, &o.DeliveryAddress, &o.Status, &o.CreatedAt,
	)
	return o, err
}