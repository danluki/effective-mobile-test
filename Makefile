.PHONY: generate migrate-create dev tests gen

ifneq (,$(wildcard ./.env))
	include .env
	export
endif

generate:
	go generate ./...

gen: generate

migrate-create:
	atlas migrate diff --env gorm

migrate-apply:
	atlas migrate apply \
	--url $(DATABASE_URL) \
	--dir "file://internal/database/migrations"

lint:
	golangci-lint run

dev:
	air

tests:
	go test -parallel=20 -covermode atomic -coverprofile=coverage.out ./...

build: generate
	rm ./build-out || true
	go build -ldflags="-s -w" -o build-out cmd/main.go
	upx -9 0q ./build-out

