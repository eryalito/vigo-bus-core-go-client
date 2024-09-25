swagger2openapi swagger.json -o openapi.yaml
openapi-generator-cli generate -i openapi.yaml -g go -o .
find . -type f -exec sed -i 's|github.com/eryalito/vigo-bus-core-go-client|github.com/eryalito/vigo-bus-core-go-client|g' {} +
go mod tidy
