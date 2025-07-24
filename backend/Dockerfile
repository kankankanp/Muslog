FROM golang:1.24-alpine AS builder

RUN apk add --no-cache git curl bash
RUN go install github.com/swaggo/swag/cmd/swag@latest

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN /go/bin/swag init -g cmd/backend/main.go

WORKDIR /app/cmd/backend
RUN go build -o main .

FROM alpine:latest
RUN apk add --no-cache bash

WORKDIR /app
COPY --from=builder /app/cmd/backend/main .
COPY --from=builder /app/.env .

EXPOSE 8080
CMD ["./main"]

COPY scripts/wait-for-it.sh /wait-for-it.sh
RUN chmod +x /wait-for-it.sh

ENTRYPOINT ["/wait-for-it.sh", "db:5432", "--", "./main"]