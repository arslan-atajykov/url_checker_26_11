package http

import (
	"encoding/json"
	"net/http"
	"sync"
)

var (
	Counter int64
	mu      sync.Mutex
)

type UrlsRequest struct {
	Urls []string `json:"urls"`
}

type UrlsResponse struct {
	UrlId int64    `json:"urls_id"`
	Urls  []string `json:"urls"`
}

func UrlHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Принимаем только POST", http.StatusBadRequest)
		return
	}

	var req UrlsRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "невалидный json", http.StatusBadRequest)
		return
	}

	if len(req.Urls) == 0 {
		http.Error(w, "отсутствуют urls", http.StatusBadRequest)
		return
	}

	mu.Lock()
	Counter++
	id := Counter
	mu.Unlock()
	resp := UrlsResponse{
		UrlId: id,
		Urls:  req.Urls,
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)

}
