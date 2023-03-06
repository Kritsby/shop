package repository

import (
	"context"
	"dev/lamoda_test/internal/model"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/zerolog/log"
	"strconv"
)

const (
	errNotFound = "one or more of the listed id's was not found or the product is out of stock"
)

type StockPsql struct {
	db *pgxpool.Pool
}

func NewStockPostgres(db *pgxpool.Pool) *StockPsql {
	return &StockPsql{db: db}
}

func (s *StockPsql) Reserve(ctx context.Context, products model.IdRequest) error {
	ok, err := s.check(ctx, products)
	if err != nil {
		return err
	}

	if !ok {
		return fmt.Errorf(errNotFound)
	}

	updateQuery := `
WITH avail AS (
    SELECT pa.storage_id,
           pa.product_id,
           ROW_NUMBER() OVER(PARTITION BY pa.product_id
               ORDER BY pa.amount DESC) AS rank
    FROM shop.product_amount AS pa
    WHERE pa.product_id = ANY(ARRAY[$1::INTEGER[]])
    AND pa.amount > 0
	AND pa.amount > pa.reserved
)
UPDATE shop.product_amount
SET reserved = reserved + 1
WHERE (storage_id, product_id) = ANY(
    SELECT storage_id, product_id
    FROM avail
    WHERE rank = 1
);`

	rows, err := s.db.Query(ctx, updateQuery, products.Ids)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var result model.Products

		err = rows.Scan(&result.Storage, &result.Product, &result.Amount)
		if err != nil {
			return err
		}

		storage := strconv.Itoa(result.Storage)
		product := strconv.Itoa(result.Product)
		amount := strconv.Itoa(result.Amount)

		log.Info().
			Str("Storage", storage).
			Str("Product", product).
			Str("Amount", amount).
			Msg("Changes")
	}

	if err = rows.Err(); err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return err
		}
	}

	return nil
}

func (s *StockPsql) ReserveRelease(ctx context.Context, products model.IdRequest) error {
	ok, err := s.check(ctx, products)
	if err != nil {
		return err
	}

	if !ok {
		return fmt.Errorf(errNotFound)
	}

	updateQuery := `
WITH avail AS (
    SELECT pa.storage_id,
           pa.product_id,
           ROW_NUMBER() OVER(PARTITION BY pa.product_id
               ORDER BY pa.reserved DESC) AS rank
    FROM shop.product_amount AS pa
    WHERE pa.product_id = ANY(ARRAY[$1::INTEGER[]])
    AND pa.reserved > 0
	AND pa.amount > pa.reserved
)
UPDATE shop.product_amount
SET reserved = reserved - 1
WHERE (storage_id, product_id) = ANY(
    SELECT storage_id, product_id
    FROM avail
    WHERE rank = 1
);`

	rows, err := s.db.Query(ctx, updateQuery, products.Ids)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var result model.Products

		err = rows.Scan(&result.Storage, &result.Product, &result.Amount)
		if err != nil {
			return err
		}

		storage := strconv.Itoa(result.Storage)
		product := strconv.Itoa(result.Product)
		amount := strconv.Itoa(result.Amount)

		log.Info().
			Str("Storage", storage).
			Str("Product", product).
			Str("Amount", amount).
			Msg("Changes")
	}

	if err = rows.Err(); err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return err
		}
	}

	return nil
}

func (s *StockPsql) GetAmount(ctx context.Context, stockId int) ([]model.Products, error) {
	query := `
SELECT
    storage_id,
    product_id,
	amount,
	reserved
FROM
	shop.product_amount WHERE storage_id = $1`

	rows, err := s.db.Query(ctx, query, stockId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []model.Products
	for rows.Next() {
		var r model.Products

		err = rows.Scan(&r.Storage, &r.Product, &r.Amount, &r.Reserved)
		if err != nil {
			return nil, err
		}

		result = append(result, r)
	}

	if err = rows.Err(); err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return nil, err
		}
	}

	return result, nil
}

func (s *StockPsql) check(ctx context.Context, products model.IdRequest) (bool, error) {
	query := `
SELECT EXISTS
    (SELECT
         *
     FROM
         shop.product_amount
     WHERE
         product_id = ANY(ARRAY[$1::INTEGER[]]) AND amount > 0);`

	var ok bool
	err := s.db.QueryRow(ctx, query, products.Ids).Scan(&ok)
	if err != nil {
		return ok, err
	}

	return ok, nil
}
