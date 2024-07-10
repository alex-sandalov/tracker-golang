package handler

import (
	"tracker-app/backend/internal/config"
	"tracker-app/backend/internal/http-server/middleware"
	"tracker-app/backend/internal/service"
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
		external := api.Group("/external")
		{
			users := external.Group("/users")
			h.initUsersAPI(users)
		}
	}

	return router
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

// initUsersAPI initializes the user API endpoints.
//
// Parameters:
// - users: The Gin router group for the user API.
func (h *Handler) initUsersAPI(users *gin.RouterGroup) {
	// GET /api/external/users: Retrieves all users.
	users.GET("/", h.GetUsers)

	// GET /api/external/users/:id/tasks/time: Retrieves labour costs by user ID.
	users.GET("/:id/tasks/time", h.GetTasksByUser)

	// POST /api/external/users/:user_id/tasks/:task_id/start: Starts a task for a user.
	users.POST("/:user_id/tasks/:task_id/start", h.StartTaskByUser)

	// POST /api/external/users/:user_id/tasks/:task_id/stop: Stops a task for a user.
	users.POST("/:user_id/tasks/:task_id/stop", h.StopTaskByUser)

	// POST /api/external/users/: Adds a new user.
	users.POST("/", h.AddUser)

	// DELETE /api/external/users/:user_id: Deletes a user by ID.
	users.DELETE("/:user_id", h.DeleteUser)

	// PATCH /api/external/users/:user_id: Updates a user by ID.
	users.PATCH("/:user_id", h.UpdateUser)
}
