package userDTO

type CreateUserRequest struct {
	Name         string `json:"name"`
	Surname      string `json:"surname"`
	Nickname     string `json:"nickname"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

type GetByIDResponse struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
}
