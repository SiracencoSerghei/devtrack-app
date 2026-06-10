package driver

import (
	"context"
	"errors"
	"strings"
)

var (
	ErrLicenseRequired = errors.New("il numero di patente è obbligatorio")
	ErrDriverNotFound  = errors.New("profilo driver non trovato")
)

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
		Status:        "AVAILABLE",
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
	status = strings.ToUpper(strings.TrimSpace(status))
	if status != "AVAILABLE" && status != "IN_TRANSIT" && status != "OFF_DUTY" {
		return errors.New("stato driver non valido")
	}
	return s.repo.UpdateStatus(ctx, id, status)
}