package postgres

import (
	"context"
	"fmt"
	"strings"
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
func (r *UserPostgres) AddUser(ctx context.Context, tx *sqlx.Tx, passportSeries, passportNumber string) (models.UserId, error) {
	var id int64

	query := `
		INSERT INTO users (passport_serie, passport_number) 
		VALUES ($1, $2) 
		RETURNING user_id
	`

	err := tx.QueryRowContext(ctx, query, passportSeries, passportNumber).Scan(&id)

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
func (r *UserPostgres) DeleteUser(ctx context.Context, tx *sqlx.Tx, id models.UserId) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE user_id = $1", tableUsers)

	_, err := tx.ExecContext(ctx, query, id.UserId)
	return err
}

// UpdateUser updates a user in the PostgreSQL database.
// It takes a context.Context and an UpdateUserRequest as parameters.
// It returns an error if the query fails.
//
// Parameters:
// - ctx: The context.Context object for the function.
// - user: A request.UpdateUserRequest object containing the user information to be updated.
//
// Returns:
// - error: An error if the query fails.
func (r *UserPostgres) UpdateUser(ctx context.Context, tx *sqlx.Tx, user request.UpdateUserRequest) error {
	var setStatements []string
	var args []interface{}

	idx := 1
	for field, value := range user.Update {
		setStatements = append(setStatements, fmt.Sprintf("%s = $%d", field, idx))
		args = append(args, value)
		idx++
	}

	query := fmt.Sprintf(
		"UPDATE %s SET %s WHERE user_id = $%d",
		tableUsers,
		strings.Join(setStatements, ", "),
		idx,
	)

	args = append(args, user.UserId.UserId)
	_, err := tx.ExecContext(ctx, query, args...)

	return err
}

// GetInfoUser retrieves user information from the PostgreSQL database based on the user ID.
// It takes a context.Context, a transaction object, and a user ID as parameters.
// It returns a UserDB object and an error.
//
// Parameters:
// - ctx: The context.Context object for the function.
// - tx: A transaction object for the database operations.
// - id: The ID of the user to retrieve information for.
//
// Returns:
// - models.UserDB: The user information retrieved from the database.
// - error: An error if the query fails.
func (r *UserPostgres) GetInfoUser(ctx context.Context, tx *sqlx.Tx, id models.UserId) (models.UserDB, error) {
	var user models.UserDB

	query := "SELECT user_id, passport_number, passport_serie FROM users WHERE user_id = $1"

	err := tx.GetContext(ctx, &user, query, id.UserId)

	return user, err
}
