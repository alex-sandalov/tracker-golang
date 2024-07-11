package repository

import (
	"context"
	"info-golang/backend/internal/http-server/request"
	"info-golang/backend/internal/models"
	"info-golang/backend/internal/repository/postgres"

	"github.com/jmoiron/sqlx"
)

type UserInterface interface {
	GetUser(ctx context.Context, user request.GetInfoRequest) (models.InfoUser, error)
}

type Repository struct {
	UserInterface
}

// NewRepository creates a new instance of the Repository struct
// with the given database connection.
//
// Parameters:
// - db: The sqlx.DB object representing the database connection.
// Returns:
// - A pointer to the Repository struct.
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		// UserInterface is an implementation of the UserInterface interface.
		// It uses the NewUserRepository function to create a new instance
		// of the UserRepository struct with the given database connection.
		UserInterface: postgres.NewUserPostgres(db),
	}
}
