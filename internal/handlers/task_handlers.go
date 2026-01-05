package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	taskdto "taskMain/internal/dto/taskDTO"
	"taskMain/internal/services"
)

type TaskHandlers interface {
	Create(w http.ResponseWriter, r *http.Request)
	Done(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
}

type taskHandlers struct {
	s services.TaskService
}

func NewTaskHandler(s services.TaskService) TaskHandlers {
	return &taskHandlers{
		s: s,
	}
}

func (h *taskHandlers) Create(w http.ResponseWriter, r *http.Request) {
	req := taskdto.CreateTaskRequest{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	taskID, err := h.s.Create(r.Context(), req.ProjectID, req.Title, req.Description)
	if err != nil {
		http.Error(w, "can't create task", http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(taskID)
}

func (h *taskHandlers) Done(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "incorrectly passed value", http.StatusBadRequest)
		return
	}

	err = h.s.Done(r.Context(), id)
	if err != nil {
		http.Error(w, "can't find this task", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("success done task!")
}

func (h *taskHandlers) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "incorrectly passed value", http.StatusBadRequest)
		return
	}

	err = h.s.Delete(r.Context(), id)
	if err != nil {
		http.Error(w, "can't find this task!", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("success delete task!")
}

func (h *taskHandlers) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "incorrectly passed value", http.StatusBadRequest)
		return
	}

	task, err := h.s.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, "can't find this task", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func (h *taskHandlers) GetAll(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "incorrectly passed value", http.StatusBadRequest)
		return
	}

	tasks, err := h.s.GetAll(r.Context(), id)
	if err != nil {
		http.Error(w, "can't find this task", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}