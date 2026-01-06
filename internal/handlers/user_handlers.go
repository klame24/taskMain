package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"taskMain/internal/dto/userDTO"
	"taskMain/internal/services"
)

type UserHandler interface {
	GetByID(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	s services.UserService
}

func NewUserHandler(s services.UserService) UserHandler {
	return &userHandler{
		s: s,
	}
}

func (h *userHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)

	user, err := h.s.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	resp := userDTO.GetUserResponse{
		Name:     user.Name,
		Surname:  user.Surname,
		Nickname: user.Nickname,
		Email:    user.Email,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
