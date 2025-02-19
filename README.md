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

1. **POST** - ```/shorten```
Тело запроса: ``` "url": "http://example.com"```
Ответ: ```"short_url": "abcde12345"```

*Пример запроса:* 
``` curl -X POST http://localhost:8080/shorten -H "Content-Type: application/json" -d '{"url": "http://apple.com"}' ```
*Получаем ответ:*
```{"short_url":"PndoKe8ytf"}```

2. Расшифровка URL

**GET** /expand/{shortURL}
Ответ:  ```"original_url": "https://example.com"```

*Пример*: ```curl http://localhost:8080/expand/PndoKe8ytf```
*Получаем ответ*: ```{"original_url":"http://apple.com"}```
3. Проверка работоспособности
**GET** /health
Ответ: OK
Пример: ```curl -X GET http://localhost:8080/health```

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
```go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest```
2. Запустите миграции:
```migrate -path ./migrations -database "postgres://user:password@localhost:5432/dbname?sslmode=disable" up```
3. Если нужно откатить последнюю миграцию:
```migrate -path ./migrations -database "postgres://user:password@localhost:5432/dbname?sslmode=disable" down 1```
### Альтернативный способ выполнения миграций через psql
Если golang-migrate недоступен, можно выполнить SQL-скрипты вручную:
```
psql -U root -d urls -f migrations/1_create_table.sql
psql -U root -d urls -f migrations/2_add_index.sql
psql -U root -d urls -f migrations/3_add_created.sql
psql -U root -d urls -f migrations/4_delete_table.sql
```
## Тестирование
В проекте добавлены Unit тесты. Для их запуска введите в терминале: ```go test ./...```

## Запуск проекта
Для начала выборите тип хранения данных (по умолчанию стоит postgres):
Перейдите в файл internal/config/config.yaml:

```
storage:
  type: "postgres"  # или "in-memory"
postgres:
  dsn: "postgres://root:1234@localhost:5432/urls?sslmode=disable"
server:
  port: 8080
```

в строчке type: в скобках укажите нужный тип хранения (postgres или in-memory)
```
docker-compose up --build
```
Для того, чтобы выключить проект нажмите Control + C или введите в терминале: 
```
docker-compose down
```
Для повторного включения:
```
docker-compose up
```
Обратите внимание, что все последующие команды будут вводиться в другом терминале

## Кодирование ссылки
В терминале введите команду, перед этим вместо "ссылка_на_сайт" вставьте свою ссылку (в кавычках):
```
curl -X POST http://localhost:8080/shorten -H "Content-Type: application/json" -d '{"url": "ссылка_на_сайт"}'
```

## Декодирование ссылки из базы данных
Чтобы раскодировать ссылку, имея сокращенную, введи в терминале:
```
curl http://localhost:8080/expand/закодированная_ссылка
```
где слова "закодированная_ссылка" замените на вашу 10-ти символьную короткую ссылку, из прошлого этапа

## Просмотр базы данных PostgreSQL
Для того, чтобы посмтреть все ссылки, которые хранятся в PostgreSQL, во время запущенной программы в терминале введите:
```
docker-compose exec postgres psql -U root -d urls
```
Далее введите:
```
SELECT * FROM urls;
```

Ниже вы можете ознакомиться с тем, что у нас получилось:
![alt text](image.png)

