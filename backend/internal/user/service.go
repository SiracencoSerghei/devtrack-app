package user

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/auth"
)

var (
	ErrInvalidInput = errors.New("tutti i campi (nome, email, password) sono obbligatori")
	ErrInvalidEmail = errors.New("il formato dell'indirizzo email non è valido")
	ErrShortPwd     = errors.New("la password deve contenere almeno 6 caratteri")

	ErrEmailAlreadyExists = errors.New("esiste già un utente con questo indirizzo email")
)

var emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

type Service struct {
	repo       Repository
	roleRepo   RoleRepository
}

func NewService(repo Repository, roleRepo RoleRepository) *Service {
	return &Service{
		repo:     repo,
		roleRepo: roleRepo,
	}
}
type RoleRepository interface {
	AssignRole(ctx context.Context, userID, roleID string) error
	GetByName(ctx context.Context, name string) (string, error)
}

func (s *Service) SignUp(ctx context.Context, name, email, password string) (User, error) {

	if name == "" || email == "" || password == "" {
		return User{}, ErrInvalidInput
	}

	if len(password) < 6 {
		return User{}, ErrShortPwd
	}

	u := User{
		Name:  name,
		Email: email,
	}

	createdUser, err := s.repo.Create(ctx, u, password)
	if err != nil {
		return User{}, err
	}

	roleID, err := s.roleRepo.GetByName(ctx, "driver")
	if err != nil {
		return User{}, err
	}

	_ = s.roleRepo.AssignRole(ctx, createdUser.ID, roleID)

	return createdUser, nil
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