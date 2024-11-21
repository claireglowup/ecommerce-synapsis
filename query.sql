-- name: GetProduct :one 
SELECT * FROM products WHERE id = $1;