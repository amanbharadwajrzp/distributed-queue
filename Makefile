# Dir where build binaries are generated. The dir should be gitignored
BUILD_OUT_DIR := "bin/"

API_MAIN_OUT       := "bin/api"
API_MAIN_FILE := "cmd/startup.go"

ABSOLUTE_PATH := $(shell pwd)


# go binary. Change this to experiment with different versions of go.
GO       = go
MODULE   = $(shell $(GO) list -m)
SERVICE  = $(shell basename $(MODULE))

# Fetch OS info
GOVERSION=$(shell go version)
UNAME_OS=$(shell go env GOOS)
UNAME_ARCH=$(shell go env GOARCH)


VERBOSE = 0
Q 		= $(if $(filter 1,$VERBOSE),,@)
M 		= $(shell printf "\033[34;1m▶\033[0m")


BIN 	 = $(CURDIR)/bin
PKGS     = $(or $(PKG),$(shell $(GO) list ./...))

$(BIN)/%: | $(BIN) ; $(info $(M) building package: $(PACKAGE)…)
	tmp=$$(mktemp -d); \
	   env GOBIN=$(BIN) go get $(PACKAGE) \
		|| ret=$$?; \
	   rm -rf $$tmp ; exit $$ret

$(BIN)/golint: PACKAGE=golang.org/x/lint/golint

GOLINT = $(BIN)/golint


.PHONY: lint
lint: | $(GOLINT) ; $(info $(M) running golint…) @ ## Run golint
	$Q $(GOLINT) -set_exit_status $(PKGS)

.PHONY: fmt
fmt: ; $(info $(M) running gofmt…) @ ## Run gofmt on all source files
	$Q $(GO) fmt $(PKGS)

.PHONY: all
all: build

.PHONY: deps
deps:
	@echo "\n + Fetching dependencies \n"
	@go install golang.org/x/lint/golint@latest
	@go install github.com/bykof/go-plantuml@latest
	@go install gotest.tools/gotestsum@latest

.PHONY: pre-build
pre-build: clean deps go-build-api

.PHONY: go-build-api ## Build the binary file for api server
go-build-api:
	@CGO_ENABLED=0 go build -v -o $(API_MAIN_OUT) $(API_MAIN_FILE)

.PHONY: go-run-api ## Run the api server
go-run-api: go-build-api
	@go run $(API_MAIN_FILE)

.PHONY: build
build: build-info pre-build ## Build the binary file for the broker and worker

.PHONY: build-info
build-info:
	@echo "\nBuild Info:\n"
	@echo "\t\033[33mOS\033[0m: $(UNAME_OS)"
	@echo "\t\033[33mArch\033[0m: $(UNAME_ARCH)"
	@echo "\t\033[33mGo Version\033[0m: $(GOVERSION)\n"

.PHONY: clean
clean: ## Remove previous builds
	@echo " + Removing cloned and generated files\n"
	@rm -rf $(API_MAIN_OUT)

