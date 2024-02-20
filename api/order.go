package api

import (
	"database/sql"
	"errors"
	"net/http"

	db "github.com/alifdwt/synapsis-backend-challenge/db/sqlc"
	"github.com/alifdwt/synapsis-backend-challenge/token"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type createOrderRequest struct {
	PaymentMethod string `json:"payment_method" binding:"required,payment_method"`
}

func (server *Server) createOrder(ctx *gin.Context) {
	var req createOrderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	cart, err := server.store.GetShoppingCartWithCartItems(ctx, authPayload.Issuer)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(errors.New("shopping cart not found")))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Get price per cart item and total price
	shoppingCart := newCartWithCartItemsResponse(cart)
	var totalPrice int64
	for i := 0; i < len(shoppingCart.CartItems); i++ {
		// Get product price
		product, _ := server.store.GetProduct(ctx, shoppingCart.CartItems[i].ProductID)
		totalPrice += product.Price * int64(shoppingCart.CartItems[i].Quantity)
	}

	createOrderArg := db.CreateOrderParams{
		UserID:        authPayload.Issuer,
		PaymentMethod: req.PaymentMethod,
		TotalCost:     totalPrice,
	}

	order, err := server.store.CreateOrder(ctx, createOrderArg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	for i := 0; i < len(shoppingCart.CartItems); i++ {
		product, _ := server.store.GetProduct(ctx, shoppingCart.CartItems[i].ProductID)
		createOrderItemArg := db.CreateOrderItemParams{
			OrderID:         order.ID,
			ProductID:       shoppingCart.CartItems[i].ProductID,
			Quantity:        int64(shoppingCart.CartItems[i].Quantity),
			PriceAtPurchase: product.Price * int64(shoppingCart.CartItems[i].Quantity),
		}
		_, err = server.store.CreateOrderItem(ctx, createOrderItemArg)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	}

	err = server.store.DeleteShoppingCart(ctx, authPayload.Issuer)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, order)
}

func (server *Server) listOrders(ctx *gin.Context) {
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	orders, err := server.store.ListOrders(ctx, authPayload.Issuer)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, orders)
}
