package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/BoneCrusher9/openmessenger/backend/internal/app"
	"github.com/BoneCrusher9/openmessenger/backend/internal/config"
	"github.com/BoneCrusher9/openmessenger/backend/internal/logger"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	logg := logger.New()

	application, err := app.New(cfg, logg)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		if err := application.Run(); err != nil {
			log.Fatal(err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	logg.Info("shutdown signal received")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := application.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	logg.Info("server stopped")
}
