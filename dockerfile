# Dockerfile
FROM golang:1.24.5-alpine
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY src/ ./src/
COPY .env ./.env
RUN go build -o main ./src
EXPOSE 8080
CMD ["./main"]