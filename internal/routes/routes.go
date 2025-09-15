package routes

import (
	"net/http"

	"github.com/alifslhdn/go-backend-boilerplate-lite/internal/controllers"
	"github.com/alifslhdn/go-backend-boilerplate-lite/internal/middlewares"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes register all app routes
func RegisterRoutes(r *gin.Engine) {
	// Default root route
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Go Backend Lite is running 🚀",
		})
	})

	// Health check
	r.GET("/health", controllers.HealthCheck)

	// Auth routes
	auth := r.Group("/auth")
	{
		auth.POST("/register", controllers.RegisterUser)
		auth.POST("/login", controllers.LoginUser)
	}

	// Protected example route
	protected := r.Group("/users")
	protected.Use(middlewares.JWTAuth())
	{
		protected.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "This is protected route"})
		})
	}
}
