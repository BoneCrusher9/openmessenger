package main

import (
	"log"

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

	application := app.New(cfg, logg)

	if err := application.Run(); err != nil {
		log.Fatal(err)
	}
}
