include .env

APP_NAME := go-boilerplate
SOURCE_PATH := ./internal/

CREATE_DOMAIN_CMD := go run ./cmd/create-domain/internal/main.go
MIGRATION_CMD := go run ./cmd/migrations/internal/main.go

DBMATE_URL := postgres://${DB_USER}:${DB_PASS}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable
DBMATE_CMD_PREFIX := dbmate --migrations-dir './migrations' --no-dump-schema
DBMATE_CMD_WITH_URL_PREFIX := ${DBMATE_CMD_PREFIX} --url ${DBMATE_URL}

# Some more documentation on this command for learning purpose:
# The `grep -E '^[a-zA-Z0-9_-]+:' Makefile`, this part finds any lines that matches as commands and its comments.
# For example: "help: ## Shows help command".
#
# The `awk` command has many instructions, so we'll split it:
# - `BEGIN { FS = ":( ##)?" };`, this sets the "file separator" to split the command and the comments.
# - `{ printf "\033[0;31m%-20s \033[0;32m%s\n", $$1, $$2 };`, this will print it as a nice looking help command.
.PHONY: help
help: ## Shows this command.
	@printf 'These are the available commands in our Makefile.\n'
	@printf '-------------------------------------------------\n'
	@grep -E '^[a-zA-Z0-9_-]+:' Makefile | awk 'BEGIN { FS = ":( ##)?" }; { printf "\033[0;31m%-20s\033[0m%s\n", $$1, $$2 };'

.PHONY: clean
clean: ## Cleans the build directory by removing all binary files.
	rm -rf build/*

.PHONY: build
build: ## Builds the app based on your operating system.
	go mod tidy -v
	go build -v -o ./build/$(APP_NAME) $(SOURCE_PATH)

.PHONY: build-prod
build-prod: ## Builds the app for production purpose.
	go mod tidy -v
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -gcflags "all=-trimpath=$(pwd)" -o ./build/$(APP_NAME)_linux_amd64 -v $(SOURCE_PATH)

.PHONY: start
start: ## Starts the app from 'build' directory.
	ENVIRONMENT=production ./build/$(APP_NAME)

.PHONY: start-prod
start-prod: ## Starts the app from 'build' directory for production.
	ENVIRONMENT=production ./build/$(APP_NAME)_linux_amd64

.PHONY: start-dev
start-dev: ## Starts the app with 'air' to allow live/hot reloading as you edit the code.
	ENVIRONMENT=development air -c ./.air.toml

.PHONY: docs-fmt
docs-fmt: ## Format the swagger annotations within the codebase.
	swag fmt -d $(SOURCE_PATH)

.PHONY: docs-gen
docs-gen: docs-fmt ## Generate swagger API documentation for this app.
	mkdir -p ./docs
	swag init -d $(SOURCE_PATH),./pkg/responses

.PHONY: create-domain
create-domain: ## Creates a domain for the app according to boilerplate (ex, make create-domain domain=finance_reports).
	$(CREATE_DOMAIN_CMD) -domain $(domain)

.PHONY: migration-new
migration-new: ## Create a new migration file (ex, migration-new name=create_accounts_table).
	${DBMATE_CMD_PREFIX} new ${name}

.PHONY: migration-status
migration-status: ## Show the migration status.
	${DBMATE_CMD_WITH_URL_PREFIX} status

.PHONY: migration-up
migration-up: ## Execute all migration files.
	${DBMATE_CMD_WITH_URL_PREFIX} up

.PHONY: migration-down
migration-down: ## Rollback 1 migration.
	${DBMATE_CMD_WITH_URL_PREFIX} down