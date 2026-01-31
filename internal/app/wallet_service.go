package app

import (
	"context"
	"pye/internal/domain"
	"pye/internal/repository"
	"time"

	"github.com/google/uuid"
)

type WalletService struct {
	repo *repository.WalletRepository
}

func NewWalletService(r *repository.WalletRepository) *WalletService {
	return &WalletService{repo: r}
}

func (s *WalletService) CreateForUser(ctx context.Context, userID string) error {
	w := &domain.Wallet{
		ID:        uuid.New().String(),
		UserID:    userID,
		Balance:   0,
		CreatedAt: time.Now().UTC(),
	}

	return s.repo.Create(ctx, w)
}

func (s *WalletService) GetBalance(ctx context.Context, userID string) (int64, error) {
	w, err := s.repo.GetByUserID(ctx, userID)
	if err != nil {
		return 0, err
	}
	return w.Balance, nil
}
