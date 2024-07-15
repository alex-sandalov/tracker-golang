package handler

import (
	"net/http"
	"tracker-app/backend/internal/http-server/request"
	"tracker-app/backend/internal/http-server/response"
	"tracker-app/backend/internal/models"

	"github.com/gin-gonic/gin"
)

// @Summary Get tasks by user
// @Description Get tasks by user
// @Tags tasks
// @Accept json
// @Produce json
// @Param user_id path int true "User ID"
// @Param request query request.GetTasksByUserRequest true "Get tasks by user request"
// @Success 200 {object} response.GetTasksByUserResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /users/{user_id}/tasks [get]
func (h *Handler) GetTasksByUser(c *gin.Context) {
	idUser, err := parseIdParam(c, "user_id")
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var req request.GetTasksByUserRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	req.UserId = models.UserId{
		UserId: idUser,
	}
	res, err := h.service.GetTasksByUser(req)
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

// @Summary Start task by user
// @Description Start task by user
// @Tags tasks
// @Accept json
// @Produce json
// @Param request body request.StartTaskRequest true "Start task request"
// @Success 200 {object} response.StartTaskResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /users/tasks [post]
func (h *Handler) StartTaskByUser(c *gin.Context) {
	var req request.StartTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.service.StartTask(req)
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}

// @Summary Stop task by user
// @Description Stop task by user
// @Tags tasks
// @Accept json
// @Produce json
// @Param request body request.StopTaskRequest true "Stop task request"
// @Success 200 {object} response.StopTaskResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /users/tasks [put]
func (h *Handler) StopTaskByUser(c *gin.Context) {
	var req request.StopTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.service.StopTask(req)
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}
