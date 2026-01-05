package repositories

import (
	"context"
	"taskMain/internal/models"
	"time"

	"github.com/jackc/pgx/v5"
)

type ProjectRepository interface {
	Create(ctx context.Context, project *models.Project) (int, error)
	Done(ctx context.Context, projectID int) error
	GetByID(ctx context.Context, projectID int) (*models.Project, error)
	Delete(ctx context.Context, projectID int) error
}

type projectRepository struct {
	db *pgx.Conn
}

func NewProjectRepository(db *pgx.Conn) ProjectRepository {
	return &projectRepository{
		db: db,
	}
}

func (r *projectRepository) Create(ctx context.Context, project *models.Project) (int, error) {
	project.CreatedAt = time.Now()
	project.StatusID = 1

	var projectID int

	sqlQuery := `
		INSERT INTO projects (owner_id, title, description, status_id, created_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id;
	`

	err := r.db.QueryRow(ctx, sqlQuery,
		project.OwnerID,
		project.Title,
		project.Description,
		project.StatusID,
		project.CreatedAt).Scan(&projectID)

	return projectID, err
}

func (r *projectRepository) Done(ctx context.Context, projectID int) error {
	sqlQuery := `
		UPDATE projects
		SET status_id=2
		WHERE projects.id=$1
	`

	_, err := r.db.Exec(ctx, sqlQuery, projectID)

	return err
}

func (r *projectRepository) GetByID(ctx context.Context, projectID int) (*models.Project, error) {
	project := models.Project{}

	sqlQuery := `
		SELECT
			owner_id, title, description, status_id, created_at
		FROM projects
		WHERE projects.id=$1
	`

	err := r.db.QueryRow(
		ctx,
		sqlQuery,
		projectID,
	).Scan(
		&project.OwnerID,
		&project.Title,
		&project.Description,
		&project.StatusID,
		&project.CreatedAt,
	)

	return &project, err
}

func (r *projectRepository) Delete(ctx context.Context, projectID int) error {
	sqlQuery := `
		DELETE FROM projects
		WHERE projects.id=$1;
	`

	_, err := r.db.Exec(ctx, sqlQuery, projectID)

	return err
}
