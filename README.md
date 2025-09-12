# 目次

- [概要](#概要)
  - [主な使用技術](#主な使用技術)
  - [システム構成図](#システム構成図)
  - [ブランチ管理](#ブランチ管理)
- [フロントエンド環境 (web-frontend)](#フロントエンド環境-web-frontend)
  - [前提条件](#前提条件)
  - [パッケージのインストール](#パッケージのインストール)
  - [API クライアントコードの自動生成](#api-クライアントコードの自動生成)
  - [ローカルサーバーの起動](#ローカルサーバーの起動)
  - [ディレクトリ構成](#ディレクトリ構成)
- [バックエンド環境 (backend)](#バックエンド環境-backend)
  - [前提条件](#前提条件-1)
  - [ローカルサーバーの起動](#ローカルサーバーの起動-1)
  - [ディレクトリ構成](#ディレクトリ構成-1)
  - [API ドキュメントの確認](#api-ドキュメントの確認)
- [コード自動整形](#コード自動整形)
- [AWS 本番環境構築 (Terraform)](#aws-本番環境構築-terraform)
  - [前提条件](#前提条件-2)
  - [デプロイ手順](#デプロイ手順)
  - [インフラ構成詳細](#インフラ構成詳細)
    - [リクエストフロー](#リクエストフロー)

# 概要
音楽理論や作曲したオリジナル曲、ギターの演奏動画を投稿したいと思い、音楽に関する技術記事を投稿できるプラットフォームを作りました。マークダウンで記事を投稿でき、タグや SpotifyAPI から検索した曲を選択して紹介する機能も作りました。また、WebSocket 通信を用いてリアルタイムチャットのできるコミュニティ機能も作成しました。インフラ〜フロントエンドまで一度自力で開発してみたいと思い、
AWS 構築を Terraform による IaC、GoによるAPIサーバー、Next.jsによるフロントエンド実装まで行いました。また、AWS ECS への自動デプロイなどの CI / CD も GitHub Actions を使用して実現しています。（運用コストが高く、現在は運用停止中です）

フロントエンドは Figma を使用してデザイン案を構想し、Next.js での UI 実装まで行いました。初めは AWS 環境でホスティングに CloudFront を使用していたため、Next.js の強みである SSR（サーバーサイドレンダリング）を使用できないという制約の中開発を進めていましたが、結局 Lambda@Edge、OpenNext 等を使用し SSR できるようなインフラ環境に変更しました。デザインの再現には Codex や Claude Code などの AI エージェントを多用し、効率的な実装に務めました。

バックエンドは Golang の Echo フレームワークを用いて、クリーンアーキテクチャを中心としたディレクトリ構造での API 開発を行いました。クリーンアーキテクチャを実際にアプリケーションコードに落とし込む技術を学ぶ非常に良い経験となりました。

今後の展望：
現状のインフラ環境では、WebSocket 通信用サーバーと RESTfulAPI 用サーバーが 1 つの ECS コンテナで稼働している状態であり、あまりスケーラブルでない構成です。
今後、記事投稿、コミュニティに加えて、他の機能を追加するとなった場合に備えて、現状のモノリシックアーキテクチャから、gRPC をマイクロサービスへと再編成したいです。

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

## インフラ構成図
![インフラ構成図](document/muslog_architechture.png)

## ブランチ管理

- main
  - 本番用のソースコードを管理するブランチ
- develop
  - 開発作業用のブランチ
  - 普段の開発はこのブランチから feature ブランチを切る
- feature
  - develop ブランチから作成するブランチ
  - 新しい機能の開発など
  - ブランチ名は`feature/{作業内容}`
    - 例: `feature/create-login-page`

# フロントエンド環境 (web-frontend)

## 前提条件

- Node.js 22.6 以上（`frontend/.node-version` 参照）
- npm 10.8 以上

## パッケージのインストール

```bash
# フロントエンド
cd frontend
npm install
```

## ローカルサーバーの起動

```bash
cd frontend
npm run dev
```
http://localhost:3000 にアクセスしてください。

## ディレクトリ構成

```
frontend/
├── openapi.yml                 # OpenAPI 定義（生成元）
├── next.config.mjs
├── src/
│   ├── app/                    # App Router ルート
│   │   ├── layout.tsx
│   │   ├── page.tsx            # トップページ
│   │   └── dashboard/
│   │       ├── layout.tsx
│       ├── page.tsx        # ダッシュボード
│       ├── me/             # プロフィール/自分の投稿
│       ├── community/      # コミュニティ
│       ├── help/           # ヘルプ
│       └── post/
│           ├── add/page.tsx
│           └── [id]/(page.tsx|edit/page.tsx)
│   ├── components/             # UI/レイアウト/モーダル/カード等
│   ├── libs/
│   │   ├── api/generated/      # openapi-types, orval 生成物
│   │   └── websocket/          # WebSocket クライアント
│   ├── contexts/               # React Context
│   ├── constants/
│   └── scss/                   # スタイル
└── package.json
```

### 作成済みページ一覧

- `/` トップ
- `/login-or-signup`
- `/dashboard`
- `/dashboard/me`
- `/dashboard/community`
- `/dashboard/help`
- `/dashboard/post/add`
- `/dashboard/post/[id]`
- `/dashboard/post/[id]/edit`

# バックエンド環境 (backend)

## 前提条件

- Go
- Docker

## ローカルサーバーの起動

プロジェクトのルートディレクトリで以下のコマンドを実行してください。

```bash
docker-compose up --build
# 起動後: http://localhost:8080
```

## ディレクトリ構成

```
backend/
├── cmd/
│   ├── backend/main.go         # Echo 起動、DI、マイグレーション、シード
│   └── scheduler/              # スケジューラ用エントリ
├── config/                     # 設定ロード
├── internal/
│   ├── adapter/                # Web/DTO/Handler 層
│   │   ├── dto/
│   │   └── handler/
│   ├── domain/                 # Entity / Repository IF
│   │   ├── entity/
│   │   └── repository/
│   ├── infrastructure/         # GORM Model/Repository 実装/Mapper
│   │   ├── model/
│   │   ├── repository/
│   │   ├── mapper/
│   │   └── logger/
│   ├── middleware/             # 認証等ミドルウェア
│   ├── seeder/                 # 初期データ投入
│   └── usecase/                # アプリケーションユースケース
├── pkg/utils/                  # 共通ユーティリティ
├── scripts/                    # 補助スクリプト
└── test/                       # テスト
```

## API ドキュメントの確認

- http://localhost:8080/swagger/index.html

### 補足（DB・Seeder・ORM）
- DB: PostgreSQL（Docker Compose）
- ORM: GORM
- マイグレーション: `cmd/backend/main.go` 起動時に `AutoMigrate`
- シード: `internal/seeder/`（GORMのスキーマに合わせて TRUNCATE 実施、重複回避）

### よく使うコマンド
- DBシェル: `docker-compose exec -it db bash && psql -U postgres -d simpleblog`
- テーブル一覧: `\\dt`
- テーブル初期化: `TRUNCATE TABLE <table> CASCADE;`

## テーブルの確認

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
    "source.fixAll.eslint": "always"
  },
  "eslint.validate": ["javascript", "javascriptreact", "typescript", "typescriptreact"],

  // JS・TS用
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
  },
  // Go用
  "[go]": {
    "editor.defaultFormatter": "golang.go"
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

### 追記推奨（今後の記載候補）
- フロント/バックの環境変数一覧（`.env.local`、`.env`）
- 認証フロー（Auth Cookie、Refresh/Access の寿命）
- API クライアント生成の運用ルール（openapi.yml の更新手順）
- Docker サービス構成（ALB/ECS/DB とローカルの差分）
- シーダー仕様（投入ユーザー、投稿件数、再実行時の挙動）

# その他開発・運用コマンド

### Gemini CLI の実行

npx https://github.com/google-gemini/gemini-cli
gemini -m "gemini-2.5-flash" --yolo

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

aws secretsmanager delete-secret --secret-id production/app_secrets --region ap-northeast-1

aws secretsmanager delete-secret --secret-id production/app_secrets --region ap-northeast-1 --force-delete-without-recovery

aws logs describe-log-groups --region ap-northeast-1 --log-
group-name-prefix "/aws/lambda/production-open-next-regional"
--query 'logGroups[].logGroupName' --output text
