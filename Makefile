lint:
	gofmt -s -w .

build:
	mkdir -p bin && go build -o bin/ ./cmd/...

swag:
	swag init -g pkg/app/handlers.go --output docs/

run_server:
	bin/server