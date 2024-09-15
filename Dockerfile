# Step 1: Build the Go application
FROM golang:1.23-alpine AS build
WORKDIR /app

# Копируем содержимое папки tender
COPY ./tender /app

# Копируем файл .env
COPY ./tender/.env /app/.env

# Загрузка зависимостей Go
RUN go mod download

# Сборка приложения
RUN go build -o tender ./main.go

# Step 2: Create a minimal image to run the binary
FROM alpine:latest
WORKDIR /app
COPY --from=build /app/tender /app/
COPY --from=build /app/.env /app/.env 
EXPOSE 8080
CMD ["./tender"]

#docker-compose down -v
#docker-compose up --build
