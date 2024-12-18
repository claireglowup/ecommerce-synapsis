// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: product.sql

package repository

import (
	"context"
)

const getProductByCategory = `-- name: GetProductByCategory :many
SELECT id, name, price, stock, category, description FROM products WHERE category = $1
`

func (q *Queries) GetProductByCategory(ctx context.Context, category string) ([]Product, error) {
	rows, err := q.db.QueryContext(ctx, getProductByCategory, category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Product{}
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Price,
			&i.Stock,
			&i.Category,
			&i.Description,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProducts = `-- name: GetProducts :many
SELECT id, name, price, stock, category, description FROM products
`

func (q *Queries) GetProducts(ctx context.Context) ([]Product, error) {
	rows, err := q.db.QueryContext(ctx, getProducts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Product{}
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Price,
			&i.Stock,
			&i.Category,
			&i.Description,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateStockProduct = `-- name: UpdateStockProduct :exec
UPDATE products
SET stock = $1
WHERE id = $1
`

func (q *Queries) UpdateStockProduct(ctx context.Context, stock int32) error {
	_, err := q.db.ExecContext(ctx, updateStockProduct, stock)
	return err
}
