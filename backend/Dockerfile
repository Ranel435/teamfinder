# Используем официальный образ Golang
FROM golang:1.23-alpine

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd/basic_server/main.go

EXPOSE 8090

CMD ["./main"] 