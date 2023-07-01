include .env

APP_NAME := go-boilerplate

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
	go build -v -o ./build/$(APP_NAME) ./internal/main.go

.PHONY: start
start: ## Starts the app from 'build' directory.
	./build/$(APP_NAME)

.PHONY: start-dev
start-dev: ## Starts the app with 'air' to allow live/hot reloading as you edit the code.
	air -c ./.air.toml

.PHONY: create-feature
create-feature: ## Creates a feature for the app according to boilerplate (ex, make create-feature name=find_todo domain=todos).
	go run ./cmd/create-feature/internal/main.go -name $(name) -domain $(domain)

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