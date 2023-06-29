include .env

APP_NAME := go-boilerplate

DBMATE_URL := postgres://${DB_USER}:${DB_PASS}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable
DBMATE_CMD_PREFIX := dbmate --migrations-dir './migrations' --no-dump-schema
DBMATE_CMD_WITH_URL_PREFIX := ${DBMATE_CMD_PREFIX} --url ${DBMATE_URL}

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

.PHONY: migration-new
migration-new:
	${DBMATE_CMD_PREFIX} new ${name}

.PHONY: migration-status
migration-status:
	${DBMATE_CMD_WITH_URL_PREFIX} status

.PHONY: migration-up
migration-up:
	${DBMATE_CMD_WITH_URL_PREFIX} up

.PHONY: migration-down
migration-down:
	${DBMATE_CMD_WITH_URL_PREFIX} down