package taskdto

type CreateTaskRequest struct {
	ProjectID   int    `json:"project_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
