package service

import (
	"context"
	"fmt"
	"log/slog"
	"tracker-app/backend/internal/http-server/request"
	"tracker-app/backend/internal/http-server/response"
	"tracker-app/backend/internal/models"
	"tracker-app/backend/internal/repository"

	"github.com/jmoiron/sqlx"
)

type UserTaskService struct {
	log       *slog.Logger
	db        *sqlx.DB
	reposUser repository.UserInterface
	reposTask repository.TaskInterfase
}

func NewUserTaskService(log *slog.Logger, db *sqlx.DB, reposUser repository.UserInterface, reposTask repository.TaskInterfase) *UserTaskService {
	return &UserTaskService{
		log:       log,
		db:        db,
		reposUser: reposUser,
		reposTask: reposTask,
	}
}

// StartTask starts a new task.
// It takes a StartTaskRequest as input and returns a StartTaskResponse and an error.
//
// taskInfo: The StartTaskRequest containing the task information.
// Returns: The StartTaskResponse and an error.
func (s *UserTaskService) StartTask(taskInfo request.StartTaskRequest) (response.StartTaskResponse, error) {
	ctx := context.Background()

	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		s.log.Error("failed to begin transaction: %s", err)
		return response.StartTaskResponse{}, err
	}

	userId := models.UserId{UserId: int64(taskInfo.UserId)}

	if _, err := s.reposUser.GetInfoUser(ctx, tx, userId); err != nil {
		s.log.Error("failed to get user info: %s", err)
		tx.Rollback()
		return response.StartTaskResponse{}, fmt.Errorf("failed to get user info")
	}

	modelTask := models.Task{
		UserId:      userId,
		Description: taskInfo.Description,
	}

	taskId, err := s.reposTask.StartTask(ctx, tx, modelTask)
	if err != nil {
		s.log.Error("failed to start task: %s", err)
		tx.Rollback()
		return response.StartTaskResponse{}, err
	}

	countTask, err := s.reposTask.GetCountTasks(ctx, tx, userId)
	if err != nil {
		s.log.Error("failed to get count tasks: %s", err)
		tx.Rollback()
		return response.StartTaskResponse{}, err
	}

	err = tx.Commit()
	if err != nil {
		s.log.Error("failed to commit transaction: %s", err)
		tx.Rollback()
		return response.StartTaskResponse{}, err
	}

	response := response.StartTaskResponse{
		TaskId:      int(taskId),
		UserId:      int(userId.UserId),
		Description: taskInfo.Description,
		CountTasks:  countTask,
	}

	return response, nil
}

// StopTask stops a task.
// It takes a StopTaskRequest as input and returns a StopTaskResponse and an error.
//
// taskInfo: The StopTaskRequest containing the task information.
// Returns: The StopTaskResponse and an error.
func (s *UserTaskService) StopTask(taskInfo request.StopTaskRequest) (response.StopTaskResponse, error) {
	ctx := context.Background()

	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		s.log.Error("failed to begin transaction: %s", err)
		return response.StopTaskResponse{}, err
	}

	task, err := s.reposTask.GetTaskById(ctx, tx, int64(taskInfo.TaskId))
	if err != nil {
		s.log.Error("failed to get task info: %s", err)
		tx.Rollback()
		return response.StopTaskResponse{}, err
	}

	if !task.Active {
		s.log.Error("task is not active")
		tx.Rollback()
		return response.StopTaskResponse{}, fmt.Errorf("task is not active")
	}

	endTime, err := s.reposTask.StopTask(ctx, tx, int64(taskInfo.TaskId))
	if err != nil {
		s.log.Error("failed to stop task: %s", err)
		tx.Rollback()
		return response.StopTaskResponse{}, err
	}

	countTask, err := s.reposTask.GetCountTasks(ctx, tx, task.UserId)
	if err != nil {
		s.log.Error("failed to get count tasks: %s", err)
		tx.Rollback()
		return response.StopTaskResponse{}, err
	}

	err = tx.Commit()
	if err != nil {
		s.log.Error("failed to commit transaction: %s", err)
		tx.Rollback()
		return response.StopTaskResponse{}, err
	}

	task.TimeStop = endTime
	response := response.StopTaskResponse{
		CountTasks: countTask,
		Task:       task,
	}

	return response, nil
}

func (s *UserTaskService) GetTasksByUser(req request.GetTasksByUserRequest) (response.GetTasksByUserResponse, error) {
	ctx := context.Background()

	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		s.log.Error("failed to begin transaction: %s", err)
		return response.GetTasksByUserResponse{}, err
	}

	statementSet := []string{"start_time > $1", "end_time < $2"}
	args := []interface{}{req.StartTime, req.EndTime}

	task, err := s.reposTask.GetTaskByUser(ctx, tx, req.Sort, statementSet, args)
	if err != nil {
		s.log.Error("failed to get task: %s", err)
		tx.Rollback()
		return response.GetTasksByUserResponse{}, err
	}

	err = tx.Commit()
	if err != nil {
		s.log.Error("error commit tx: %s", err)
		tx.Rollback()
		return response.GetTasksByUserResponse{}, err
	}

	responseTask := response.GetTasksByUserResponse{
		CountTasks: len(task),
		Tasks:      task,
	}

	return responseTask, nil
}
