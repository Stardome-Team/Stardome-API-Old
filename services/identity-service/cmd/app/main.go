package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/Blac-Panda/Stardome-API/services/identity-service/pkg/database"

	"github.com/Blac-Panda/Stardome-API/services/identity-service/internal/auth"
	"github.com/Blac-Panda/Stardome-API/services/identity-service/internal/config"
	"github.com/Blac-Panda/Stardome-API/services/identity-service/pkg/log"

	_ "github.com/lib/pq"
)

var version = "0.0.1"
var flagConfig = flag.String("config", "./config/local.yml", "Path to configuration file")

func main() {
	flag.Parse()

	// create application logger tag with application version
	logger := log.New().With(nil, "version", version)

	// load application configurations
	cfg, err := config.Load(*flagConfig)

	if err != nil {
		logger.Errorf("Failed to load application configuration: %s", err)
		os.Exit(1)
	}

	// connect to the database
	db, err := database.OpenConnection(cfg)

	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}

	srv := &http.Server{
		Addr:    ":1010",
		Handler: buildHandler(cfg, db, logger),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error(err)
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logger.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Errorf("Server Shutdown: %s", err)
	}
	logger.Info("Server exiting")
}

func buildHandler(cfg *config.Config, db *database.DB, logger log.Logger) *gin.Engine {
	router := gin.Default()

	group := router.Group("/api")

	auth.CreateHandlers(group, auth.NewService(auth.NewRepository(db, logger), logger), logger)

	return router
}
