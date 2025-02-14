# используем офиц образ Go
FROM golang:1.22
# Рабочая директория внутри контейнера
WORKDIR /app
# Копируем файлы
COPY go.mod go.sum ./
# бинарник
RUN go mod download
# Копируем код
COPY . .
# Сборка
RUN go build -o url-shortener ./cmd/main.go
# Порт
EXPOSE 8080
# Запуск приложения
CMD [ "./url-shortener" ]


