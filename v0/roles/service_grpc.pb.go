// Copyright 2025 Canonical Ltd.
// SPDX-License-Identifier: AGPL-3.0

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: v0/roles/service.proto

package roles

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	RolesService_ListRoles_FullMethodName              = "/identity.platform.api.roles.RolesService/ListRoles"
	RolesService_GetRole_FullMethodName                = "/identity.platform.api.roles.RolesService/GetRole"
	RolesService_CreateRole_FullMethodName             = "/identity.platform.api.roles.RolesService/CreateRole"
	RolesService_UpdateRole_FullMethodName             = "/identity.platform.api.roles.RolesService/UpdateRole"
	RolesService_RemoveRole_FullMethodName             = "/identity.platform.api.roles.RolesService/RemoveRole"
	RolesService_ListRoleEntitlements_FullMethodName   = "/identity.platform.api.roles.RolesService/ListRoleEntitlements"
	RolesService_UpdateRoleEntitlements_FullMethodName = "/identity.platform.api.roles.RolesService/UpdateRoleEntitlements"
	RolesService_RemoveRoleEntitlement_FullMethodName  = "/identity.platform.api.roles.RolesService/RemoveRoleEntitlement"
	RolesService_GetRoleGroups_FullMethodName          = "/identity.platform.api.roles.RolesService/GetRoleGroups"
)

