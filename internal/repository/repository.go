package repository

import (
	"analytics/internal/db"
	"database/sql"
	"fmt"
	"log"
	"time"
)

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
func InsertEvent(name string, timestamp time.Time) error {
	// Подключение к базе данных
	conn, err := db.Connect()
	if err != nil {
		return fmt.Errorf("ошибка при подключении к базе данных: %v", err)
	}
	defer func(conn *sql.DB) {
		err := conn.Close()
		if err != nil {
			log.Printf("Ошибка при закрытии соединения с базой данных: %v", err)
		}
	}(conn)

	// Начало транзакции
	tx, err := conn.Begin()
	if err != nil {
		return fmt.Errorf("ошибка при начале транзакции: %v", err)
	}
	defer func(tx *sql.Tx) {
		err := tx.Commit()
		if err != nil {
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
