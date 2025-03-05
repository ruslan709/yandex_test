package orchestrator

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Calculate(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Expression string `json:"expression"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusUnprocessableEntity)
		return
	}

	id, err := h.service.AddExpression(req.Expression)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"id": id})
}

func (h *Handler) GetExpressions(w http.ResponseWriter, r *http.Request) {
	expressions := h.service.GetExpressions()
	json.NewEncoder(w).Encode(map[string]interface{}{"expressions": expressions})
}

func (h *Handler) GetExpressionByID(w http.ResponseWriter, r *http.Request, idStr string) {
	id, err := strconv.Atoi(idStr) // Преобразуем строку в число
	if err != nil {
		http.Error(w, "Invalid expression ID", http.StatusBadRequest)
		return
	}

	expression, err := h.service.GetExpressionByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{"expression": expression})
}

func (h *Handler) GetTask(w http.ResponseWriter, r *http.Request) {
	task, err := h.service.GetTask()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{"task": task})
}

func (h *Handler) PostTaskResult(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ID     int     `json:"id"`
		Result float64 `json:"result"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusUnprocessableEntity)
		return
	}

	if err := h.service.PostTaskResult(req.ID, req.Result); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}
