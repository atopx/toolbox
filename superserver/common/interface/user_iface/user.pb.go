// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.2
// source: data/protos/user.proto

package user_iface

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

type UserStatus int32

const (
	UserStatus_USER_NORMAL          UserStatus = 0 // 正常
	UserStatus_USER_STATUS_DISABLED UserStatus = 1 // 禁用
	UserStatus_USER_STATUS_BLACK    UserStatus = 2 // 黑名单
)

// Enum value maps for UserStatus.
var (
	UserStatus_name = map[int32]string{
		0: "USER_NORMAL",
		1: "USER_STATUS_DISABLED",
		2: "USER_STATUS_BLACK",
	}
	UserStatus_value = map[string]int32{
		"USER_NORMAL":          0,
		"USER_STATUS_DISABLED": 1,
		"USER_STATUS_BLACK":    2,
	}
)

func (x UserStatus) Enum() *UserStatus {
	p := new(UserStatus)
	*p = x
	return p
}

func (x UserStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_data_protos_user_proto_enumTypes[0].Descriptor()
}

func (UserStatus) Type() protoreflect.EnumType {
	return &file_data_protos_user_proto_enumTypes[0]
}

func (x UserStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserStatus.Descriptor instead.
func (UserStatus) EnumDescriptor() ([]byte, []int) {
	return file_data_protos_user_proto_rawDescGZIP(), []int{0}
}

type UserTokenStatus int32

const (
	UserTokenStatus_USER_TOKEN_NORMAL   UserTokenStatus = 0 //  正常
	UserTokenStatus_USER_TOKEN_DISABLED UserTokenStatus = 1 // 已禁用
	UserTokenStatus_USER_TOKEN_EXPIRED  UserTokenStatus = 2 // 已过期
)

// Enum value maps for UserTokenStatus.
var (
	UserTokenStatus_name = map[int32]string{
		0: "USER_TOKEN_NORMAL",
		1: "USER_TOKEN_DISABLED",
		2: "USER_TOKEN_EXPIRED",
	}
	UserTokenStatus_value = map[string]int32{
		"USER_TOKEN_NORMAL":   0,
		"USER_TOKEN_DISABLED": 1,
		"USER_TOKEN_EXPIRED":  2,
	}
)

func (x UserTokenStatus) Enum() *UserTokenStatus {
	p := new(UserTokenStatus)
	*p = x
	return p
}

func (x UserTokenStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserTokenStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_data_protos_user_proto_enumTypes[1].Descriptor()
}

func (UserTokenStatus) Type() protoreflect.EnumType {
	return &file_data_protos_user_proto_enumTypes[1]
}

func (x UserTokenStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserTokenStatus.Descriptor instead.
func (UserTokenStatus) EnumDescriptor() ([]byte, []int) {
	return file_data_protos_user_proto_rawDescGZIP(), []int{1}
}

type RoleNature int32

const (
	RoleNature_ROLE_DEFAULT RoleNature = 0 // 初始化角色
	RoleNature_ROLE_SYSTEM  RoleNature = 1 // 系统角色
	RoleNature_ROLE_CUSTOM  RoleNature = 2 // 自定义角色
)

// Enum value maps for RoleNature.
var (
	RoleNature_name = map[int32]string{
		0: "ROLE_DEFAULT",
		1: "ROLE_SYSTEM",
		2: "ROLE_CUSTOM",
	}
	RoleNature_value = map[string]int32{
		"ROLE_DEFAULT": 0,
		"ROLE_SYSTEM":  1,
		"ROLE_CUSTOM":  2,
	}
)

func (x RoleNature) Enum() *RoleNature {
	p := new(RoleNature)
	*p = x
	return p
}

func (x RoleNature) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RoleNature) Descriptor() protoreflect.EnumDescriptor {
	return file_data_protos_user_proto_enumTypes[2].Descriptor()
}

func (RoleNature) Type() protoreflect.EnumType {
	return &file_data_protos_user_proto_enumTypes[2]
}

func (x RoleNature) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RoleNature.Descriptor instead.
func (RoleNature) EnumDescriptor() ([]byte, []int) {
	return file_data_protos_user_proto_rawDescGZIP(), []int{2}
}

var File_data_protos_user_proto protoreflect.FileDescriptor

var file_data_protos_user_proto_rawDesc = []byte{
	0x0a, 0x16, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x4e, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4e,
	0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x55, 0x53, 0x45, 0x52, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10,
	0x01, 0x12, 0x15, 0x0a, 0x11, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x42, 0x4c, 0x41, 0x43, 0x4b, 0x10, 0x02, 0x2a, 0x59, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x15, 0x0a, 0x11, 0x55,
	0x53, 0x45, 0x52, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c,
	0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e,
	0x5f, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x55,
	0x53, 0x45, 0x52, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x45, 0x58, 0x50, 0x49, 0x52, 0x45,
	0x44, 0x10, 0x02, 0x2a, 0x40, 0x0a, 0x0a, 0x52, 0x6f, 0x6c, 0x65, 0x4e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c,
	0x54, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x53, 0x59, 0x53, 0x54,
	0x45, 0x4d, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x43, 0x55, 0x53,
	0x54, 0x4f, 0x4d, 0x10, 0x02, 0x42, 0x1d, 0x5a, 0x1b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x66, 0x61, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_data_protos_user_proto_rawDescOnce sync.Once
	file_data_protos_user_proto_rawDescData = file_data_protos_user_proto_rawDesc
)

func file_data_protos_user_proto_rawDescGZIP() []byte {
	file_data_protos_user_proto_rawDescOnce.Do(func() {
		file_data_protos_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_data_protos_user_proto_rawDescData)
	})
	return file_data_protos_user_proto_rawDescData
}

var file_data_protos_user_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_data_protos_user_proto_goTypes = []interface{}{
	(UserStatus)(0),      // 0: UserStatus
	(UserTokenStatus)(0), // 1: UserTokenStatus
	(RoleNature)(0),      // 2: RoleNature
}
var file_data_protos_user_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_data_protos_user_proto_init() }
func file_data_protos_user_proto_init() {
	if File_data_protos_user_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_data_protos_user_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_data_protos_user_proto_goTypes,
		DependencyIndexes: file_data_protos_user_proto_depIdxs,
		EnumInfos:         file_data_protos_user_proto_enumTypes,
	}.Build()
	File_data_protos_user_proto = out.File
	file_data_protos_user_proto_rawDesc = nil
	file_data_protos_user_proto_goTypes = nil
	file_data_protos_user_proto_depIdxs = nil
}