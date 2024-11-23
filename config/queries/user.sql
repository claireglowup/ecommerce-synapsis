-- name: RegisterUser :exec
INSERT INTO users (name, email, password) 
VALUES ($1, $2, $3);

-- name: Login :one
SELECT id, password FROM users WHERE email = $1;

-- name: GetCartByUserId :many
SELECT p.name, p.price, p.category, p.description, ci.quantity 
FROM carts 
JOIN cart_items as ci on carts.id = ci.cart_id
JOIN products as p on ci.product_id = p.id
WHERE carts.user_id = $1;