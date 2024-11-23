-- name: GetProducts :many
SELECT * FROM products;

-- name: GetProductByCategory :many
SELECT * FROM products WHERE category = $1;