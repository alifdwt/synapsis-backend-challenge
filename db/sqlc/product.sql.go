// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: product.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createProduct = `-- name: CreateProduct :one
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
RETURNING id, user_id, title, description, price, category_id, created_at, updated_at
`

type CreateProductParams struct {
	ID          string         `json:"id"`
	UserID      string         `json:"user_id"`
	Title       string         `json:"title"`
	Description sql.NullString `json:"description"`
	Price       int64          `json:"price"`
	CategoryID  string         `json:"category_id"`
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error) {
	row := q.db.QueryRowContext(ctx, createProduct,
		arg.ID,
		arg.UserID,
		arg.Title,
		arg.Description,
		arg.Price,
		arg.CategoryID,
	)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Description,
		&i.Price,
		&i.CategoryID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteProduct = `-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = $1
`

func (q *Queries) DeleteProduct(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteProduct, id)
	return err
}

const getProduct = `-- name: GetProduct :one
SELECT products.id, user_id, title, description, price, category_id, created_at, updated_at, categories.id, name FROM products
INNER JOIN categories ON products.category_id = categories.id
WHERE products.id = $1
ORDER BY products.created_at DESC
LIMIT 1
`

type GetProductRow struct {
	ID          string         `json:"id"`
	UserID      string         `json:"user_id"`
	Title       string         `json:"title"`
	Description sql.NullString `json:"description"`
	Price       int64          `json:"price"`
	CategoryID  string         `json:"category_id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	ID_2        string         `json:"id_2"`
	Name        string         `json:"name"`
}

func (q *Queries) GetProduct(ctx context.Context, id string) (GetProductRow, error) {
	row := q.db.QueryRowContext(ctx, getProduct, id)
	var i GetProductRow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Description,
		&i.Price,
		&i.CategoryID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ID_2,
		&i.Name,
	)
	return i, err
}

const listProducts = `-- name: ListProducts :many
SELECT products.id, user_id, title, description, price, category_id, created_at, updated_at, categories.id, name FROM products
INNER JOIN categories ON products.category_id = categories.id
ORDER BY products.created_at DESC
LIMIT $1
OFFSET $2
`

type ListProductsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type ListProductsRow struct {
	ID          string         `json:"id"`
	UserID      string         `json:"user_id"`
	Title       string         `json:"title"`
	Description sql.NullString `json:"description"`
	Price       int64          `json:"price"`
	CategoryID  string         `json:"category_id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	ID_2        string         `json:"id_2"`
	Name        string         `json:"name"`
}

func (q *Queries) ListProducts(ctx context.Context, arg ListProductsParams) ([]ListProductsRow, error) {
	rows, err := q.db.QueryContext(ctx, listProducts, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListProductsRow{}
	for rows.Next() {
		var i ListProductsRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Title,
			&i.Description,
			&i.Price,
			&i.CategoryID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.ID_2,
			&i.Name,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateProduct = `-- name: UpdateProduct :one
UPDATE products
SET id = $3, title = $4, description = $5, price = $6, category_id = $7
WHERE id = $1
    AND user_id = $2
RETURNING id, user_id, title, description, price, category_id, created_at, updated_at
`

type UpdateProductParams struct {
	ID          string         `json:"id"`
	UserID      string         `json:"user_id"`
	ID_2        string         `json:"id_2"`
	Title       string         `json:"title"`
	Description sql.NullString `json:"description"`
	Price       int64          `json:"price"`
	CategoryID  string         `json:"category_id"`
}

func (q *Queries) UpdateProduct(ctx context.Context, arg UpdateProductParams) (Product, error) {
	row := q.db.QueryRowContext(ctx, updateProduct,
		arg.ID,
		arg.UserID,
		arg.ID_2,
		arg.Title,
		arg.Description,
		arg.Price,
		arg.CategoryID,
	)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Description,
		&i.Price,
		&i.CategoryID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}