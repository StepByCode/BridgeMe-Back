//go:build tools

package main

//go:generate oapi-codegen --config=api/codegen.yaml api/openapi.yaml
func main() {}