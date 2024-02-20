-- name: CreateUser :one
INSERT INTO users (
    username,
    hashed_password,
    full_name,
    email
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: GetUserWithProducts :one
SELECT * FROM users_with_products
WHERE username = $1 LIMIT 1;

-- name: ListUserWithProducts :many
SELECT * FROM users_with_products
ORDER BY created_at
LIMIT $1
OFFSET $2;