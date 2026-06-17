package driver

import (
	"context"
	"strings"
)

// Робимо стійкі типізовані Енапи замість сирих рядків
type Status string

const (
	StatusAvailable Status = "AVAILABLE"
	StatusInTransit Status = "IN_TRANSIT"
	StatusOffDuty   Status = "OFF_DUTY"
)

type DomainError struct {
	Status  int
	Message string
}

func (e DomainError) Error() string { return e.Message }
func (e DomainError) APIError() (int, string) { return e.Status, e.Message }

var (
	ErrLicenseRequired = DomainError{Status: 400, Message: "il numero di patente è obbligatorio"}
	ErrDriverNotFound  = DomainError{Status: 404, Message: "profilo driver non trovato"}
	ErrInvalidStatus   = DomainError{Status: 400, Message: "stato driver non valido"}
)

type ServiceInterface interface {
	CreateProfile(ctx context.Context, userID, license, phone string) (Driver, error)
	GetProfileByUserID(ctx context.Context, userID string) (Driver, error)
	UpdateDriverStatus(ctx context.Context, id string, status string) error
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateProfile(ctx context.Context, userID, license, phone string) (Driver, error) {
	license = strings.TrimSpace(license)
	if license == "" {
		return Driver{}, ErrLicenseRequired
	}

	d := Driver{
		UserID:        userID,
		LicenseNumber: license,
		Phone:         strings.TrimSpace(phone),
		Status:        string(StatusAvailable),
	}

	return s.repo.Create(ctx, d)
}

func (s *Service) GetProfileByUserID(ctx context.Context, userID string) (Driver, error) {
	d, err := s.repo.GetByUserID(ctx, userID)
	if err != nil {
		return Driver{}, ErrDriverNotFound
	}
	return d, nil
}

func (s *Service) UpdateDriverStatus(ctx context.Context, id string, status string) error {
	typedStatus := Status(strings.ToUpper(strings.TrimSpace(status)))

	// Елегантна та безпечна перевірка енапу
	switch typedStatus {
	case StatusAvailable, StatusInTransit, StatusOffDuty:
		return s.repo.UpdateStatus(ctx, id, string(typedStatus))
	default:
		return ErrInvalidStatus
	}
}