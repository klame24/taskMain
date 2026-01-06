package routes

import (
	"net/http"
	"taskMain/internal/handlers"
)

func SetupUserRoutes(h handlers.UserHandler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{id}", h.GetByID)

	return mux
}
