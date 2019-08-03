CHAP := $(shell echo ${CHAP})
ARGS := $(shell echo ${ARGS})

.PHONY: test
test: lint ## Runs the tests and higlights race conditions
	GO111MODULE=auto go test -race $$(go list ./...)

.PHONY: lint
lint: ## Runs more than 20 different linters using golangci-lint to ensure consistency in code.
ifeq ($(shell which golangci-lint),)
	brew install golangci/tap/golangci-lint
endif
	golangci-lint run

.PHONY: run
run:
	GO111MODULE=auto go run $(CHAP)/main.go $(ARGS)