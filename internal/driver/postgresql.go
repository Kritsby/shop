package driver

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"shop/internal/config"
)

func NewPostgresPool(cfg config.Postgres) (*pgxpool.Pool, error) {
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.PgUser, cfg.PgPassword, cfg.PgHost, cfg.PgPort, cfg.PgDb)
	pool, err := pgxpool.Connect(context.Background(), url)
	if err != nil {
		return nil, err
	}
	return pool, nil
}
