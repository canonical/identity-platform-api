// Copyright 2025 Canonical Ltd.
// SPDX-License-Identifier: AGPL-3.0

module github.com/canonical/identity-platform-api

go 1.25.0

require (
	buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go v1.36.12-20260825204119-511051f7f437.2
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.30.0
	google.golang.org/genproto/googleapis/api v0.0.0-20260904194346-d0f1323225a4
	google.golang.org/grpc v1.83.1
	google.golang.org/protobuf v1.36.12
)

require (
	golang.org/x/net v0.57.0 // indirect
	golang.org/x/sys v0.47.0 // indirect
	golang.org/x/text v0.40.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260825221802-da73d73af1c5 // indirect
)
