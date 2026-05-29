package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/auth"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) SignUp(ctx context.Context, name, email, password string) (User, error) {
	if name == "" || email == "" || password == "" {
		return User{}, errors.New("tutti i campi sono obbligatori")
	}
	
	u := User{Name: name, Email: email}
	return s.repo.Create(ctx, u, password)
}

func (s *Service) Login(ctx context.Context, email, password string) (string, User, error) {
	u, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return "", User{}, errors.New("credenziali non valide")
	}

	if !auth.CheckPasswordHash(password, u.PasswordHash) {
		return "", User{}, errors.New("credenziali non valide")
	}

	token, err := auth.GenerateToken(u.ID, u.Email)
	if err != nil {
		return "", User{}, fmt.Errorf("failed to generate token: %w", err)
	}

	return token, u, nil
}

func (s *Service) GetAll(ctx context.Context) ([]User, error) {
	return s.repo.GetAll(ctx)
}