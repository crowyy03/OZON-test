# ozon-test
# Сервис сокращения ссылок на Go

## Запуск проекта
Для начала выборите тип хранения данных (по умолчанию стоит postgres):
Перейдите в файл internal/config/config.yaml:
```bash
storage:
  type: "postgres"  # или "in-memory"
postgres:
  dsn: "postgres://root:1234@localhost:5432/urls?sslmode=disable"
server:
  port: 8080
```
в строчке type: в скобках укажите нужный тип хранения (postgres или in-memory)
```bash
docker-compose up --build
```
Для того, чтобы выключить проект нажмите Control + C или введите в терминале: 
```bash
docker-compose down
```
Для повторного включения:
```bash
docker-compose up
```
Обратите внимание, что все последующие команды будут вводиться в другом терминале
## Кодирование ссылки
В терминале введите команду, перед этим вместо "ссылка_на_сайт" вставьте свою ссылку (в кавычках):
```bash
curl -X POST http://localhost:8080/shorten -H "Content-Type: application/json" -d '{"url": "ссылка_на_сайт"}'
```

## Декодирование ссылки из базы данных
Чтобы раскодировать ссылку, имея сокращенную, введи в терминале:
```bash
curl http://localhost:8080/expand/закодированная_ссылка
```
где слова "закодированная_ссылка" замените на вашу 10-ти символьную короткую ссылку, из прошлого этапа

## Просмотр базы данных PostgreSQL
Для того, чтобы посмтреть все ссылки, которые хранятся в PostgreSQL, во время запущенной программы в терминале введите:
```bash 
docker-compose exec postgres psql -U root -d urls
```
Далее введите:
```bash
SELECT * FROM urls;
```

Ниже вы можете ознакомиться с тем, что у нас получилось:
![alt text](image.png)