GOLANGCI_LINT_VERSION := 1.50.0
GOLANGCI_LINT := bin/golangci-lint-$(GOLANGCI_LINT_VERSION)/golangci-lint
PKGSITE_PORT := 3030
PKGSITE := $(shell command -v pkgsite 2> /dev/null)

all: help

## test: run tests
.PHONY: test
test:
	@echo "🧪 Testing the application"
	@go test -race ./...

## coverage: create coverage report in HTML format
.PHONY: coverage
coverage:
	@echo "📜 Creating coverage report in HTML format"
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out

## lint: lint the application
.PHONY: lint
lint: ${GOLANGCI_LINT}
	@echo "🕵️  Linting code"
	@${GOLANGCI_LINT} run

.PHONY: help
help:
	@echo "📓 Run one of the following commands using make <command>"
	@echo
	@cat Makefile | grep "^##" | column -t -s ":" | sed "s/##/  /"
	@echo

## doc: start documentation server
.PHONY: doc
doc:
ifndef PKGSITE
	@echo "❌ Missing binary pkgsite"
	@echo "Install with: 'go install golang.org/x/pkgsite/cmd/pkgsite@latest'"
	@echo "\tand make sure it is avaliable in PATH"
	exit 1
endif
	@echo "📜 Starting documentation server"
	@echo "Go to http://localhost:${PKGSITE_PORT}/github.com/philiplinell/sh to see documentation"
	@pkgsite -http localhost:${PKGSITE_PORT}

${GOLANGCI_LINT}:
	@echo "Please install golang ci lint version ${GOLANGCI_LINT_VERSION}"
	@echo "into ${GOLANGCI_LINT}"
	@echo
	@echo "You should be able to find it here: https://github.com/golangci/golangci-lint/releases/tag/v${GOLANGCI_LINT_VERSION}"
	@exit 1
