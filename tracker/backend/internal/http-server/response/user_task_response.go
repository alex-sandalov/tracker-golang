package response

import "tracker-app/backend/internal/models"

type StartTaskResponse struct {
	TaskId      int    `json:"taskId"`
	UserId      int    `json:"userId"`
	Description string `json:"description"`
	CountTasks  int    `json:"countTasks"`
}

type StopTaskResponse struct {
	CountTasks int `json:"countTasks"`
	models.Task
}

type GetTasksByUserResponse struct {
	CountTasks int           `json:"countTasks"`
	Tasks      []models.Task `json:"tasks"`
}
