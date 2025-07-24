# ビルドステージ
FROM golang:1.24-alpine AS builder
WORKDIR /app

# 依存関係をインストール
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# ソースコード全体をコピーしてビルド
COPY ./cmd/backend/main.go ./cmd/backend/
COPY . . 
WORKDIR /app/cmd/backend
RUN go build -o main .

# 実行ステージ（最小限のイメージ）
FROM golang:1.24-alpine
WORKDIR /app

# ビルドしたバイナリと環境設定をコピー
COPY --from=builder /app/cmd/backend/main .
COPY --from=builder /app/.env .

EXPOSE 8080
CMD ["./main"]