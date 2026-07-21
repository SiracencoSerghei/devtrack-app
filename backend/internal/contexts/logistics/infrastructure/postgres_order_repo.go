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

func (r *PostgresRepository) Search(ctx context.Context, queryStr string) ([]domain.Order, error) {
	// Пошук за будь-яким збігом у pickup_address, delivery_address, status або ID
	query := `
	SELECT id, customer_id, pickup_address, delivery_address, status, created_at
	FROM orders
	WHERE pickup_address ILIKE $1 
	   OR delivery_address ILIKE $1 
	   OR status ILIKE $1 
	   OR id::text ILIKE $1
	ORDER BY created_at DESC
	`
	searchTerm := "%" + queryStr + "%"
	rows, err := r.db.Query(ctx, query, searchTerm)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []domain.Order
	for rows.Next() {
		var o domain.Order
		if err := rows.Scan(&o.ID, &o.CustomerID, &o.PickupAddress, &o.DeliveryAddress, &o.Status, &o.CreatedAt); err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}

	return orders, nil
}