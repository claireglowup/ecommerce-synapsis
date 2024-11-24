-- name: GetCartByUserId :many
SELECT ci.id, p.name, p.price, p.category, p.description, ci.quantity
FROM carts 
JOIN cart_items as ci on carts.id = ci.cart_id
JOIN products as p on ci.product_id = p.id
WHERE carts.user_id = $1;

-- name: AddIdUserToCart :one
INSERT INTO carts (user_id) 
VALUES ($1)
RETURNING id;

-- name: AddProductToCartItems :exec
INSERT INTO cart_items (cart_id, product_id, quantity)
VALUES ($1, $2, $3);

-- name: DeleteProductOnCartById :exec
DELETE FROM cart_items WHERE id = $1 AND cart_id = $2;

-- name: GetCartIdByUserId :one
SELECT id 
FROM carts
WHERE user_id = $1;