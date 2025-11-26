package main

import (
	"log"
	"net/http"
	urlhttp "url_checker/internal/http"
	"url_checker/internal/repo"
)

func main() {
	repository, err := repo.NewFileRepo("data/tasks.json")
	if err != nil {
		log.Fatal(err)
	}

	handler := urlhttp.NewHandler(repository)
	router := urlhttp.NewRouter(handler)

	log.Println("Сервер запущен на :8080")

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
