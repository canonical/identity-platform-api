// Copyright 2025 Canonical Ltd.
// SPDX-License-Identifier: AGPL-3.0

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: v0/groups/service.proto

package groups

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
	GroupsService_ListGroups_FullMethodName              = "/identity.platform.api.groups.GroupsService/ListGroups"
	GroupsService_GetGroup_FullMethodName                = "/identity.platform.api.groups.GroupsService/GetGroup"
	GroupsService_CreateGroup_FullMethodName             = "/identity.platform.api.groups.GroupsService/CreateGroup"
	GroupsService_UpdateGroup_FullMethodName             = "/identity.platform.api.groups.GroupsService/UpdateGroup"
	GroupsService_RemoveGroup_FullMethodName             = "/identity.platform.api.groups.GroupsService/RemoveGroup"
	GroupsService_ListGroupEntitlements_FullMethodName   = "/identity.platform.api.groups.GroupsService/ListGroupEntitlements"
	GroupsService_UpdateGroupEntitlements_FullMethodName = "/identity.platform.api.groups.GroupsService/UpdateGroupEntitlements"
	GroupsService_RemoveGroupEntitlement_FullMethodName  = "/identity.platform.api.groups.GroupsService/RemoveGroupEntitlement"
	GroupsService_GetGroupRoles_FullMethodName           = "/identity.platform.api.groups.GroupsService/GetGroupRoles"
	GroupsService_UpdateGroupRoles_FullMethodName        = "/identity.platform.api.groups.GroupsService/UpdateGroupRoles"
	GroupsService_RemoveGroupRole_FullMethodName         = "/identity.platform.api.groups.GroupsService/RemoveGroupRole"
	GroupsService_GetGroupIdentities_FullMethodName      = "/identity.platform.api.groups.GroupsService/GetGroupIdentities"
	GroupsService_UpdateGroupIdentities_FullMethodName   = "/identity.platform.api.groups.GroupsService/UpdateGroupIdentities"
	GroupsService_RemoveGroupIdentity_FullMethodName     = "/identity.platform.api.groups.GroupsService/RemoveGroupIdentity"
)

// GroupsServiceClient is the client API for GroupsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GroupsServiceClient interface {
	ListGroups(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListGroupsResp, error)
	GetGroup(ctx context.Context, in *GetGroupReq, opts ...grpc.CallOption) (*GetGroupResp, error)
	CreateGroup(ctx context.Context, in *CreateGroupReq, opts ...grpc.CallOption) (*CreateGroupResp, error)
	UpdateGroup(ctx context.Context, in *UpdateGroupReq, opts ...grpc.CallOption) (*UpdateGroupResp, error)
	RemoveGroup(ctx context.Context, in *RemoveGroupReq, opts ...grpc.CallOption) (*RemoveGroupResp, error)
	ListGroupEntitlements(ctx context.Context, in *ListGroupEntitlementsReq, opts ...grpc.CallOption) (*ListGroupEntitlementsResp, error)
	UpdateGroupEntitlements(ctx context.Context, in *UpdateGroupEntitlementsReq, opts ...grpc.CallOption) (*UpdateGroupEntitlementsResp, error)
	RemoveGroupEntitlement(ctx context.Context, in *RemoveGroupEntitlementReq, opts ...grpc.CallOption) (*RemoveGroupEntitlementResp, error)
	GetGroupRoles(ctx context.Context, in *GetGroupRolesReq, opts ...grpc.CallOption) (*GetGroupRolesResp, error)
	UpdateGroupRoles(ctx context.Context, in *UpdateGroupRolesReq, opts ...grpc.CallOption) (*UpdateGroupRolesResp, error)
	RemoveGroupRole(ctx context.Context, in *RemoveGroupRoleReq, opts ...grpc.CallOption) (*RemoveGroupRoleResp, error)
	GetGroupIdentities(ctx context.Context, in *GetGroupIdentitiesReq, opts ...grpc.CallOption) (*GetGroupIdentitiesResp, error)
	UpdateGroupIdentities(ctx context.Context, in *UpdateGroupIdentitiesReq, opts ...grpc.CallOption) (*UpdateGroupIdentitiesResp, error)
	RemoveGroupIdentity(ctx context.Context, in *RemoveGroupIdentityReq, opts ...grpc.CallOption) (*RemoveGroupIdentityResp, error)
}

type groupsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGroupsServiceClient(cc grpc.ClientConnInterface) GroupsServiceClient {
	return &groupsServiceClient{cc}
}

