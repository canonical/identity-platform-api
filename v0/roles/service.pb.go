// Copyright 2025 Canonical Ltd.
// SPDX-License-Identifier: AGPL-3.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: v0/roles/service.proto

package roles

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_v0_roles_service_proto protoreflect.FileDescriptor

var file_v0_roles_service_proto_rawDesc = []byte{
	0x0a, 0x16, 0x76, 0x30, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x72, 0x6f, 0x6c, 0x65, 0x73, 0x1a, 0x14, 0x76, 0x30, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67,
	0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xed, 0x1b, 0x0a, 0x0c, 0x52, 0x6f, 0x6c, 0x65, 0x73,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xdc, 0x02, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x2a, 0x2e,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x8a, 0x02, 0x92, 0x41, 0xf1, 0x01,
	0x4a, 0x10, 0x0a, 0x03, 0x32, 0x30, 0x30, 0x12, 0x09, 0x0a, 0x07, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x4a, 0x45, 0x0a, 0x03, 0x34, 0x30, 0x31, 0x12, 0x3e, 0x0a, 0x0c, 0x55, 0x6e, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x12, 0x2e, 0x0a, 0x2c, 0x1a, 0x2a, 0x2e,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4a, 0x42, 0x0a, 0x03, 0x34, 0x30, 0x33,
	0x12, 0x3b, 0x0a, 0x09, 0x46, 0x6f, 0x72, 0x62, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x12, 0x2e, 0x0a,
	0x2c, 0x1a, 0x2a, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4a, 0x52, 0x0a,
	0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x47, 0x0a, 0x15, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x20, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x2e, 0x0a, 0x2c, 0x1a, 0x2a, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x30,
	0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0xe5, 0x02, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x52, 0x6f,
	0x6c, 0x65, 0x12, 0x27, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x28, 0x2e, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x86, 0x02, 0x92, 0x41, 0xe8, 0x01, 0x4a, 0x07, 0x0a, 0x03,
	0x32, 0x30, 0x30, 0x12, 0x00, 0x4a, 0x45, 0x0a, 0x03, 0x34, 0x30, 0x31, 0x12, 0x3e, 0x0a, 0x0c,
	0x55, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x12, 0x2e, 0x0a, 0x2c,
	0x1a, 0x2a, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4a, 0x42, 0x0a, 0x03,
	0x34, 0x30, 0x33, 0x12, 0x3b, 0x0a, 0x09, 0x46, 0x6f, 0x72, 0x62, 0x69, 0x64, 0x64, 0x65, 0x6e,
	0x12, 0x2e, 0x0a, 0x2c, 0x1a, 0x2a, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x4a, 0x52, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x47, 0x0a, 0x15, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x20, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x2e, 0x0a, 0x2c, 0x1a, 0x2a, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x30, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0xef,
	0x02, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x2a, 0x2e,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x2b, 0x2e, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x87, 0x02, 0x92, 0x41, 0xe8, 0x01, 0x4a, 0x07, 0x0a,
	0x03, 0x32, 0x30, 0x31, 0x12, 0x00, 0x4a, 0x45, 0x0a, 0x03, 0x34, 0x30, 0x31, 0x12, 0x3e, 0x0a,
	0x0c, 0x55, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x12, 0x2e, 0x0a,
	0x2c, 0x1a, 0x2a, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4a, 0x42, 0x0a,
	0x03, 0x34, 0x30, 0x33, 0x12, 0x3b, 0x0a, 0x09, 0x46, 0x6f, 0x72, 0x62, 0x69, 0x64, 0x64, 0x65,
	0x6e, 0x12, 0x2e, 0x0a, 0x2c, 0x1a, 0x2a, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x4a, 0x52, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x47, 0x0a, 0x15,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x20,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x2e, 0x0a, 0x2c, 0x1a, 0x2a, 0x2e, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x3a, 0x04, 0x72, 0x6f, 0x6c,
	0x65, 0x22, 0x0d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x30, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x73,
	0x12, 0xb5, 0x03, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x12,
	0x2a, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x2b, 0x2e, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0xcd, 0x02, 0x92, 0x41, 0xa9, 0x02, 0x4a,
	0x45, 0x0a, 0x03, 0x34, 0x30, 0x31, 0x12, 0x3e, 0x0a, 0x0c, 0x55, 0x6e, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x12, 0x2e, 0x0a, 0x2c, 0x1a, 0x2a, 0x2e, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4a, 0x42, 0x0a, 0x03, 0x34, 0x30, 0x33, 0x12, 0x3b, 0x0a,
	0x09, 0x46, 0x6f, 0x72, 0x62, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x12, 0x2e, 0x0a, 0x2c, 0x1a, 0x2a,
	0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4a, 0x48, 0x0a, 0x03, 0x35, 0x30,
	0x31, 0x12, 0x41, 0x0a, 0x0f, 0x4e, 0x6f, 0x74, 0x20, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x65, 0x64, 0x12, 0x2e, 0x0a, 0x2c, 0x1a, 0x2a, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x4a, 0x52, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12,
	0x47, 0x0a, 0x15, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x20, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x20, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x2e, 0x0a, 0x2c, 0x1a, 0x2a, 0x2e, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x3a, 0x04,
	0x72, 0x6f, 0x6c, 0x65, 0x32, 0x12, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x30, 0x2f, 0x72, 0x6f,
	0x6c, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0xee, 0x02, 0x0a, 0x0a, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x2a, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x72, 0x6f, 0x6c, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x6f, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x1a, 0x2b, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x6f, 0x6c, 0x65,
	0x73, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x22, 0x86, 0x02, 0x92, 0x41, 0xe8, 0x01, 0x4a, 0x07, 0x0a, 0x03, 0x32, 0x30, 0x30, 0x12, 0x00,
	0x4a, 0x45, 0x0a, 0x03, 0x34, 0x30, 0x31, 0x12, 0x3e, 0x0a, 0x0c, 0x55, 0x6e, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x12, 0x2e, 0x0a, 0x2c, 0x1a, 0x2a, 0x2e, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4a, 0x42, 0x0a, 0x03, 0x34, 0x30, 0x33, 0x12, 0x3b,
	0x0a, 0x09, 0x46, 0x6f, 0x72, 0x62, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x12, 0x2e, 0x0a, 0x2c, 0x1a,
	0x2a, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4a, 0x52, 0x0a, 0x07, 0x64,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x47, 0x0a, 0x15, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x20, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12,
	0x2e, 0x0a, 0x2c, 0x1a, 0x2a, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x14, 0x2a, 0x12, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x30, 0x2f, 0x72,
	0x6f, 0x6c, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0xd3, 0x03, 0x0a, 0x14, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x12, 0x34, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x73,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x35, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22,
	0xcd, 0x02, 0x92, 0x41, 0xa2, 0x02, 0x4a, 0x07, 0x0a, 0x03, 0x32, 0x30, 0x30, 0x12, 0x00, 0x4a,
	0x45, 0x0a, 0x03, 0x34, 0x30, 0x31, 0x12, 0x3e, 0x0a, 0x0c, 0x55, 0x6e, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x12, 0x2e, 0x0a, 0x2c, 0x1a, 0x2a, 0x2e, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4a, 0x42, 0x0a, 0x03, 0x34, 0x30, 0x33, 0x12, 0x3b, 0x0a,
	0x09, 0x46, 0x6f, 0x72, 0x62, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x12, 0x2e, 0x0a, 0x2c, 0x1a, 0x2a,
	0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4a, 0x52, 0x0a, 0x07, 0x64, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x47, 0x0a, 0x15, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x20, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x2e,
	0x0a, 0x2c, 0x1a, 0x2a, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x72, 0x38,
	0x0a, 0x36, 0x0a, 0x12, 0x58, 0x2d, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x2d, 0x50, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x42, 0x61, 0x73, 0x65, 0x36, 0x34, 0x20, 0x65,
	0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x20, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x20, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x12, 0x1f,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x30, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x2f, 0x7b, 0x69,
	0x64, 0x7d, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12,
	0x94, 0x03, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x36, 0x2e, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x6f, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0xa9, 0x02, 0x92, 0x41, 0xe8,
	0x01, 0x4a, 0x07, 0x0a, 0x03, 0x32, 0x30, 0x30, 0x12, 0x00, 0x4a, 0x45, 0x0a, 0x03, 0x34, 0x30,
	0x31, 0x12, 0x3e, 0x0a, 0x0c, 0x55, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65,
	0x64, 0x12, 0x2e, 0x0a, 0x2c, 0x1a, 0x2a, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x4a, 0x42, 0x0a, 0x03, 0x34, 0x30, 0x33, 0x12, 0x3b, 0x0a, 0x09, 0x46, 0x6f, 0x72, 0x62,
	0x69, 0x64, 0x64, 0x65, 0x6e, 0x12, 0x2e, 0x0a, 0x2c, 0x1a, 0x2a, 0x2e, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4a, 0x52, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x12, 0x47, 0x0a, 0x15, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x20, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x20, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x2e, 0x0a, 0x2c, 0x1a, 0x2a, 0x2e,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x37, 0x3a,
	0x14, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x50, 0x61, 0x74,
	0x63, 0x68, 0x52, 0x65, 0x71, 0x32, 0x1f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x30, 0x2f, 0x72,
	0x6f, 0x6c, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x8c, 0x03, 0x0a, 0x15, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x52, 0x6f, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x35, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x2e, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0xa3, 0x02, 0x92, 0x41, 0xe8, 0x01, 0x4a, 0x07, 0x0a, 0x03, 0x32, 0x30, 0x30, 0x12, 0x00, 0x4a,
	0x45, 0x0a, 0x03, 0x34, 0x30, 0x31, 0x12, 0x3e, 0x0a, 0x0c, 0x55, 0x6e, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x12, 0x2e, 0x0a, 0x2c, 0x1a, 0x2a, 0x2e, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4a, 0x42, 0x0a, 0x03, 0x34, 0x30, 0x33, 0x12, 0x3b, 0x0a,
	0x09, 0x46, 0x6f, 0x72, 0x62, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x12, 0x2e, 0x0a, 0x2c, 0x1a, 0x2a,
	0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4a, 0x52, 0x0a, 0x07, 0x64, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x47, 0x0a, 0x15, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x20, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x2e,
	0x0a, 0x2c, 0x1a, 0x2a, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x31, 0x2a, 0x2f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x30, 0x2f, 0x72, 0x6f,
	0x6c, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x7d, 0x12, 0xfe, 0x02, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x2d, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x72, 0x6f, 0x6c, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x2e, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72,
	0x6f, 0x6c, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x8d, 0x02, 0x92, 0x41, 0xe8, 0x01, 0x4a, 0x07, 0x0a,
	0x03, 0x32, 0x30, 0x30, 0x12, 0x00, 0x4a, 0x45, 0x0a, 0x03, 0x34, 0x30, 0x31, 0x12, 0x3e, 0x0a,
	0x0c, 0x55, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x12, 0x2e, 0x0a,
	0x2c, 0x1a, 0x2a, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4a, 0x42, 0x0a,
	0x03, 0x34, 0x30, 0x33, 0x12, 0x3b, 0x0a, 0x09, 0x46, 0x6f, 0x72, 0x62, 0x69, 0x64, 0x64, 0x65,
	0x6e, 0x12, 0x2e, 0x0a, 0x2c, 0x1a, 0x2a, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x4a, 0x52, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x47, 0x0a, 0x15,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x20,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x2e, 0x0a, 0x2c, 0x1a, 0x2a, 0x2e, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x30, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x2f,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x6e, 0x6f, 0x6e, 0x69, 0x63, 0x61, 0x6c, 0x2f, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2d, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_v0_roles_service_proto_goTypes = []interface{}{
	(*emptypb.Empty)(nil),             // 0: google.protobuf.Empty
	(*GetRoleReq)(nil),                // 1: identity.platform.api.roles.GetRoleReq
	(*CreateRoleReq)(nil),             // 2: identity.platform.api.roles.CreateRoleReq
	(*UpdateRoleReq)(nil),             // 3: identity.platform.api.roles.UpdateRoleReq
	(*RemoveRoleReq)(nil),             // 4: identity.platform.api.roles.RemoveRoleReq
	(*ListRoleEntitlementsReq)(nil),   // 5: identity.platform.api.roles.ListRoleEntitlementsReq
	(*UpdateRoleEntitlementsReq)(nil), // 6: identity.platform.api.roles.UpdateRoleEntitlementsReq
	(*RemoveRoleEntitlementReq)(nil),  // 7: identity.platform.api.roles.RemoveRoleEntitlementReq
	(*GetRoleGroupsReq)(nil),          // 8: identity.platform.api.roles.GetRoleGroupsReq
	(*ListRolesResp)(nil),             // 9: identity.platform.api.roles.ListRolesResp
	(*GetRoleResp)(nil),               // 10: identity.platform.api.roles.GetRoleResp
	(*CreateRoleResp)(nil),            // 11: identity.platform.api.roles.CreateRoleResp
	(*UpdateRoleResp)(nil),            // 12: identity.platform.api.roles.UpdateRoleResp
	(*RemoveRoleResp)(nil),            // 13: identity.platform.api.roles.RemoveRoleResp
	(*ListRoleEntitlementsResp)(nil),  // 14: identity.platform.api.roles.ListRoleEntitlementsResp
	(*GetRoleGroupsResp)(nil),         // 15: identity.platform.api.roles.GetRoleGroupsResp
}
var file_v0_roles_service_proto_depIdxs = []int32{
	0,  // 0: identity.platform.api.roles.RolesService.ListRoles:input_type -> google.protobuf.Empty
	1,  // 1: identity.platform.api.roles.RolesService.GetRole:input_type -> identity.platform.api.roles.GetRoleReq
	2,  // 2: identity.platform.api.roles.RolesService.CreateRole:input_type -> identity.platform.api.roles.CreateRoleReq
	3,  // 3: identity.platform.api.roles.RolesService.UpdateRole:input_type -> identity.platform.api.roles.UpdateRoleReq
	4,  // 4: identity.platform.api.roles.RolesService.RemoveRole:input_type -> identity.platform.api.roles.RemoveRoleReq
	5,  // 5: identity.platform.api.roles.RolesService.ListRoleEntitlements:input_type -> identity.platform.api.roles.ListRoleEntitlementsReq
	6,  // 6: identity.platform.api.roles.RolesService.UpdateRoleEntitlements:input_type -> identity.platform.api.roles.UpdateRoleEntitlementsReq
	7,  // 7: identity.platform.api.roles.RolesService.RemoveRoleEntitlement:input_type -> identity.platform.api.roles.RemoveRoleEntitlementReq
	8,  // 8: identity.platform.api.roles.RolesService.GetRoleGroups:input_type -> identity.platform.api.roles.GetRoleGroupsReq
	9,  // 9: identity.platform.api.roles.RolesService.ListRoles:output_type -> identity.platform.api.roles.ListRolesResp
	10, // 10: identity.platform.api.roles.RolesService.GetRole:output_type -> identity.platform.api.roles.GetRoleResp
	11, // 11: identity.platform.api.roles.RolesService.CreateRole:output_type -> identity.platform.api.roles.CreateRoleResp
	12, // 12: identity.platform.api.roles.RolesService.UpdateRole:output_type -> identity.platform.api.roles.UpdateRoleResp
	13, // 13: identity.platform.api.roles.RolesService.RemoveRole:output_type -> identity.platform.api.roles.RemoveRoleResp
	14, // 14: identity.platform.api.roles.RolesService.ListRoleEntitlements:output_type -> identity.platform.api.roles.ListRoleEntitlementsResp
	0,  // 15: identity.platform.api.roles.RolesService.UpdateRoleEntitlements:output_type -> google.protobuf.Empty
	0,  // 16: identity.platform.api.roles.RolesService.RemoveRoleEntitlement:output_type -> google.protobuf.Empty
	15, // 17: identity.platform.api.roles.RolesService.GetRoleGroups:output_type -> identity.platform.api.roles.GetRoleGroupsResp
	9,  // [9:18] is the sub-list for method output_type
	0,  // [0:9] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_v0_roles_service_proto_init() }
func file_v0_roles_service_proto_init() {
	if File_v0_roles_service_proto != nil {
		return
	}
	file_v0_roles_model_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v0_roles_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v0_roles_service_proto_goTypes,
		DependencyIndexes: file_v0_roles_service_proto_depIdxs,
	}.Build()
	File_v0_roles_service_proto = out.File
	file_v0_roles_service_proto_rawDesc = nil
	file_v0_roles_service_proto_goTypes = nil
	file_v0_roles_service_proto_depIdxs = nil
}
