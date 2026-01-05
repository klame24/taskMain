package models

type User struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Surname      string `json:"surname"`
	Nickname     string `json:"nickname"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}
