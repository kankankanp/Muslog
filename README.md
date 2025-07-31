# 概要
Simple Blogは、Next.jsとGoを使用して構築されたブログプラットフォームです。

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
- データベース：PostgreSQL (Docker Compose経由)
- API仕様：Swagger (Swag)

## システム構成図
```mermaid
graph TD
    subgraph "Browser"
        A[Next.js Frontend]
    end

    subgraph "Server"
        B[Go Backend API]
        C[MySQL Database]
    end

    A -- HTTP Request --> B
    B -- CRUD --> C
```

## インフラ構成図 (AWS with Terraform)

このプロジェクトのインフラはAWS上にTerraformを用いて構築されています。主要なサービスは以下の通りです。

- **VPC**: ネットワークの論理的な分離
- **Public Subnet**: インターネットからのアクセスを許可するリソース（EC2インスタンス）
- **Private Subnet**: 内部からのアクセスのみを許可するリソース（RDSデータベース）
- **EC2**: Goバックエンドアプリケーションのホスティング
- **RDS (PostgreSQL)**: データベースサービス
- **S3 (Frontend)**: Next.jsフロントエンドの静的ファイルホスティング
- **S3 (Uploads)**: ユーザーがアップロードするファイルの保存
- **CloudFront**: フロントエンドコンテンツのCDN

```mermaid
graph TD
    subgraph User
        User[User]
    end

    subgraph AWS Cloud
        subgraph Global Services
            CloudFront[CloudFront]
            S3Frontend[S3 Bucket (Frontend)]
            S3Uploads[S3 Bucket (Uploads)]
        end

        subgraph VPC
            subgraph Public Subnet
                EC2[EC2 Instance (Go Backend)]
            end

            subgraph Private Subnet
                RDS[RDS (PostgreSQL)]
            end

            EC2 -- DB Connection --> RDS
        end

        User -- HTTPS --> CloudFront
        CloudFront -- Get Objects --> S3Frontend
        CloudFront -- API Requests --> EC2
        EC2 -- Read/Write --> S3Uploads
    end

    classDef default fill:#fff,stroke:#333,stroke-width:2px;
    classDef user fill:#f9f,stroke:#333,stroke-width:2px;
    classDef cdn fill:#f9f,stroke:#333,stroke-width:2px;
    classDef s3 fill:#ccf,stroke:#333,stroke-width:2px;
    classDef ec2 fill:#bbf,stroke:#333,stroke-width:2px;
    classDef rds fill:#bfb,stroke:#333,stroke-width:2px;
    classDef vpc fill:#eee,stroke:#333,stroke-width:1px;
    classDef subnet fill:#ddd,stroke:#333,stroke-width:1px;

    class User user;
    class CloudFront cdn;
    class S3Frontend,S3Uploads s3;
    class EC2 ec2;
    class RDS rds;
```

## ブランチ管理
- main
  - 本番用のソースコードを管理するブランチ
- develop
  - 開発作業用のブランチ
  - 普段の開発はこのブランチからfeatureブランチを切る
- feature
  - developブランチから作成するブランチ
  - 新しい機能の開発など
  - ブランチ名は`feature/{issue番号}-{作業内容}`
     - 例: `feature/1-create-login-page`

# フロントエンド環境 (web-frontend)
## 前提条件
- Node.js (v18以上推奨)
- npm

## パッケージのインストール
```bash
# web-frontendディレクトリに移動
cd web-frontend

# パッケージのインストール
npm install
```

## APIクライアントコードの自動生成
バックエンドサーバーを起動した状態で、以下のコマンドを実行すると、`src/app/api/`配下にAPIクライアントコードが生成されます。

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
web-frontend/
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

## APIドキュメントの確認
ローカルサーバーを起動後、以下のURLにアクセスするとSwagger UIでAPIドキュメントを確認できます。
- http://localhost:8080/swagger/index.html

# コード自動整形
VSCodeの拡張機能を検索し、
- Prettier - Code Formatter
- ESLint
- Go
をインストール後、VSCodeの`settings.json`に以下を記述してください。

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
  }
}
```
設定後、ファイル保存時にコード整形とインポート整理・不要インポート削除が実行されます。




### Gemini CLIの実行
npx https://github.com/google-gemini/gemini-cli
gemini -m "gemini-2.5-flash"



### DBのマイグレーション・シーディング
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

### ゲストログイン
NbekDfg@QNqqGvl.info