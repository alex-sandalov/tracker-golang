package handler

import (
	"net/http"
	"tracker-app/backend/internal/http-server/request"
	"tracker-app/backend/internal/http-server/response"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetUsers(c *gin.Context) {

}

func (h *Handler) GetTasksByUser(c *gin.Context) {

}

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

func (h *Handler) StopTaskByUser(c *gin.Context) {

}
