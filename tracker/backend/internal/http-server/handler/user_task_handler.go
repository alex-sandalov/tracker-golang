package handler

import (
	"net/http"
	"tracker-app/backend/internal/http-server/request"
	"tracker-app/backend/internal/http-server/response"
	"tracker-app/backend/internal/models"

	"github.com/gin-gonic/gin"
)

// GetTasksByUser handles the request to get tasks for a specific user.
// It expects a query parameter for the user ID.
// The function returns a list of tasks for the user and an error if there is any issue.
//
// Parameters:
//   - c: The Gin context.
//
// Returns:
//   - None.
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

// StartTaskByUser handles the request to start a new task for a user.
// It expects a JSON request body of type StartTaskRequest.
// The function returns a StartTaskResponse and an error if the task fails to start.
//
// Parameters:
//   - c: The Gin context.
//
// Returns:
//   - None.
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

// StopTaskByUser handles the request to stop a task for a user.
// It expects a JSON request body of type StopTaskRequest.
// The function returns a StopTaskResponse and an error if the task fails to stop.
//
// Parameters:
//   - c: The Gin context.
//
// Returns:
//   - None.
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
