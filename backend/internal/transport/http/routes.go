package http

import "github.com/gin-gonic/gin"

func (h *Handler) RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")

	api.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
			"name":   "OpenMessenger",
		})
	})

	auth := api.Group("/auth")
	{
		auth.POST("/register", h.Register)
	}
}
