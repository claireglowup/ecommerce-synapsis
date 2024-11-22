-- name: GetProductById :one 
SELECT * FROM products WHERE id = $1;