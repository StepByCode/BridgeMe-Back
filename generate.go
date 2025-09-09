//go:build tools

package main

import (
	_ "github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen"
)

//go:generate oapi-codegen --config=api/codegen.yaml api/openapi.yaml > internal/interfaces/generated/types.go
func main() {}