package db

import (
	"database/sql"
	"log"

	_ "github.com/ClickHouse/clickhouse-go"
)

// Connect открывает соединение с ClickHouse и возвращает объект *sql.DB.
func Connect() (*sql.DB, error) {
	// Параметры подключения к ClickHouse
	connectParams := "tcp://localhost:9000?username=default&password=&database=default"

	// Открываем соединение
	db, err := sql.Open("clickhouse", connectParams)
	if err != nil {
		log.Fatal(err)
	}

	return db, err
}
