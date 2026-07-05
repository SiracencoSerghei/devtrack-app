package application

import (
	"context"
)

type RoleService interface {
	GetByName(ctx context.Context, name string) (string, error)
	AssignRole(ctx context.Context, userID, roleID string) error
	GetByUserID(ctx context.Context, userID string) ([]string, error)
}

type Service struct {
	repo RoleService
}

func NewService(repo RoleService) *Service {
	return &Service{repo: repo}
}

func (s *Service) AssignDefaultRoles(ctx context.Context, userID string, isFirstUser bool) error {
	roles := []string{"driver"}

	if isFirstUser {
		roles = append(roles, "dispatcher", "admin")
	}

	for _, roleName := range roles {
		roleID, err := s.repo.GetByName(ctx, roleName)
		if err != nil {
			continue
		}
		_ = s.repo.AssignRole(ctx, userID, roleID)
	}

	return nil
}

func (s *Service) GetUserRoles(ctx context.Context, userID string) ([]string, error) {
	return s.repo.GetByUserID(ctx, userID)
}