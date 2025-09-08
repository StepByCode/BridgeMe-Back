# BridgeMe - Backend

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white) ![Docker](https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white) ![Terraform](https://img.shields.io/badge/Terraform-7B42BC?style=for-the-badge&logo=terraform&logoColor=white) ![GitHub Actions](https://img.shields.io/badge/GitHub%20Actions-2088FF?style=for-the-badge&logo=githubactions&logoColor=white) ![MongoDB](https://img.shields.io/badge/MongoDB-47A248?style=for-the-badge&logo=mongodb&logoColor=white) ![GitHub commit activity](https://img.shields.io/github/commit-activity/m/dokkiichan/BridgeMe-Back?style=for-the-badge)

## プロジェクト概要
このプロジェクトは、NFCキーホルダーをスマートフォンにかざすだけで、プロフィールをウェブブラウザに表示するシンプルなサービスです。バックエンドはGo言語で実装されており、クリーンアーキテクチャの原則に従っています。

## 開発環境のセットアップ

### 前提条件
- Docker
- Docker Compose

### 環境構築手順

1.  **(初回セットアップまたはクリーンな状態から始める場合)** 以前にコンテナを起動したことがある場合は、既存のデータボリュームを削除してクリーンな状態から始めることを推奨します。
    ```bash
    docker compose down -v
    ```

2.  **リポジトリのクローン**
    ```bash
    git clone https://github.com/StepByCode/BridgeMe.git
    cd BridgeMe-Back
    ```

2.  **.env ファイルの作成**
    プロジェクトのルートディレクトリに `.env` ファイルを作成し、以下の内容を記述してください。
    ```env
    MONGO_DB_ROOT_USERNAME=root
    MONGO_DB_ROOT_PASSWORD=password
    ME_CONFIG_BASICAUTH_USERNAME=admin
    ME_CONFIG_BASICAUTH_PASSWORD=changeme
    ```
    *   `MONGO_DB_ROOT_USERNAME`, `MONGO_DB_ROOT_PASSWORD`: MongoDBの認証情報です。
    *   `ME_CONFIG_BASICAUTH_USERNAME`, `ME_CONFIG_BASICAUTH_PASSWORD`: Mongo Expressの管理画面にログインするための認証情報です。`changeme` は任意のパスワードに変更してください。

3.  **Docker Compose の起動**
    以下のコマンドを実行すると、MongoDBとGoバックエンド、Mongo Expressが起動します。
    ```bash
    docker compose up --build
    ```
    初回起動時やコード変更時には `--build` オプションが必要です。

## トラブルシューティング: `backend` または `mongo-express` がMongoDBに接続できない場合

`backend`サービスや`mongo-express`がMongoDBへの認証エラーで起動しない場合、MongoDBの初期ユーザー情報とアプリケーションが使用する認証情報が一致していない可能性があります。これは、MongoDBのデータボリュームに古いユーザー情報が残っている場合に発生します。

**解決策:**

1.  **MongoDBのデータボリュームを削除し、コンテナを再作成します。**
    これにより、`.env`ファイルで指定したユーザー名とパスワードでMongoDBが初期化されます。

    ```bash
    docker compose down -v
    ```

2.  **`.env`ファイルの内容が、使用したいユーザー名とパスワードで正しく設定されていることを確認してください。**
    例:
    ```env
    MONGO_DB_ROOT_USERNAME=your_desired_username
    MONGO_DB_ROOT_PASSWORD=your_desired_password
    ME_CONFIG_BASICAUTH_USERNAME=your_desired_username
    ME_CONFIG_BASICAUTH_PASSWORD=your_desired_password
    ```
    **注意:** `ME_CONFIG_BASICAUTH_USERNAME` と `ME_CONFIG_BASICAUTH_PASSWORD` は、Mongo Expressの管理画面ログイン用です。MongoDBの認証情報とは別に設定することも可能ですが、混乱を避けるため同じ値にすることをお勧めします。

3.  **再度、コンテナを起動します。**

    ```bash
    docker compose up --build
    ```

## アプリケーションの実行

`docker compose up` コマンドで全てのサービスが起動します。

-   **Goバックエンド:** `http://localhost:8080`
-   **MongoDB:** `localhost:2701` (Dockerコンテナ内部からアクセス)
-   **Mongo Express:** `http://localhost:8081`

## デプロイ (Deployment)

本番環境へのデプロイは、GitHub Actionsを利用して自動化されています。デプロイフローは以下の通りです。

1.  **プルリクエストの作成**: 変更を`main`ブランチに対するプルリクエストとして作成します。
2.  **CIの実行**: プルリクエストが作成されると、自動的にCI（テスト）が実行されます。
3.  **マージ**: CIが成功し、レビューで承認されると、プルリクエストを`main`ブランチにマージできます。（ブランチ保護ルールにより、CIの失敗したコードはマージできません。）
4.  **デプロイの開始**: `main`ブランチへのマージをトリガーに、デプロイワークフローが自動的に開始されます。
5.  **手動承認**: `Terraform Apply`のステップで、GitHub ActionsのUI上で**手動承認**が求められます。
6.  **デプロイ実行**: 承認されると、Terraformが本番サーバーに変更を適用します。

## APIドキュメント (OpenAPI / Swagger UI)

Goバックエンドが起動している状態で、以下のURLにアクセスするとAPIドキュメント（Swagger UI）を確認できます。

-   **Swagger UI:** `http://localhost:8082`

## Mongo Express (MongoDB管理画面)

MongoDBが起動している状態で、以下のURLにアクセスするとMongo Expressの管理画面にログインできます。

-   **Mongo Express:** `http://localhost:8081`
    *   **ユーザー名:** `admin`
    *   **パスワード:** `.env` ファイルで設定した `ME_CONFIG_BASICAUTH_PASSWORD` の値（デフォルトは `changeme`）

## APIエンドポイント

| メソッド | パス           | 説明           |
| :------- | :------------- | :------------- |
| `POST`   | `/profiles`    | プロフィールの作成 |
| `GET`    | `/profiles/{id}` | 特定のプロフィールを取得 |
| `GET`    | `/profiles`    | 全てのプロフィールを取得 |

### バリデーションルール
- `POST /profiles`
  - `name`: 必須項目 (required)
  - `affiliation`: 必須項目 (required)
  - `bio`: 必須項目 (required)

## バックエンド構成

```mermaid
graph TD
    subgraph System Overview
        subgraph backend_container [Backend Container]
            direction LR
            cmd_main[cmd/main.go] --> internal_interfaces[internal/interfaces]
            internal_interfaces --> internal_usecase[internal/usecase]
            internal_usecase --> internal_domain[internal/domain]
            internal_usecase --> internal_infrastructure[internal/infrastructure]
            internal_interfaces --> internal_infrastructure
        end

        mongo_container[Mongo Container]
        mongo_express_container[Mongo Express Container]
        swagger_ui_container[Swagger UI Container]

        backend_container --> mongo_container
        mongo_express_container --> mongo_container
        backend_container -- serves API --> swagger_ui_container
    end
```

