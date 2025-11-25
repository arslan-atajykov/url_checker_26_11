package main

import (
	"log"
	"net/http"
	urlhttp "url_checker/internal/http"
	"url_checker/internal/jobs"
	"url_checker/internal/repo"
)

func main() {
	repository := repo.NewMemoryRepo()

	jobQueue := jobs.NewJobQueue(repository, 100)
	jobQueue.StartWorker()

	handler := urlhttp.NewHandler(repository)
	router := urlhttp.NewRouter(handler)

	log.Println("Сервер запущен на :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
