package application

import (
	"context"
	"log/slog"
	"regexp"
	"strings"

	"github.com/SiracencoSerghei/devtrack-app/backend/internal/contexts/identity/domain"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/shared/auth"
)

var emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

type DomainError struct {
	Status  int
	Message string
}

func (e DomainError) Error() string {
	return e.Message
}

func (e DomainError) APIError() (int, string) {
	return e.Status, e.Message
}

var (
	ErrInvalidInput = DomainError{
		Status:  400,
		Message: "tutti i campi (nome, email, password) sono obbligatori",
	}
	ErrInvalidEmail = DomainError{
		Status:  400,
		Message: "il formato dell'indirizzo email non è valido",
	}
	ErrShortPwd = DomainError{
		Status:  400,
		Message: "la password deve contenere almeno 6 caratteri",
	}
	ErrEmailAlreadyExists = DomainError{
		Status:  409,
		Message: "esiste già un utente con questo indirizzo email",
	}
	ErrUnauthorized = DomainError{
		Status:  401,
		Message: "credenziali non valide",
	}
)

type ServiceInterface interface {
	SignUp(ctx context.Context, name, email, password string) (domain.User, error)
	Login(ctx context.Context, email, password string) (string, domain.User, error)
	GetAll(ctx context.Context) ([]domain.User, error)
}

type AccessService interface {
	AssignDefaultRoles(ctx context.Context, userID string, isFirstUser bool) error
	GetUserRoles(ctx context.Context, userID string) ([]string, error)
}

type Service struct {
	repo         domain.Repository
	tokenManager *auth.TokenManager
	access       AccessService
}

// ПРАВИЛЬНИЙ КОНСТРУКТОР: Тепер залежність access обов'язкова та ініціалізується!
func NewService(
	repo domain.Repository,
	tm *auth.TokenManager,
	access AccessService,
) *Service {
	return &Service{
		repo:         repo,
		tokenManager: tm,
		access:       access,
	}
}

func (s *Service) SignUp(ctx context.Context, name, email, password string) (domain.User, error) {
	name = strings.TrimSpace(name)
	email = strings.ToLower(strings.TrimSpace(email))

	if name == "" || email == "" || password == "" {
		return domain.User{}, ErrInvalidInput
	}
	if !emailRegex.MatchString(email) {
		return domain.User{}, ErrInvalidEmail
	}
	if len(password) < 6 {
		return domain.User{}, ErrShortPwd
	}

	user := domain.User{
		Name:  name,
		Email: email,
	}

	count, err := s.repo.Count(ctx)
	if err != nil {
		return domain.User{}, err
	}

	isFirstUser := count == 0
	roles := []string{"driver"}

	if isFirstUser {
		roles = append(roles, "dispatcher", "admin")
	}

	created, err := s.repo.CreateWithRoles(ctx, user, password, roles)
	if err != nil {
		slog.Error("user creation failed", "error", err)
		return domain.User{}, err
	}

	return created, nil
}

func (s *Service) Login(ctx context.Context, email, password string) (string, domain.User, error) {
	email = strings.ToLower(strings.TrimSpace(email))

	u, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return "", domain.User{}, ErrUnauthorized
	}

	if !auth.CheckPasswordHash(password, u.PasswordHash) {
		return "", domain.User{}, ErrUnauthorized
	}

	// Більше не падає в panic! Об'єкт s.access тепер гарантовано існує
	roles, err := s.access.GetUserRoles(ctx, u.ID)
	if err != nil {
		return "", domain.User{}, err
	}

	u.Roles = roles

	token, err := s.tokenManager.GenerateToken(u.ID, u.Email, u.Roles)
	if err != nil {
		return "", domain.User{}, err
	}

	return token, u, nil
}

func (s *Service) GetAll(ctx context.Context) ([]domain.User, error) {
	return s.repo.GetAll(ctx)
}