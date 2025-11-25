package http

import (
	"encoding/json"
	"net/http"
	"url_checker/internal/model"
	"url_checker/internal/pdf"
)

type ReportRequest struct {
	LinksList []int64 `json:"links_list"`
}

func (h *Handler) ReportHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Принимаем только POST", http.StatusBadRequest)
		return
	}

	var req ReportRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "невалидный json", http.StatusBadRequest)
		return
	}

	if len(req.LinksList) == 0 {
		http.Error(w, "links_list пуст", http.StatusBadRequest)
		return
	}

	var tasks []model.Task
	for _, id := range req.LinksList {
		task, ok := h.repo.GetTask(id)
		if ok {
			tasks = append(tasks, task)
		}
	}

	if len(tasks) == 0 {
		http.Error(w, "задачи не найдены", http.StatusNotFound)
		return
	}

	fileBytes, err := pdf.GeneratePDF(tasks)
	if err != nil {
		http.Error(w, "ошибка PDF генерации", 500)
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Write(fileBytes)
}
