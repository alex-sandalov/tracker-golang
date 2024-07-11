package response

import (
	"log"

	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Message string `json:"message"`
}

// NewErrorResponse creates a new error response and aborts the request with the
// specified status code and error message.
//
// Parameters:
// - c: The Gin context.
// - statusCode: The HTTP status code.
// - message: The error message.
func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	log.Default().Println(message)

	c.AbortWithStatusJSON(statusCode, errorResponse{
		Message: message,
	})
}
