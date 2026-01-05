package projectdto

type CreateProjectRequest struct {
	OwnerID     int    `json:"owner_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
