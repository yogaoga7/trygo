package routes

import (
	"trygo/internal/handlers"
	"trygo/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.Engine, userHandler *handlers.UserHandler) {
	userRoutes := r.Group("/users")
	{
		// Rute ini menggunakan AuthMiddleware
		userRoutes.POST("", middleware.AuthMiddleware, userHandler.CreateUser)

		// Rute ini tidak menggunakan middleware
		userRoutes.GET("/:id", userHandler.GetUser)

		// Rute ini juga menggunakan AuthMiddleware
		userRoutes.PUT("/:id", middleware.AuthMiddleware, userHandler.UpdateUser)

		// Rute ini tanpa middleware
		userRoutes.DELETE("/:id", userHandler.DeleteUser)
	}
}
