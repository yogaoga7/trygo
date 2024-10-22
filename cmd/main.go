package main

import (
	"log"
	"trygo/config"
	"trygo/internal/handlers"
	"trygo/internal/repositories"
	"trygo/internal/routes"
	"trygo/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// Muat konfigurasi
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Repository dan Service
	userRepo := repositories.NewUserRepository(cfg.DB)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	// Gin router
	r := gin.Default()

	// Setup routes
	routes.SetupUserRoutes(r, userHandler)

	// Jalankan server
	r.Run(":8080")
}
