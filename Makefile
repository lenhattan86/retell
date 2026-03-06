APP_NAME=api-server

.PHONY: build run lint test docker

build:
	go build -o $(APP_NAME) main.go

run: build
	./$(APP_NAME)

lint:
	golangci-lint run ./...

test:
	go test ./...

docker:
	docker build -t lenhatan86/$(APP_NAME):latest .
