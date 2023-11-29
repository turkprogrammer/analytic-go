package router

import (
	"analytics/internal/handler"
	"github.com/gorilla/mux"
	"net/http"
)

// NewRouter создает новый роутер и определяет маршруты.
func NewRouter() *mux.Router {
	r := mux.NewRouter()

	// Обработчик для корневого пути
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Добро пожаловать в сервис аналитики", http.StatusOK)
	}).Methods("GET")

	// Обработчик для отправки событий аналитики
	r.HandleFunc("/track", handler.TrackHandler).Methods("POST")

	return r
}
