package middleware

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/labstack/echo/v4"
)

// Auth0のドメインとAudienceを設定
var (
	auth0Domain   = os.Getenv("AUTH0_DOMAIN")
	auth0Audience = os.Getenv("AUTH0_AUDIENCE")
)

// CustomClaimsは、IDトークンから抽出されるクレームを定義
type CustomClaims struct {
	Scope string `json:"scope"`
}

// ValidateはCustomClaimsが有効かどうかを検証
func (c CustomClaims) Validate(ctx context.Context) error {
	// スコープの検証など、追加のバリデーションロジックをここに追加
	return nil
}

// NewAuthMiddlewareはJWTを検証する新しいミドルウェアを作成
func NewAuthMiddleware() (*jwtmiddleware.JWTMiddleware, error) {
	// JWTの検証に使うJwks URI
	jwksURI := fmt.Sprintf("https://%s/.well-known/jwks.json", auth0Domain)
	parsedURL, err := url.Parse(jwksURI)
	if err != nil {
		return nil, err // URLのパースに失敗した場合はエラーを返す
	}

	// JWKSリゾーバーの作成
	provider := jwks.NewCachingProvider(parsedURL, 5*time.Minute)

	// JWT検証オブジェクトの作成
	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		jwksURI,
		[]string{auth0Audience},
		validator.WithAllowedClockSkew(30*time.Second),
	)
	if err != nil {
		return nil, err
	}

	// JWTミドルウェアの作成
	middleware := jwtmiddleware.New(jwtValidator.ValidateToken)
	return middleware, nil
}

// Auth0ミドルウェアをEchoミドルウェアに変換するアダプター
func Auth0EchoMiddleware(authMiddleware *jwtmiddleware.JWTMiddleware) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// EchoのContextから http.ResponseWriter と http.Request を取得
			w := c.Response().Writer
			r := c.Request()

			// authMiddleware.CheckJWT を実行
			// 認証が成功したら次のハンドラを呼び出す
			authMiddleware.CheckJWT(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// 認証が成功したら、元のEchoハンドラを呼び出す
				c.SetRequest(r) // リクエストが変更されている可能性があるので更新
				c.SetResponse(echo.NewResponse(w, c.Echo())) // レスポンスも更新
				next(c)
			})).ServeHTTP(w, r)

			return nil // ここでエラーハンドリングが必要になる可能性あり
		}
	}
}
