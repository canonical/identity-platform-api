# Copyright 2025 Canonical Ltd.
# SPDX-License-Identifier: AGPL-3.0

GO_GEN_FOLDER=go
OPENAPI_GEN_FOLDER=openapi
GO=go
GO_BIN=admin-ui-api
BUF_BIN=buf

clean:
	rm -rf $(GO_GEN_FOLDER)/pkg
	rm -rf $(OPENAPI_GEN_FOLDER)/*
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
	$(GO) run convert.go
.PHONY: openapi-v3
