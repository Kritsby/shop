package service

import (
	"context"
	"dev/lamoda_test/internal/model"
)

//go:generate mockgen -source=interfaces.go -destination=mocks/mock.go

type Stocker interface {
	Reserve(ctx context.Context, products model.IdRequest) error
	ReserveRelease(ctx context.Context, products model.IdRequest) error
	GetAmount(ctx context.Context, stock int) ([]model.Products, error)
}
