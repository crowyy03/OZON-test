# ozon-test
# Сервис сокращения ссылок на Go

## Запуск проекта
```bash
docker-compose up --build
## Кодирование ссылки
curl -X POST http://localhost:8080/shorten -H "Content-Type: application/json" -d '{"url": "https://good.com"}'

