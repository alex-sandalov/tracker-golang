package handler

import (
	"info-golang/backend/internal/http-server/request"
	"info-golang/backend/internal/http-server/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetInfo handles the GET /info endpoint. It retrieves user information based on the provided request parameters.
//
// Parameters:
// - c: The Gin context.
func (h *Handler) GetInfo(c *gin.Context) {
	var req request.GetInfoRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.service.UserInterface.GetUserInfo(req)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response.NewGetUserInfoResponse(user))
}
