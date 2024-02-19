-- name: CreateCartItem :one
INSERT INTO cart_items (
    cart_id,
    product_id,
    quantity
) VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: GetCartItem :one
SELECT * FROM cart_items
WHERE id = $1;

-- name: GetCartItemsByUserID :many
SELECT * FROM cart_items
WHERE cart_id = $1;

-- name: UpdateCartItem :one
UPDATE cart_items
SET quantity = $2
WHERE cart_id = $1 AND product_id = $3
RETURNING *;

-- name: DeleteCartItem :exec
DELETE FROM cart_items
WHERE cart_id = $1 AND product_id = $2;