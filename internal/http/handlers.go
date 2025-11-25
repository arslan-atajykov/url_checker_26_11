package http

import (
	"encoding/json"
	"net/http"
	"url_checker/internal/repo"
)

type Handler struct {
	repo repo.Repository
}

func NewHandler(r repo.Repository) *Handler {
	return &Handler{repo: r}
}

type LinksRequest struct {
	Links []string `json:"Links"`
}

type LinksResponse struct {
	LinkId int64    `json:"Links_id"`
	Links  []string `json:"Links"`
}

func (h *Handler) LinkHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Принимаем только POST", http.StatusBadRequest)
		return
	}

	var req LinksRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "невалидный json", http.StatusBadRequest)
		return
	}

	if len(req.Links) == 0 {
		http.Error(w, "отсутствуют Links", http.StatusBadRequest)
		return
	}

	task := h.repo.CreateTask(req.Links)
	resp := LinksResponse{
		LinkId: task.ID,
		Links:  req.Links,
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)

}
