MKFILE_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
ROOT := $(dir $(MKFILE_PATH))
GOBIN ?= $(ROOT)/tools/bin
ENV_PATH = PATH=$(GOBIN):$(PATH)
LINTER_NAME := golangci-lint
LINTER_VERSION := v1.60.2

all: build

.PHONY: build
build:
	go build -mod=vendor

.PHONY: test
test:
	go test ./...

.PHONY: vendor
vendor:
	go mod tidy
	go mod vendor

.PHONY: install-linter
install-linter:
	if [ ! -f $(GOBIN)/$(LINTER_VERSION)/$(LINTER_NAME) ]; then \
		echo INSTALLING $(GOBIN)/$(LINTER_VERSION)/$(LINTER_NAME) $(LINTER_VERSION) ; \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOBIN)/$(LINTER_VERSION) $(LINTER_VERSION) ; \
		echo DONE ; \
	fi

.PHONY: lint
lint: install-linter
	$(GOBIN)/$(LINTER_VERSION)/$(LINTER_NAME) run --config .golangci.yml

