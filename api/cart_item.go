package api

import (
	"database/sql"
	"encoding/json"
	"net/http"

	db "github.com/alifdwt/synapsis-backend-challenge/db/sqlc"
	"github.com/alifdwt/synapsis-backend-challenge/token"
	"github.com/gin-gonic/gin"
)

type deleteCartItemRequest struct {
	ProductID string `uri:"productId" binding:"required"`
}

// deleteCartItem godoc
// @Summary      Delete cart item
// @Description  Delete cart item
// @Tags         cart
// @Accept       json
// @Produce      json
// @Param        productId path string true "Product ID"
// @Success      200 {object} cartWithCartItemsResponse
// @Security	 BearerAuth
// @Router       /cart-items/{productId} [delete]
func (server *Server) deleteCartItem(ctx *gin.Context) {
	var req deleteCartItemRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

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

	var cartItems []db.CartItem
	err = json.Unmarshal(cart.CartItems, &cartItems)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Find the cart item where the product id matches
	var cartItem db.CartItem
	for _, item := range cartItems {
		if item.ProductID == req.ProductID {
			cartItem = item
			break
		}
	}

	if cartItem.ProductID == "" {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	arg := db.DeleteCartItemParams{
		CartID:    cart.ID,
		ProductID: req.ProductID,
	}

	err = server.store.DeleteCartItem(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	shoppingCart, _ := server.store.GetShoppingCartWithCartItems(ctx, authPayload.Issuer)
	ctx.JSON(http.StatusOK, shoppingCart)
}
