# Используем официальный образ Golang
FROM golang:1.23-alpine

# Устанавливаем необходимые зависимости
RUN apk add --no-cache git

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы go.mod и go.sum
COPY go.mod go.sum ./

# Скачиваем зависимости
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем приложение
RUN go build -o main ./cmd/basic_server/main.go

# Открываем порт
EXPOSE 8090

# Запускаем приложение
CMD ["./main"] 