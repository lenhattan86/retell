APP_NAME=api-server
BUILDER_IMAGE=lenhattan86/golang:1.21-alpine
APP_IMAGE=lenhattan86/$(APP_NAME):latest

.PHONY: build run lint test docker docker-builder

build:
	go build -o $(APP_NAME) main.go

run: build
	./$(APP_NAME)

lint:
	golangci-lint run ./...

test:
	go test ./...

docker-builder:
	docker build -f Dockerfile.builder -t $(BUILDER_IMAGE) .
	docker push $(BUILDER_IMAGE)

docker:
	docker build -t lenhattan86/$(APP_IMAGE):latest .
	docker push lenhattan86/$(APP_IMAGE):latest
