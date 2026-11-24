// Copyright 2025 Canonical Ltd.
// SPDX-License-Identifier: AGPL-3.0

module github.com/canonical/identity-platform-api

go 1.26.0

require (
	buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go v1.36.12-20260825204119-511051f7f437.2
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.34.0
	google.golang.org/genproto/googleapis/api v0.0.0-20261123191846-1e707ab9fbc2
	google.golang.org/grpc v1.85.0
	google.golang.org/protobuf v1.36.12
)

require (
	golang.org/x/net v0.62.0 // indirect
	golang.org/x/sys v0.49.0 // indirect
	golang.org/x/text v0.44.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20261118221215-c8830f628b3f // indirect
)
