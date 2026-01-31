package repository

import (
	"context"
	"pye/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, u *domain.User) error {
	query := `
	INSERT INTO users (id, name, email, password_hash, created_at)
	VALUES ($1,$2,$3,$4,$5)
	`
	_, err := r.db.Exec(ctx, query,
		u.ID, u.Name, u.Email, u.PasswordHash, u.CreatedAt,
	)
	return err
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	u := &domain.User{}
	query := `SELECT id,name,email,password_hash,created_at FROM users WHERE email=$1`

	err := r.db.QueryRow(ctx, query, email).
		Scan(&u.ID, &u.Name, &u.Email, &u.PasswordHash, &u.CreatedAt)

	return u, err
}
