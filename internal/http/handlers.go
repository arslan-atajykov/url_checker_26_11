package http

import (
	"encoding/json"
	"net/http"
	"url_checker/internal/checker"
	"url_checker/internal/model"
	"url_checker/internal/repo"
)

type Handler struct {
	repo repo.Repository
}

func NewHandler(r repo.Repository) *Handler {
	return &Handler{
		repo: r,
	}
}

type LinksRequest struct {
	Links []string `json:"links"`
}

type LinksResponse struct {
	Links    map[string]string `json:"links"`
	LinksNum int64             `json:"links_num"`
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

	result := make(map[string]string)
	linkStructs := make([]model.LinkStruct, len(req.Links))

	for i, url := range req.Links {
		status := checker.CheckURL(url)

		result[url] = status

		linkStructs[i] = model.LinkStruct{
			URL:     url,
			Lstatus: model.LStatus(status),
		}
	}

	task := h.repo.CreateTaskWithLinks(linkStructs)

	resp := LinksResponse{
		Links:    result,
		LinksNum: task.ID,
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}
