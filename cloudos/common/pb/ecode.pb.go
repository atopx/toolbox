// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.24.2
// source: ecode.proto

package pb

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
	ECode_SUCCESS ECode = 0 // 请求成功
	// 客户端错误
	ECode_BadRequest ECode = 40000000
	// 无效参数
	ECode_InvalidUsername ECode = 40000001 // 无效用户名
	ECode_InvalidPassword ECode = 40000002 // 无效密码
	ECode_InvalidName     ECode = 40000003 // 无效的名称
	ECode_InvalidDomain   ECode = 40000004 // 无效的作用域
	// 资源冲突
	ECode_AlreadyExistName ECode = 40001001 // 名称已存在
	// auth
	ECode_AuthTokenInvalid ECode = 40010001 // 无效Token
	ECode_AuthTokenExpired ECode = 40010002 // Token过期
	ECode_UserDisabled     ECode = 40010003 // 用户被禁用
	// 禁止访问
	ECode_Forbidden ECode = 40030000 // 禁止访问
	// not found
	ECode_RecordNotFound ECode = 40040000 // 请求资源不存在
	ECode_NotFoundUser   ECode = 40040001 // 用户不存在
	ECode_NotFoundToken  ECode = 40040002 // Token不存在
	ECode_NotFoundId     ECode = 40040003 // ID不存在
	// 服务端错误
	ECode_ServerInternalError ECode = 50000000
	ECode_Unimplemented       ECode = 50000001
)

// Enum value maps for ECode.
var (
	ECode_name = map[int32]string{
		0:        "SUCCESS",
		40000000: "BadRequest",
		40000001: "InvalidUsername",
		40000002: "InvalidPassword",
		40000003: "InvalidName",
		40000004: "InvalidDomain",
		40001001: "AlreadyExistName",
		40010001: "AuthTokenInvalid",
		40010002: "AuthTokenExpired",
		40010003: "UserDisabled",
		40030000: "Forbidden",
		40040000: "RecordNotFound",
		40040001: "NotFoundUser",
		40040002: "NotFoundToken",
		40040003: "NotFoundId",
		50000000: "ServerInternalError",
		50000001: "Unimplemented",
	}
	ECode_value = map[string]int32{
		"SUCCESS":             0,
		"BadRequest":          40000000,
		"InvalidUsername":     40000001,
		"InvalidPassword":     40000002,
		"InvalidName":         40000003,
		"InvalidDomain":       40000004,
		"AlreadyExistName":    40001001,
		"AuthTokenInvalid":    40010001,
		"AuthTokenExpired":    40010002,
		"UserDisabled":        40010003,
		"Forbidden":           40030000,
		"RecordNotFound":      40040000,
		"NotFoundUser":        40040001,
		"NotFoundToken":       40040002,
		"NotFoundId":          40040003,
		"ServerInternalError": 50000000,
		"Unimplemented":       50000001,
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
	return file_ecode_proto_enumTypes[0].Descriptor()
}

func (ECode) Type() protoreflect.EnumType {
	return &file_ecode_proto_enumTypes[0]
}

func (x ECode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ECode.Descriptor instead.
func (ECode) EnumDescriptor() ([]byte, []int) {
	return file_ecode_proto_rawDescGZIP(), []int{0}
}

var File_ecode_proto protoreflect.FileDescriptor

var file_ecode_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x65,
	0x63, 0x6f, 0x64, 0x65, 0x2a, 0xfa, 0x02, 0x0a, 0x05, 0x45, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0b,
	0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0a, 0x42,
	0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x10, 0x80, 0xb4, 0x89, 0x13, 0x12, 0x16,
	0x0a, 0x0f, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x10, 0x81, 0xb4, 0x89, 0x13, 0x12, 0x16, 0x0a, 0x0f, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x10, 0x82, 0xb4, 0x89, 0x13, 0x12, 0x12,
	0x0a, 0x0b, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x10, 0x83, 0xb4,
	0x89, 0x13, 0x12, 0x14, 0x0a, 0x0d, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x44, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x10, 0x84, 0xb4, 0x89, 0x13, 0x12, 0x17, 0x0a, 0x10, 0x41, 0x6c, 0x72, 0x65,
	0x61, 0x64, 0x79, 0x45, 0x78, 0x69, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x10, 0xe9, 0xbb, 0x89,
	0x13, 0x12, 0x17, 0x0a, 0x10, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x10, 0x91, 0x82, 0x8a, 0x13, 0x12, 0x17, 0x0a, 0x10, 0x41, 0x75,
	0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x10, 0x92,
	0x82, 0x8a, 0x13, 0x12, 0x13, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x44, 0x69, 0x73, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x10, 0x93, 0x82, 0x8a, 0x13, 0x12, 0x10, 0x0a, 0x09, 0x46, 0x6f, 0x72, 0x62,
	0x69, 0x64, 0x64, 0x65, 0x6e, 0x10, 0xb0, 0x9e, 0x8b, 0x13, 0x12, 0x15, 0x0a, 0x0e, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x4e, 0x6f, 0x74, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x10, 0xc0, 0xec, 0x8b,
	0x13, 0x12, 0x13, 0x0a, 0x0c, 0x4e, 0x6f, 0x74, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x55, 0x73, 0x65,
	0x72, 0x10, 0xc1, 0xec, 0x8b, 0x13, 0x12, 0x14, 0x0a, 0x0d, 0x4e, 0x6f, 0x74, 0x46, 0x6f, 0x75,
	0x6e, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x10, 0xc2, 0xec, 0x8b, 0x13, 0x12, 0x11, 0x0a, 0x0a,
	0x4e, 0x6f, 0x74, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x64, 0x10, 0xc3, 0xec, 0x8b, 0x13, 0x12,
	0x1a, 0x0a, 0x13, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x80, 0xe1, 0xeb, 0x17, 0x12, 0x14, 0x0a, 0x0d, 0x55,
	0x6e, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x65, 0x64, 0x10, 0x81, 0xe1, 0xeb,
	0x17, 0x42, 0x13, 0x5a, 0x11, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ecode_proto_rawDescOnce sync.Once
	file_ecode_proto_rawDescData = file_ecode_proto_rawDesc
)

func file_ecode_proto_rawDescGZIP() []byte {
	file_ecode_proto_rawDescOnce.Do(func() {
		file_ecode_proto_rawDescData = protoimpl.X.CompressGZIP(file_ecode_proto_rawDescData)
	})
	return file_ecode_proto_rawDescData
}

var file_ecode_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ecode_proto_goTypes = []interface{}{
	(ECode)(0), // 0: ecode.ECode
}
var file_ecode_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ecode_proto_init() }
func file_ecode_proto_init() {
	if File_ecode_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ecode_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ecode_proto_goTypes,
		DependencyIndexes: file_ecode_proto_depIdxs,
		EnumInfos:         file_ecode_proto_enumTypes,
	}.Build()
	File_ecode_proto = out.File
	file_ecode_proto_rawDesc = nil
	file_ecode_proto_goTypes = nil
	file_ecode_proto_depIdxs = nil
}
