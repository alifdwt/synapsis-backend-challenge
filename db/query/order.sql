-- name: CreateOrder :one
INSERT INTO orders (
    user_id,
    payment_method,
    total_cost
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetOrder :one
SELECT * FROM order_with_order_items
WHERE id = $1 LIMIT 1;

-- name: ListOrders :many
SELECT * FROM order_with_order_items
WHERE user_id = $1
ORDER BY order_date DESC;