# URL Shortener API

## Описание проекта
Этот сервис ползволяет сокращать длинные URL-адреса, а также получать их оригинальные версии

## Как генерируются ссылки?
1. Генерируется случайный уникальный индентификатор каждой ссылки
2. Идентификатор присваивается к оригинальной URL, а после сохраняется в хранилище 

## Структура проекта
```
OZON/
│── cmd/                  # Главная точка входа в приложение
│── internal/
│   ├── api/              # Реализация API (обработчики запросов)
│   ├── shortener/        # Логика сокращения ссылок
│   ├── storage/          # Взаимодействие с базой данных
│── tests/                # Интеграционные и unit-тесты
│── migrations/           # Миграции базы данных
│── config/               # Конфигурационные файлы
│── README.md             # Документация
```

## API EndPoints

1. **POST** - `/shorten`
- Тело запроса: `"url": "http://example.com"`
- Ответ: `"short_url": "abcde12345"`

*Пример запроса:* 
```bash
curl -X POST http://localhost:8080/shorten -H "Content-Type: application/json" -d '{"url": "http://apple.com"}'
```

*Получаем ответ:*
```json
{"short_url":"PndoKe8ytf"}
```

2. **GET** `/expand/{shortURL}`

- Ответ:  `"original_url": "https://example.com"`

*Пример:* 
```bash 
curl http://localhost:8080/expand/PndoKe8ytf
```

*Получаем ответ:* ```{"original_url":"http://apple.com"}```

3. Проверка работоспособности

**GET** `/health`
- Ответ: `OK`
- *Пример:* 
```bash
curl -X GET http://localhost:8080/health
```

## Обработка ошибок
 - 400 Bad Request – некорректный формат запроса
 - 404 Not Found – ссылка не найдена
 - 500 Internal Server Error – внутренняя ошибка сервера
 
## Масштабируемость и устойчивость
- Поддержка работы с несколькими пользователями одновременно
- Оптимизация хранилища для быстрой обработки запросов
- Возможность работы на долгий срок без деградации

## Миграции базы данных
Миграции используются для обновления структуры базы данных
### Как запустить миграции?
1. Установите golang-migrate, если он не установлен:
  ```bash
  go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
  ```
2. Запустите миграции:
  ```bash
  migrate -path ./migrations -database "postgres://user:password@localhost:5432/dbname?sslmode=disable" up
  ```
3. Если нужно откатить последнюю миграцию:
  ```bash
  migrate -path ./migrations -database "postgres://user:password@localhost:5432/dbname?sslmode=disable" down 1
  ```
### Альтернативный способ выполнения миграций через psql
Если golang-migrate недоступен, можно выполнить SQL-скрипты вручную:
```
psql -U root -d urls -f migrations/1_create_table.sql
psql -U root -d urls -f migrations/2_add_index.sql
psql -U root -d urls -f migrations/3_add_created.sql
psql -U root -d urls -f migrations/4_delete_table.sql
```
## Тестирование
В проекте также реализованы модульные тесты. Для их запуска введите в терминале:  
```bash 
go test ./...
```

## Запуск проекта
Для начала выберите тип хранилища данных (по умолчанию используется PostgreSQL):
Перейдите в файл `internal/config/config.yaml`:
```yaml
storage:
  type: "postgres"  # или "in-memory"
postgres:
  dsn: "postgres://root:1234@localhost:5432/urls?sslmode=disable"
server:
  port: 8080
```
В строке `type:` укажите необходимый тип хранилища (postgres или in-memory).

Запустите проект:
```bash
docker-compose up --build
```
Для завершения работы проекта нажмите Control + C или введите в терминале: 
```bash
docker-compose down
```
Для повторного запуска:
```bash
docker-compose up
```
Обратите внимание, что все последующие команды будут вводиться в другом терминале.

## Просмотр базы данных PostgreSQL
Для просмотра всех ссылок, хранящихся в PostgreSQL, во время работы программы в терминале введите:
```bash
docker-compose exec postgres psql -U root -d urls
```
Затем введите:
```sql
SELECT * FROM urls;
```

Ниже представлено, что было достигнуто:
![alt text](image.png)

