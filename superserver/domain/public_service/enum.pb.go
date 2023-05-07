// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.2
// source: public/enum.proto

package public_service

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	common "superserver/domain/public/common"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Enum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Value int32  `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	Desc  string `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`
}

func (x *Enum) Reset() {
	*x = Enum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_enum_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Enum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Enum) ProtoMessage() {}

func (x *Enum) ProtoReflect() protoreflect.Message {
	mi := &file_public_enum_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Enum.ProtoReflect.Descriptor instead.
func (*Enum) Descriptor() ([]byte, []int) {
	return file_public_enum_proto_rawDescGZIP(), []int{0}
}

func (x *Enum) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Enum) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *Enum) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

type Enums struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key  string  `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Data []*Enum `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *Enums) Reset() {
	*x = Enums{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_enum_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Enums) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Enums) ProtoMessage() {}

func (x *Enums) ProtoReflect() protoreflect.Message {
	mi := &file_public_enum_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Enums.ProtoReflect.Descriptor instead.
func (*Enums) Descriptor() ([]byte, []int) {
	return file_public_enum_proto_rawDescGZIP(), []int{1}
}

func (x *Enums) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Enums) GetData() []*Enum {
	if x != nil {
		return x.Data
	}
	return nil
}

type ListEnumParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *common.Header `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
}

func (x *ListEnumParams) Reset() {
	*x = ListEnumParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_enum_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEnumParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEnumParams) ProtoMessage() {}

func (x *ListEnumParams) ProtoReflect() protoreflect.Message {
	mi := &file_public_enum_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEnumParams.ProtoReflect.Descriptor instead.
func (*ListEnumParams) Descriptor() ([]byte, []int) {
	return file_public_enum_proto_rawDescGZIP(), []int{2}
}

func (x *ListEnumParams) GetHeader() *common.Header {
	if x != nil {
		return x.Header
	}
	return nil
}

type ListEnumReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *common.ReplyHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Data   []*Enums            `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ListEnumReply) Reset() {
	*x = ListEnumReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_enum_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEnumReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEnumReply) ProtoMessage() {}

func (x *ListEnumReply) ProtoReflect() protoreflect.Message {
	mi := &file_public_enum_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEnumReply.ProtoReflect.Descriptor instead.
func (*ListEnumReply) Descriptor() ([]byte, []int) {
	return file_public_enum_proto_rawDescGZIP(), []int{3}
}

func (x *ListEnumReply) GetHeader() *common.ReplyHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *ListEnumReply) GetData() []*Enums {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_public_enum_proto protoreflect.FileDescriptor

var file_public_enum_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x1a, 0x13, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x44, 0x0a, 0x04, 0x45, 0x6e, 0x75, 0x6d,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65,
	0x73, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x22, 0x43,
	0x0a, 0x05, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x38, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x75, 0x6d, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x26, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x22, 0x67, 0x0a,
	0x0d, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2b,
	0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x23, 0x5a, 0x21, 0x73, 0x75, 0x70, 0x65, 0x72, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_public_enum_proto_rawDescOnce sync.Once
	file_public_enum_proto_rawDescData = file_public_enum_proto_rawDesc
)

func file_public_enum_proto_rawDescGZIP() []byte {
	file_public_enum_proto_rawDescOnce.Do(func() {
		file_public_enum_proto_rawDescData = protoimpl.X.CompressGZIP(file_public_enum_proto_rawDescData)
	})
	return file_public_enum_proto_rawDescData
}

var file_public_enum_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_public_enum_proto_goTypes = []interface{}{
	(*Enum)(nil),               // 0: public_service.Enum
	(*Enums)(nil),              // 1: public_service.Enums
	(*ListEnumParams)(nil),     // 2: public_service.ListEnumParams
	(*ListEnumReply)(nil),      // 3: public_service.ListEnumReply
	(*common.Header)(nil),      // 4: public.Header
	(*common.ReplyHeader)(nil), // 5: public.ReplyHeader
}
var file_public_enum_proto_depIdxs = []int32{
	0, // 0: public_service.Enums.data:type_name -> public_service.Enum
	4, // 1: public_service.ListEnumParams.header:type_name -> public.Header
	5, // 2: public_service.ListEnumReply.header:type_name -> public.ReplyHeader
	1, // 3: public_service.ListEnumReply.data:type_name -> public_service.Enums
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_public_enum_proto_init() }
func file_public_enum_proto_init() {
	if File_public_enum_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_public_enum_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Enum); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_public_enum_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Enums); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_public_enum_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEnumParams); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_public_enum_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEnumReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_public_enum_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_public_enum_proto_goTypes,
		DependencyIndexes: file_public_enum_proto_depIdxs,
		MessageInfos:      file_public_enum_proto_msgTypes,
	}.Build()
	File_public_enum_proto = out.File
	file_public_enum_proto_rawDesc = nil
	file_public_enum_proto_goTypes = nil
	file_public_enum_proto_depIdxs = nil
}