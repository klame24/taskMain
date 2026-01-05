package models

import "time"

type Project struct {
	ID          int       `json:"id"`
	OwnerID     int       `json:"owner_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StatusID    int       `json:"status_id"`
	CreatedAt   time.Time `json:"created_at"`
}
