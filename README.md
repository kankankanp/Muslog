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
├── .gitignore
├── docker-compose.yml
├── Dockerfile
├── Dockerfile.local
├── Dockerfile.scheduler
├── go.mod
├── go.sum
├── README.md
├── cmd/
│   ├── backend/
│   │   └── main.go         # アプリケーションのエントリーポイント
│   └── scheduler/
│       └── main.go         # スケジューラーのエントリーポイント
├── config/
│   └── config.go           # アプリケーション設定 (DB接続情報、ポート番号など)
├── internal/
│   ├── domain/
│   │   ├── entities/
│   │   │   └── user.go                         # ドメインエンティティ
│   │   ├── services/
│   │   │   └── user_service.go                 # ドメインサービス
│   │   └── repositories/
│   │       └── user_repository_interface.go    # リポジトリインターフェース
│   ├── usecases/
│   │   ├── input/
│   │   │   └── update_user_name_input.go       # ユースケース入力DTO
│   │   ├── output/
│   │   │   └── update_user_name_output.go      # ユースケース出力DTO
│   │   └── update_user_name_interactor.go      # ユースケース実装
│   ├── infrastructure/
│   │   ├── repositories/
│   │   │   └── user_repository.go              # リポジトリ実装
│   │   ├── models/
│   │   │   └── user_model.go                   # DBモデル
│   │   ├── database/
│   │   │   └── db.go                           # DB接続初期化
│   │   ├── router/
│   │   │   └── router.go                       # Echoルーター設定
│   │   └── server/
│   │       └── server.go                       # Echoサーバー起動
│   └── interfaces/
│       ├── controllers/
│       │   └── user_controller.go              # HTTPハンドラ
│       ├── presenters/
│       │   ├── error_presenter.go              # エラーレスポンス整形
│       │   └── user_presenter.go               # レスポンス整形
│       └── middlewares/
│           └── auth_middleware.go              # 認証ミドルウェア
├── pkg/
│   └── utils/
│       └── parser.go                           # 共通ユーティリティ
├── scripts/
│   └── wait-for-it.sh                          # スクリプト
└── test/
    └── blog_handler_test.go                    # テスト```
```
---

## 各ディレクトリの役割

### 1. `backend/`（プロジェクトルート）
- `.gitignore`：Gitのバージョン管理から除外するファイルを指定します。  
- `docker-compose.yml`：Docker Composeの設定ファイル。開発環境のコンテナ定義（DB、バックエンドなど）を記述します。  
- `Dockerfile`：アプリケーションのDockerイメージをビルドするための定義ファイル。  
- `Dockerfile.local`：ローカル開発用のDockerイメージ定義。  
- `Dockerfile.scheduler`：スケジューラー用のDockerイメージ定義。  
- `go.mod`：Goモジュールファイル。プロジェクトの依存関係（例：`github.com/labstack/echo/v4` など）を管理します。  
- `go.sum`：`go.mod` に記載されたモジュールのチェックサム情報。  
- `README.md`：バックエンドに関する説明ドキュメント。  
- `cmd/`：アプリケーションのエントリーポイントを含む実行可能ファイルを配置します。  
  - `backend/main.go`：メインアプリケーションのエントリーポイント。依存関係を初期化し、Echoサーバーを起動します。  
  - `scheduler/main.go`：スケジューラーのエントリーポイント。  
- `config/`：アプリケーション全体の設定（データベース接続情報、ポート番号、APIキーなど）をロード・管理します。  
- `pkg/`：プロジェクト全体で再利用可能な共通ユーティリティ関数やヘルパー（例：バリデーション、文字列操作、時間処理など）を配置します。  
  - `utils/parser.go`：パーサー関連のユーティリティ関数。  
- `scripts/`：各種スクリプトファイル（例：`wait-for-it.sh`）。  
- `test/`：テストコードを配置します。  
  - `blog_handler_test.go`：ブログハンドラのテスト。  

### 2. `internal/`（アプリケーション内部実装）
このディレクトリ内のコードは、`backend` モジュール内でのみインポートされることを意図しています。外部から内部実装が漏れることを防ぎます。

#### `internal/domain/`（ドメイン層）
- `entities/`：アプリケーションのビジネスエンティティ（例：`User` 構造体）。ビジネスルールと状態をカプセル化します。  
- `services/`：複数のエンティティにまたがるビジネスロジックや、特定のエンティティに属さないロジックを扱うドメインサービス。  
- `repositories/`：ドメイン層が依存するリポジトリインターフェース。データ永続化に関する抽象的な操作を定義します。  

#### `internal/usecases/`（ユースケース層／アプリケーション層）
- `input/`：ユースケースの入力データ構造（DTO）を定義します。  
- `output/`：ユースケースの出力データ構造（DTO）を定義します。  
- `*_interactor.go`：特定のユースケースのビジネスロジックを実装するインターアクター。ドメイン層のサービスやリポジトリを呼び出して処理を実行します。  

#### `internal/infrastructure/`（インフラ層）
- `repositories/`：`internal/domain/repositories` で定義されたインターフェースの実装。DBや外部APIとのやり取りを担当します。  
- `models/`：DBのテーブル構造に対応するモデル。リポジトリがDBマッピングで使用します。  
- `database/`：DB接続の初期化と管理。  
- `router/`：Echoのルーティング設定。`main.go` から呼び出され、コントローラーをHTTPパスとメソッドにマッピングします。  
- `server/`：Echoサーバーの起動ロジックをカプセル化します。  

#### `internal/interfaces/`（プレゼンテーション層／アダプター層）
- `controllers/`：HTTPリクエストを受け取り、ユースケースを呼び出してレスポンスを返すハンドラ。`echo.Context` を使用します。  
- `presenters/`：  
  - `error_presenter.go`：発生したエラーをクライアントに返す形式（JSONなど）に整形します。  
  - `user_presenter.go`：ユースケース出力をHTTPレスポンス形式に変換します。  
- `middlewares/`：Echoのミドルウェア（例：認証、ロギング、CORSなど）を定義します。  

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
gemini -m "gemini-2.5-flash-lite" --yolo
- 会話履歴の保存
/chat save my-session

### Claude code の実行

claude --dangerously-skip-permissions

## AWS リソースの停止・開始

本番環境のコストを抑えるために、ECS と RDS を手動で停止・開始できます。

### 停止

```bash
# 手動でリソースを停止（ECSとRDS両方）
aws ssm start-automation-execution \
    --document-name "production-run-scheduler-task" \
    --parameters "Action=stop"
```

### 開始

```bash
# 手動でリソースを開始（ECSとRDS両方）
aws ssm start-automation-execution \
    --document-name "production-run-scheduler-task" \
    --parameters "Action=start"
```

### 実行状況の確認

```bash
# 実行状況を確認
aws ssm describe-automation-executions \
    --filters "Key=DocumentName,Values=production-run-scheduler-task" \
    --max-results 3

# ECSサービスの状態確認
aws ecs describe-services \
    --cluster "production-cluster" \
    --services "production-backend-service" \
    --query "services[0].{ServiceName:serviceName,DesiredCount:desiredCount,RunningCount:runningCount}"

# RDSクラスターの状態確認
aws rds describe-db-clusters \
    --query "DBClusters[?contains(DBClusterIdentifier, 'production')].{Identifier:DBClusterIdentifier,Status:Status}"
```

**注意事項:**
- 停止: ECS の desired count が 0 になり、RDS が stopping → stopped 状態になります
- 開始: RDS が starting → available 状態になり、ECS の desired count が 2 に戻ります  
- RDS の停止・開始は数分かかることがあります

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
