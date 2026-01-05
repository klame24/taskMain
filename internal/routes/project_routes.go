package routes

import (
	"net/http"
	"taskMain/internal/handlers"
)

func SetupProjectRoutes(h handlers.ProjectHandlers, h1 handlers.TaskHandlers) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /", h.Create)
	mux.HandleFunc("PUT /{id}", h.Done)
	mux.HandleFunc("GET /{id}", h.GetByID)
	mux.HandleFunc("DELETE /{id}", h.Delete)
	mux.HandleFunc("GET /{id}/tasks", h1.GetAll)

	return mux
}
