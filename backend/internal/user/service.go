package user

import (
	"context"
	"fmt"
	"regexp"

	"github.com/SiracencoSerghei/devtrack-app/backend/internal/auth"
)

var emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

// 🌟 ВИПРАВЛЕНО: Структуровані Enterprise помилки, що само-мапляться на HTTP
type DomainError struct {
	Status  int
	Message string
}

func (e DomainError) Error() string { return e.Message }
func (e DomainError) APIError() (int, string) { return e.Status, e.Message }

var (
	ErrInvalidInput       = DomainError{Status: 400, Message: "tutti i campi (nome, email, password) sono obbligatori"}
	ErrInvalidEmail       = DomainError{Status: 400, Message: "il formato dell'indirizzo email non è valido"}
	ErrShortPwd           = DomainError{Status: 400, Message: "la password deve contenere almeno 6 caratteri"}
	ErrEmailAlreadyExists = DomainError{Status: 409, Message: "esiste già un utente con questo indirizzo email"}
	ErrUnauthorized       = DomainError{Status: 401, Message: "credenziali non valide"}
)

// 🌟 ВИПРАВЛЕНО: Оголошення сервісного інтерфейсу для Mock-тестів
type ServiceInterface interface {
	SignUp(ctx context.Context, name, email, password string) (User, error)
	Login(ctx context.Context, email, password string) (string, User, error)
	GetAll(ctx context.Context) ([]User, error)
}

type Service struct {
	repo         Repository
	roleRepo     RoleRepository
	tokenManager *auth.TokenManager
}

func NewService(repo Repository, roleRepo RoleRepository, tm *auth.TokenManager) *Service {
	return &Service{repo: repo, roleRepo: roleRepo, tokenManager: tm}
}

func (s *Service) SignUp(ctx context.Context, name, email, password string) (User, error) {
	if name == "" || email == "" || password == "" {
		return User{}, ErrInvalidInput
	}
	if len(password) < 6 {
		return User{}, ErrShortPwd
	}
	if !emailRegex.MatchString(email) {
		return User{}, ErrInvalidEmail
	}

	totalUsers, err := s.repo.Count(ctx)
	if err != nil {
		totalUsers = 0
	}

	u := User{Name: name, Email: email}
	createdUser, err := s.repo.Create(ctx, u, password)
	if err != nil {
		return User{}, err
	}

	rolesToAssign := []string{"driver"}
	if totalUsers == 0 {
		rolesToAssign = []string{"driver", "dispatcher", "admin"}
	}

	for _, roleName := range rolesToAssign {
		roleID, err := s.roleRepo.GetByName(ctx, roleName)
		if err == nil {
			_ = s.roleRepo.AssignRole(ctx, createdUser.ID, roleID)
		}
	}

	createdUser.Roles, _ = s.roleRepo.GetByUserID(ctx, createdUser.ID)
	return createdUser, nil
}

func (s *Service) Login(ctx context.Context, email, password string) (string, User, error) {
	u, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return "", User{}, ErrUnauthorized
	}

	if !auth.CheckPasswordHash(password, u.PasswordHash) {
		return "", User{}, ErrUnauthorized
	}

	token, err := s.tokenManager.GenerateToken(u.ID, u.Email, u.Roles)
	if err != nil {
		return "", User{}, fmt.Errorf("failed to generate token: %w", err)
	}

	return token, u, nil
}

func (s *Service) GetAll(ctx context.Context) ([]User, error) {
	return s.repo.GetAll(ctx)
}