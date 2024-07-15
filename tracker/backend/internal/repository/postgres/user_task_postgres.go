package postgres

import (
	"context"
	"fmt"
	"strings"
	"time"
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
		"INSERT INTO %s (user_id, description, active, start_time, end_time) VALUES ($1, $2, $3, NOW(), NOW()) RETURNING task_id",
		tableUserTasks,
	)

	err := tx.QueryRowxContext(ctx, query, task.UserId.UserId, task.Description, true).Scan(&idTask)

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

// StopTask stops a task in the PostgreSQL database by updating the 'active' field to false.
//
// Parameters:
//   - ctx: The context.Context object for the function.
//   - tx: A transaction object for the database operations.
//   - taskId: The ID of the task to be stopped.
//
// Returns:
//   - The end time of the stopped task and an error if the query fails.
func (u *UserTaskRepository) StopTask(ctx context.Context, tx *sqlx.Tx, taskId int64) (time.Time, error) {
	query := fmt.Sprintf("UPDATE %s SET active = $1, end_time = NOW() WHERE task_id = $2 RETURNING end_time", tableUserTasks)

	endTime := time.Now()
	err := tx.QueryRowxContext(ctx, query, false, taskId).Scan(&endTime)

	return endTime, err
}

// GetTaskById retrieves a task from the PostgreSQL database based on the task ID.
//
// ctx: The context.Context object for the function.
// tx: A transaction object for the database operations.
// id: The ID of the task to retrieve.
// Returns: The Task object and an error if the query fails.
func (u *UserTaskRepository) GetTaskById(ctx context.Context, tx *sqlx.Tx, id int64) (models.Task, error) {
	var taskInfo models.Task
	query := fmt.Sprintf("SELECT * FROM %s WHERE task_id = $1", tableUserTasks)
	err := tx.GetContext(ctx, &taskInfo, query, id)
	return taskInfo, err
}

// GetTaskByUser retrieves tasks based on the specified conditions for a user.
// It takes a context.Context, a transaction object, the type of sorting, a slice of statements, and a slice of arguments as parameters.
// It returns a slice of Task objects and an error.
//
// ctx: The context.Context object for the function.
// tx: A transaction object for the database operations.
// typeSort: The type of sorting to apply to the tasks.
// statementSet: A slice of strings representing conditions to filter the tasks.
// args: A slice of interface{} containing the arguments for the query.
// Returns: A slice of Task objects and an error.
func (u *UserTaskRepository) GetTaskByUser(ctx context.Context, tx *sqlx.Tx, typeSort string, statementSet []string, args []interface{}) ([]models.Task, error) {
	var tasks []models.Task
	query := fmt.Sprintf(`
		SELECT task_id, user_id, description, active, start_time, end_time,
			end_time - start_time AS duration
		FROM %s
		WHERE %s
		ORDER BY duration %s`, tableUserTasks, strings.Join(statementSet, " AND "), typeSort)
	err := tx.SelectContext(ctx, &tasks, query, args...)
	return tasks, err
}
