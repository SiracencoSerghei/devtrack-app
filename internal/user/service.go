package user

import (
	"context"
	"errors"
	"regexp"
	"strings"
)

var emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)

type Service struct {
	repo UserRepository
}

func NewService(repo UserRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, name, email string) (User, error) {
	select {
	case <-ctx.Done():
		return User{}, ctx.Err()
	default:
	}

	name = strings.TrimSpace(name)
	email = strings.TrimSpace(strings.ToLower(email))

	if name == "" || email == "" {
		return User{}, errors.New("name and email are required")
	}

	if !emailRegex.MatchString(email) {
		return User{}, errors.New("invalid email format")
	}

	return s.repo.Create(User{Name: name, Email: email})
}

func (s *Service) GetAll(ctx context.Context) ([]User, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}
	return s.repo.GetAll(), nil
}