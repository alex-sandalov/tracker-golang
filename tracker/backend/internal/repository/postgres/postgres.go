package postgres

import (
	"fmt"
	"tracker-app/backend/internal/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	tableUsers = "users"
)

// NewPostrgesDb creates a new PostgreSQL database connection.
//
// cfg: A DataBase struct containing the database connection details.
// Returns: A sqlx.DB object and an error if there was any issue with creating the connection.
func NewPostrgesDb(cfg config.Database) (*sqlx.DB, error) {
	// Construct the connection string.
	connectStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Name, cfg.Password, cfg.SSLMode)

	// Open the database connection.
	db, err := sqlx.Open("postgres", connectStr)
	if err != nil {
		return nil, err
	}

	// Ping the database to check if the connection is successful.
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// Return the database connection.
	return db, nil
}
