-- name: RegisterUser :exec
INSERT INTO users (name, email, password) 
VALUES ($1, $2, $3);

-- name: Login :one
SELECT * FROM users WHERE email = $1;