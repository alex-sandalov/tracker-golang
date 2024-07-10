package request

import (
	"effective-mobile-golang/backend/internal/models"
	"time"
)

type GetUsersRequest struct {
	models.User
	Page    int `json:"page" form:"page" binding:"min=1"`
	PerPage int `json:"perPage" form:"per_page" binding:"min=1"`
}

type GetTasksByUserRequest struct {
	UserId    int       `json:"userId" form:"user_id" binding:"required"`
	StartTime time.Time `json:"startTime" form:"start_time" binding:"required"`
	EndTime   time.Time `json:"endTime" form:"end_time" binding:"required"`
	Sort      string    `json:"sort" form:"sort" default:"desc"`
}
