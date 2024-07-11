package request

import (
	"time"
	"tracker-app/backend/internal/models"
)

type GetUsersRequest struct {
	models.User
	Page    int `json:"page" form:"page" binding:"min=1"`
	PerPage int `json:"perPage" form:"perPage" binding:"min=1"`
}

type GetTasksByUserRequest struct {
	UserId    int       `json:"userId" form:"userId" binding:"required"`
	StartTime time.Time `json:"startTime" form:"startTime" binding:"required"`
	EndTime   time.Time `json:"endTime" form:"endTime" binding:"required"`
	Sort      string    `json:"sort" form:"sort" default:"desc"`
}
