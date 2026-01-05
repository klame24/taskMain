package services

import (
	"context"
	"taskMain/internal/models"
	"taskMain/internal/repositories"
)

type TaskService interface {
	Create(ctx context.Context, projectID int, title, description string) (int, error)
	Done(ctx context.Context, taskID int) error
	Delete(ctx context.Context, taskID int) error
}

type taskService struct {
	r repositories.TaskRepository
}

func NewTaskService(r repositories.TaskRepository) TaskService {
	return &taskService{
		r: r,
	}
}

func (s *taskService) Create(ctx context.Context, projectID int, title, description string) (int, error) {
	task := models.Task{
		ProjectID:   projectID,
		Title:       title,
		Description: description,
	}
	taskID, err := s.r.Create(ctx, &task)

	return taskID, err
}

func (s *taskService) Done(ctx context.Context, taskID int) error {
	err := s.r.Done(ctx, taskID)

	return err
}

func (s *taskService) Delete(ctx context.Context, taskID int) error {
	err := s.r.Delete(ctx, taskID)

	return err
}
