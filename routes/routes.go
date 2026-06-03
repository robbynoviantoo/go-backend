package routes

import (
	"go-backend/handler"
	"go-backend/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	// Add CORS middleware
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173", "http://localhost:5174"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	r.Use(cors.New(config))

	r.POST("/register", handler.Register)
	r.POST("/login", handler.Login)

	auth := r.Group("/api")
	auth.Use(middleware.AuthMiddleware())

	auth.GET("/users", handler.GetUsers)
	auth.GET("/me", handler.Me)
	auth.PUT("/users/:id", handler.UpdateUser)
	auth.DELETE("/users/:id", handler.DeleteUser)

	auth.POST("/items", handler.CreateItem)
	// auth.GET("/items", handler.GetItems)
	auth.GET("/items", handler.GetAllItems)
	auth.DELETE("/items/:id", handler.DeleteItem)

	auth.POST("/product-categories", handler.CreateProductCategory)
	auth.GET("/product-categories", handler.GetProductCategories)
	auth.PUT("/product-categories/:id", handler.UpdateProductCategory)
	auth.DELETE("/product-categories/:id", handler.DeleteProductCategory)

	auth.POST("/products", handler.CreateProduct)
	auth.GET("/products", handler.GetProducts)
	auth.GET("/products/:id", handler.GetProductByID)
	auth.PUT("/products/:id", handler.UpdateProduct)
	auth.DELETE("/products/:id", handler.DeleteProduct)

	return r
}
