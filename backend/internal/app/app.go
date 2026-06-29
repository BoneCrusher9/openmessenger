package app

import (
	"fmt"
	"log/slog"

	"github.com/gin-gonic/gin"

	"github.com/BoneCrusher9/openmessenger/backend/internal/config"
)

type App struct {
	Config *config.Config
	Logger *slog.Logger
	Router *gin.Engine
}

func New(cfg *config.Config, logger *slog.Logger) *App {
	router := gin.New()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
			"name":   "OpenMessenger",
		})
	})

	return &App{
		Config: cfg,
		Logger: logger,
		Router: router,
	}
}

func (a *App) Run() error {
	addr := fmt.Sprintf("%s:%d", a.Config.Server.Host, a.Config.Server.Port)

	a.Logger.Info("server started", "address", addr)

	return a.Router.Run(addr)
}
