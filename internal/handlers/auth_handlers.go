package handlers

import (
	"encoding/json"
	"net/http"
	authdto "taskMain/internal/dto/authDTO"
	"taskMain/internal/dto/userDTO"
	"taskMain/internal/services"
)

type AuthHandler interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}

type authHandlers struct {
	s services.AuthService
}

func NewAuthHandler(s services.AuthService) AuthHandler {
	return &authHandlers{
		s: s,
	}
}

func (h *authHandlers) Register(w http.ResponseWriter, r *http.Request) {
	req := userDTO.CreateUserRequest{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	userID, err := h.s.Register(r.Context(), req)
	if err != nil {
		http.Error(w, "can't create user", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userID)
}

func (h *authHandlers) Login(w http.ResponseWriter, r *http.Request) {
	req := authdto.LoginRequest{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	response, err := h.s.Login(r.Context(), req)
	if err != nil {
		http.Error(w, "login failed: "+err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
