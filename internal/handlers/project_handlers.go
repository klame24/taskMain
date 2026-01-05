package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	projectdto "taskMain/internal/dto/projectDTO"
	"taskMain/internal/services"
)

type ProjectHandlers interface {
	Create(w http.ResponseWriter, r *http.Request)
	Done(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type projectHandlers struct {
	s services.ProjectService
}

func NewProjectHandlers(s services.ProjectService) ProjectHandlers {
	return &projectHandlers{
		s: s,
	}
}

func (h *projectHandlers) Create(w http.ResponseWriter, r *http.Request) {
	req := projectdto.CreateProjectRequest{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	projectID, err := h.s.Create(r.Context(), req.OwnerID, req.Title, req.Description)
	fmt.Println(err)
	if err != nil {
		http.Error(w, "cant create project", http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projectID)
}

func (h *projectHandlers) Done(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "incorrectly passed value", http.StatusBadRequest)
		return
	}

	err = h.s.Done(r.Context(), id)
	if err != nil {
		http.Error(w, "can't find this project", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("success update!")
}

func (h *projectHandlers) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "incorrectly passed value", http.StatusBadRequest)
		return
	}

	project, err := h.s.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, "can't find this project", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(project)
}

func (h *projectHandlers) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "incorrectly passed value", http.StatusBadRequest)
		return
	}

	err = h.s.Delete(r.Context(), id)
	if err != nil {
		http.Error(w, "cant't find this project", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("success delete!")
}
