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

func (h *Handler) GetUsers(c *gin.Context) {
	var req request.GetUsersRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.service.GetUsers(req)
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

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

	ctx := context.Background()
	user, err := h.service.GetInfoUser(ctx, passportSeries, passportNumber)
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := h.service.AddUser(passportSeries, passportNumber)
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user.UserId = userId

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
	id, err := parseIdParam(c, "user_id")
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

// UpdateUser handles the request to update a user.
//
// This function expects a JSON request body of type UpdateUserRequest,
// which contains the updated user information.
// The function updates the user with the given ID.
//
// Parameters:
//   - c: The Gin context.
//
// Returns:
//   - None.
func (h *Handler) UpdateUser(c *gin.Context) {
	id, err := parseIdParam(c, "user_id")
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var req request.UpdateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	req.UserId = models.UserId{UserId: id}

	infoUser, err := h.service.UpdateUser(req)
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, infoUser)
}

// parseIdParam parses an integer id from the request parameter.
//
// Parameters:
//   - c: The Gin context.
//   - nameId: The name of the id parameter in the URL.
//
// Returns:
//   - id: The parsed integer id.
//   - err: An error if the parsing fails.
func parseIdParam(c *gin.Context, nameId string) (int64, error) {
	idParam := c.Param(nameId)
	id, err := strconv.ParseInt(idParam, 10, 64)
	return id, err
}
