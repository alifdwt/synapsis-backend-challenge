package api

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"time"

	db "github.com/alifdwt/synapsis-backend-challenge/db/sqlc"
	"github.com/alifdwt/synapsis-backend-challenge/token"
	"github.com/alifdwt/synapsis-backend-challenge/util"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type createProductRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Price       int32  `json:"price" binding:"required"`
	CategoryID  string `json:"category_id" binding:"required"`
}

type productResponse struct {
	ID          string      `json:"id"`
	UserID      string      `json:"user_id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Price       int32       `json:"price"`
	Category    db.Category `json:"category"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

// createProduct godoc
// @Summary      Create new product
// @Description  Add a new product
// @Tags         product
// @Accept       json
// @Produce      json
// @Param        product  body      createProductRequest  true  "Product"
// @Success      200   {object}  productResponse
// @Security	 BearerAuth
// @Router       /products [post]
func (server *Server) createProduct(ctx *gin.Context) {
	var req createProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	_, err := server.store.GetCategory(ctx, req.CategoryID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(fmt.Errorf("category not found: %s", req.CategoryID)))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.CreateProductParams{
		ID:          util.ConvertTitleToId(req.Title),
		UserID:      authPayload.Issuer,
		Title:       req.Title,
		Description: req.Description,
		Price:       int64(req.Price),
		CategoryID:  req.CategoryID,
	}

	product, err := server.store.CreateProduct(ctx, arg)
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

	response := productResponse{
		ID:          product.ID,
		UserID:      product.UserID,
		Title:       product.Title,
		Description: product.Description,
		Price:       int32(product.Price),
		Category: db.Category{
			ID:   product.CategoryID,
			Name: cases.Title(language.English).String(product.CategoryID),
		},
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, response)
}

type getProductRequest struct {
	ID string `uri:"id" binding:"required,min=1"`
}

// getProduct godoc
// @Summary      Get product
// @Description  Get product
// @Tags         product
// @Accept       json
// @Produce      json
// @Param        id  path      string  true  "Product ID"
// @Success      200   {object}  productResponse
// @Router       /products/{id} [get]
func (server *Server) getProduct(ctx *gin.Context) {
	var req getProductRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	product, err := server.store.GetProduct(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := productResponse{
		ID:          product.ID,
		UserID:      product.UserID,
		Title:       product.Title,
		Description: product.Description,
		Price:       int32(product.Price),
		Category: db.Category{
			ID:   product.ID_2,
			Name: product.Name,
		},
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, response)
}

type listProductsRequest struct {
	PageID   int32 `form:"page_id"`
	PageSize int32 `form:"page_size"`
}

// listProducts godoc
// @Summary      List products
// @Description  List products
// @Tags         product
// @Accept       json
// @Produce      json
// @Param        page_id   query     int32  false  "Page ID"
// @Param        page_size query     int32  false  "Page Size"
// @Success      200       {object}  []productResponse
// @Router       /products [get]
func (server *Server) listProducts(ctx *gin.Context) {
	var req listProductsRequest
	req.PageID = 1
	req.PageSize = 10
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListProductsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	products, err := server.store.ListProducts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	var listProductsResponse []productResponse
	for _, product := range products {
		listProductsResponse = append(listProductsResponse, productResponse{
			ID:          product.ID,
			UserID:      product.UserID,
			Title:       product.Title,
			Description: product.Description,
			Price:       int32(product.Price),
			Category: db.Category{
				ID:   product.ID_2,
				Name: product.Name,
			},
			CreatedAt: product.CreatedAt,
			UpdatedAt: product.UpdatedAt,
		})
	}

	ctx.JSON(http.StatusOK, listProductsResponse)
}

type updateProductURI struct {
	ID string `uri:"id" binding:"required,min=1"`
}

type updateProductRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int32  `json:"price"`
	CategoryID  string `json:"category_id"`
}

// type updateProductResponse struct {
//     ID          string `json:"id"`
//     UserID      string `json:"user_id"`
//     Title       string `json:"title"`
// 	Description string `json:"description"`
// 	Price       int32  `json:"price"`
// 	CategoryID  string `json:"category_id"`
// }

// updateProduct godoc
// @Summary      Update product
// @Description  Update product
// @Tags         product
// @Accept       json
// @Produce      json
// @Param        id    path      string  true  "Product ID"
// @Success      200   {object}  productResponse
// @Security	 BearerAuth
// @Router       /products/{id} [put]
func (server *Server) updateProduct(ctx *gin.Context) {
	var uri updateProductURI
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var req updateProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	product, err := server.store.GetProduct(ctx, uri.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if product.UserID != authPayload.Issuer {
		err := errors.New("product doesn't belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	_, err = server.store.GetCategory(ctx, req.CategoryID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.UpdateProductParams{
		ID:          uri.ID,
		ID_2:        util.ConvertTitleToId(req.Title),
		UserID:      authPayload.Issuer,
		Title:       req.Title,
		Description: req.Description,
		Price:       int64(req.Price),
		CategoryID:  req.CategoryID,
	}

	updatedProduct, err := server.store.UpdateProduct(ctx, arg)
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

	response := productResponse{
		ID:          updatedProduct.ID,
		UserID:      updatedProduct.UserID,
		Title:       updatedProduct.Title,
		Description: updatedProduct.Description,
		Price:       int32(updatedProduct.Price),
		Category: db.Category{
			ID:   updatedProduct.CategoryID,
			Name: cases.Title(language.English).String(updatedProduct.CategoryID),
		},
		CreatedAt: updatedProduct.CreatedAt,
		UpdatedAt: updatedProduct.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, response)
}

// deleteProduct godoc
// @Summary      Delete product
// @Description  Delete product
// @Tags         product
// @Accept       json
// @Produce      json
// @Param        id  path      string  true  "Product ID"
// @Success      200   {object}  productResponse
// @Security	 BearerAuth
// @Router       /products/{id} [delete]
func (server *Server) deleteProduct(ctx *gin.Context) {
	var req getProductRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	product, err := server.store.GetProduct(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if product.UserID != authPayload.Issuer {
		err := errors.New("product doesn't belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	err = server.store.DeleteProduct(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := productResponse{
		ID:          product.ID,
		UserID:      product.UserID,
		Title:       product.Title,
		Description: product.Description,
		Price:       int32(product.Price),
		Category: db.Category{
			ID:   product.ID_2,
			Name: product.Name,
		},
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, response)
}
