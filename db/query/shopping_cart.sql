-- name: CreateShoppingCart :one
INSERT INTO shopping_cart (
    user_id
) VALUES (
    $1
)
RETURNING *;

-- name: GetShoppingCart :one
SELECT * FROM shopping_cart
WHERE user_id = $1
LIMIT 1;

-- name: DeleteShoppingCart :exec
DELETE FROM shopping_cart
WHERE user_id = $1;