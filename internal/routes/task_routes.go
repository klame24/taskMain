package routes

import (
	"net/http"
	"taskMain/internal/handlers"
)

func SetupTaskRoutes(h handlers.TaskHandlers) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /", h.Create)
	mux.HandleFunc("PUT /{id}", h.Done)
	mux.HandleFunc("DELETE /{id}", h.Delete)
	
	return mux
}