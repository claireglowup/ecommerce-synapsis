// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package sqlc

import (
	"context"
)

const getProduct = `-- name: GetProduct :one
SELECT id, name, price, stock, description FROM products WHERE id = $1
`

func (q *Queries) GetProduct(ctx context.Context, id int32) (Product, error) {
	row := q.db.QueryRowContext(ctx, getProduct, id)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Price,
		&i.Stock,
		&i.Description,
	)
	return i, err
}