func (c *groupsServiceClient) ListGroups(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListGroupsResp, error) {
	out := new(ListGroupsResp)
	err := c.cc.Invoke(ctx, GroupsService_ListGroups_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupsServiceClient) GetGroup(ctx context.Context, in *GetGroupReq, opts ...grpc.CallOption) (*GetGroupResp, error) {
	out := new(GetGroupResp)
	err := c.cc.Invoke(ctx, GroupsService_GetGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupsServiceClient) CreateGroup(ctx context.Context, in *CreateGroupReq, opts ...grpc.CallOption) (*CreateGroupResp, error) {
	out := new(CreateGroupResp)
	err := c.cc.Invoke(ctx, GroupsService_CreateGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupsServiceClient) UpdateGroup(ctx context.Context, in *UpdateGroupReq, opts ...grpc.CallOption) (*UpdateGroupResp, error) {
	out := new(UpdateGroupResp)
	err := c.cc.Invoke(ctx, GroupsService_UpdateGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupsServiceClient) RemoveGroup(ctx context.Context, in *RemoveGroupReq, opts ...grpc.CallOption) (*RemoveGroupResp, error) {
	out := new(RemoveGroupResp)
	err := c.cc.Invoke(ctx, GroupsService_RemoveGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupsServiceClient) ListGroupEntitlements(ctx context.Context, in *ListGroupEntitlementsReq, opts ...grpc.CallOption) (*ListGroupEntitlementsResp, error) {
	out := new(ListGroupEntitlementsResp)
	err := c.cc.Invoke(ctx, GroupsService_ListGroupEntitlements_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupsServiceClient) UpdateGroupEntitlements(ctx context.Context, in *UpdateGroupEntitlementsReq, opts ...grpc.CallOption) (*UpdateGroupEntitlementsResp, error) {
	out := new(UpdateGroupEntitlementsResp)
	err := c.cc.Invoke(ctx, GroupsService_UpdateGroupEntitlements_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupsServiceClient) RemoveGroupEntitlement(ctx context.Context, in *RemoveGroupEntitlementReq, opts ...grpc.CallOption) (*RemoveGroupEntitlementResp, error) {
	out := new(RemoveGroupEntitlementResp)
	err := c.cc.Invoke(ctx, GroupsService_RemoveGroupEntitlement_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupsServiceClient) GetGroupRoles(ctx context.Context, in *GetGroupRolesReq, opts ...grpc.CallOption) (*GetGroupRolesResp, error) {
	out := new(GetGroupRolesResp)
	err := c.cc.Invoke(ctx, GroupsService_GetGroupRoles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupsServiceClient) UpdateGroupRoles(ctx context.Context, in *UpdateGroupRolesReq, opts ...grpc.CallOption) (*UpdateGroupRolesResp, error) {
	out := new(UpdateGroupRolesResp)
	err := c.cc.Invoke(ctx, GroupsService_UpdateGroupRoles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupsServiceClient) RemoveGroupRole(ctx context.Context, in *RemoveGroupRoleReq, opts ...grpc.CallOption) (*RemoveGroupRoleResp, error) {
	out := new(RemoveGroupRoleResp)
	err := c.cc.Invoke(ctx, GroupsService_RemoveGroupRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupsServiceClient) GetGroupIdentities(ctx context.Context, in *GetGroupIdentitiesReq, opts ...grpc.CallOption) (*GetGroupIdentitiesResp, error) {
	out := new(GetGroupIdentitiesResp)
	err := c.cc.Invoke(ctx, GroupsService_GetGroupIdentities_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupsServiceClient) UpdateGroupIdentities(ctx context.Context, in *UpdateGroupIdentitiesReq, opts ...grpc.CallOption) (*UpdateGroupIdentitiesResp, error) {
	out := new(UpdateGroupIdentitiesResp)
	err := c.cc.Invoke(ctx, GroupsService_UpdateGroupIdentities_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupsServiceClient) RemoveGroupIdentity(ctx context.Context, in *RemoveGroupIdentityReq, opts ...grpc.CallOption) (*RemoveGroupIdentityResp, error) {
	out := new(RemoveGroupIdentityResp)
	err := c.cc.Invoke(ctx, GroupsService_RemoveGroupIdentity_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GroupsServiceServer is the server API for GroupsService service.
// All implementations must embed UnimplementedGroupsServiceServer
// for forward compatibility
type GroupsServiceServer interface {
	ListGroups(context.Context, *emptypb.Empty) (*ListGroupsResp, error)
	GetGroup(context.Context, *GetGroupReq) (*GetGroupResp, error)
	CreateGroup(context.Context, *CreateGroupReq) (*CreateGroupResp, error)
	UpdateGroup(context.Context, *UpdateGroupReq) (*UpdateGroupResp, error)
	RemoveGroup(context.Context, *RemoveGroupReq) (*RemoveGroupResp, error)
	ListGroupEntitlements(context.Context, *ListGroupEntitlementsReq) (*ListGroupEntitlementsResp, error)
	UpdateGroupEntitlements(context.Context, *UpdateGroupEntitlementsReq) (*UpdateGroupEntitlementsResp, error)
	RemoveGroupEntitlement(context.Context, *RemoveGroupEntitlementReq) (*RemoveGroupEntitlementResp, error)
	GetGroupRoles(context.Context, *GetGroupRolesReq) (*GetGroupRolesResp, error)
	UpdateGroupRoles(context.Context, *UpdateGroupRolesReq) (*UpdateGroupRolesResp, error)
	RemoveGroupRole(context.Context, *RemoveGroupRoleReq) (*RemoveGroupRoleResp, error)
	GetGroupIdentities(context.Context, *GetGroupIdentitiesReq) (*GetGroupIdentitiesResp, error)
	UpdateGroupIdentities(context.Context, *UpdateGroupIdentitiesReq) (*UpdateGroupIdentitiesResp, error)
	RemoveGroupIdentity(context.Context, *RemoveGroupIdentityReq) (*RemoveGroupIdentityResp, error)
	mustEmbedUnimplementedGroupsServiceServer()
}

// UnimplementedGroupsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGroupsServiceServer struct {
}

func (UnimplementedGroupsServiceServer) ListGroups(context.Context, *emptypb.Empty) (*ListGroupsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGroups not implemented")
}
func (UnimplementedGroupsServiceServer) GetGroup(context.Context, *GetGroupReq) (*GetGroupResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroup not implemented")
}
func (UnimplementedGroupsServiceServer) CreateGroup(context.Context, *CreateGroupReq) (*CreateGroupResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGroup not implemented")
}
func (UnimplementedGroupsServiceServer) UpdateGroup(context.Context, *UpdateGroupReq) (*UpdateGroupResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGroup not implemented")
}
func (UnimplementedGroupsServiceServer) RemoveGroup(context.Context, *RemoveGroupReq) (*RemoveGroupResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveGroup not implemented")
}
func (UnimplementedGroupsServiceServer) ListGroupEntitlements(context.Context, *ListGroupEntitlementsReq) (*ListGroupEntitlementsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGroupEntitlements not implemented")
}
func (UnimplementedGroupsServiceServer) UpdateGroupEntitlements(context.Context, *UpdateGroupEntitlementsReq) (*UpdateGroupEntitlementsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGroupEntitlements not implemented")
}
func (UnimplementedGroupsServiceServer) RemoveGroupEntitlement(context.Context, *RemoveGroupEntitlementReq) (*RemoveGroupEntitlementResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveGroupEntitlement not implemented")
}
func (UnimplementedGroupsServiceServer) GetGroupRoles(context.Context, *GetGroupRolesReq) (*GetGroupRolesResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroupRoles not implemented")
}
func (UnimplementedGroupsServiceServer) UpdateGroupRoles(context.Context, *UpdateGroupRolesReq) (*UpdateGroupRolesResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGroupRoles not implemented")
}
func (UnimplementedGroupsServiceServer) RemoveGroupRole(context.Context, *RemoveGroupRoleReq) (*RemoveGroupRoleResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveGroupRole not implemented")
}
func (UnimplementedGroupsServiceServer) GetGroupIdentities(context.Context, *GetGroupIdentitiesReq) (*GetGroupIdentitiesResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroupIdentities not implemented")
}
func (UnimplementedGroupsServiceServer) UpdateGroupIdentities(context.Context, *UpdateGroupIdentitiesReq) (*UpdateGroupIdentitiesResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGroupIdentities not implemented")
}
func (UnimplementedGroupsServiceServer) RemoveGroupIdentity(context.Context, *RemoveGroupIdentityReq) (*RemoveGroupIdentityResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveGroupIdentity not implemented")
}
func (UnimplementedGroupsServiceServer) mustEmbedUnimplementedGroupsServiceServer() {}

// UnsafeGroupsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GroupsServiceServer will
// result in compilation errors.
type UnsafeGroupsServiceServer interface {
	mustEmbedUnimplementedGroupsServiceServer()
}

func RegisterGroupsServiceServer(s grpc.ServiceRegistrar, srv GroupsServiceServer) {
	s.RegisterService(&GroupsService_ServiceDesc, srv)
}

func _GroupsService_ListGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupsServiceServer).ListGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupsService_ListGroups_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupsServiceServer).ListGroups(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupsService_GetGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupsServiceServer).GetGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupsService_GetGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupsServiceServer).GetGroup(ctx, req.(*GetGroupReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupsService_CreateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGroupReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupsServiceServer).CreateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupsService_CreateGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupsServiceServer).CreateGroup(ctx, req.(*CreateGroupReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupsService_UpdateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGroupReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupsServiceServer).UpdateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupsService_UpdateGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupsServiceServer).UpdateGroup(ctx, req.(*UpdateGroupReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupsService_RemoveGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveGroupReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupsServiceServer).RemoveGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupsService_RemoveGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupsServiceServer).RemoveGroup(ctx, req.(*RemoveGroupReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupsService_ListGroupEntitlements_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGroupEntitlementsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupsServiceServer).ListGroupEntitlements(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupsService_ListGroupEntitlements_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupsServiceServer).ListGroupEntitlements(ctx, req.(*ListGroupEntitlementsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupsService_UpdateGroupEntitlements_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGroupEntitlementsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupsServiceServer).UpdateGroupEntitlements(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupsService_UpdateGroupEntitlements_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupsServiceServer).UpdateGroupEntitlements(ctx, req.(*UpdateGroupEntitlementsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupsService_RemoveGroupEntitlement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveGroupEntitlementReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupsServiceServer).RemoveGroupEntitlement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupsService_RemoveGroupEntitlement_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupsServiceServer).RemoveGroupEntitlement(ctx, req.(*RemoveGroupEntitlementReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupsService_GetGroupRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupRolesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupsServiceServer).GetGroupRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupsService_GetGroupRoles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupsServiceServer).GetGroupRoles(ctx, req.(*GetGroupRolesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupsService_UpdateGroupRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGroupRolesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupsServiceServer).UpdateGroupRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupsService_UpdateGroupRoles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupsServiceServer).UpdateGroupRoles(ctx, req.(*UpdateGroupRolesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupsService_RemoveGroupRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveGroupRoleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupsServiceServer).RemoveGroupRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupsService_RemoveGroupRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupsServiceServer).RemoveGroupRole(ctx, req.(*RemoveGroupRoleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupsService_GetGroupIdentities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupIdentitiesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupsServiceServer).GetGroupIdentities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupsService_GetGroupIdentities_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupsServiceServer).GetGroupIdentities(ctx, req.(*GetGroupIdentitiesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupsService_UpdateGroupIdentities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGroupIdentitiesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupsServiceServer).UpdateGroupIdentities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupsService_UpdateGroupIdentities_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupsServiceServer).UpdateGroupIdentities(ctx, req.(*UpdateGroupIdentitiesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupsService_RemoveGroupIdentity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveGroupIdentityReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupsServiceServer).RemoveGroupIdentity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupsService_RemoveGroupIdentity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupsServiceServer).RemoveGroupIdentity(ctx, req.(*RemoveGroupIdentityReq))
	}
	return interceptor(ctx, in, info, handler)
}

// GroupsService_ServiceDesc is the grpc.ServiceDesc for GroupsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GroupsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "identity.platform.api.groups.GroupsService",
	HandlerType: (*GroupsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListGroups",
			Handler:    _GroupsService_ListGroups_Handler,
		},
		{
			MethodName: "GetGroup",
			Handler:    _GroupsService_GetGroup_Handler,
		},
		{
			MethodName: "CreateGroup",
			Handler:    _GroupsService_CreateGroup_Handler,
		},
		{
			MethodName: "UpdateGroup",
			Handler:    _GroupsService_UpdateGroup_Handler,
		},
		{
			MethodName: "RemoveGroup",
			Handler:    _GroupsService_RemoveGroup_Handler,
		},
		{
			MethodName: "ListGroupEntitlements",
			Handler:    _GroupsService_ListGroupEntitlements_Handler,
		},
		{
			MethodName: "UpdateGroupEntitlements",
			Handler:    _GroupsService_UpdateGroupEntitlements_Handler,
		},
		{
			MethodName: "RemoveGroupEntitlement",
			Handler:    _GroupsService_RemoveGroupEntitlement_Handler,
		},
		{
			MethodName: "GetGroupRoles",
			Handler:    _GroupsService_GetGroupRoles_Handler,
		},
		{
			MethodName: "UpdateGroupRoles",
			Handler:    _GroupsService_UpdateGroupRoles_Handler,
		},
		{
			MethodName: "RemoveGroupRole",
			Handler:    _GroupsService_RemoveGroupRole_Handler,
		},
		{
			MethodName: "GetGroupIdentities",
			Handler:    _GroupsService_GetGroupIdentities_Handler,
		},
		{
			MethodName: "UpdateGroupIdentities",
			Handler:    _GroupsService_UpdateGroupIdentities_Handler,
		},
		{
			MethodName: "RemoveGroupIdentity",
			Handler:    _GroupsService_RemoveGroupIdentity_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v0/groups/service.proto",
}
