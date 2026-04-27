set shell := ["bash", "-c"]

app_name   := "go-boilerplate"
src_rest   := "./internal/apps/rest"
src_models := "./internal/models/mysql"
src_dtos   := "./internal/dtos"
pkg_err    := "./pkg/customerror"

# OS Detection
os_var := if os() == "windows" {
    "windows"
} else {
    if os() == "macos" {
        "darwin"
    } else {
        "linux"
    }
}

bin_ext := if os() == "windows" {
    ".exe"
} else {
    ""
}

# Default target: list all commands
default:
    @just --list

# Clean the build directory
clean:
    rm -rf build/*

# Format and generate swagger docs
docs:
    go install github.com/swaggo/swag/cmd/swag@latest
    swag fmt -d {{src_rest}}
    mkdir -p {{src_rest}}/docs
    swag init -g main.go -d {{src_rest}},{{src_models}},{{src_dtos}},{{pkg_err}} -o {{src_rest}}/docs

# Build REST API (usage: just build-rest target=./my-bin)
build-rest target=('./build/' + app_name + '-rest' + bin_ext): docs
    go mod download -x
    GOOS={{os_var}} go build -v -o {{target}} {{src_rest}}

# Build REST API for production (Linux AMD64)
build-prod target=('./build/' + app_name + '-rest_linux_amd64'):
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -trimpath -ldflags="-s -w" -tags=prod -o {{target}} {{src_rest}}

# Start REST API with live reloading (Air)
start-dev:
    go install github.com/air-verse/air@latest
    air -c ./.air.rest.toml

# Database migrations (usage: just migration-new name=add_users)
migration action name="":
    go run ./cmd/migration/main.go -action={{action}} -name={{name}}

# Run tests (usage: just test args="-v")
test args="":
    go test {{args}} ./pkg/...

# Run tests with coverage report
test-coverage:
    go test -coverprofile=coverage.out ./pkg/...
    go tool cover -html=coverage.out