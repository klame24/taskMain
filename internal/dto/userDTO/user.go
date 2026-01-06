package userDTO

type CreateUserRequest struct {
	Name         string `json:"name"`
	Surname      string `json:"surname"`
	Nickname     string `json:"nickname"`
	Email        string `json:"email"`
	Password string `json:"password"`
}

type GetUserResponse struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
}