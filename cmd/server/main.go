package main

import (
	"fmt"
	"os"

	"github.com/alifslhdn/go-backend-boilerplate-lite/internal/config"
	"github.com/alifslhdn/go-backend-boilerplate-lite/internal/models"
	"github.com/alifslhdn/go-backend-boilerplate-lite/internal/routes"
	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env
	err := godotenv.Load()
	if err != nil {
		fmt.Println("⚠️  No .env file found")
	}

	// Connect DB
	config.ConnectDB()

	// Auto migrate models
	config.DB.AutoMigrate(&models.User{})

	// Setup Gin
	r := gin.Default()
	routes.RegisterRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("🚀 Server running on http://localhost:" + port)
	r.Run(":" + port)
}
