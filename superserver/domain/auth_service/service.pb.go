// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: auth/service.proto

package auth_service

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_auth_service_proto protoreflect.FileDescriptor

var file_auth_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x1a, 0x0f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x61, 0x75, 0x74, 0x68, 0x2f,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x18, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x72, 0x6f, 0x6c, 0x65,
	0x5f, 0x72, 0x65, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x85, 0x0d, 0x0a, 0x0b, 0x41,
	0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x08, 0x4c, 0x69,
	0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x1a, 0x1b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x4e, 0x0a, 0x0b, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x1f, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x1a, 0x1e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x5d, 0x0a, 0x10, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x24, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x23, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x45, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x1c, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x6f, 0x6c, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x1b, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f,
	0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x4e, 0x0a, 0x0b, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x1f, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c,
	0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x1e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x6f,
	0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x5d, 0x0a, 0x10, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x24, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x1a, 0x23, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c,
	0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x5a, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x66, 0x12, 0x23, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x66, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x22,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x66, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x63, 0x0a, 0x12, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x66, 0x12, 0x26, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x66, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x1a, 0x25, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x52,
	0x65, 0x66, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x72, 0x0a, 0x17, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x52,
	0x65, 0x66, 0x12, 0x2b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x66, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a,
	0x2a, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x6f, 0x6c, 0x65, 0x52, 0x65, 0x66, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x4b, 0x0a, 0x0a, 0x4c,
	0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x1e, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x1d, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x54, 0x0a, 0x0d, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x21, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x20, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x63,
	0x0a, 0x12, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x12, 0x26, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x25, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x57, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x21, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x60, 0x0a, 0x11,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x25, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x24, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x6f,
	0x0a, 0x16, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x1a, 0x29, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x54, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x21, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x1a, 0x20, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x5d, 0x0a, 0x10, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x24, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a,
	0x23, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x6c, 0x0a, 0x15, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x29, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x28, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x42, 0x21, 0x5a, 0x1f, 0x73, 0x75, 0x70, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_auth_service_proto_goTypes = []interface{}{
	(*ListUserParams)(nil),                // 0: auth_service.ListUserParams
	(*OperateUserParams)(nil),             // 1: auth_service.OperateUserParams
	(*BatchOperateUserParams)(nil),        // 2: auth_service.BatchOperateUserParams
	(*ListRoleParams)(nil),                // 3: auth_service.ListRoleParams
	(*OperateRoleParams)(nil),             // 4: auth_service.OperateRoleParams
	(*BatchOperateRoleParams)(nil),        // 5: auth_service.BatchOperateRoleParams
	(*ListUserRoleRefParams)(nil),         // 6: auth_service.ListUserRoleRefParams
	(*OperateUserRoleRefParams)(nil),      // 7: auth_service.OperateUserRoleRefParams
	(*BatchOperateUserRoleRefParams)(nil), // 8: auth_service.BatchOperateUserRoleRefParams
	(*ListAccessParams)(nil),              // 9: auth_service.ListAccessParams
	(*OperateAccessParams)(nil),           // 10: auth_service.OperateAccessParams
	(*BatchOperateAccessParams)(nil),      // 11: auth_service.BatchOperateAccessParams
	(*ListPermissionParams)(nil),          // 12: auth_service.ListPermissionParams
	(*OperatePermissionParams)(nil),       // 13: auth_service.OperatePermissionParams
	(*BatchOperatePermissionParams)(nil),  // 14: auth_service.BatchOperatePermissionParams
	(*ListAuthTokenParams)(nil),           // 15: auth_service.ListAuthTokenParams
	(*OperateAuthTokenParams)(nil),        // 16: auth_service.OperateAuthTokenParams
	(*BatchOperateAuthTokenParams)(nil),   // 17: auth_service.BatchOperateAuthTokenParams
	(*ListUserReply)(nil),                 // 18: auth_service.ListUserReply
	(*OperateUserReply)(nil),              // 19: auth_service.OperateUserReply
	(*BatchOperateUserReply)(nil),         // 20: auth_service.BatchOperateUserReply
	(*ListRoleReply)(nil),                 // 21: auth_service.ListRoleReply
	(*OperateRoleReply)(nil),              // 22: auth_service.OperateRoleReply
	(*BatchOperateRoleReply)(nil),         // 23: auth_service.BatchOperateRoleReply
	(*ListUserRoleRefReply)(nil),          // 24: auth_service.ListUserRoleRefReply
	(*OperateUserRoleRefReply)(nil),       // 25: auth_service.OperateUserRoleRefReply
	(*BatchOperateUserRoleRefReply)(nil),  // 26: auth_service.BatchOperateUserRoleRefReply
	(*ListAccessReply)(nil),               // 27: auth_service.ListAccessReply
	(*OperateAccessReply)(nil),            // 28: auth_service.OperateAccessReply
	(*BatchOperateAccessReply)(nil),       // 29: auth_service.BatchOperateAccessReply
	(*ListPermissionReply)(nil),           // 30: auth_service.ListPermissionReply
	(*OperatePermissionReply)(nil),        // 31: auth_service.OperatePermissionReply
	(*BatchOperatePermissionReply)(nil),   // 32: auth_service.BatchOperatePermissionReply
	(*ListAuthTokenReply)(nil),            // 33: auth_service.ListAuthTokenReply
	(*OperateAuthTokenReply)(nil),         // 34: auth_service.OperateAuthTokenReply
	(*BatchOperateAuthTokenReply)(nil),    // 35: auth_service.BatchOperateAuthTokenReply
}
var file_auth_service_proto_depIdxs = []int32{
	0,  // 0: auth_service.AuthService.ListUser:input_type -> auth_service.ListUserParams
	1,  // 1: auth_service.AuthService.OperateUser:input_type -> auth_service.OperateUserParams
	2,  // 2: auth_service.AuthService.BatchOperateUser:input_type -> auth_service.BatchOperateUserParams
	3,  // 3: auth_service.AuthService.ListRole:input_type -> auth_service.ListRoleParams
	4,  // 4: auth_service.AuthService.OperateRole:input_type -> auth_service.OperateRoleParams
	5,  // 5: auth_service.AuthService.BatchOperateRole:input_type -> auth_service.BatchOperateRoleParams
	6,  // 6: auth_service.AuthService.ListUserRoleRef:input_type -> auth_service.ListUserRoleRefParams
	7,  // 7: auth_service.AuthService.OperateUserRoleRef:input_type -> auth_service.OperateUserRoleRefParams
	8,  // 8: auth_service.AuthService.BatchOperateUserRoleRef:input_type -> auth_service.BatchOperateUserRoleRefParams
	9,  // 9: auth_service.AuthService.ListAccess:input_type -> auth_service.ListAccessParams
	10, // 10: auth_service.AuthService.OperateAccess:input_type -> auth_service.OperateAccessParams
	11, // 11: auth_service.AuthService.BatchOperateAccess:input_type -> auth_service.BatchOperateAccessParams
	12, // 12: auth_service.AuthService.ListPermission:input_type -> auth_service.ListPermissionParams
	13, // 13: auth_service.AuthService.OperatePermission:input_type -> auth_service.OperatePermissionParams
	14, // 14: auth_service.AuthService.BatchOperatePermission:input_type -> auth_service.BatchOperatePermissionParams
	15, // 15: auth_service.AuthService.ListAuthToken:input_type -> auth_service.ListAuthTokenParams
	16, // 16: auth_service.AuthService.OperateAuthToken:input_type -> auth_service.OperateAuthTokenParams
	17, // 17: auth_service.AuthService.BatchOperateAuthToken:input_type -> auth_service.BatchOperateAuthTokenParams
	18, // 18: auth_service.AuthService.ListUser:output_type -> auth_service.ListUserReply
	19, // 19: auth_service.AuthService.OperateUser:output_type -> auth_service.OperateUserReply
	20, // 20: auth_service.AuthService.BatchOperateUser:output_type -> auth_service.BatchOperateUserReply
	21, // 21: auth_service.AuthService.ListRole:output_type -> auth_service.ListRoleReply
	22, // 22: auth_service.AuthService.OperateRole:output_type -> auth_service.OperateRoleReply
	23, // 23: auth_service.AuthService.BatchOperateRole:output_type -> auth_service.BatchOperateRoleReply
	24, // 24: auth_service.AuthService.ListUserRoleRef:output_type -> auth_service.ListUserRoleRefReply
	25, // 25: auth_service.AuthService.OperateUserRoleRef:output_type -> auth_service.OperateUserRoleRefReply
	26, // 26: auth_service.AuthService.BatchOperateUserRoleRef:output_type -> auth_service.BatchOperateUserRoleRefReply
	27, // 27: auth_service.AuthService.ListAccess:output_type -> auth_service.ListAccessReply
	28, // 28: auth_service.AuthService.OperateAccess:output_type -> auth_service.OperateAccessReply
	29, // 29: auth_service.AuthService.BatchOperateAccess:output_type -> auth_service.BatchOperateAccessReply
	30, // 30: auth_service.AuthService.ListPermission:output_type -> auth_service.ListPermissionReply
	31, // 31: auth_service.AuthService.OperatePermission:output_type -> auth_service.OperatePermissionReply
	32, // 32: auth_service.AuthService.BatchOperatePermission:output_type -> auth_service.BatchOperatePermissionReply
	33, // 33: auth_service.AuthService.ListAuthToken:output_type -> auth_service.ListAuthTokenReply
	34, // 34: auth_service.AuthService.OperateAuthToken:output_type -> auth_service.OperateAuthTokenReply
	35, // 35: auth_service.AuthService.BatchOperateAuthToken:output_type -> auth_service.BatchOperateAuthTokenReply
	18, // [18:36] is the sub-list for method output_type
	0,  // [0:18] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_auth_service_proto_init() }
func file_auth_service_proto_init() {
	if File_auth_service_proto != nil {
		return
	}
	file_auth_user_proto_init()
	file_auth_role_proto_init()
	file_auth_access_proto_init()
	file_auth_token_proto_init()
	file_auth_permission_proto_init()
	file_auth_user_role_ref_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_auth_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_service_proto_goTypes,
		DependencyIndexes: file_auth_service_proto_depIdxs,
	}.Build()
	File_auth_service_proto = out.File
	file_auth_service_proto_rawDesc = nil
	file_auth_service_proto_goTypes = nil
	file_auth_service_proto_depIdxs = nil
}
