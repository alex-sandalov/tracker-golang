package httpserver

import (
	"context"
	"info-golang/backend/internal/config"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

// Run starts the HTTP server with the provided configuration and handler.
//
// Parameters:
// - cfg: The server configuration.
// - handler: The HTTP handler.
//
// Returns:
// - error: An error if the server fails to start.
func (s *Server) Run(cfg config.Server, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + cfg.Port, // Listen on the specified port.
		Handler:        handler,        // Use the provided handler.
		MaxHeaderBytes: cfg.MaxHeaderBytes,
		ReadTimeout:    cfg.ReadTimeout,
		WriteTimeout:   cfg.WriteTimeout,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
