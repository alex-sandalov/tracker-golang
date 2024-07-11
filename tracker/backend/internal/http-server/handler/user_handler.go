package handler

import (
	"context"
	"strconv"
	"tracker-app/backend/internal/http-server/request"
	"tracker-app/backend/internal/http-server/response"
	"tracker-app/backend/internal/lib"
	"tracker-app/backend/internal/models"

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

	userId, err := h.service.AddUser(passportSeries, passportNumber)
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	ctx := context.Background()
	user, err := h.service.GetInfoUser(ctx, userId, passportSeries, passportNumber)
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

// DeleteUser handles the request to delete a user.
//
// This function expects a JSON request body of type DeleteUserRequest,
// which contains the ID of the user to be deleted.
// The function deletes the user with the given ID.
//
// Parameters:
//   - c: The Gin context.
//
// Returns:
//   - None.
func (h *Handler) DeleteUser(c *gin.Context) {
	idParam := c.Param("user_id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.DeleteUser(request.DeleteUserRequest{
		UserId: models.UserId{UserId: id},
	})

	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) UpdateUser(c *gin.Context) {

}
