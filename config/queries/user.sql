-- name: RegisterUser :exec
INSERT INTO users (name, email, password) 
VALUES ($1, $2, $3);

-- name: Login :one
SELECT id, password FROM users WHERE email = $1;

-- name: GetUserByID :one 
SELECT name, email FROM users WHERE id = $1;