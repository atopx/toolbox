// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: proto/defines/ecode.proto

package ecode_iface

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ECode int32

const (
	ECode_NOERROR          ECode = 0     // 默认无错误
	ECode_AUTH_NOTIN       ECode = 40100 // 账户未登录
	ECode_AUTH_INVALID     ECode = 40102 // 无效Token
	ECode_AUTH_EXPIRED     ECode = 40103 // 账户已过期
	ECode_AUTH_LIMIT       ECode = 40104 // 限制登录
	ECode_ACCESS_FORBIDDEN ECode = 40300 // 禁止访问
)

// Enum value maps for ECode.
var (
	ECode_name = map[int32]string{
		0:     "NOERROR",
		40100: "AUTH_NOTIN",
		40102: "AUTH_INVALID",
		40103: "AUTH_EXPIRED",
		40104: "AUTH_LIMIT",
		40300: "ACCESS_FORBIDDEN",
	}
	ECode_value = map[string]int32{
		"NOERROR":          0,
		"AUTH_NOTIN":       40100,
		"AUTH_INVALID":     40102,
		"AUTH_EXPIRED":     40103,
		"AUTH_LIMIT":       40104,
		"ACCESS_FORBIDDEN": 40300,
	}
)

func (x ECode) Enum() *ECode {
	p := new(ECode)
	*p = x
	return p
}

func (x ECode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ECode) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_defines_ecode_proto_enumTypes[0].Descriptor()
}

func (ECode) Type() protoreflect.EnumType {
	return &file_proto_defines_ecode_proto_enumTypes[0]
}

func (x ECode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ECode.Descriptor instead.
func (ECode) EnumDescriptor() ([]byte, []int) {
	return file_proto_defines_ecode_proto_rawDescGZIP(), []int{0}
}

var File_proto_defines_ecode_proto protoreflect.FileDescriptor

var file_proto_defines_ecode_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x73, 0x2f,
	0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x78, 0x0a, 0x05, 0x45,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x4e, 0x4f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10,
	0x00, 0x12, 0x10, 0x0a, 0x0a, 0x41, 0x55, 0x54, 0x48, 0x5f, 0x4e, 0x4f, 0x54, 0x49, 0x4e, 0x10,
	0xa4, 0xb9, 0x02, 0x12, 0x12, 0x0a, 0x0c, 0x41, 0x55, 0x54, 0x48, 0x5f, 0x49, 0x4e, 0x56, 0x41,
	0x4c, 0x49, 0x44, 0x10, 0xa6, 0xb9, 0x02, 0x12, 0x12, 0x0a, 0x0c, 0x41, 0x55, 0x54, 0x48, 0x5f,
	0x45, 0x58, 0x50, 0x49, 0x52, 0x45, 0x44, 0x10, 0xa7, 0xb9, 0x02, 0x12, 0x10, 0x0a, 0x0a, 0x41,
	0x55, 0x54, 0x48, 0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x10, 0xa8, 0xb9, 0x02, 0x12, 0x16, 0x0a,
	0x10, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x46, 0x4f, 0x52, 0x42, 0x49, 0x44, 0x44, 0x45,
	0x4e, 0x10, 0xec, 0xba, 0x02, 0x42, 0x13, 0x5a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65,
	0x63, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x66, 0x61, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_defines_ecode_proto_rawDescOnce sync.Once
	file_proto_defines_ecode_proto_rawDescData = file_proto_defines_ecode_proto_rawDesc
)

func file_proto_defines_ecode_proto_rawDescGZIP() []byte {
	file_proto_defines_ecode_proto_rawDescOnce.Do(func() {
		file_proto_defines_ecode_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_defines_ecode_proto_rawDescData)
	})
	return file_proto_defines_ecode_proto_rawDescData
}

var file_proto_defines_ecode_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_defines_ecode_proto_goTypes = []interface{}{
	(ECode)(0), // 0: ECode
}
var file_proto_defines_ecode_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_defines_ecode_proto_init() }
func file_proto_defines_ecode_proto_init() {
	if File_proto_defines_ecode_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_defines_ecode_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_defines_ecode_proto_goTypes,
		DependencyIndexes: file_proto_defines_ecode_proto_depIdxs,
		EnumInfos:         file_proto_defines_ecode_proto_enumTypes,
	}.Build()
	File_proto_defines_ecode_proto = out.File
	file_proto_defines_ecode_proto_rawDesc = nil
	file_proto_defines_ecode_proto_goTypes = nil
	file_proto_defines_ecode_proto_depIdxs = nil
}
