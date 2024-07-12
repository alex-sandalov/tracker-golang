package response

type StartTaskResponse struct {
	TaskId      int    `json:"taskId"`
	UserId      int    `json:"userId"`
	Description string `json:"description"`
	CountTasks  int    `json:"countTasks"`
}
