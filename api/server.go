package api

import (
	db "github.com/alifdwt/synapsis-backend-challenge/db/sqlc"
	"github.com/alifdwt/synapsis-backend-challenge/token"
	"github.com/alifdwt/synapsis-backend-challenge/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.TokenSymetricKey)
	if err != nil {
		return nil, err
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("payment_method", validPaymentMethod)
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)
	router.GET("/products/:id", server.getProduct)
	router.GET("/products", server.listProducts)
	router.GET("/categories", server.listCategories)
	router.GET("/categories/:id", server.getCategory)
	router.GET("/users/:id", server.getUser)
	router.GET("/users", server.listUsers)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	authRoutes.POST("/products", server.createProduct)
	authRoutes.PUT("/products/:id", server.updateProduct)
	authRoutes.DELETE("/products/:id", server.deleteProduct)
	authRoutes.POST("/categories", server.createCategory)
	authRoutes.GET("/cart", server.getCart)
	authRoutes.POST("/cart", server.createCart)
	authRoutes.DELETE("/cart", server.deleteCart)
	authRoutes.POST("/order", server.createOrder)
	authRoutes.GET("/orders", server.listOrders)
	authRoutes.DELETE("/cart-items/:productId", server.deleteCartItem)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
