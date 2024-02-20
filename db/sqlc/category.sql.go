// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: category.sql

package db

import (
	"context"
	"encoding/json"
)

const createCategory = `-- name: CreateCategory :one
INSERT INTO categories (
    id,
    name
) VALUES (
    $1, $2
) RETURNING id, name
`

type CreateCategoryParams struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (q *Queries) CreateCategory(ctx context.Context, arg CreateCategoryParams) (Category, error) {
	row := q.db.QueryRowContext(ctx, createCategory, arg.ID, arg.Name)
	var i Category
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const deleteCategory = `-- name: DeleteCategory :exec
DELETE FROM categories
WHERE id = $1
`

func (q *Queries) DeleteCategory(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteCategory, id)
	return err
}

const getCategory = `-- name: GetCategory :one
SELECT id, name, products FROM categories_with_products
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetCategory(ctx context.Context, id string) (CategoriesWithProduct, error) {
    row := q.db.QueryRowContext(ctx, getCategory, id)
    var i CategoriesWithProduct
    var productsJSON string
    err := row.Scan(&i.ID, &i.Name, &productsJSON)
    if err != nil {
        return CategoriesWithProduct{}, err
    }

    // Unmarshal JSON ke slice produk
    err = json.Unmarshal([]byte(productsJSON), &i.Products)
    if err != nil {
        return CategoriesWithProduct{}, err
    }

    return i, nil
}

const listCategories = `-- name: ListCategories :many
SELECT id, name, products FROM categories_with_products
ORDER BY name
LIMIT $1
OFFSET $2
`

type ListCategoriesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListCategories(ctx context.Context, arg ListCategoriesParams) ([]CategoriesWithProduct, error) {
    rows, err := q.db.QueryContext(ctx, listCategories, arg.Limit, arg.Offset)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var categories []CategoriesWithProduct

    for rows.Next() {
        var category CategoriesWithProduct
        var productsJSON string

        if err := rows.Scan(&category.ID, &category.Name, &productsJSON); err != nil {
            return nil, err
        }

        // Unmarshal JSON ke dalam slice produk
        var products []Product
        err := json.Unmarshal([]byte(productsJSON), &products)
        if err != nil {
            return nil, err
        }

        category.Products = products
        categories = append(categories, category)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return categories, nil
}


const updateCategory = `-- name: UpdateCategory :one
UPDATE categories
SET id = $2, name = $3
WHERE id = $1
RETURNING id, name
`

type UpdateCategoryParams struct {
	ID   string `json:"id"`
	ID_2 string `json:"id_2"`
	Name string `json:"name"`
}

func (q *Queries) UpdateCategory(ctx context.Context, arg UpdateCategoryParams) (Category, error) {
	row := q.db.QueryRowContext(ctx, updateCategory, arg.ID, arg.ID_2, arg.Name)
	var i Category
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}