// RolesServiceClient is the client API for RolesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RolesServiceClient interface {
	ListRoles(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListRolesResp, error)
	GetRole(ctx context.Context, in *GetRoleReq, opts ...grpc.CallOption) (*GetRoleResp, error)
	CreateRole(ctx context.Context, in *CreateRoleReq, opts ...grpc.CallOption) (*CreateRoleResp, error)
	UpdateRole(ctx context.Context, in *UpdateRoleReq, opts ...grpc.CallOption) (*UpdateRoleResp, error)
	RemoveRole(ctx context.Context, in *RemoveRoleReq, opts ...grpc.CallOption) (*RemoveRoleResp, error)
	ListRoleEntitlements(ctx context.Context, in *ListRoleEntitlementsReq, opts ...grpc.CallOption) (*ListRoleEntitlementsResp, error)
	UpdateRoleEntitlements(ctx context.Context, in *UpdateRoleEntitlementsReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RemoveRoleEntitlement(ctx context.Context, in *RemoveRoleEntitlementReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetRoleGroups(ctx context.Context, in *GetRoleGroupsReq, opts ...grpc.CallOption) (*GetRoleGroupsResp, error)
}

type rolesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRolesServiceClient(cc grpc.ClientConnInterface) RolesServiceClient {
	return &rolesServiceClient{cc}
}

func (c *rolesServiceClient) ListRoles(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListRolesResp, error) {
	out := new(ListRolesResp)
	err := c.cc.Invoke(ctx, RolesService_ListRoles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesServiceClient) GetRole(ctx context.Context, in *GetRoleReq, opts ...grpc.CallOption) (*GetRoleResp, error) {
	out := new(GetRoleResp)
	err := c.cc.Invoke(ctx, RolesService_GetRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesServiceClient) CreateRole(ctx context.Context, in *CreateRoleReq, opts ...grpc.CallOption) (*CreateRoleResp, error) {
	out := new(CreateRoleResp)
	err := c.cc.Invoke(ctx, RolesService_CreateRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesServiceClient) UpdateRole(ctx context.Context, in *UpdateRoleReq, opts ...grpc.CallOption) (*UpdateRoleResp, error) {
	out := new(UpdateRoleResp)
	err := c.cc.Invoke(ctx, RolesService_UpdateRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesServiceClient) RemoveRole(ctx context.Context, in *RemoveRoleReq, opts ...grpc.CallOption) (*RemoveRoleResp, error) {
	out := new(RemoveRoleResp)
	err := c.cc.Invoke(ctx, RolesService_RemoveRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesServiceClient) ListRoleEntitlements(ctx context.Context, in *ListRoleEntitlementsReq, opts ...grpc.CallOption) (*ListRoleEntitlementsResp, error) {
	out := new(ListRoleEntitlementsResp)
	err := c.cc.Invoke(ctx, RolesService_ListRoleEntitlements_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesServiceClient) UpdateRoleEntitlements(ctx context.Context, in *UpdateRoleEntitlementsReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, RolesService_UpdateRoleEntitlements_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesServiceClient) RemoveRoleEntitlement(ctx context.Context, in *RemoveRoleEntitlementReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, RolesService_RemoveRoleEntitlement_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesServiceClient) GetRoleGroups(ctx context.Context, in *GetRoleGroupsReq, opts ...grpc.CallOption) (*GetRoleGroupsResp, error) {
	out := new(GetRoleGroupsResp)
	err := c.cc.Invoke(ctx, RolesService_GetRoleGroups_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RolesServiceServer is the server API for RolesService service.
// All implementations must embed UnimplementedRolesServiceServer
// for forward compatibility
type RolesServiceServer interface {
	ListRoles(context.Context, *emptypb.Empty) (*ListRolesResp, error)
	GetRole(context.Context, *GetRoleReq) (*GetRoleResp, error)
	CreateRole(context.Context, *CreateRoleReq) (*CreateRoleResp, error)
	UpdateRole(context.Context, *UpdateRoleReq) (*UpdateRoleResp, error)
	RemoveRole(context.Context, *RemoveRoleReq) (*RemoveRoleResp, error)
	ListRoleEntitlements(context.Context, *ListRoleEntitlementsReq) (*ListRoleEntitlementsResp, error)
	UpdateRoleEntitlements(context.Context, *UpdateRoleEntitlementsReq) (*emptypb.Empty, error)
	RemoveRoleEntitlement(context.Context, *RemoveRoleEntitlementReq) (*emptypb.Empty, error)
	GetRoleGroups(context.Context, *GetRoleGroupsReq) (*GetRoleGroupsResp, error)
	mustEmbedUnimplementedRolesServiceServer()
}

// UnimplementedRolesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRolesServiceServer struct {
}

func (UnimplementedRolesServiceServer) ListRoles(context.Context, *emptypb.Empty) (*ListRolesResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRoles not implemented")
}
func (UnimplementedRolesServiceServer) GetRole(context.Context, *GetRoleReq) (*GetRoleResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRole not implemented")
}
func (UnimplementedRolesServiceServer) CreateRole(context.Context, *CreateRoleReq) (*CreateRoleResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRole not implemented")
}
func (UnimplementedRolesServiceServer) UpdateRole(context.Context, *UpdateRoleReq) (*UpdateRoleResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRole not implemented")
}
func (UnimplementedRolesServiceServer) RemoveRole(context.Context, *RemoveRoleReq) (*RemoveRoleResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveRole not implemented")
}
func (UnimplementedRolesServiceServer) ListRoleEntitlements(context.Context, *ListRoleEntitlementsReq) (*ListRoleEntitlementsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRoleEntitlements not implemented")
}
func (UnimplementedRolesServiceServer) UpdateRoleEntitlements(context.Context, *UpdateRoleEntitlementsReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRoleEntitlements not implemented")
}
func (UnimplementedRolesServiceServer) RemoveRoleEntitlement(context.Context, *RemoveRoleEntitlementReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveRoleEntitlement not implemented")
}
func (UnimplementedRolesServiceServer) GetRoleGroups(context.Context, *GetRoleGroupsReq) (*GetRoleGroupsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoleGroups not implemented")
}
func (UnimplementedRolesServiceServer) mustEmbedUnimplementedRolesServiceServer() {}

// UnsafeRolesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RolesServiceServer will
// result in compilation errors.
type UnsafeRolesServiceServer interface {
	mustEmbedUnimplementedRolesServiceServer()
}

func RegisterRolesServiceServer(s grpc.ServiceRegistrar, srv RolesServiceServer) {
	s.RegisterService(&RolesService_ServiceDesc, srv)
}

func _RolesService_ListRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServiceServer).ListRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RolesService_ListRoles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServiceServer).ListRoles(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _RolesService_GetRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServiceServer).GetRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RolesService_GetRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServiceServer).GetRole(ctx, req.(*GetRoleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RolesService_CreateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServiceServer).CreateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RolesService_CreateRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServiceServer).CreateRole(ctx, req.(*CreateRoleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RolesService_UpdateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRoleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServiceServer).UpdateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RolesService_UpdateRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServiceServer).UpdateRole(ctx, req.(*UpdateRoleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RolesService_RemoveRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveRoleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServiceServer).RemoveRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RolesService_RemoveRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServiceServer).RemoveRole(ctx, req.(*RemoveRoleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RolesService_ListRoleEntitlements_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRoleEntitlementsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServiceServer).ListRoleEntitlements(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RolesService_ListRoleEntitlements_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServiceServer).ListRoleEntitlements(ctx, req.(*ListRoleEntitlementsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RolesService_UpdateRoleEntitlements_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRoleEntitlementsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServiceServer).UpdateRoleEntitlements(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RolesService_UpdateRoleEntitlements_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServiceServer).UpdateRoleEntitlements(ctx, req.(*UpdateRoleEntitlementsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RolesService_RemoveRoleEntitlement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveRoleEntitlementReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServiceServer).RemoveRoleEntitlement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RolesService_RemoveRoleEntitlement_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServiceServer).RemoveRoleEntitlement(ctx, req.(*RemoveRoleEntitlementReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RolesService_GetRoleGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoleGroupsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServiceServer).GetRoleGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RolesService_GetRoleGroups_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServiceServer).GetRoleGroups(ctx, req.(*GetRoleGroupsReq))
	}
	return interceptor(ctx, in, info, handler)
}

// RolesService_ServiceDesc is the grpc.ServiceDesc for RolesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RolesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "identity.platform.api.roles.RolesService",
	HandlerType: (*RolesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListRoles",
			Handler:    _RolesService_ListRoles_Handler,
		},
		{
			MethodName: "GetRole",
			Handler:    _RolesService_GetRole_Handler,
		},
		{
			MethodName: "CreateRole",
			Handler:    _RolesService_CreateRole_Handler,
		},
		{
			MethodName: "UpdateRole",
			Handler:    _RolesService_UpdateRole_Handler,
		},
		{
			MethodName: "RemoveRole",
			Handler:    _RolesService_RemoveRole_Handler,
		},
		{
			MethodName: "ListRoleEntitlements",
			Handler:    _RolesService_ListRoleEntitlements_Handler,
		},
		{
			MethodName: "UpdateRoleEntitlements",
			Handler:    _RolesService_UpdateRoleEntitlements_Handler,
		},
		{
			MethodName: "RemoveRoleEntitlement",
			Handler:    _RolesService_RemoveRoleEntitlement_Handler,
		},
		{
			MethodName: "GetRoleGroups",
			Handler:    _RolesService_GetRoleGroups_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v0/roles/service.proto",
}
