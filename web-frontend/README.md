# 開発者向け情報

## 前提条件

- Node.js(v22.6 以上)
- npm(v10.8 以上)
- Git(v2.39 以上)

## 依存パッケージのインストール

```npm install```

## ローカルサーバーの起動

```npm run dev```
サーバーを起動後、http://localhost:3000 でアクセスしてください。

# 利用者向け情報
![日記](/public/readme/intro1.png)
![ブログ検索画面](/public/readme/intro2.png)
この Web アプリケーションは、「1 日を象徴する音楽を添えた日記」を投稿できるブログサイトです。

## 使用技術

### フロントエンド

- 言語：TypeScript
- フレームワーク： Next.js 14.2.15(App Router)
- スタイル： TailwindCSS, Sass
- 状態管理： Redux Toolkit

### バックエンド

- 言語：TypeScript
- フレームワーク：Next.js 14.2.15(App Router)
- データベース：PostgreSQL(Supabase Database)

- 認証：Auth.js

### 開発環境・インフラ

- IDE：Visual Studio Code
- ホスティング：Vercel
- バージョン管理：Git、GitHub

### システム構成図

![システム構成図](/public/readme/system-configuration.png)
