K6_VERSION ?= latest
K6_DOCS_PATH ?= ./k6-docs

.PHONY: help lint test test-gh prepare

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-15s %s\n", $$1, $$2}'

lint: ## Run linters
	golangci-lint run ./...

test: test-gh ## Run all tests
	go test -race -count=1 ./...

test-gh: ## Run GitHub Actions script tests
	@for t in .github/scripts/*_test.sh; do echo "=== $$t ==="; bash "$$t" || exit 1; done

prepare: ## Prepare docs bundle for a version
	go run ./cmd/prepare --k6-version=$(K6_VERSION) --k6-docs-path=$(K6_DOCS_PATH)

.DEFAULT_GOAL := help
