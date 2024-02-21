package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	db "github.com/alifdwt/synapsis-backend-challenge/db/sqlc"
	"github.com/alifdwt/synapsis-backend-challenge/util"
	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type userResponse struct {
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	FullName  string    `json:"full_name"`
	CreatedAt time.Time `json:"created_at"`
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
		Username:  user.Username,
		Email:     user.Email,
		FullName:  user.FullName,
		CreatedAt: user.CreatedAt,
	}
}

type userWithProductsResponse struct {
	Username  string       `json:"username"`
	Email     string       `json:"email"`
	FullName  string       `json:"full_name"`
	CreatedAt time.Time    `json:"created_at"`
	Products  []db.Product `json:"products"`
}

func newUserWithProductsResponse(user db.UsersWithProduct) userWithProductsResponse {
	var product []db.Product
	err := json.Unmarshal(user.Products, &product)
	if err != nil {
		return userWithProductsResponse{}
	}
	if product[0].ID == "" {
		product = []db.Product{}
	}
	return userWithProductsResponse{
		Username:  user.Username,
		Email:     user.Email,
		FullName:  user.FullName,
		CreatedAt: user.CreatedAt,
		Products:  product,
	}
}

// createUser godoc
// @Summary      Create new user
// @Description  Add a new user
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        user  body      createUserRequest  true  "User"
// @Success      200   {object}  userResponse
// @Router       /users [post]
func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Username:       req.Username,
		FullName:       req.FullName,
		Email:          req.Email,
		HashedPassword: hashedPassword,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			ctx.JSON(http.StatusForbidden, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := newUserResponse(user)
	ctx.JSON(http.StatusOK, response)
}

type loginUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

type loginUserResponse struct {
	AccessToken string                   `json:"access_token"`
	User        userWithProductsResponse `json:"user"`
}

// loginUser godoc
// @Summary      Login user
// @Description  Login user
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        user  body      loginUserRequest  true  "User"
// @Success      200   {object}  loginUserResponse
// @Router       /users/login [post]
func (server *Server) loginUser(ctx *gin.Context) {
	var req loginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	user, err := server.store.GetUserWithProducts(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = util.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	accessToken, err := server.tokenMaker.CreateToken(
		user.Username,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := loginUserResponse{
		AccessToken: accessToken,
		User:        newUserWithProductsResponse(user),
	}
	ctx.JSON(http.StatusOK, response)
}

type getUserRequest struct {
	ID string `uri:"id" binding:"required,min=1"`
}

// getUser godoc
// @Summary      Get user
// @Description  Get user
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        user  path      string  true  "User"
// @Success      200   {object}  userWithProductsResponse
// @Router       /users/{id} [get]
func (server *Server) getUser(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUserWithProducts(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, newUserWithProductsResponse(user))
}

type listUsersRequest struct {
	PageID   int32 `form:"page_id"`
	PageSize int32 `form:"page_size"`
}

// listUsers godoc
// @Summary      List users
// @Description  List users
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        user  query     listUsersRequest  true  "User"
// @Success      200   {array}   userWithProductsResponse
// @Router       /users [get]
func (server *Server) listUsers(ctx *gin.Context) {
	var req listUsersRequest
	req.PageID = 1
	req.PageSize = 10
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListUserWithProductsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	users, err := server.store.ListUserWithProducts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	var listUsersResponse []userWithProductsResponse
	for _, user := range users {
		listUsersResponse = append(listUsersResponse, newUserWithProductsResponse(user))
	}

	ctx.JSON(http.StatusOK, listUsersResponse)
}
