SPEC_PATH = api/openapi3.yaml
APISERVER_PATH = cmd/finder_web/app/apiserver
APISERVER_PORT = 8080

default: apiserver

docker: apiserver
	docker build . -t flight-finder
	
runcli: 
	go run cmd/finder_cli/main.go -flights_data=./assets

runweb: 
	GIN_MODE=release go run cmd/finder_web/main.go -flights_data=./assets -web_data=./web

buildweb: 
	CGO_ENABLED=0 go build -o server cmd/finder_web/main.go

test: apiserver
	go vet ./...
	go test ./...
	
# Generate Go server stub from OpenAPI3 specification
# see: https://github.com/OpenAPITools/openapi-generator
# see: https://openapi-generator.tech/docs/generators/go-gin-server
# see: https://openapi-generator.tech/docs/generators/go
apiserver: ${SPEC_PATH}
	docker run \
		--rm \
		-u $(shell id -u ${USER}):$(shell id -g ${USER}) \
		-v "$(shell pwd):/local" \
		openapitools/openapi-generator-cli:latest generate \
		-i "/local/${SPEC_PATH}" \
		-g go-gin-server \
		--additional-properties=enumClassPrefix=true,serverPort=${APISERVER_PORT} \
		--package-name "apiserver" \
		-o "/local/${APISERVER_PATH}"
	rm "${APISERVER_PATH}/main.go" # Remove default main.go
	rm "${APISERVER_PATH}/go/api_default.go" # Remove default handlers
	sed "/\/\/ Index/,+3  d" -i "${APISERVER_PATH}/go/routers.go" # Remove default "Index" handler


.PHONY: ${APISERVER_PATH}