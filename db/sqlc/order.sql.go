// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: order.sql

package db

import (
	"context"
)

const createOrder = `-- name: CreateOrder :one
INSERT INTO orders (
    user_id,
    payment_method,
    total_cost
) VALUES (
    $1, $2, $3
) RETURNING id, user_id, payment_method, total_cost, order_date
`

type CreateOrderParams struct {
	UserID        string `json:"user_id"`
	PaymentMethod string `json:"payment_method"`
	TotalCost     int64  `json:"total_cost"`
}

func (q *Queries) CreateOrder(ctx context.Context, arg CreateOrderParams) (Order, error) {
	row := q.db.QueryRowContext(ctx, createOrder, arg.UserID, arg.PaymentMethod, arg.TotalCost)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.PaymentMethod,
		&i.TotalCost,
		&i.OrderDate,
	)
	return i, err
}

const getOrder = `-- name: GetOrder :one
SELECT id, user_id, payment_method, total_cost, order_date, order_items FROM order_with_order_items
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetOrder(ctx context.Context, id string) (OrderWithOrderItem, error) {
	row := q.db.QueryRowContext(ctx, getOrder, id)
	var i OrderWithOrderItem
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.PaymentMethod,
		&i.TotalCost,
		&i.OrderDate,
		&i.OrderItems,
	)
	return i, err
}

const listOrders = `-- name: ListOrders :many
SELECT id, user_id, payment_method, total_cost, order_date, order_items FROM order_with_order_items
WHERE user_id = $1
ORDER BY order_date DESC
`

func (q *Queries) ListOrders(ctx context.Context, userID string) ([]OrderWithOrderItem, error) {
	rows, err := q.db.QueryContext(ctx, listOrders, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []OrderWithOrderItem{}
	for rows.Next() {
		var i OrderWithOrderItem
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.PaymentMethod,
			&i.TotalCost,
			&i.OrderDate,
			&i.OrderItems,
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
