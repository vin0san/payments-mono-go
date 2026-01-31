package repository

import (
	"context"

	"pye/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type WalletRepository struct {
	db *pgxpool.Pool
}

func NewWalletRepository(db *pgxpool.Pool) *WalletRepository {
	return &WalletRepository{db: db}
}

func (r *WalletRepository) Create(ctx context.Context, w *domain.Wallet) error {
	query := `
	INSERT INTO wallets (id, user_id, balance, created_at)
	VALUES ($1,$2,$3,$4)
	`
	_, err := r.db.Exec(ctx, query,
		w.ID, w.UserID, w.Balance, w.CreatedAt,
	)
	return err
}

func (r *WalletRepository) GetByUserID(ctx context.Context, userID string) (*domain.Wallet, error) {
	w := &domain.Wallet{}

	query := `
	SELECT id,user_id,balance,created_at
	FROM wallets
	WHERE user_id=$1
	`
	err := r.db.QueryRow(ctx, query, userID).
		Scan(&w.ID, &w.UserID, &w.Balance, &w.CreatedAt)
	return w, err
}
