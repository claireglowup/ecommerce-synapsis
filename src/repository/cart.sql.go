// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: cart.sql

package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const addIdUserToCart = `-- name: AddIdUserToCart :one
INSERT INTO carts (user_id) 
VALUES ($1)
RETURNING id
`

func (q *Queries) AddIdUserToCart(ctx context.Context, userID uuid.NullUUID) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, addIdUserToCart, userID)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const addProductToCartItems = `-- name: AddProductToCartItems :exec
INSERT INTO cart_items (cart_id, product_id, quantity)
VALUES ($1, $2, $3)
`

type AddProductToCartItemsParams struct {
	CartID    uuid.NullUUID `json:"cart_id"`
	ProductID uuid.NullUUID `json:"product_id"`
	Quantity  int32         `json:"quantity"`
}

func (q *Queries) AddProductToCartItems(ctx context.Context, arg AddProductToCartItemsParams) error {
	_, err := q.db.ExecContext(ctx, addProductToCartItems, arg.CartID, arg.ProductID, arg.Quantity)
	return err
}

const getCartByUserId = `-- name: GetCartByUserId :many
SELECT p.name, p.price, p.category, p.description, ci.quantity 
FROM carts 
JOIN cart_items as ci on carts.id = ci.cart_id
JOIN products as p on ci.product_id = p.id
WHERE carts.user_id = $1
`

type GetCartByUserIdRow struct {
	Name        string         `json:"name"`
	Price       string         `json:"price"`
	Category    string         `json:"category"`
	Description sql.NullString `json:"description"`
	Quantity    int32          `json:"quantity"`
}

func (q *Queries) GetCartByUserId(ctx context.Context, userID uuid.NullUUID) ([]GetCartByUserIdRow, error) {
	rows, err := q.db.QueryContext(ctx, getCartByUserId, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetCartByUserIdRow{}
	for rows.Next() {
		var i GetCartByUserIdRow
		if err := rows.Scan(
			&i.Name,
			&i.Price,
			&i.Category,
			&i.Description,
			&i.Quantity,
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