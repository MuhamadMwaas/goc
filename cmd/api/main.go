package main

import (
	"context"
	"database/sql"
	"os" // New import
	"time"

	"github.com/futek/donation-campaign/internal/application"
	"github.com/futek/donation-campaign/internal/config"
	"github.com/futek/donation-campaign/internal/infrastructure/http"
	"github.com/futek/donation-campaign/internal/infrastructure/persistence/postgres"
	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/rs/zerolog" // New import
	"github.com/rs/zerolog/log"
)

// @title Donation Campaign API
// @version 1.0
// @description This is the API for the Donation Campaign application.
// @host localhost:8080
// @BasePath /
func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	cfg, err := config.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load configuration")
	}

	db, err := sql.Open("pgx", cfg.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to connect to database")
	}
	defer db.Close()

	userRepo := postgres.NewUserRepository(db)
	userService := application.NewUserService(userRepo)

	r := http.NewRouter(userService)

	// This is a sample health check, we need to move it to a proper handler
	// @Summary Health check
	// @Description Check if the API is running
	// @Tags health
	// @Accept  json
	// @Produce  json
	// @Success 200 {object} map[string]string
	// @Failure 500 {object} map[string]string
	// @Router /health [get]
	r.GET("/health", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), 1*time.Second)
		defer cancel()

		if err := db.PingContext(ctx); err != nil {
			c.JSON(500, gin.H{"status": "unhealthy"})
			return
		}

		c.JSON(200, gin.H{"status": "healthy"})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
