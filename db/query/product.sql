-- name: CreateProduct :one
INSERT INTO products (
    id,
    user_id,
    title,
    description,
    price,
    category_id
) VALUES (
    $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: GetProduct :one
SELECT * FROM products
INNER JOIN categories ON products.category_id = categories.id
WHERE products.id = $1
ORDER BY products.created_at DESC
LIMIT 1;

-- name: ListProducts :many
SELECT * FROM products
INNER JOIN categories ON products.category_id = categories.id
ORDER BY products.created_at DESC
LIMIT $1
OFFSET $2;

-- name: UpdateProduct :one
UPDATE products
SET id = $3, title = $4, description = $5, price = $6, category_id = $7
WHERE id = $1
    AND user_id = $2
RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = $1;