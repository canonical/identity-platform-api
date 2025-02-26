# Copyright 2025 Canonical Ltd.
# SPDX-License-Identifier: AGPL-3.0

.ONESHELL:

GO_GEN_FOLDER=v0
OPENAPI_GEN_FOLDER=openapi
GO=go
GO_BIN=admin-ui-api
BUF_BIN=buf

clean:
	find $(GO_GEN_FOLDER) -not \
		\( -name 'entrypoint.go' -or \
		   -name 'go.mod' -or \
		   -name 'go.sum' -or \
		   -name "$(GO_GEN_FOLDER)" \
	    \) -delete
.PHONY: clean

generate:
	$(BUF_BIN) generate
.PHONY: generate


vendor:
	cd $(GO_GEN_FOLDER)
	$(GO) mod vendor
	cd -
.PHONY: vendor

tidy:
	cd $(GO_GEN_FOLDER)
	$(GO) mod tidy
	cd -
.PHONY: tidy

install-buf:
	# need sudo to put executable inside /usr/local/bin
	sudo install-buf.sh
.PHONY: install-buf

build:
	$(GO) build . -o $(GO_BIN)
.PHONY: build

openapi-v3:
	cd convert
	$(GO) run convert.go
.PHONY: openapi-v3
