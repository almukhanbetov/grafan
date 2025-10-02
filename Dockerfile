# Этап сборки
FROM golang:1.24-alpine AS build
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -o grafan .

# Этап запуска
FROM alpine:latest
WORKDIR /app
COPY --from=build /app/grafan .
EXPOSE 8080
CMD ["./grafan"]
