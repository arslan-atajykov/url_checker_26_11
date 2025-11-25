package main

import (
	"log"
	"net/http"
	urlhttp "url_checker/internal/http"
)

func main() {
	router := urlhttp.NewRouter()
	log.Println("Сервер запущен на :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
