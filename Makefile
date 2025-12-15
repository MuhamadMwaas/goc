
.PHONY: build test migrate migrate-down up down run lint mock swagger

build:
	@echo "Building the application..."
	go build -o bin/application cmd/api/main.go

test: mock
	@echo "Running tests..."
	go test ./...

migrate:
	@echo "Running database migrations..."
	go run cmd/migrate/main.go -command up

migrate-down:
	@echo "Reverting database migrations..."
	go run cmd/migrate/main.go -command down

up:
	@echo "Starting docker-compose..."
	docker-compose up -d

down:
	@echo "Stopping docker-compose..."
	docker-compose down

run:
	@echo "Starting docker-compose with build..."
	docker-compose up --build

lint:
	@echo "Running linter..."
	golangci-lint run ./...

mock:
	@echo "Checking for mockery..."
	@if ! go list -f '{{.Dir}}' github.com/vektra/mockery/v2 > /dev/null 2>&1; then \
		echo "mockery not found, installing..."; \
		go get github.com/vektra/mockery/v2; \
	fi
	@echo "Generating mocks..."
	go run github.com/vektra/mockery/v2 --all --keeptree --output ./mocks

swagger:
	@echo "Generating OpenAPI documentation..."
	swag init -g cmd/api/main.go -o ./openapi --parseDependency
