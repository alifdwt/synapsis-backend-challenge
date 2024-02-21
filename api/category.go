package api

import (
	"database/sql"
	"encoding/json"
	"net/http"

	db "github.com/alifdwt/synapsis-backend-challenge/db/sqlc"
	"github.com/alifdwt/synapsis-backend-challenge/util"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type createCategoryRequest struct {
	Name string `json:"name"`
}

type categoryResponse struct {
	ID       string       `json:"id"`
	Name     string       `json:"name"`
	Products []db.Product `json:"products"`
}

func newCategoryResponse(category db.CategoriesWithProduct) categoryResponse {
	var products []db.Product
	err := json.Unmarshal(category.Products, &products)
	if err != nil {
		return categoryResponse{}
	}
	if products[0].ID == "" {
		products = []db.Product{}
	}
	return categoryResponse{
		ID:       category.ID,
		Name:     category.Name,
		Products: products,
	}
}

// createCategory godoc
// @Summary      Create new category
// @Description  Add a new category
// @Tags         category
// @Accept       json
// @Produce      json
// @Param        category  body      createCategoryRequest  true  "Category"
// @Success      200   {object}  categoryResponse
// @Security	 BearerAuth
// @Router       /categories [post]
func (server *Server) createCategory(ctx *gin.Context) {
	var req createCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateCategoryParams{
		ID:   util.ConvertTitleToId(req.Name),
		Name: req.Name,
	}

	category, err := server.store.CreateCategory(ctx, arg)
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

	ctx.JSON(http.StatusOK, category)
}

type getCategoryRequest struct {
	ID string `uri:"id" binding:"required"`
}

// getCategory godoc
// @Summary      Get category
// @Description  Get category
// @Tags         category
// @Accept       json
// @Produce      json
// @Param        id  path      string  true  "Category ID"
// @Success      200   {object}  categoryResponse
// @Router       /categories/{id} [get]
func (server *Server) getCategory(ctx *gin.Context) {
	var req getCategoryRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	category, err := server.store.GetCategory(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, newCategoryResponse(category))
}

type listCategoriesRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

// listCategories godoc
// @Summary      List categories
// @Description  List categories
// @Tags         category
// @Accept       json
// @Produce      json
// @Param        category  query     listCategoriesRequest  true  "Category"
// @Success      200   {array}   categoryResponse
// @Router       /categories [get]
func (server *Server) listCategories(ctx *gin.Context) {
	var req listCategoriesRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListCategoriesParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	categories, err := server.store.ListCategories(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	var categoriesResponse []categoryResponse
	for _, category := range categories {
		categoriesResponse = append(categoriesResponse, newCategoryResponse(category))
	}

	ctx.JSON(http.StatusOK, categoriesResponse)
}
