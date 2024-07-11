package postgres

import (
	"context"
	"fmt"
	"info-golang/backend/internal/http-server/request"
	"info-golang/backend/internal/models"

	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}

// NewUserPostgres creates a new instance of the UserPostgres struct.
//
// db: The sqlx.DB object representing the PostgreSQL database connection.
// Returns: A pointer to the UserPostgres struct.
func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{
		db: db,
	}
}

// GetUser retrieves user information from the database based on the provided passport
// series and number.
//
// ctx: The context.Context object for the request.
// user: The request.GetInfoRequest object containing the passport series and number.
// Returns: The models.InfoUser object representing the retrieved user information
// and an error if the retrieval fails.
func (u *UserPostgres) GetUser(ctx context.Context, user request.GetInfoRequest) (models.InfoUser, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE passport_serie = $1 AND passport_number = $2", tableUsers)

	var infoUser models.InfoUser

	err := u.db.GetContext(ctx, &infoUser, query, user.PassportSerie, user.PassportNumber)

	return infoUser, err
}
