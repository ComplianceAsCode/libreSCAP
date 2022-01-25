PROJECT_NAME := "libreSCAP"

PKG := "github.com/ComplianceAsCode/$(PROJECT_NAME)"

PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)

GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)

.PHONY: all dep lint vet test test-coverage build clean

all: build

dep: ## Get the dependencies
	@go mod download

lint: ## Lint the Golang files
	@golint -set_exit_status ${PKG_LIST}

vet: ## Run go vet
	@go vet ${PKG_LIST}

test: ## Run unit tests
	@go test -short ${PKG_LIST}

test-coverage: ## Run tests with coverage reports
	@go test -short -coverprofile cover.out -covermode=atomic ${PKG_LIST}

build: dep ## Build the binary file
	@go build -i -o build/main ${PKG}

clean: ## Remove previous build
	@rm -f ${PROJECT_NAME)/build

help: ## Display help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
