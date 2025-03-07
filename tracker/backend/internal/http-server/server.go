package httpserver

import (
	"context"
	"tracker-app/backend/internal/config"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(cfg config.Server, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + cfg.Port,
		Handler:        handler,
		MaxHeaderBytes: cfg.MaxHeaderBytes,
		ReadTimeout:    cfg.ReadTimeout,
		WriteTimeout:   cfg.WriteTimeout,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
