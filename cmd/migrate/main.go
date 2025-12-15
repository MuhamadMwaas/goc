package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/futek/donation-campaign/internal/config"
	"github.com/futek/donation-campaign/internal/infrastructure/migration"
)

func main() {
	var migrationsURL, command string

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v\n", err)
	}

	flag.StringVar(&migrationsURL, "migrations-url", "file://migration", "Migrations folder URL")
	flag.StringVar(&command, "command", "up", "Migration command (up or down)")
	flag.Parse()

	if err := migration.Run(cfg.DBSource, migrationsURL, command); err != nil {
		fmt.Fprintf(os.Stderr, "Migration failed: %v\n", err)
		os.Exit(1)
	}
}
