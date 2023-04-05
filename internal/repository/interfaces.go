package repository

import (
	"context"
	"shop/internal/model"
)

type Repository interface {
	Reserve(ctx context.Context, products model.IdRequest) error
	ReserveRelease(ctx context.Context, products model.IdRequest) error
	GetAmount(ctx context.Context, stock int) ([]model.Products, error)
}
