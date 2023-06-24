APP_NAME := go-boilerplate

.PHONY: build
build:
	go build -v -o ./build/${APP_NAME} ./internal/main.go

.PHONY: start
start:
	./build/${APP_NAME}

.PHONY: start-dev
start-dev:
	air -c ./.air.toml

.PHONY: create-feature
create-feature:
	go run ./cmd/create-feature/main.go -name ${name} -domain ${domain}