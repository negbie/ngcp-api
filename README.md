# ngcp-api-client

swagger generate client -f Swagger20.json

For this generation to compile you need to have some packages in your GOPATH:

	* github.com/go-openapi/errors
	* github.com/go-openapi/runtime
	* github.com/go-openapi/runtime/client
	* github.com/go-openapi/strfmt

You can get these now with: go get -u -f ./...

oapi-codegen -generate "types,client,spec" OpenApi3Json.json > clientOpenAPI.go
java -jar openapi-generator-cli.jar generate -g go -o api3 -i raw_swagger.json
