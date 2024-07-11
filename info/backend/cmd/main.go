package main

import (
	"context"
	"info-golang/backend/internal/config"
	httpserver "info-golang/backend/internal/http-server"
	"info-golang/backend/internal/http-server/handler"
	"info-golang/backend/internal/repository"
	"info-golang/backend/internal/repository/postgres"
	"info-golang/backend/internal/service"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.MustLoad()

	db, err := postgres.NewPostrgesDb(cfg.Database)
	if err != nil {
		log.Fatalf("failed to connect database: %s", err)
	}
	log := setupLogger()

	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handler := handler.NewHandler(service)

	srv := new(httpserver.Server)
	go func() {
		if err := srv.Run(cfg.Server, handler.InitRoutes(&cfg.CORS, log)); err != nil {
			log.Error("failed to run http server: %s", err)
		}
	}()

	log.Info("server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Info("shutdown server")
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Error("failed to shutdown: %s", err)
	}

	log.Info("server exiting")
}

// setupLogger initializes and configures the logger.
//
// It returns a pointer to the newly created logger.
func setupLogger() *slog.Logger {
	var log *slog.Logger

	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	})

	log = slog.New(handler)

	return log
}
