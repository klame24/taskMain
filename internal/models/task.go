package models

import "time"

type Task struct {
	ID          int       `json:"id"`
	ProjectID   int       `json:"project_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StatusID    int       `json:"status_id"`
	CreatedAt   time.Time `json:"created_at"`
}
