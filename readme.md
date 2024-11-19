# Проект: Сервер для Neo4j БД

## Разворачивание на VDS (без контейнеров)

1. **Установить зависимости**:

   - Установите [Go](https://go.dev/) (не ниже версии 1.22).
   - Установите [Neo4j](https://neo4j.com/download/) и запустите его.

2. **Настроить переменные окружения**:  
   Создайте файл `.env` в корне проекта на основе примера:

   ```env
   NEO4J_URI=bolt://localhost:7687
   NEO4J_USER=neo4j
   NEO4J_PASSWORD=lV8OPmIWJoBS89
   AUTH_TOKEN=YOUR_AUTH_TOKEN
   HTTP_PORT=:8199
   ```

3. **Запустить сервер**:
   ```bash
   go run ./cmd/server/main.go
   ```

Сервер будет доступен по адресу `http://localhost:8199`.

---

## Запуск тестов

```bash
go test ./...
```

---

## Структура проекта

- **cmd/server** — точка входа в приложение.
- **internal/config** — работа с конфигурацией.
- **internal/delivery/http** — HTTP-обработчики и middleware.
- **internal/models** — модели данных для работы с Neo4j.
- **internal/repository** — слой для взаимодействия с базой данных.
- **internal/service** — бизнес-логика и тесты.

---

Готово! 🚀
