package repositories

import (
	"context"
	"taskMain/internal/models"
	"time"

	"github.com/jackc/pgx/v5"
)

type TaskRepository interface {
	Create(ctx context.Context, task *models.Task) (int, error)
	Done(ctx context.Context, taskID int) error
	Delete(ctx context.Context, taskID int) error
	GetByID(ctx context.Context, taskID int) (*models.Task, error)
	GetAll(ctx context.Context, projectID int) ([]*models.Task, error)
}

type taskRepository struct {
	db *pgx.Conn
}

func NewTaskRepository(db *pgx.Conn) TaskRepository {
	return &taskRepository{
		db: db,
	}
}

func (r *taskRepository) Create(ctx context.Context, task *models.Task) (int, error) {
	var taskID int

	task.CreatedAt = time.Now()
	task.StatusID = 1

	sqlQuery := `
		INSERT INTO tasks (project_id, title, description, status_id, created_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id;
	`
	err := r.db.QueryRow(
		ctx,
		sqlQuery,
		task.ProjectID,
		task.Title,
		task.Description,
		task.StatusID,
		task.CreatedAt).Scan(&taskID)

	return taskID, err
}

func (r *taskRepository) Done(ctx context.Context, taskID int) error {
	sqlQuery := `
		UPDATE tasks
		SET status_id=2
		WHERE tasks.id=$1;
	`

	_, err := r.db.Exec(ctx, sqlQuery, taskID)

	return err
}

func (r *taskRepository) Delete(ctx context.Context, taskID int) error {
	sqlQuery := `
		DELETE FROM tasks
		WHERE tasks.id=$1;
	`

	_, err := r.db.Exec(ctx, sqlQuery, taskID)

	return err
}

func (r *taskRepository) GetByID(ctx context.Context, taskID int) (*models.Task, error) {
	task := models.Task{}
	task.ID = taskID

	sqlQuery := `
		SELECT
			project_id, title, description, status_id, created_at
		FROM tasks
		WHERE tasks.id=$1;
	`

	err := r.db.QueryRow(ctx, sqlQuery, taskID).Scan(
		&task.ProjectID,
		&task.Title,
		&task.Description,
		&task.StatusID,
		&task.CreatedAt,
	)

	return &task, err
}

func (r *taskRepository) GetAll(ctx context.Context, projectID int) ([]*models.Task, error) {
	var tasks []*models.Task

	sqlQuery := `
		SELECT
			id, project_id, title, description, status_id, created_at
		FROM tasks
		WHERE project_id=$1;
	`
	rows, err := r.db.Query(ctx, sqlQuery, projectID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		task := models.Task{}

		err := rows.Scan(
			&task.ID,
			&task.ProjectID,
			&task.Title,
			&task.Description,
			&task.StatusID,
			&task.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, &task)
	}

	return tasks, nil
}
