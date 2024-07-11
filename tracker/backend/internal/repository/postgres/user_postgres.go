package postgres

import (
	"context"
	"fmt"
	"tracker-app/backend/internal/http-server/request"
	"tracker-app/backend/internal/models"

	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}

// NewUserPostgres creates a new instance of UserPostgres.
// It takes a *sqlx.DB as a parameter and returns a pointer to UserPostgres.
//
// Parameters:
// - db: A *sqlx.DB object representing the connection to the PostgreSQL database.
//
// Returns:
// - *UserPostgres: A pointer to UserPostgres.
func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

// AddUser adds a new user to the PostgreSQL database.
// It takes a context.Context, passport series, and passport number as parameters.
// It returns a UserId and an error.
//
// Parameters:
// - ctx: The context.Context object for the function.
// - passportSeries: The passport series of the user.
// - passportNumber: The passport number of the user.
//
// Returns:
// - models.UserId: The ID of the newly added user.
// - error: An error if the query fails.
func (r *UserPostgres) AddUser(ctx context.Context, passportSeries, passportNumber string) (models.UserId, error) {
	var id int64

	query := `
		INSERT INTO users (passport_serie, passport_number) 
		VALUES ($1, $2) 
		RETURNING user_id
	`

	err := r.db.QueryRowxContext(ctx, query, passportSeries, passportNumber).Scan(&id)

	return models.UserId{
		UserId: id,
	}, err
}

// DeleteUser deletes a user from the PostgreSQL database.
// It takes a context.Context and a user ID as parameters.
// It returns an error if the query fails.
//
// Parameters:
// - ctx: The context.Context object for the function.
// - id: The ID of the user to be deleted.
//
// Returns:
// - error: An error if the query fails.
func (r *UserPostgres) DeleteUser(ctx context.Context, id models.UserId) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE user_id = $1", tableUsers)
	
	_, err := r.db.ExecContext(ctx, query, id.UserId)

	return err
}

func (r *UserPostgres) UpdateUser(ctx context.Context, user request.UpdateUserRequest) error {
	return nil
}
