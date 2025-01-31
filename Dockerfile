# Используем официальный образ Go
FROM golang:1.22-alpine

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы проекта
COPY . .

# Скачиваем зависимости
RUN go mod download

# Собираем приложение
RUN go build -o main .

# Открываем порт для приложения
EXPOSE 8080

# Команда для запуска приложения
CMD ["./main"]