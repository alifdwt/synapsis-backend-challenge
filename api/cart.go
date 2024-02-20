package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	db "github.com/alifdwt/synapsis-backend-challenge/db/sqlc"
	"github.com/alifdwt/synapsis-backend-challenge/token"
	"github.com/gin-gonic/gin"
)

type CartRequest struct {
	UserID string `json:"user_id"`
}

type createCartRequest struct {
	ProductID string `json:"product_id"`
}

type cartWithCartItemsResponse struct {
	ID        string        `json:"id"`
	UserID    string        `json:"user_id"`
	CreatedAt time.Time     `json:"created_at"`
	CartItems []db.CartItem `json:"cart_items"`
}

func newCartWithCartItemsResponse(cart db.ShoppingCartWithCartItem) cartWithCartItemsResponse {
	var cartItems []db.CartItem
	err := json.Unmarshal(cart.CartItems, &cartItems)
	if err != nil {
		return cartWithCartItemsResponse{}
	}
	return cartWithCartItemsResponse{
		ID:        cart.ID,
		UserID:    cart.UserID,
		CreatedAt: cart.CreatedAt,
		CartItems: cartItems,
	}
}

func (server *Server) createCart(ctx *gin.Context) {
	var req createCartRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	cart, err := server.store.GetShoppingCartWithCartItems(ctx, authPayload.Issuer)

	if err != nil {
		if err == sql.ErrNoRows {
			cart, err := server.store.CreateShoppingCart(ctx, authPayload.Issuer)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, errorResponse(err))
				return
			}
			createCartArg := db.CreateCartItemParams{
				CartID:    cart.ID,
				ProductID: req.ProductID,
				Quantity:  1,
			}
			_, err = server.store.CreateCartItem(ctx, createCartArg)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, errorResponse(err))
				return
			}
			shoppingCart, _ := server.store.GetShoppingCartWithCartItems(ctx, authPayload.Issuer)
			ctx.JSON(http.StatusOK, shoppingCart)
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	shoppingCart := newCartWithCartItemsResponse(cart)
	productFound := false
	for i := 0; i < len(shoppingCart.CartItems); i++ {
		if shoppingCart.CartItems[i].ProductID == req.ProductID {
			shoppingCart.CartItems[i].Quantity = shoppingCart.CartItems[i].Quantity + 1
			updateCartArg := db.UpdateCartItemParams{
				CartID:    shoppingCart.CartItems[i].CartID,
				Quantity:  shoppingCart.CartItems[i].Quantity,
				ProductID: req.ProductID,
			}
			_, err = server.store.UpdateCartItem(ctx, updateCartArg)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, errorResponse(err))
				return
			}
			productFound = true
			break
		}
	}

	if !productFound {
		createCartArg := db.CreateCartItemParams{
			CartID:    shoppingCart.ID,
			ProductID: req.ProductID,
			Quantity:  1,
		}
		_, err = server.store.CreateCartItem(ctx, createCartArg)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	}

	shoppingCart2, _ := server.store.GetShoppingCartWithCartItems(ctx, authPayload.Issuer)
	ctx.JSON(http.StatusOK, shoppingCart2)
}

// type getCartRequest struct {
// 	UserID string `uri:"user_id" binding:"required"`
// }

func (server *Server) getCart(ctx *gin.Context) {
	// var req getCartRequest
	// if err := ctx.ShouldBindUri(&req); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, errorResponse(err))
	// 	return
	// }

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	cart, err := server.store.GetShoppingCartWithCartItems(ctx, authPayload.Issuer)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, cart)
}

func (server *Server) deleteCart(ctx *gin.Context) {
	// var req getCartRequest
	// if err := ctx.ShouldBindUri(&req); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, errorResponse(err))
	// 	return
	// }

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	cart, err := server.store.GetShoppingCartWithCartItems(ctx, authPayload.Issuer)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = server.store.DeleteShoppingCart(ctx, cart.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, cart)
}
