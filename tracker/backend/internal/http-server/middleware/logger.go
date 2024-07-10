package middleware

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware(log *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		entry := log.With(
			slog.String("method", c.Request.Method),
			slog.String("path", c.Request.URL.Path),
			slog.String("remote_addr", c.Request.RemoteAddr),
			slog.String("ip", c.ClientIP()),
		)

		entry.Info("request started")

		c.Next()
	}
}
