package main

import (
	"log"
	"net/http"
	urlhttp "url_checker/internal/http"
	"url_checker/internal/repo"
)

func main() {
	repo := repo.NewMemoryRepo()
	handler := urlhttp.NewHandler(repo)
	router := urlhttp.NewRouter(handler)

	log.Println("Сервер запущен на :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
