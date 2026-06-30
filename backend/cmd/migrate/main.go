package main

import (
	"log"

	"github.com/BoneCrusher9/openmessenger/backend/internal/config"
	"github.com/BoneCrusher9/openmessenger/backend/internal/migrations"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Applying migrations...")

	if err := migrations.Run(cfg.Database); err != nil {
		log.Fatal(err)
	}

	log.Println("Migrations applied successfully.")
}
