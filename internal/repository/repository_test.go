// Ваш файл repository_test.go

package repository

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestInsertEvent(t *testing.T) {
	// Создаем фейковую базу данных
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("не удалось создать фейковую базу данных: %v", err)
	}
	defer func() {
		_ = db.Close()
	}()

	// Ожидаем запрос на вставку в базу данных
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO events").
		WithArgs("test_event", sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Вызываем функцию InsertEvent
	err = InsertEvent("test_event", time.Now(), db)

	// Убеждаемся, что все ожидаемые запросы были выполнены
	assert.NoError(t, mock.ExpectationsWereMet())
}
