package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck simple health endpoint
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "Go Backend Lite is running 🚀",
	})
}
