package driver

import (
	"context"
	"dev/lamoda_test/internal/config"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

func NewPostgresPool(cfg config.Postgres) (*pgxpool.Pool, error) {
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.PgUser, cfg.PgPassword, cfg.PgHost, cfg.PgPort, cfg.PgDb)
	pool, err := pgxpool.Connect(context.Background(), url)
	if err != nil {
		return nil, err
	}
	return pool, nil
}
