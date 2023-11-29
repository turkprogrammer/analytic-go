# Analytics Service

Этот проект представляет собой минимальный сервис аналитики на языке программирования Go с использованием базы данных ClickHouse.

## Зависимости

github.com/ClickHouse/clickhouse-go: Драйвер ClickHouse для Go.
github.com/gorilla/mux: Мощный маршрутизатор HTTP для Go.

## Запустите проект:

go mod tidy
go run cmd/main.go

Ваш сервис будет доступен по адресу http://localhost:8080.

## Как проверить работоспособность
    curl -X POST http://localhost:8080/track -d "event=page_view"
Или выполните аналогичный POST запрос в Postman.

    http://localhost:8080/track?event=page_view

Проверьте, что получаете ответ "Событие успешно зарегистрировано".

Проверьте базу данных ClickHouse, чтобы убедиться, что событие было записано. Используйте средства администрирования ClickHouse или запросы SQL.

## Структура проекта

```plaintext
analytics/
|-- cmd/
|   `-- main.go
|-- internal/
|   |-- db/
|   |   `-- db.go
|   |-- handler/
|   |   `-- handler.go
|   |-- repository/
|   |   `-- repository.go
|   `-- router/
|       `-- router.go
|-- go.mod
|-- go.sum
`-- README.md

cmd: Этот каталог содержит главный файл main.go, который является точкой входа в приложение.
internal: Здесь расположены внутренние пакеты вашего приложения.
db: Код для взаимодействия с базой данных (например, подключение к ClickHouse).
handler: Обработчики HTTP-запросов.
repository: Код для работы с базой данных (например, вставка событий).
router: Конфигурация маршрутизатора HTTP.
go.mod и go.sum: Файлы для управления зависимостями проекта.

