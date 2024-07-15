package handler

import (
	"context"
	"strconv"
	_ "tracker-app/backend/docs"
	"tracker-app/backend/internal/http-server/request"
	"tracker-app/backend/internal/http-server/response"
	"tracker-app/backend/internal/lib"
	"tracker-app/backend/internal/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get users
// @Description Get users
// @Tags users
// @Accept json
// @Produce json
// @Param request query request.GetUsersRequest true "Get users request"
// @Success 200 {object} response.GetUsersResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /users [get]
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

// @Summary Add user
// @Description Add user
// @Tags users
// @Accept json
// @Produce json
// @Param request body request.AddUserRequest true "Add user request"
// @Success 200 {object} models.User
// @Failure 400 {object} response.ErrorResponse
// @Router /users [post]
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

// @Summary Delete user
// @Description Delete user
// @Tags users
// @Accept json
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200
// @Failure 400 {object} response.ErrorResponse
// @Router /users/{user_id} [delete]
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

// @Summary Update user
// @Description Update user
// @Tags users
// @Accept json
// @Produce json
// @Param user_id path int true "User ID"
// @Param request body request.UpdateUserRequest true "Update user request"
// @Success 200 {object} response.UpdateUserResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /users/{user_id} [put]
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

func parseIdParam(c *gin.Context, nameId string) (int64, error) {
	idParam := c.Param(nameId)
	id, err := strconv.ParseInt(idParam, 10, 64)
	return id, err
}
