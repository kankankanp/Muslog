# Backend (Echo + GORM)

## ディレクトリ構成

```
backend/
├── cmd/backend/           # エントリポイント(main.go)
├── internal/
│   ├── handler/           # Echoハンドラ
│   ├── service/           # ビジネスロジック
│   ├── repository/        # DBアクセス
│   └── model/             # ドメインモデル
├── config/                # 設定管理
├── docs/                  # ドキュメント
├── test/                  # テスト
├── pkg/                   # 外部公開用パッケージ
├── go.mod
├── go.sum
└── README.md
```

## 起動方法

1. 必要な環境変数を設定（.env ファイル or docker-compose）
2. `go mod tidy` で依存解決
3. `go run cmd/backend/main.go` で起動

## 主な API

- `/api/blog` ... ブログ記事の CRUD
- `/api/user` ... ユーザー情報取得

## 開発 Tips

- 各層ごとに責務を分離
- テストは`test/`配下に追加
- 設定は`config/`で一元管理
