package handler

import (
	"analytics/internal/repository"
	"fmt"
	"log"
	"net/http"
	"time"
)

// TrackHandler обрабатывает запросы на отслеживание событий.
func TrackHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	// Получаем имя события из параметра запроса
	eventName := r.FormValue("event")
	if eventName == "" {
		http.Error(w, "Параметр 'event' отсутствует", http.StatusBadRequest)
		return
	}

	// Вставляем событие в базу данных
	err := repository.InsertEvent(eventName, time.Now())
	if err != nil {
		log.Printf("Ошибка при записи в базу данных: %v", err)
		http.Error(w, "Ошибка при записи в базу данных", http.StatusInternalServerError)
		return
	}

	_, err = fmt.Fprint(w, "Событие успешно зарегистрировано")
	if err != nil {
		return
	}
}
