// Copyright 2025 Canonical Ltd.
// SPDX-License-Identifier: AGPL-3.0

package api

import (
	"context"
	"flag"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/grpclog"
)

var (
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:9090", "gRPC server endpoint")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()

	// Uncomment when initial APIs are actually added
	//
	//opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	//
	//err := RegisterYourServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	//if err != nil {
	//	return err
	//}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(":8081", mux)
}

func main() {
	flag.Parse()

	if err := run(); err != nil {
		grpclog.Fatal(err)
	}
}
