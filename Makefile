VERSION_PKG := $(shell go list -m)/internal/version
GIT_COMMIT ?= $(shell git rev-parse --short HEAD 2>/dev/null)$(shell test -n "`git status --porcelain 2>/dev/null`" && echo "+CHANGES")
LDFLAGS := -X '$(VERSION_PKG).GitCommit=$(GIT_COMMIT)' -X '$(VERSION_PKG).BuildTime=$(shell date -u +%FT%TZ)'

export DATABASE_URL ?= postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable

.PHONY: run build test e2e lint generate docker

run:
	go run -race -ldflags "$(LDFLAGS)" . server

build:
	CGO_ENABLED=0 go build -ldflags "$(LDFLAGS)" -o bin/app .

test:
	go test -race ./...

e2e:
	go test -race -tags e2e ./...

lint:
	golangci-lint run ./...

generate:
	go generate ./...

docker:
	docker build --build-arg GIT_COMMIT=$(GIT_COMMIT) -t $(notdir $(shell go list -m)) .
