package repository

import (
	"context"
	"dev/lamoda_test/internal/model"
)

type Repository interface {
	Reserve(ctx context.Context, products model.Ids) error
	ReserveRelease(ctx context.Context, products model.Ids) error
	GetAmount(ctx context.Context, stock int) ([]model.Products, error)
}
