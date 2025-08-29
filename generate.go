//go:build tools

package main

//go:generate oapi-codegen --config=codegen.yaml docs/openapi.yaml
func main() {}