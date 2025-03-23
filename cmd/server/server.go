package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	app "github.com/Sergey-Polishchenko/simple-api/internal/application"
	"github.com/Sergey-Polishchenko/simple-api/internal/config"
	repo "github.com/Sergey-Polishchenko/simple-api/internal/infrastructure/postgres"
	"github.com/Sergey-Polishchenko/simple-api/internal/interfaces/http"
	httpserver "github.com/Sergey-Polishchenko/simple-api/internal/interfaces/http/server"
	"github.com/Sergey-Polishchenko/simple-api/internal/pkg/logger"
)

func main() {
	logger := logger.NewZapLogger()

	env, err := config.Load()
	if err != nil {
		logger.Error("can't load .env", "error", err)
		return
	}
	port := fmt.Sprintf(":%s", env.Port)

	db, err := gorm.Open(postgres.Open(env.DB.ConnString()), &gorm.Config{})
	if err != nil {
		logger.Error("can't connect to postgres database", "error", err)
		return
	}

	repo, err := repo.NewUserRepo(db)
	if err != nil {
		logger.Error("can't automigrate postgres database", "error", err)
		return
	}

	app := app.NewUserApp(repo, logger)
	router := http.NewRouter(app)
	httpServer := httpserver.New(port, router)

	errChan := make(chan error, 1)

	go func() {
		logger.Info("Starting HTTP server", "address", port)
		if err := httpServer.Start(); err != nil {
			errChan <- fmt.Errorf("HTTP server error: %w", err)
		}
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case <-interrupt:
		logger.Info("Received interrupt signal, shutting down...")
	case err := <-errChan:
		logger.Error("Server error", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := httpServer.Stop(ctx); err != nil {
		logger.Error("HTTP shutdown error", err)
	}

	logger.Info("Server stopped")
}
