package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/alifdwt/synapsis-backend-challenge/util"
	"github.com/stretchr/testify/require"
)

func createRandomCartItem(t *testing.T, cartId string) CartItem {
	category := createRandomCategory(t)

	product := createRandomProduct(t, category.ID)

	arg := CreateCartItemParams{
		CartID:    cartId,
		ProductID: product.ID,
		Quantity:  util.RandomInt(1, 10),
	}

	cartItem, err := testQueries.CreateCartItem(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, cartItem)

	require.Equal(t, arg.CartID, cartItem.CartID)
	require.Equal(t, arg.ProductID, cartItem.ProductID)
	require.Equal(t, arg.Quantity, cartItem.Quantity)

	return cartItem
}

func TestCreateCartItem(t *testing.T) {
	shoppingCart := createRandomShoppingCart(t)
	createRandomCartItem(t, shoppingCart.ID)
}

func TestGetCartItemsByUserID(t *testing.T) {
	shoppingCart := createRandomShoppingCart(t)
	var lastCartItem CartItem
	for i := 0; i < 10; i++ {
		lastCartItem = createRandomCartItem(t, shoppingCart.ID)
	}

	cartItems, err := testQueries.GetCartItemsByUserID(context.Background(), lastCartItem.CartID)
	require.NoError(t, err)
	require.NotEmpty(t, cartItems)

	for _, cartItem := range cartItems {
		require.NotEmpty(t, cartItem)
		require.Equal(t, lastCartItem.CartID, cartItem.CartID)
	}
}

func TestUpdateCartItem(t *testing.T) {
	shoppingCart := createRandomShoppingCart(t)
	cartItem1 := createRandomCartItem(t, shoppingCart.ID)

	arg := UpdateCartItemParams{
		CartID:    cartItem1.CartID,
		Quantity:  util.RandomInt(1, 10),
		ProductID: cartItem1.ProductID,
	}

	cartItem2, err := testQueries.UpdateCartItem(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, cartItem2)

	require.NotEqual(t, cartItem1.Quantity, cartItem2.Quantity)
}

func TestDeleteCartItem(t *testing.T) {
	shoppingCart := createRandomShoppingCart(t)
	cartItem1 := createRandomCartItem(t, shoppingCart.ID)

	arg := DeleteCartItemParams{
		CartID:    cartItem1.CartID,
		ProductID: cartItem1.ProductID,
	}

	err := testQueries.DeleteCartItem(context.Background(), arg)
	require.NoError(t, err)

	cartItem2, err := testQueries.GetCartItem(context.Background(), cartItem1.CartID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, cartItem2)
}
