.PHONY:
.DEFAULT_GOAL := build

lint:
	golangci-lint run

fmt:
	gofmt -s -w .

build:
	go mod download && CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/main.go

run: build
	./app
