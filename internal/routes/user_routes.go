package routes

import (
	"net/http"
	"taskMain/internal/handlers"
)

func SetupRoutes(h handlers.UserHandler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /user/{id}", h.GetByID)
	mux.HandleFunc("POST /user", h.Create)

	return mux
}
