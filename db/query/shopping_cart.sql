-- name: CreateShoppingCart :one
INSERT INTO shopping_carts (
    user_id
) VALUES (
    $1
)
RETURNING *;

-- name: GetShoppingCart :one
SELECT * FROM shopping_carts
WHERE user_id = $1
LIMIT 1;

-- name: GetShoppingCartWithCartItems :one
SELECT * FROM shopping_cart_with_cart_items
WHERE user_id = $1
LIMIT 1;

-- name: DeleteShoppingCart :exec
DELETE FROM shopping_carts
WHERE user_id = $1;