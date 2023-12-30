package task

import (
	"encoding/json"
	"fmt"
	"net/http"
	"todoapp/internal/usecase"
)

type Handler struct {
	useCase UseCases
}

func NewHandler(u UseCases) *Handler {
	return &Handler{useCase: u}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateTaskInputDto
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := h.useCase.Create.Execute(input)
	if err != nil {
		fmt.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		return
	}
}

func (h *Handler) FindAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		return
	}

	output, err := h.useCase.FindAll.Execute()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		return
	}
}
