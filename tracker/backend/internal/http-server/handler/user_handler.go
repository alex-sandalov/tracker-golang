package handler

import (
	"effective-mobile-golang/backend/internal/http-server/request"
	"effective-mobile-golang/backend/internal/http-server/response"
	"effective-mobile-golang/backend/internal/lib"

	"net/http"

	"github.com/gin-gonic/gin"
)

// AddUser handles the request to add a new user.
// It expects a JSON request body of type AddUserRequest.
// The function parses the passport number from the request and returns the parsed values.
//
// Parameters:
//   - c: The Gin context.
//
// Returns:
//   - None.
func (h *Handler) AddUser(c *gin.Context) {
	var req request.AddUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	passportSeries, passportNumber, err := lib.ParsePassport(req.PassportNumber)
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"passportSeries": passportSeries,
		"passportNumber": passportNumber,
	})
}

func (h *Handler) DeleteUser(c *gin.Context) {

}

func (h *Handler) UpdateUser(c *gin.Context) {

}
