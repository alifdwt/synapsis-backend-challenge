package db

import (
	"context"
	"testing"

	"github.com/alifdwt/synapsis-backend-challenge/util"
	"github.com/stretchr/testify/require"
)

func createRandomOrderItem(t *testing.T, orderId string) OrderItem {
	category := createRandomCategory(t)

	product := createRandomProduct(t, category.ID)

	arg := CreateOrderItemParams{
		OrderID:         orderId,
		ProductID:       product.ID,
		Quantity:        util.RandomInt(1, 10),
		PriceAtPurchase: util.RandomMoney(),
	}

	orderItem, err := testQueries.CreateOrderItem(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, orderItem)

	require.Equal(t, arg.OrderID, orderItem.OrderID)
	require.Equal(t, arg.ProductID, orderItem.ProductID)
	require.Equal(t, arg.Quantity, orderItem.Quantity)
	require.Equal(t, arg.PriceAtPurchase, orderItem.PriceAtPurchase)

	return orderItem
}

func TestCreateOrderItem(t *testing.T) {
	order := createRandomOrder(t)
	createRandomOrderItem(t, order.ID)
}

func TestListOrderItemsByOrderId(t *testing.T) {
	order := createRandomOrder(t)
	orderItem := createRandomOrderItem(t, order.ID)

	orderItems, err := testQueries.ListOrderItemsByOrderID(context.Background(), orderItem.OrderID)
	require.NoError(t, err)
	require.NotEmpty(t, orderItems)

	for _, orderItem := range orderItems {
		require.NotEmpty(t, orderItem)
		require.Equal(t, order.ID, orderItem.OrderID)
	}
}
