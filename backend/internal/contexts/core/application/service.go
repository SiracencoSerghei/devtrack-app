package application

import (
	"context"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/contexts/core/domain"
)

type DomainError struct {
	Status  int
	Message string
}

func (e DomainError) Error() string   { return e.Message }
func (e DomainError) APIError() (int, string) { return e.Status, e.Message }

var ErrNameRequired = DomainError{Status: 400, Message: "first and last name are required"}

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
		return domain.Employee{}, ErrNameRequired // Передає чистий 400 Bad Request
	}
	return s.repo.CreateEmployee(ctx, emp)
}

func (s *Service) GetProfile(ctx context.Context, userID string) (domain.Employee, error) {
	return s.repo.GetEmployeeByUserID(ctx, userID)
}