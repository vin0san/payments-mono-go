package app

import (
	"context"
	"errors"
	"strings"
	"time"

	"pye/internal/domain"
	"pye/internal/repository"
	"pye/pkg/security"

	"github.com/google/uuid"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(r *repository.UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) CreateUser(ctx context.Context, name, email string) (*domain.User, error) {
	u := &domain.User{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		CreatedAt: time.Now().UTC(),
	}

	if err := s.repo.Create(ctx, u); err != nil {
		return nil, err
	}

	return u, nil
}

func (s *UserService) Register(ctx context.Context, name, email, password string) (*domain.User, error) {
	hash, err := security.HashPassword(password)
	if err != nil {
		return nil, err
	}

	u := &domain.User{
		ID:           uuid.New().String(),
		Name:         name,
		Email:        email,
		PasswordHash: hash,
		CreatedAt:    time.Now().UTC(),
	}

	err = s.repo.Create(ctx, u)
	if err != nil {
		if strings.Contains(err.Error(), "users_email_key") {
			return nil, errors.New("email already exists")
		}
		return nil, err
	}

	return u, nil
}

func (s *UserService) Login(ctx context.Context, email, password string) (string, error) {
	u, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	if err := security.CheckPassword(u.PasswordHash, password); err != nil {
		return "", err
	}

	return security.GenerateToken(u.ID)
}
