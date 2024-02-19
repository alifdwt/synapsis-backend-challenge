// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: order_item.sql

package db

import (
	"context"
)

const createOrderItem = `-- name: CreateOrderItem :one
INSERT INTO order_items (
    order_id,
    product_id,
    quantity,
    price_at_purchase
) VALUES (
    $1, $2, $3, $4
)
RETURNING id, order_id, product_id, quantity, price_at_purchase
`

type CreateOrderItemParams struct {
	OrderID         string `json:"order_id"`
	ProductID       string `json:"product_id"`
	Quantity        int64  `json:"quantity"`
	PriceAtPurchase int64  `json:"price_at_purchase"`
}

func (q *Queries) CreateOrderItem(ctx context.Context, arg CreateOrderItemParams) (OrderItem, error) {
	row := q.db.QueryRowContext(ctx, createOrderItem,
		arg.OrderID,
		arg.ProductID,
		arg.Quantity,
		arg.PriceAtPurchase,
	)
	var i OrderItem
	err := row.Scan(
		&i.ID,
		&i.OrderID,
		&i.ProductID,
		&i.Quantity,
		&i.PriceAtPurchase,
	)
	return i, err
}

const listOrderItemsByOrderID = `-- name: ListOrderItemsByOrderID :many
SELECT id, order_id, product_id, quantity, price_at_purchase FROM order_items
WHERE order_id = $1
`

func (q *Queries) ListOrderItemsByOrderID(ctx context.Context, orderID string) ([]OrderItem, error) {
	rows, err := q.db.QueryContext(ctx, listOrderItemsByOrderID, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []OrderItem{}
	for rows.Next() {
		var i OrderItem
		if err := rows.Scan(
			&i.ID,
			&i.OrderID,
			&i.ProductID,
			&i.Quantity,
			&i.PriceAtPurchase,
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
