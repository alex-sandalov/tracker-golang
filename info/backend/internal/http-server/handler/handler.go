package handler

import (
	"info-golang/backend/internal/config"
	"info-golang/backend/internal/http-server/middleware"
	"info-golang/backend/internal/service"
	"log/slog"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

// NewHandler initializes a new instance of Handler with the provided service.
//
// Parameters:
// - service: The service to be used by the handler.
// - Returns: The new Handler instance.
func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

// InitRoutes initializes the routes for the handler.
//
// Parameters:
// - cfg: The CORS configuration.
// - Returns: The Gin router.
func (h *Handler) InitRoutes(cfg *config.CORS, log *slog.Logger) *gin.Engine {
	router := gin.Default()

	h.initCORS(cfg, router)
	h.initLogger(log, router)

	api := router.Group("/api")
	{
		internal := api.Group("/internal")
		{
			h.initInternal(internal)
		}
	}

	return router
}

// initInternal initializes the internal API routes for the handler.
//
// Parameters:
// - router: The Gin router group for the internal API.
func (h *Handler) initInternal(router *gin.RouterGroup) {
	// GET /api/internal/info
	router.GET("/info", h.GetInfo)
}

// initCORS initializes CORS middleware with the provided configuration.
//
// Parameters:
// - cfg: The CORS configuration.
// - router: The Gin router.
func (h *Handler) initCORS(cfg *config.CORS, router *gin.Engine) {
	router.Use(cors.New(cors.Config{
		AllowOrigins:     cfg.AllowedOrigins,
		AllowMethods:     cfg.AllowedMethods,
		AllowHeaders:     cfg.AllowedHeaders,
		AllowCredentials: cfg.AllowedCredentials,
		MaxAge:           cfg.MaxAge,
	}))
}

// initLogger initializes the Gin router with a logger middleware.
//
// The logger middleware logs each request's information, including the method, path,
// remote address, and IP address.
//
// Parameters:
// - log: The logger instance.
// - router: The Gin router.
func (h *Handler) initLogger(log *slog.Logger, router *gin.Engine) {
	router.Use(middleware.LoggerMiddleware(log))
}
