package services

import (
	"context"
	"taskMain/internal/models"
	"taskMain/internal/repositories"
)

type ProjectService interface {
	Create(ctx context.Context, ownerID int, title, description string) (int, error)
	Done(ctx context.Context, projectID int) error
	GetByID(ctx context.Context, projectID int) (*models.Project, error)
	Delete(ctx context.Context, projectID int) error
}

type projectService struct {
	r repositories.ProjectRepository
}

func NewProjectService(r repositories.ProjectRepository) ProjectService {
	return &projectService{
		r: r,
	}
}

func (s *projectService) Create(
	ctx context.Context, ownerID int,
	title, description string,
) (int, error) {
	project := models.Project{
		OwnerID:     ownerID,
		Title:       title,
		Description: description,
	}

	projectID, err := s.r.Create(ctx, &project)

	return projectID, err
}

func (s *projectService) Done(ctx context.Context, projectID int) error {
	err := s.r.Done(ctx, projectID)

	return err
}

func (s *projectService) GetByID(ctx context.Context, projectID int) (*models.Project, error) {
	project, err := s.r.GetByID(ctx, projectID)

	return project, err
}

func (s *projectService) Delete(ctx context.Context, projectID int) error {
	err := s.r.Delete(ctx, projectID)

	return err
}
