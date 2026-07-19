package application

import (
	"context"
	"strings"

	"github.com/SiracencoSerghei/devtrack-app/backend/internal/contexts/fleet/domain"
)

type DomainError struct {
	Status  int
	Message string
}

func (e DomainError) Error() string { return e.Message }
func (e DomainError) APIError() (int, string) { return e.Status, e.Message }

var (
	ErrLicenseRequired   = DomainError{Status: 400, Message: "il numero di patente è obbligatorio"}
	ErrDriverNotFound    = DomainError{Status: 404, Message: "profilo driver non trovato"}
	ErrInvalidStatus     = DomainError{Status: 400, Message: "stato driver non valido"}
	ErrDriverAlreadyExists = DomainError{Status: 409, Message: "un profilo driver per questo utente esiste già"}
)

type ServiceInterface interface {
	CreateProfile(ctx context.Context, userID, license, phone string) (domain.Driver, error)
	GetProfileByUserID(ctx context.Context, userID string) (domain.Driver, error)
	UpdateDriverStatus(ctx context.Context, id string, status string) error
}

type Service struct {
	repo domain.Repository
}

func NewService(repo domain.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateProfile(ctx context.Context, userID, license, phone string) (domain.Driver, error) {
	license = strings.TrimSpace(license)
	if license == "" {
		return domain.Driver{}, ErrLicenseRequired
	}

	d := domain.Driver{
		UserID:        userID,
		LicenseNumber: license,
		Phone:         strings.TrimSpace(phone),
		Status:        domain.StatusAvailable,
	}

	return s.repo.Create(ctx, d)
}

func (s *Service) GetProfileByUserID(ctx context.Context, userID string) (domain.Driver, error) {
	d, err := s.repo.GetByUserID(ctx, userID)
	if err != nil {
		return domain.Driver{}, ErrDriverNotFound
	}
	return d, nil
}

func (s *Service) UpdateDriverStatus(ctx context.Context, id string, status string) error {
	typedStatus := domain.DriverStatus(strings.ToUpper(strings.TrimSpace(status)))

	switch typedStatus {
	case domain.StatusAvailable, domain.StatusInTransit, domain.StatusOffDuty:
		return s.repo.UpdateStatus(ctx, id, string(typedStatus))
	default:
		return ErrInvalidStatus
	}
}