GOLANGCI_LINT_VERSION := 1.50.0
GOLANGCI_LINT := bin/golangci-lint-$(GOLANGCI_LINT_VERSION)/golangci-lint

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

${GOLANGCI_LINT}:
	@echo "Please install golang ci lint version ${GOLANGCI_LINT_VERSION}"
	@echo "into ${GOLANGCI_LINT}"
	@echo
	@echo "You should be able to find it here: https://github.com/golangci/golangci-lint/releases/tag/v${GOLANGCI_LINT_VERSION}"
	@exit 1
