package service

import (
	"context"
	"dev/lamoda_test/internal/model"
	"dev/lamoda_test/internal/repository"
)

type Stock struct {
	repo repository.Repository
}

func NewStock(repo repository.Repository) *Stock {
	return &Stock{
		repo: repo,
	}
}

func (s *Stock) Reserve(ctx context.Context, products model.Ids) error {
	err := s.repo.Reserve(ctx, products)
	if err != nil {
		return err
	}

	return nil
}
func (s *Stock) ReserveRelease(ctx context.Context, products model.Ids) error {
	err := s.repo.ReserveRelease(ctx, products)
	if err != nil {
		return err
	}
	return nil
}

func (s *Stock) GetAmount(ctx context.Context, stock int) ([]model.Products, error) {
	result, err := s.repo.GetAmount(ctx, stock)
	if err != nil {
		return nil, err
	}
	return result, nil
}
