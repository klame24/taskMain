package authdto

import "taskMain/internal/dto/userDTO"

type LoginRequest struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
	User         userDTO.GetUserResponse
}
