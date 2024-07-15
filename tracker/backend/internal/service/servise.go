package service

import (
	"context"
	"log/slog"
	"tracker-app/backend/internal/http-server/request"
	"tracker-app/backend/internal/http-server/response"
	"tracker-app/backend/internal/models"
	"tracker-app/backend/internal/repository"

	"github.com/jmoiron/sqlx"
)

type UserInterface interface {
	GetInfoUser(ctx context.Context, passportSeries, passportNumber string) (models.User, error)
	AddUser(passportSeries, passportNumber string) (models.UserId, error)
	GetUsers(req request.GetUsersRequest) (response.GetUsersResponse, error)
	DeleteUser(id request.DeleteUserRequest) error
	UpdateUser(user request.UpdateUserRequest) (response.UpdateUserResponse, error)
}

type UserTaskInterfase interface {
	StartTask(task request.StartTaskRequest) (response.StartTaskResponse, error)
	StopTask(task request.StopTaskRequest) (response.StopTaskResponse, error)
	GetTasksByUser(req request.GetTasksByUserRequest) (response.GetTasksByUserResponse, error)
}

type Service struct {
	UserInterface
	UserTaskInterfase
}

func NewService(log *slog.Logger, repos *repository.Repository, db *sqlx.DB) *Service {
	return &Service{
		UserInterface:     NewUserService(log, repos.UserInterface, db),
		UserTaskInterfase: NewUserTaskService(log, db, repos.UserInterface, repos.TaskInterfase),
	}
}
