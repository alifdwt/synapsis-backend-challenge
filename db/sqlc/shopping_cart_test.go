package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomShoppingCart(t *testing.T) ShoppingCart {
	user := createRandomUser(t)

	shoppingCart, err := testQueries.CreateShoppingCart(context.Background(), user.Username)
	require.NoError(t, err)
	require.NotEmpty(t, shoppingCart)

	require.Equal(t, user.Username, shoppingCart.UserID)
	require.NotZero(t, shoppingCart.CreatedAt)
	require.NotZero(t, shoppingCart.ID)

	return shoppingCart
}

func TestCreateShoppingCart(t *testing.T) {
	createRandomShoppingCart(t)
}

func TestGetShoppingCart(t *testing.T) {
	shoppingCart1 := createRandomShoppingCart(t)

	shoppingCart2, err := testQueries.GetShoppingCart(context.Background(), shoppingCart1.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, shoppingCart2)

	require.Equal(t, shoppingCart1.UserID, shoppingCart2.UserID)
	require.Equal(t, shoppingCart1.CreatedAt, shoppingCart2.CreatedAt)
	require.Equal(t, shoppingCart1.ID, shoppingCart2.ID)

	require.WithinDuration(t, shoppingCart1.CreatedAt, shoppingCart2.CreatedAt, time.Second)
}

func TestDeleteShoppingCart(t *testing.T) {
	shoppingCart1 := createRandomShoppingCart(t)

	err := testQueries.DeleteShoppingCart(context.Background(), shoppingCart1.UserID)
	require.NoError(t, err)

	shoppingCart2, err := testQueries.GetShoppingCart(context.Background(), shoppingCart1.UserID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, shoppingCart2)
}
