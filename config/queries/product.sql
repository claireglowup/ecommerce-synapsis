-- name: GetProducts :many
SELECT * FROM products;

-- name: GetProductByCategory :many
SELECT * FROM products WHERE category = $1;

-- name: UpdateStockProduct :exec
UPDATE products
SET stock = $1
WHERE id = $1;