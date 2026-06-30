package http

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

func New(logger *slog.Logger, handler *Handler) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	router.Use(
		gin.Recovery(),
		LoggerMiddleware(logger),
	)

	handler.RegisterRoutes(router)

	return router
}
