package repository

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

// Repository представляет собой репозиторий для взаимодействия с базой данных.
type Repository struct {
	DB *sql.DB
}

// Connect открывает соединение с ClickHouse и возвращает объект *sql.DB.
func Connect() (*sql.DB, error) {
	// Параметры подключения к ClickHouse
	connectParams := "tcp://localhost:9000?username=default&password=&database=default"

	// Открываем соединение
	database, err := sql.Open("clickhouse", connectParams)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Устанавливаем диалект ClickHouse
	_, err = database.Exec("SET sql_dialect = 'default'")
	if err != nil {
		log.Printf("Ошибка при установке диалекта ClickHouse: %v", err)
		return nil, err
	}

	// Проверяем, что соединение действительно установлено
	err = database.Ping()
	if err != nil {
		log.Printf("Ошибка при проверке соединения с ClickHouse: %v", err)
		return nil, err
	}

	return database, nil
}

// InsertEvent вставляет событие в базу данных.
func InsertEvent(name string, timestamp time.Time, db *sql.DB) error {
	// Начало транзакции
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("ошибка при начале транзакции: %v", err)
	}
	defer func(tx *sql.Tx) {
		if err := tx.Commit(); err != nil {
			log.Printf("Ошибка при коммите транзакции: %v", err)
		}
	}(tx)

	// Выполнение запроса на вставку в рамках транзакции
	_, err = tx.Exec("INSERT INTO events (event_name, event_time) VALUES (?, ?)", name, timestamp)
	if err != nil {
		return fmt.Errorf("ошибка при вставке события в базу данных: %v", err)
	}

	return nil
}
