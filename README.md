# 概要

Muslog は、総合的音楽プラットフォーム

## 主な使用技術

- フロントエンド
  - 言語：TypeScript
  - フレームワーク：Next.js(App Router)
  - スタイル：SCSS, Tailwind CSS
  - コード整形・静的解析：Prettier・ESLint
- バックエンド
  - 言語：Go
  - フレームワーク：Echo
  - コンテナ：Docker
  - ORM：GORM
- データベース：PostgreSQL (Docker Compose 経由)
- API 仕様：Swagger (Swag)

## システム構成図

```
                 ┌──────────────────────────────┐
                 │          CloudFront          │
                 │ (CDN / キャッシュ / HTTPS)     │ 
                 └─────┬─────────────┬──────────┘
                       │             │
         HTML/CSS/JS   │             │  /api/*
                       │             │
     ┌─────────────────▼───┐   ┌────▼─────────────────┐
     │ S3（フロント配信用） │   │  ALB（API用ロードバランサ）│
     └─────────────────────┘   └──────────┬───────────┘
                                           │
                                 ┌─────────▼─────────┐
                                 │ ECS (Fargate)     │
                                 │ Go Echo APIサーバ  │
                                 └─────────┬─────────┘
                                           │
                            ┌──────────────▼──────────────┐
                            │   RDS PostgreSQL (DB)       │
                            └────────────────────────────┘

【画像アップロード/配信】
ユーザー → (PUT: 署名付きURL) → S3（画像・音源保存）
CloudFront → (GET) → S3（画像・音源）

【コンテナ配布】
GitHub Actions → (docker push) → ECR → ECSでpullして起動

【監視/ログ】
ALB → CloudWatch（アクセスログ・メトリクス）
ECS → CloudWatch（アプリログ・メトリクス）
```

## ブランチ管理

- main
  - 本番用のソースコードを管理するブランチ
- develop
  - 開発作業用のブランチ
  - 普段の開発はこのブランチから feature ブランチを切る
- feature
  - develop ブランチから作成するブランチ
  - 新しい機能の開発など
  - ブランチ名は`feature/{issue番号}-{作業内容}`
    - 例: `feature/1-create-login-page`

# フロントエンド環境 (web-frontend)

## 前提条件

- Node.js (v18 以上推奨)
- npm

## パッケージのインストール

```bash
# web-frontendディレクトリに移動
cd web-frontend

# パッケージのインストール
npm install
```

## API クライアントコードの自動生成

バックエンドサーバーを起動した状態で、以下のコマンドを実行すると、`src/app/api/`配下に API クライアントコードが生成されます。

```bash
# web-frontendディレクトリにいることを確認
# openapi.jsonの更新
curl http://localhost:8080/swagger/doc.json -o openapitools/openapi.json

# OpenAPI GeneratorによるTypeScriptクライアントの生成
npm run openapi-gen
```

## ローカルサーバーの起動

`npm run dev`を実行後、 http://localhost:3000 にアクセスしてください。

## ディレクトリ構成

```
frontend/
├── src/
│   ├── app/                # ページルーティング定義
│   │   ├── api/            # 自動生成したAPIクライアントコード
│   │   ├── components/     # UIコンポーネント群
│   │   ├── hooks/          # React用のカスタムフック
│   │   └── libs/           # 汎用ロジックやユーティリティ関数
│   └── scss/               # スタイル定義
├── openapitools/
│   └── openapi.json        # バックエンドから取得したOpenAPIスキーマファイル
└── package.json            # 依存パッケージ・スクリプト定義
```

# バックエンド環境 (backend)

## 前提条件

- Go
- Docker
- Docker Compose

## ローカルサーバーの起動

プロジェクトのルートディレクトリで以下のコマンドを実行してください。
`docker-compose up --build` を実行後、 http://localhost:8080 にアクセスするとバックエンドサーバーが起動します。

## ディレクトリ構成

```
backend/
├── cmd/
│   └── backend/
│       └── main.go         # アプリのエントリーポイント
├── config/                 # 設定ファイル関連
├── internal/
│   ├── handler/            # HTTPリクエストの処理
│   ├── model/              # データ構造の定義
│   ├── repository/         # DBとのやり取り
│   ├── seeder/             # 初期データ投入
│   └── service/            # ビジネスロジック
├── go.mod                  # 依存パッケージ管理
├── Dockerfile              # Dockerイメージ定義
└── ...
```

## API ドキュメントの確認

ローカルサーバーを起動後、以下の URL にアクセスすると Swagger UI で API ドキュメントを確認できます。

- http://localhost:8080/swagger/index.html

# コード自動整形

VSCode の拡張機能を検索し、

- Prettier - Code Formatter
- ESLint
- Go
  をインストール後、VSCode の`settings.json`に以下を記述してください。

```json
{
  // 共通設定
  "editor.formatOnSave": true,
  "editor.codeActionsOnSave": {
    "source.organizeImports": "always",
    "source.fixAll": "always"
  },

  // Go用
  "[go]": {
    "editor.defaultFormatter": "golang.go"
  },

  // TypeScript/JavaScript用
  "[typescript]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  },
  "[typescriptreact]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  },
  "[javascript]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  },
  "[javascriptreact]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  }
}
```

設定後、ファイル保存時にコード整形とインポート整理・不要インポート削除が実行されます。

# AWS 本番環境構築 (Terraform)

## 前提条件

- AWS CLI 設定済み
- Terraform (v1.0 以上推奨)
- 適切な IAM ロール・権限

## デプロイ手順

```bash
# 本番環境のTerraformディレクトリに移動
cd terraform/environments/production

# Terraform初期化
terraform init

# 実行計画確認
terraform plan

# インフラ構築
terraform apply
```

## インフラ構成詳細

- **Route 53**: ドメイン名の DNS 管理 (example.com)
- **CloudFront**: CDN によるキャッシュ・HTTPS 配信
- **ALB**: API 用ロードバランサー (/api/\* へのリクエストルーティング)
- **S3**: 静的ファイル配信 (HTML/CSS/JS) + 画像・音源保存
- **ECS Fargate**: Go アプリケーションのコンテナ実行
- **ECR**: Docker イメージレジストリ
- **RDS PostgreSQL**: データベース
- **CloudWatch**: ログ・メトリクス監視

### リクエストフロー

1. **静的ファイル**: ユーザー → CloudFront → S3
2. **API リクエスト**: ユーザー → CloudFront → ALB → ECS Fargate
3. **画像アップロード**: ユーザー → 署名付き URL → S3 → CloudFront 配信
4. **コンテナデプロイ**: GitHub Actions → docker build → ECR → ECS

# 開発・運用コマンド

### Gemini CLI の実行

npx https://github.com/google-gemini/gemini-cli
gemini -m "gemini-2.5-flash" --yolo

### Claude code の実行

claude --dangerously-skip-permissions

### DB のマイグレーション・シーディング

```
# backendコンテナに直接入る
docker-compose exec -it backend bash
# マイグレーションファイルの作成

# マイグレーションの実行

# シードデータの追加
```

### テーブルの確認

```
# dbコンテナに直接入る
docker-compose exec -it db bash
# PostgreSQLに移動
psql -U postgres -d simpleblog
# 特定のテーブルの中身を確認
SELECT * FROM <テーブル名>;
# 特定のテーブルの中身を削除
TRUNCATE TABLE <テーブル名> CASCADE;
# テーブル一覧を確認
\dt:
```
