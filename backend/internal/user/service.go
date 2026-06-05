package user

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strings"
	
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/auth"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/db"
)

var (
	ErrInvalidInput = errors.New("tutti i campi (nome, email, password) sono obbligatori")
	ErrInvalidEmail = errors.New("il formato dell'indirizzo email non è valido")
	ErrShortPwd     = errors.New("la password deve contenere almeno 6 caratteri")
)

var emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

type Service struct {
    repo db.UserRepository
}

func NewService(repo db.UserRepository) *Service {
    return &Service{repo: repo}
}

func (s *Service) SignUp(ctx context.Context, name, email, password string) (User, error) {

	name = strings.TrimSpace(name)
	email = strings.ToLower(strings.TrimSpace(email))

	if name == "" || email == "" || password == "" {
		return User{}, ErrInvalidInput
	}

	if !emailRegex.MatchString(email) {
		return User{}, ErrInvalidEmail
	}

	if len(password) < 6 {
		return User{}, ErrShortPwd
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