package routes

import (
	"net/http"
	"taskMain/internal/handlers"
)

func SetupAuthRoutes(h handlers.AuthHandler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /register", h.Register)
	mux.HandleFunc("POST /login", h.Login)

	return mux
}
