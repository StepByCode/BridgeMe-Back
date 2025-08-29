# Spec-First移行計画メモ

## 目的

プロジェクトの将来的な大規模化に備え、現在のCode-FirstアプローチからSpec-Firstアプローチへ移行する。
API仕様書を唯一の信頼できる情報源 (Single Source of Truth) とし、チーム間の認識齟齬の解消、及びフロントエンドチームとの並行開発の促進を目指す。

## 移行手順

以下に、明日実施するタスクの計画を記す。

### 1. コード生成ツールの導入

- OpenAPI仕様書からGoのコードを生成するため、`oapi-codegen` をインストールする。

  ```bash
  go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
  ```

### 2. API仕様書の整備

- `docs/openapi.yaml` を、コード生成のマスターデータとして再整備する。
- 全てのエンドポイント、リクエスト/レスポンスの型、コンポーネントが正確に定義されていることを確認する。

### 3. コード生成コマンドの準備

- `go:generate` ディレクティブを使用して、コード生成を簡単に再現できるようにする。
- 例えば、プロジェクトルートに `generate.go` のようなファイルを作成し、以下を記述する。

  ```go
  //go:generate oapi-codegen --config=codegen.yaml docs/openapi.yaml
  ```

- 上記コマンドで参照する `codegen.yaml` 設定ファイルを作成し、出力先パッケージやファイル名を指定する。（例: `interfaces/generated/` 以下に出力）

### 4. 実装の修正

- `oapi-codegen` で自動生成されたサーバーインターフェースを、`interfaces/controllers/profile_controller.go` が実装するようにリファクタリングする。
- これまで手書きで定義していたリクエスト/レスポンス用の構造体を、自動生成された型に置き換える。

### 5. 既存コードのクリーンアップ

- 自動生成されたコードによって不要になった、手書きの型定義や構造体をプロジェクトから削除する。
- `swaggo/swag` 関連のコメント (`// @Summary` など) は不要になるため、合わせて削除する。

---

**担当:** Gemini
**期日:** 明日
