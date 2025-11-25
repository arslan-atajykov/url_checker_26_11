package http

import (
	"encoding/json"
	"net/http"
	"url_checker/internal/jobs"
	"url_checker/internal/repo"
)

type Handler struct {
	repo repo.Repository
	jobs *jobs.JobQueue
}

func NewHandler(r repo.Repository, j *jobs.JobQueue) *Handler {
	return &Handler{
		repo: r,
		jobs: j,
	}
}

type LinksRequest struct {
	Links []string `json:"links"`
}

type LinksResponse struct {
	Links    []string `json:"links"`
	LinksNum int64    `json:"links_num"`
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
		http.Error(w, "отсутствуют links", http.StatusBadRequest)
		return
	}

	task := h.repo.CreateTask(req.Links)

	h.jobs.Submit(task.ID)

	resp := LinksResponse{
		LinksNum: task.ID,
		Links:    req.Links,
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}
