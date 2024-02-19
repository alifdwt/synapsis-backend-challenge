package db

import (
	"context"
	"testing"
	"time"

	"github.com/alifdwt/synapsis-backend-challenge/util"
	"github.com/stretchr/testify/require"
)

func createRandomOrder(t *testing.T) Order {
	user := createRandomUser(t)
	arg := CreateOrderParams{
		UserID:        user.Username,
		PaymentMethod: "COD",
		TotalCost:     util.RandomMoney(),
	}

	order, err := testQueries.CreateOrder(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, order)

	require.Equal(t, arg.UserID, order.UserID)
	require.Equal(t, arg.PaymentMethod, order.PaymentMethod)
	require.Equal(t, arg.TotalCost, order.TotalCost)

	require.NotZero(t, order.OrderDate)

	return order
}

func TestCreateOrder(t *testing.T) {
	createRandomOrder(t)
}

func TestGetOrder(t *testing.T) {
	order1 := createRandomOrder(t)

	order2, err := testQueries.GetOrder(context.Background(), order1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, order2)

	require.Equal(t, order1.ID, order2.ID)
	require.Equal(t, order1.UserID, order2.UserID)
	require.Equal(t, order1.PaymentMethod, order2.PaymentMethod)
	require.Equal(t, order1.TotalCost, order2.TotalCost)

	require.WithinDuration(t, order1.OrderDate, order2.OrderDate, time.Second)
}

func TestListOrderByUserId(t *testing.T) {
	order := createRandomOrder(t)

	orders, err := testQueries.ListOrderByUserId(context.Background(), order.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, orders)

	for _, order := range orders {
		require.NotEmpty(t, order)
		require.Equal(t, order.UserID, order.UserID)
	}
}
