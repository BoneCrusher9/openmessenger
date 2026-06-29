package http

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

func New(logger *slog.Logger) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	router.Use(
		gin.Recovery(),
		LoggerMiddleware(logger),
	)

	RegisterRoutes(router)

	return router
}
