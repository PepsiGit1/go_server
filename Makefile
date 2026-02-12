.PHONY: run build test clean install help

# Variables
BINARY_NAME=server
MAIN_PATH=cmd/api/main.go

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

install: ## Install dependencies
	go mod download
	go mod tidy

run: ## Run the application
	go run $(MAIN_PATH)

build: ## Build the application
	go build -o bin/$(BINARY_NAME) $(MAIN_PATH)

test: ## Run tests
	go test -v ./...

test-coverage: ## Run tests with coverage
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

clean: ## Clean build files
	rm -rf bin/
	rm -f coverage.out

fmt: ## Format code
	go fmt ./...

lint: ## Run linter
	golangci-lint run

dev: ## Run in development mode with hot reload (requires air)
	air

docker-build: ## Build docker image
	docker build -t go-server-api .

docker-run: ## Run docker container
	docker run -p 8080:8080 go-server-api
