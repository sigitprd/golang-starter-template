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