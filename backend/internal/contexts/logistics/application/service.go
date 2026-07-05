package application

import (
	"context"
	"strings"

	"github.com/SiracencoSerghei/devtrack-app/backend/internal/contexts/logistics/domain"
)

type DomainError struct {
	Status  int
	Message string
}

func (e DomainError) Error() string   { return e.Message }
func (e DomainError) APIError() (int, string) { return e.Status, e.Message }

var (
	ErrAddressesRequired = DomainError{Status: 400, Message: "pickup and delivery addresses are required"}
	ErrOrderNotFound     = DomainError{Status: 404, Message: "order not found"}
)

type Repository interface {
	Create(ctx context.Context, o domain.Order) (domain.Order, error)
	GetByID(ctx context.Context, id string) (domain.Order, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateOrder(ctx context.Context, customerID, pickup, delivery string) (domain.Order, error) {
	pickup = strings.TrimSpace(pickup)
	delivery = strings.TrimSpace(delivery)

	if pickup == "" || delivery == "" {
		return domain.Order{}, ErrAddressesRequired
	}

	o := domain.Order{
		CustomerID:      customerID,
		PickupAddress:   pickup,
		DeliveryAddress: delivery,
		Status:          "PENDING",
	}

	return s.repo.Create(ctx, o)
}

func (s *Service) GetOrder(ctx context.Context, id string) (domain.Order, error) {
	o, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return domain.Order{}, ErrOrderNotFound
	}
	return o, nil
}