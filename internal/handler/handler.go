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

	// Подключаемся к базе данных
	db, err := repository.Connect()
	if err != nil {
		log.Printf("Ошибка при подключении к базе данных: %v", err)
		http.Error(w, "Ошибка при подключении к базе данных", http.StatusInternalServerError)
		return
	}
	defer func() {
		err := db.Close()
		if err != nil {
			log.Printf("Ошибка при закрытии соединения с базой данных: %v", err)
		}
	}()

	// Вставляем событие в базу данных
	err = repository.InsertEvent(eventName, time.Now(), db)
	if err != nil {
		log.Printf("Ошибка при записи в базу данных: %v", err)
		http.Error(w, "Ошибка при записи в базу данных", http.StatusInternalServerError)
		return
	}

	_, err = fmt.Fprint(w, "Событие успешно зарегистрировано")
	if err != nil {
		log.Printf("Ошибка при отправке ответа: %v", err)
	}
}
