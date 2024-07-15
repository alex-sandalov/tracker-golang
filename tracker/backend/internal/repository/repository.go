package repository

import (
	"context"
	"time"
	"tracker-app/backend/internal/http-server/request"
	"tracker-app/backend/internal/models"
	"tracker-app/backend/internal/repository/postgres"

	"github.com/jmoiron/sqlx"
)

type UserInterface interface {
	AddUser(ctx context.Context, tx *sqlx.Tx, passportSeries, passportNumber string) (models.UserId, error)
	DeleteUser(ctx context.Context, tx *sqlx.Tx, id models.UserId) error
	UpdateUser(ctx context.Context, tx *sqlx.Tx, user request.UpdateUserRequest) error
	GetInfoUser(ctx context.Context, tx *sqlx.Tx, id models.UserId) (models.UserDB, error)
	GetUsers(ctx context.Context, tx *sqlx.Tx, statamentSet []string, args []interface{}, limit, offset int) ([]models.User, error)
	GetCountUsersFilters(ctx context.Context, tx *sqlx.Tx, statamentSet []string, args []interface{}) (int, error)
}

type TaskInterfase interface {
	StartTask(ctx context.Context, tx *sqlx.Tx, task models.Task) (int64, error)
	GetCountTasks(ctx context.Context, tx *sqlx.Tx, userId models.UserId) (int, error)
	StopTask(ctx context.Context, tx *sqlx.Tx, taskId int64) (time.Time, error)
	GetTaskById(ctx context.Context, tx *sqlx.Tx, id int64) (models.Task, error)
	GetTaskByUser(ctx context.Context, tx *sqlx.Tx, typeSort string, statamentSet []string, args []interface{}) ([]models.Task, error)
}

type Repository struct {
	UserInterface
	TaskInterfase
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		UserInterface: postgres.NewUserPostgres(db),
		TaskInterfase: postgres.NewUserTaskRepository(db),
	}
}
