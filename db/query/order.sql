-- name: CreateOrder :one
INSERT INTO orders (
    user_id,
    payment_method,
    total_cost
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetOrder :one
SELECT * FROM orders
WHERE id = $1 LIMIT 1;

-- name: ListOrderByUserId :many
SELECT * FROM orders
WHERE user_id = $1
ORDER BY order_date DESC;