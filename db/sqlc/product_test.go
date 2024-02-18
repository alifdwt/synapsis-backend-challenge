package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/alifdwt/synapsis-backend-challenge/util"
	"github.com/stretchr/testify/require"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func createRandomProduct(t *testing.T, categoryId string) Product {
	user := createRandomUser(t)

	productName := util.RandomString(6)
	arg := CreateProductParams{
		ID:     productName,
		UserID: user.Username,
		Title:  cases.Title(language.English).String(productName),
		Description: sql.NullString{
			String: util.RandomString(6),
			Valid:  true,
		},
		Price:      util.RandomMoney(),
		CategoryID: categoryId,
	}

	product, err := testQueries.CreateProduct(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, product)

	require.Equal(t, arg.ID, product.ID)
	require.Equal(t, arg.UserID, product.UserID)
	require.Equal(t, arg.Title, product.Title)
	require.Equal(t, arg.Description.String, product.Description.String)
	require.Equal(t, arg.Price, product.Price)
	require.Equal(t, arg.CategoryID, product.CategoryID)

	require.NotZero(t, product.CreatedAt)

	return product
}

func TestCreateProduct(t *testing.T) {
	category := createRandomCategory(t)
	createRandomProduct(t, category.ID)
}

func TestGetProduct(t *testing.T) {
	category := createRandomCategory(t)
	product1 := createRandomProduct(t, category.ID)

	product2, err := testQueries.GetProduct(context.Background(), product1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, product2)

	require.Equal(t, product1.ID, product2.ID)
	require.Equal(t, product1.UserID, product2.UserID)
	require.Equal(t, product1.Title, product2.Title)
	require.Equal(t, product1.Description.String, product2.Description.String)
	require.Equal(t, product1.Price, product2.Price)
	require.Equal(t, product1.CategoryID, product2.CategoryID)

	require.WithinDuration(t, product1.CreatedAt, product2.CreatedAt, time.Second)
}

func TestListProducts(t *testing.T) {
	category := createRandomCategory(t)
	for i := 0; i < 10; i++ {
		createRandomProduct(t, category.ID)
	}

	arg := ListProductsParams{
		Limit:  5,
		Offset: 0,
	}

	products, err := testQueries.ListProducts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, products)

	for _, product := range products {
		require.NotEmpty(t, product)
	}
}

func TestUpdateProduct(t *testing.T) {
	category := createRandomCategory(t)
	product1 := createRandomProduct(t, category.ID)

	productName := util.RandomString(6)
	arg := UpdateProductParams{
		ID:     product1.ID,
		ID_2:   productName,
		UserID: product1.UserID,
		Title:  cases.Title(language.English).String(productName),
		Price:  util.RandomMoney(),
		Description: sql.NullString{
			String: util.RandomString(20),
			Valid:  true,
		},
		CategoryID: product1.CategoryID,
	}

	product2, err := testQueries.UpdateProduct(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, product2)

	require.NotEqual(t, product1.ID, product2.ID)
	require.NotEqual(t, product1.Title, product2.Title)
	require.NotEqual(t, product1.Price, product2.Price)
	require.NotEqual(t, product1.Description.String, product2.Description.String)
	require.Equal(t, product1.CategoryID, product2.CategoryID)
	require.Equal(t, product1.UserID, product2.UserID)

	require.NotZero(t, product1.UpdatedAt)
	require.WithinDuration(t, product2.CreatedAt, product2.CreatedAt, time.Second)
}

func TestDeleteProduct(t *testing.T) {
	category := createRandomCategory(t)
	product1 := createRandomProduct(t, category.ID)

	err := testQueries.DeleteProduct(context.Background(), product1.ID)
	require.NoError(t, err)

	product2, err := testQueries.GetProduct(context.Background(), product1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, product2)
}
