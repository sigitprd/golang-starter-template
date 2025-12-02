include .env
export

build:
	@mkdir -p $(APP_BIN_DIR)
	GO111MODULE=on CGO_ENABLED=0 go build -o $(APP_BIN_DIR)/$(APP_NAME) ./cmd/server/main.go

build-linux:
	@mkdir -p $(APP_BIN_DIR)
	GOOS=linux GOARCH=amd64 GO111MODULE=on CGO_ENABLED=0 go build -o $(APP_BIN_DIR)/$(APP_NAME)-linux ./cmd/server/main.go

build-mac:
	@mkdir -p $(APP_BIN_DIR)
	GOOS=darwin GOARCH=amd64 GO111MODULE=on CGO_ENABLED=0 go build -o $(APP_BIN_DIR)/$(APP_NAME)-mac ./cmd/server/main.go

build-windows:
	@mkdir -p $(APP_BIN_DIR)
	GOOS=windows GOARCH=amd64 GO111MODULE=on CGO_ENABLED=0 go build -o $(APP_BIN_DIR)/$(APP_NAME).exe ./cmd/server/main.go

migrate-new:
	@read -p "Migration name: " name; \
	timestamp=$$(date +%Y%m%d%H%M%S); \
	touch ./migrations/$${timestamp}_$${name}.up.sql; \
	touch ./migrations/$${timestamp}_$${name}.down.sql; \
	echo "✅ Created: $${timestamp}_$${name}.[up|down].sql"

migrate-up:
	migrate -path ./migrations -database "postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" up

migrate-down:
	migrate -path ./migrations -database "postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" down 1

migrate-drop:
	migrate -path ./migrations -database "postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" drop -f

run:
	go run ./cmd/server/main.go

test:
	go test -v -race -cover ./...

test-coverage:
	go test -v -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "✅ Coverage report generated: coverage.html"

lint:
	@command -v golangci-lint >/dev/null 2>&1 || { echo "golangci-lint not installed. Run: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; exit 1; }
	golangci-lint run ./...

fmt:
	go fmt ./...
	@command -v goimports >/dev/null 2>&1 && goimports -w . || echo "⚠️  goimports not installed. Run: go install golang.org/x/tools/cmd/goimports@latest"

tidy:
	go mod tidy

clean:
	rm -rf $(APP_BIN_DIR)
	rm -f coverage.out coverage.html

help:
	@echo "Available targets:"
	@echo "  make run              - Run the application"
	@echo "  make build            - Build the application"
	@echo "  make build-linux      - Build for Linux"
	@echo "  make build-mac        - Build for macOS"
	@echo "  make build-windows    - Build for Windows"
	@echo "  make test             - Run tests"
	@echo "  make test-coverage    - Run tests with coverage report"
	@echo "  make lint             - Run linter"
	@echo "  make fmt              - Format code"
	@echo "  make tidy             - Tidy dependencies"
	@echo "  make migrate-new      - Create new migration"
	@echo "  make migrate-up       - Run migrations up"
	@echo "  make migrate-down     - Run migrations down"
	@echo "  make migrate-drop     - Drop all migrations"
	@echo "  make clean            - Clean build artifacts"

.PHONY: build build-linux build-mac build-windows migrate-new migrate-up migrate-down migrate-drop run test test-coverage lint fmt tidy clean help