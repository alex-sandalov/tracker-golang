package postgres

import (
	"context"
	"fmt"
	"tracker-app/backend/internal/models"

	"github.com/jmoiron/sqlx"
)

type UserTaskRepository struct {
	db *sqlx.DB
}

// NewUserTaskRepository creates a new UserTaskRepository instance.
//
// db: A sqlx.DB object representing the database connection.
// Returns: A UserTaskRepository object.
func NewUserTaskRepository(db *sqlx.DB) *UserTaskRepository {
	return &UserTaskRepository{db: db}
}

// StartTask starts a new task in the PostgreSQL database.
// It takes a context.Context, a transaction object, and a Task object as parameters.
// It returns the ID of the newly created task and an error if the query fails.
//
// ctx: The context.Context object for the function.
// tx: A transaction object for the database operations.
// task: A Task object representing the task to be started.
// Returns: The ID of the newly created task and an error.
func (u *UserTaskRepository) StartTask(ctx context.Context, tx *sqlx.Tx, task models.Task) (int64, error) {
	var idTask int64

	query := fmt.Sprintf(
		"INSERT INTO %s (user_id, description, start_time) VALUES ($1, $2, NOW()) RETURNING task_id",
		tableUserTasks,
	)

	err := tx.QueryRowxContext(ctx, query, task.UserId.UserId, task.Description).Scan(&idTask)

	return idTask, err
}

// GetCountTasks retrieves the count of tasks associated with the given user.
//
// ctx: The context.Context object for the function.
// tx: A transaction object for the database operations.
// userId: The ID of the user to get the count of tasks for.
// Returns: The count of tasks associated with the user and an error.
func (u *UserTaskRepository) GetCountTasks(ctx context.Context, tx *sqlx.Tx, userId models.UserId) (int, error) {
	var count int
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE user_id = $1", tableUserTasks)
	err := tx.QueryRowxContext(ctx, query, userId.UserId).Scan(&count)
	return count, err
}
