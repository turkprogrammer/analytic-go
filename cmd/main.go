package main

import (
	"analytics/internal/router"
	"log"
	"net/http"
)

func main() {
	r := router.NewRouter()

	log.Println("Сервер запущен на :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
