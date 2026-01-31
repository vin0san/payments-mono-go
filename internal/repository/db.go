package repository

import (
	"context"
	"fmt"
	"pye/internal/config"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgres(cfg config.DBConfig) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.User,
		cfg.Pass,
		cfg.Host,
		cfg.Port,
		cfg.Name,
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return pgxpool.New(ctx, dsn)
}
