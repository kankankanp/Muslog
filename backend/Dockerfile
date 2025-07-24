FROM golang:1.24-alpine AS builder

RUN apk add --no-cache git curl

# swag CLI
RUN go install github.com/swaggo/swag/cmd/swag@latest

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN /go/bin/swag init -g cmd/backend/main.go

WORKDIR /app/cmd/backend
RUN go build -o main .

# 最小限の実行イメージ
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/cmd/backend/main .
COPY --from=builder /app/.env .

EXPOSE 8080
CMD ["./main"]