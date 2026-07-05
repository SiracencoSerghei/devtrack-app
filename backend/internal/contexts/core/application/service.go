package application

import (
	"context"
	"errors"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/contexts/core/domain"
)

type Repository interface {
	CreateEmployee(ctx context.Context, emp domain.Employee) (domain.Employee, error)
	GetEmployeeByUserID(ctx context.Context, userID string) (domain.Employee, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) RegisterEmployee(ctx context.Context, emp domain.Employee) (domain.Employee, error) {
	if emp.FirstName == "" || emp.LastName == "" {
		return domain.Employee{}, errors.New("first and last name are required")
	}
	return s.repo.CreateEmployee(ctx, emp)
}

func (s *Service) GetProfile(ctx context.Context, userID string) (domain.Employee, error) {
	return s.repo.GetEmployeeByUserID(ctx, userID)
}