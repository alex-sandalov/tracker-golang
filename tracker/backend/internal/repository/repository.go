package repository

import (
	"context"
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
}

type Repository struct {
	UserInterface
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		UserInterface: postgres.NewUserPostgres(db),
	}
}
