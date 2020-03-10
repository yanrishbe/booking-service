.PHONY: lint test

lint:
	golangci-lint run --config .golangci.yml

test:
	make start_images
	go test -timeout 30s -cover ./...
	make stop_images

start_images:
	make stop_images
	docker run --rm -d -p 27017:27017 --name mongodb mongo:4.2

stop_images:
	docker rm -f mongodb || true

#swagger: swagger-spec | swagger-validate
#
#swagger-spec:
#	go mod vendor
#	env GO111MODULE=off SWAGGER_GENERATE_EXTENSION=false swagger -q generate spec -m -o ./swagger.json
#
#swagger-validate:
#	swagger -q validate ./swagger.json