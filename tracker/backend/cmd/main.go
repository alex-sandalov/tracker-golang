package main

import (
	"context"
	"effective-mobile-golang/backend/internal/config"
	httpserver "effective-mobile-golang/backend/internal/http-server"
	"effective-mobile-golang/backend/internal/http-server/handler"
	"effective-mobile-golang/backend/internal/service"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger()

	service := service.NewService()
	handler := handler.NewHandler(service)

	fmt.Println(cfg.Server)
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

func setupLogger() *slog.Logger {
	var log *slog.Logger

	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	})

	log = slog.New(handler)

	return log
}
