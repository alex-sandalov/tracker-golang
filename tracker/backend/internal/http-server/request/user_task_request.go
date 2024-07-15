package request

import (
	"time"
	"tracker-app/backend/internal/models"
)

type GetTasksByUserRequest struct {
	UserId    models.UserId
	StartTime time.Time `json:"startTime" form:"startTime" binding:"required"`
	EndTime   time.Time `json:"endTime" form:"endTime" binding:"required"`
	Sort      string    `json:"sort" form:"sort" default:"desc"`
}

type StartTaskRequest struct {
	UserId      int    `json:"userId" form:"userId" binding:"required"`
	Description string `json:"description" form:"description"`
}

type StopTaskRequest struct {
	TaskId int `json:"taskId" form:"taskId" binding:"required"`
	UserId int `json:"userId" form:"userId" binding:"required"`
}
