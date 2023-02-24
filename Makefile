lint:
	gofmt -s -w .

run:
	go buid PinGo/cmd/server

swag:
	swag init -g pkg/app/handlers.go --output docs/
