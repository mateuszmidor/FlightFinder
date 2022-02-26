SPEC_PATH = api/openapi3.yaml
APISERVER_PATH = cmd/finder_web/app/apiserver

default: apiserver

docker:
	docker build . -t flight-finder
	
rundocker:
	docker run --rm --name=flight-finder -p 8080:80 flight-finder:latest
	
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
		-v "$(shell pwd):/build" \
		--entrypoint=/bin/bash \
		openapitools/openapi-generator-cli:latest -c "cd /build && scripts/build_openapi3.sh"

.PHONY: ${APISERVER_PATH}