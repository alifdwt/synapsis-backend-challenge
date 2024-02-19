// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: shopping_cart.sql

package db

import (
	"context"
)

const createShoppingCart = `-- name: CreateShoppingCart :one
INSERT INTO shopping_cart (
    user_id
) VALUES (
    $1
)
RETURNING id, user_id, created_at
`

func (q *Queries) CreateShoppingCart(ctx context.Context, userID string) (ShoppingCart, error) {
	row := q.db.QueryRowContext(ctx, createShoppingCart, userID)
	var i ShoppingCart
	err := row.Scan(&i.ID, &i.UserID, &i.CreatedAt)
	return i, err
}

const deleteShoppingCart = `-- name: DeleteShoppingCart :exec
DELETE FROM shopping_cart
WHERE user_id = $1
`

func (q *Queries) DeleteShoppingCart(ctx context.Context, userID string) error {
	_, err := q.db.ExecContext(ctx, deleteShoppingCart, userID)
	return err
}

const getShoppingCart = `-- name: GetShoppingCart :one
SELECT id, user_id, created_at FROM shopping_cart
WHERE user_id = $1
LIMIT 1
`

func (q *Queries) GetShoppingCart(ctx context.Context, userID string) (ShoppingCart, error) {
	row := q.db.QueryRowContext(ctx, getShoppingCart, userID)
	var i ShoppingCart
	err := row.Scan(&i.ID, &i.UserID, &i.CreatedAt)
	return i, err
}