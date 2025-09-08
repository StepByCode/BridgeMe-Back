# 開発者向けドキュメント (Development)

## API Spec-First開発

本プロジェクトは、APIの仕様を最初に定義する **Spec-First** の開発アプローチを採用しています。
API仕様は、プロジェクトルートの `docs/openapi.yaml` にOpenAPI 3.0形式で記述されています。このファイルが、APIに関する唯一の信頼できる情報源（Single Source of Truth）です。

### API仕様の変更フロー

1.  **`docs/openapi.yaml` を編集する**
    -   エンドポイントの追加、リクエスト/レスポンスの型、パラメータなどを修正します。

2.  **コードを自動生成する**
    -   `openapi.yaml` の変更をGoのコードに反映させるため、プロジェクトルートで以下のコマンドを実行します。

    ```bash
    go generate
    ```

    このコマンドは、`generate.go` ファイルに記述された `//go:generate` ディレクティブをトリガーします。
    `oapi-codegen` が `codegen.yaml` の設定に基づいて、`interfaces/generated/` ディレクトリに型定義とサーバーインターフェースを生成・更新します。

3.  **実装を修正する**
    -   コード生成によって `interfaces/generated/types.go` のサーバーインターフェース (`ServerInterface`) が変更された場合、`interfaces/controllers/profile_controller.go` がそのインターフェースを満たすように修正する必要があります。
    -   例えば、新しいエンドポイントを追加した場合は、コントローラーにそのエンドポイントに対応するメソッドを実装します。

## サーバーの起動

以下のコマンドで、開発サーバーを起動できます。

```bash
go run main.go
```

## データベース (Database)

MongoDBを使用しています。
接続情報は環境変数 `MONGODB_URI` で設定します。

## デプロイ (Deployment)

本番環境へのデプロイはGitHub Actionsを通じて行われます。
詳細なデプロイフローについては、プロジェクトルートの `README.md` を参照してください。
