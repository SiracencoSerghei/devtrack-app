package user

import (
	"context"
	"errors"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, name, email string) (User, error) {
	select {
	case <-ctx.Done():
		return User{}, ctx.Err()
	default:
	}

	if name == "" || email == "" {
		return User{}, errors.New("name and email are required")
	}

	if s.repo.ExistsByEmail(email) {
		return User{}, errors.New("email already exists")
	}

	user := User{
		ID:    s.repo.NextID(),
		Name:  name,
		Email: email,
	}

	if err := s.repo.Save(user); err != nil {
		return User{}, err
	}

	return user, nil
}

func (s *Service) GetAll(ctx context.Context) ([]User, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	return s.repo.GetAll(), nil
}