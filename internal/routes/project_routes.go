package routes

import (
	"net/http"
	"taskMain/internal/handlers"
)

func SetupProjectRoutes(h handlers.ProjectHandlers) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /", h.Create)
	mux.HandleFunc("PUT /{id}", h.Done)
	mux.HandleFunc("GET /{id}", h.GetByID)
	mux.HandleFunc("DELETE /{id}", h.Delete)

	return mux
}
