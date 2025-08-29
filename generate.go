//go:build tools

package main

//go:generate oapi-codegen --config=codegen.yaml api/openapi.yaml
func main() {}