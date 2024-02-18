-- name: CreateCategory :one
INSERT INTO categories (
    id,
    name
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetCategory :one
SELECT * FROM categories_with_products
WHERE id = $1 LIMIT 1;

-- name: ListCategories :many
SELECT * FROM categories_with_products
ORDER BY name
LIMIT $1
OFFSET $2;

-- name: UpdateCategory :one
UPDATE categories
SET id = $2, name = $3
WHERE id = $1
RETURNING *;

-- name: DeleteCategory :exec
DELETE FROM categories
WHERE id = $1;