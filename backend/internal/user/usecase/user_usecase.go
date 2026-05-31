package usecase

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"time"

	"github.com/SiracencoSerghei/devtrack-app/backend/internal/auth"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/domain"
	"github.com/go-playground/validator/v10"
)

var (
	ErrValidationFailed = errors.New("i dati inseriti non sono validi")
	ErrInvalidAuth      = errors.New("credenziali non valide")
)

type SignUpInput struct {
	Name     string `json:"name" validate:"required,min=2,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserUseCase struct {
	repo     domain.UserRepository
	validate *validator.Validate
}

func NewUserUseCase(repo domain.UserRepository) *UserUseCase {
	return &UserUseCase{
		repo:     repo,
		validate: validator.New(),
	}
}

func (u *UserUseCase) SignUp(ctx context.Context, input SignUpInput) (domain.User, error) {
	
	if err := u.validate.Struct(input); err != nil {
		return domain.User{}, ErrValidationFailed
	}

	
	hashedPassword, err := auth.HashPassword(input.Password)
	if err != nil {
		return domain.User{}, errors.New("errore durante la messa in sicurezza della password")
	}

	userEntity := domain.User{
		Name:  input.Name,
		Email: input.Email,
	}

	return u.repo.Create(ctx, userEntity, hashedPassword)
}

func (u *UserUseCase) Login(ctx context.Context, input LoginInput) (string, string, domain.User, error) {
	if err := u.validate.Struct(input); err != nil {
		return "", "", domain.User{}, ErrValidationFailed
	}

	user, err := u.repo.GetByEmail(ctx, input.Email)
	if err != nil {
		return "", "", domain.User{}, ErrInvalidAuth
	}

	if !auth.CheckPasswordHash(input.Password, user.PasswordHash) {
		return "", "", domain.User{}, ErrInvalidAuth
	}

	accessToken, err := auth.GenerateToken(user.ID, user.Email)
	if err != nil {
		return "", "", domain.User{}, err
	}

	refreshToken, err := generateRandomString(32)
	if err != nil {
		return "", "", domain.User{}, err
	}

	session := domain.Session{
		UserID:       user.ID,
		RefreshToken: refreshToken,
		ExpiresAt:    time.Now().Add(7 * 24 * time.Hour),
	}

	if err := u.repo.CreateSession(ctx, session); err != nil {
		return "", "", domain.User{}, err
	}

	return accessToken, refreshToken, user, nil
}

func (u *UserUseCase) GetAll(ctx context.Context) ([]domain.User, error) {
	return u.repo.GetAll(ctx)
}

func generateRandomString(n int) (string, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}