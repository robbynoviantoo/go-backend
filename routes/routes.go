package routes

import (
	"go-backend/handler"
	"go-backend/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.POST("/register", handler.Register)
	r.POST("/login", handler.Login)

	auth := r.Group("/api")
	auth.Use(middleware.AuthMiddleware())

	auth.GET("/users", handler.GetUsers)
	auth.PUT("/users/:id", handler.UpdateUser)
	auth.DELETE("/users/:id", handler.DeleteUser)

	auth.POST("/items", handler.CreateItem)
	auth.GET("/items", handler.GetItems)
	auth.DELETE("/items/:id", handler.DeleteItem)

	return r
}
