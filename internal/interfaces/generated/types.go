package generated

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// ProfileInput defines model for ProfileInput.
type ProfileInput struct {
	Affiliation *string `json:"affiliation,omitempty"`
	Bio         *string `json:"bio,omitempty"`
	InstagramId *string `json:"instagramId,omitempty"`
	Name        *string `json:"name,omitempty"`
	TwitterId   *string `json:"twitterId,omitempty"`
}

// Profile defines model for Profile.
type Profile struct {
	Affiliation *string `json:"affiliation,omitempty"`
	Bio         *string `json:"bio,omitempty"`
	CreatedAt   *string `json:"createdAt,omitempty"`
	Id          *string `json:"id,omitempty"`
	InstagramId *string `json:"instagramId,omitempty"`
	Name        *string `json:"name,omitempty"`
	TwitterId   *string `json:"twitterId,omitempty"`
}

// CreateProfileJSONRequestBody defines body for CreateProfile for application/json ContentType.
type CreateProfileJSONRequestBody = ProfileInput

// UpdateProfileJSONRequestBody defines body for UpdateProfile for application/json ContentType.
type UpdateProfileJSONRequestBody = ProfileInput

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// 全プロフィールの一覧を取得する
	// (GET /profiles)
	GetProfiles(ctx echo.Context) error
	// 新しいプロフィールを作成する
	// (POST /profiles)
	CreateProfile(ctx echo.Context) error
	// 指定したIDのプロフィール情報を取得する
	// (GET /profiles/{id})
	GetProfileById(ctx echo.Context, id openapi_types.UUID) error
	// 指定したIDのプロフィール情報を更新する
	// (PUT /profiles/{id})
	UpdateProfile(ctx echo.Context, id openapi_types.UUID) error
	// 指定したIDのプロフィール情報を削除する
	// (DELETE /profiles/{id})
	DeleteProfile(ctx echo.Context, id openapi_types.UUID) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetProfiles converts echo context to params.
func (w *ServerInterfaceWrapper) GetProfiles(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetProfiles(ctx)
	return err
}

// CreateProfile converts echo context to params.
func (w *ServerInterfaceWrapper) CreateProfile(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.CreateProfile(ctx)
	return err
}

// GetProfileById converts echo context to params.
func (w *ServerInterfaceWrapper) GetProfileById(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id openapi_types.UUID

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	return w.Handler.GetProfileById(ctx, id)
}

// UpdateProfile converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateProfile(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id openapi_types.UUID

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	return w.Handler.UpdateProfile(ctx, id)
}

// DeleteProfile converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteProfile(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id openapi_types.UUID

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	return w.Handler.DeleteProfile(ctx, id)
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/profiles", wrapper.GetProfiles)
	router.POST(baseURL+"/profiles", wrapper.CreateProfile)
	router.GET(baseURL+"/profiles/:id", wrapper.GetProfileById)
	router.PUT(baseURL+"/profiles/:id", wrapper.UpdateProfile)
	router.DELETE(baseURL+"/profiles/:id", wrapper.DeleteProfile)
}
