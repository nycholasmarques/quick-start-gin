APP_NAME=devnych
GO_CMD=go
ATLAS=atlas
SQLC=sqlc
SWAG=swag

server:
	air \
	--build.cmd "go build -o tmp/main ./cmd/devnych" \
	--build.bin "tmp/main" \
	--build.delay "100" \
	--build.include_ext "go" \
	--build.stop_on_error "false" \
	--misc.clean_on_exit true

dev:
	make server

run:
	$(GO_CMD) run cmd/app/main.go

build:
	$(GO_CMD) build -o bin/$(APP_NAME) cmd/app/main.go

test:
	$(GO_CMD) test ./... -cover

sqlc:
	$(SQLC) generate

migrate.diff:
	$(ATLAS) migrate diff --env development

migrate:
	$(ATLAS) migrate apply --env development

migrate.status:
	$(ATLAS) migrate status --env development

migrate.lint:
	$(ATLAS) migrate lint --env development

migrate.hash:
	$(ATLAS) migrate hash --env development

fmt:
	$(GO_CMD) fmt ./...

docs:
	$(SWAG) init -g cmd/devnych/main.go -o internal/docs