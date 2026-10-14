package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	apiGroup := r.Group("/api/v1")
	{
		apiGroup.GET("/status", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "online"})
		})
	}
}
