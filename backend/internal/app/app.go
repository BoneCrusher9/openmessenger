package app

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/BoneCrusher9/openmessenger/backend/internal/config"
	"github.com/BoneCrusher9/openmessenger/backend/internal/database"
	transporthttp "github.com/BoneCrusher9/openmessenger/backend/internal/transport/http"
)

type App struct {
	Config *config.Config
	Logger *slog.Logger
	Router *gin.Engine
	Server *http.Server
	DB     *pgxpool.Pool
}

func New(cfg *config.Config, logger *slog.Logger) (*App, error) {
	db, err := database.New(cfg.Database)
	if err != nil {
		return nil, err
	}

	logger.Info("connected to PostgreSQL")

	router := transporthttp.New(logger)

	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)

	server := &http.Server{
		Addr:              addr,
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	return &App{
		Config: cfg,
		Logger: logger,
		Router: router,
		Server: server,
		DB:     db,
	}, nil
}

func (a *App) Run() error {
	a.Logger.Info("HTTP server started", "address", a.Server.Addr)

	err := a.Server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

func (a *App) Shutdown(ctx context.Context) error {
	a.Logger.Info("shutting down HTTP server")

	if err := a.Server.Shutdown(ctx); err != nil {
		return err
	}

	if a.DB != nil {
		a.DB.Close()
		a.Logger.Info("PostgreSQL connection closed")
	}

	return nil
}
