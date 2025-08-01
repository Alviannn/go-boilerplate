APP_NAME := go-boilerplate

SOURCE_PATH           := ./internal
SOURCE_DTOS_PATH      := $(SOURCE_PATH)/dtos
SOURCE_MODELS_PATH    := $(SOURCE_PATH)/models
SOURCE_REST_PATH      := $(SOURCE_PATH)/apps/rest
PKG_CUSTOM_ERROR_PATH := ./pkg/customerror

MIGRATION_CMD := go run ./cmd/migration/main.go

TEST_PATH_LIST            := ./pkg/...
TEST_COVERAGE_OUTPUT_FILE := test-coverage.out

GOOS_VAR := linux
BIN_EXT  :=

ifeq ($(OS), Windows_NT)
	GOOS_VAR := windows
	BIN_EXT := .exe
endif

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
	@grep -E '^[a-zA-Z0-9_-]+:' Makefile | awk 'BEGIN { FS = " ##" }; { printf "\033[0;31m%-20s\033[0m%s\n", $$1, $$2 };'

.PHONY: clean
clean: ## Cleans the build directory by removing all binary files.
	rm -rf build/*

.PHONY: docs-fmt
docs-fmt: ## Format the swagger annotations within the codebase.
	command -v swag >/dev/null 2>&1 || go install github.com/swaggo/swag/cmd/swag@latest
	swag fmt -d $(SOURCE_REST_PATH)

.PHONY: docs-gen
docs-gen: docs-fmt ## Generate swagger API documentation for this app.
	mkdir -p $(SOURCE_REST_PATH)/docs
	swag init -d\
		$(SOURCE_REST_PATH),\
	$(shell find $(SOURCE_MODELS_PATH)/* -type d | tr '\n' ',' | sed 's/,$$//'),\
	$(SOURCE_DTOS_PATH),\
	$(PKG_CUSTOM_ERROR_PATH)\
		-o $(SOURCE_REST_PATH)/docs

# --------------------------------------v REST API v-------------------------------------- #

.PHONY: build-rest
build-rest: docs-gen ## Builds REST API app based on your operating system.
	go mod tidy -v
	GOOS=$(GOOS_VAR) go build -v -o ./build/$(APP_NAME)-rest$(BIN_EXT) $(SOURCE_REST_PATH)

.PHONY: build-rest-prod
build-rest-prod: ## Builds REST API app for production purpose.
	go mod download -x
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -trimpath -ldflags="-s -w" -tags=prod -o $(or $(target), ./build/$(APP_NAME)-rest_linux_amd64) $(SOURCE_REST_PATH)

.PHONY: start-rest
start-rest: ## Starts REST API app from 'build' directory.
	ENVIRONMENT=production ./build/$(APP_NAME)-rest

.PHONY: start-rest-prod
start-rest-prod: ## Starts REST API app from 'build' directory for production.
	ENVIRONMENT=production ./build/$(APP_NAME)-rest_linux_amd64

.PHONY: start-rest-dev
start-rest-dev: ## Starts REST API app with 'air' to allow live/hot reloading as you edit the code.
	command -v air >/dev/null 2>&1 || go install github.com/air-verse/air@latest
	ENVIRONMENT=development air -c ./.air.rest.toml

# --------------------------------------v DB MIGRATIONS v-------------------------------------- #

.PHONY: migration-new
migration-new: ## Create a new migration file (ex, migration-new name=create_accounts_table).
	$(MIGRATION_CMD) -action=new -name=$(name)

.PHONY: migration-status
migration-status: ## Show the migration status.
	$(MIGRATION_CMD) -action=status

.PHONY: migration-up
migration-up: ## Execute all migration files.
	$(MIGRATION_CMD) -action=up

.PHONY: migration-down
migration-down: ## Rollback 1 migration.
	$(MIGRATION_CMD) -action=down

# --------------------------------------v TEST v-------------------------------------- #

.PHONY: test
test: ## Run unit tests.
	go test $(args) $(TEST_PATH_LIST)

.PHONY: test-coverage
test-coverage: ## Run unit tests with coverage.
	go test $(args) -coverprofile=$(TEST_COVERAGE_OUTPUT_FILE) $(TEST_PATH_LIST)
	go tool cover -html=$(TEST_COVERAGE_OUTPUT_FILE)
	rm $(TEST_COVERAGE_OUTPUT_FILE)