GOOSE := $(shell \
	if command -v goose >/dev/null 2>&1; then \
		command -v goose; \
	elif [ -n "$$(go env GOBIN)" ] && [ -x "$$(go env GOBIN)/goose" ]; then \
		printf "%s/goose" "$$(go env GOBIN)"; \
	elif [ -x "$$(go env GOPATH)/bin/goose" ]; then \
		printf "%s/bin/goose" "$$(go env GOPATH)"; \
	else \
		printf "goose"; \
	fi)

ifneq (,$(wildcard .env))
include .env
endif

DB_HOST ?= localhost
DB_PORT ?= 5432
DB_USER ?= postgres
DB_PASSWORD ?= postgres
DB_NAME ?= pinger
DB_SSL_MODE ?= disable
DATABASE_URL ?= postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL_MODE)

.PHONY: help build run dev lint create-migration migrate-up migrate-down docs-generate install-tools

help:
	@echo "Available commands:"
	@echo "  make build       - Build the application"
	@echo "  make run         - Run the application"
	@echo "  make dev         - Run the application in development mode"
	@echo "  make lint        - Run linter on the codebase"
	@echo "  make format      - Format the code and re-arrange imports"
	@echo "  make create-migration name=create_monitors - Create database migration files"
	@echo "  make migrate-up  - Apply database migrations"
	@echo "  make migrate-down- Rollback database migrations"


build:
	go build -o bin/app ./cmd/api

run:
	go run ./cmd/api

dev:
	go run ./cmd/api

lint: format
	golangci-lint run ./...

format:
	@gofmt -s -w .
	@goimports -w .

install-tools:
	@echo "Installing tools..."
	@go install github.com/pressly/goose/v3/cmd/goose@latest
	@go install mvdan.cc/goimports/v2@latest
	@go install github.com/segmentio/golines@latest

create-migration:
	@if [ -z "$(name)" ]; then echo "usage: make create-migration name=create_monitors"; exit 1; fi
	@mkdir -p internal/database/migrations
	$(GOOSE) -dir internal/database/migrations create $(name) sql

migrate-up:
	$(GOOSE) -dir internal/database/migrations postgres "$(DATABASE_URL)" up

migrate-down:
	$(GOOSE) -dir internal/database/migrations postgres "$(DATABASE_URL)" down

docker-up:
	docker compose -f docker/docker-compose.yml up -d

docker-down:
	docker compose -f docker/docker-compose.yml down
