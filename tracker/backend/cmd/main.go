package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"tracker-app/backend/internal/config"
	httpserver "tracker-app/backend/internal/http-server"
	"tracker-app/backend/internal/http-server/handler"
	"tracker-app/backend/internal/repository"
	"tracker-app/backend/internal/repository/postgres"
	"tracker-app/backend/internal/service"
)

func main() {
	cfg := config.MustLoad()

	db, err := postgres.NewPostrgesDb(cfg.Database)
	if err != nil {
		log.Fatalf("failed to connect to database: %s", err)
	}

	log := setupLogger()

	repos := repository.NewRepository(db)
	service := service.NewService(log, repos, db)
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

// setupLogger initializes and returns a new logger instance.
func setupLogger() *slog.Logger {
	var log *slog.Logger

	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	})

	log = slog.New(handler)

	return log
}
