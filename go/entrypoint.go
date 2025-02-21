// Copyright 2025 Canonical Ltd.
// SPDX-License-Identifier: AGPL-3.0

package main

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/canonical/identity-platform-api/pkg/roles"
)

type roleService struct {
	roles.UnimplementedRolesServiceServer
}

func (r *roleService) ListRoles(ctx context.Context, empty *emptypb.Empty) (*roles.ListRolesResp, error) {
	log.Println("ListRoles")
	return &roles.ListRolesResp{}, nil
}

func (r *roleService) GetRole(ctx context.Context, req *roles.GetRoleReq) (*roles.ListRolesResp, error) {
	log.Println("GetRole")
	return &roles.ListRolesResp{}, nil
}

func (r *roleService) CreateRole(ctx context.Context, req *roles.CreateRoleReq) (*roles.ListRolesResp, error) {
	log.Println("CreateRole")
	return &roles.ListRolesResp{}, nil
}

func (r *roleService) UpdateRole(ctx context.Context, req *roles.UpdateRoleReq) (*roles.UpdateRoleResp, error) {
	log.Println("UpdateRole")
	return &roles.UpdateRoleResp{}, nil
}

func (r *roleService) RemoveRole(ctx context.Context, req *roles.RemoveRoleReq) (*roles.RemoveRoleResp, error) {
	log.Println("RemoveRole")
	return &roles.RemoveRoleResp{}, nil
}

func (r *roleService) ListRoleEntitlements(ctx context.Context, req *roles.ListRoleEntitlementsReq) (*roles.ListRoleEntitlementsResp, error) {
	log.Println("ListRoleEntitlements")
	return &roles.ListRoleEntitlementsResp{}, nil
}

func (r *roleService) UpdateRoleEntitlements(ctx context.Context, req *roles.UpdateRoleEntitlementsReq) (*emptypb.Empty, error) {
	log.Println("UpdateRoleEntitlements")
	return &emptypb.Empty{}, nil
}

func (r *roleService) RemoveRoleEntitlement(ctx context.Context, req *roles.RemoveRoleEntitlementReq) (*emptypb.Empty, error) {
	log.Println("RemoveRoleEntitlement")
	return &emptypb.Empty{}, nil
}

func (r *roleService) GetRoleGroups(ctx context.Context, req *roles.GetRoleGroupsReq) (*roles.GetRoleGroupsResp, error) {
	log.Println("GetRoleGroups")
	return &roles.GetRoleGroupsResp{}, nil
}

func HTTPHeaderMatcher(key string) (string, bool) {
	switch key {
	case "X-Token-Pagination", "Link":
		return key, true
	default:
		return runtime.DefaultHeaderMatcher(key)
	}
}

func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux(runtime.WithIncomingHeaderMatcher(HTTPHeaderMatcher))

	err := roles.RegisterRolesServiceHandlerServer(ctx, mux, new(roleService))
	if err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(":8081", mux)
}

func main() {
	if err := run(); err != nil {
		grpclog.Fatal(err)
	}
}
