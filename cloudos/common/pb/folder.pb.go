// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: folder.proto

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

type DomainFolder int32

const (
	DomainFolder_NONE DomainFolder = 0
	DomainFolder_NOTE DomainFolder = 1
	DomainFolder_FILE DomainFolder = 2
)

// Enum value maps for DomainFolder.
var (
	DomainFolder_name = map[int32]string{
		0: "NONE",
		1: "NOTE",
		2: "FILE",
	}
	DomainFolder_value = map[string]int32{
		"NONE": 0,
		"NOTE": 1,
		"FILE": 2,
	}
)

func (x DomainFolder) Enum() *DomainFolder {
	p := new(DomainFolder)
	*p = x
	return p
}

func (x DomainFolder) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DomainFolder) Descriptor() protoreflect.EnumDescriptor {
	return file_folder_proto_enumTypes[0].Descriptor()
}

func (DomainFolder) Type() protoreflect.EnumType {
	return &file_folder_proto_enumTypes[0]
}

func (x DomainFolder) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DomainFolder.Descriptor instead.
func (DomainFolder) EnumDescriptor() ([]byte, []int) {
	return file_folder_proto_rawDescGZIP(), []int{0}
}

type Folder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ParentId   int64  `protobuf:"varint,3,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	Domain     string `protobuf:"bytes,4,opt,name=domain,proto3" json:"domain,omitempty"`
	CreateTime int64  `protobuf:"varint,1001,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime int64  `protobuf:"varint,1002,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	DeleteTime int64  `protobuf:"varint,1003,opt,name=delete_time,json=deleteTime,proto3" json:"delete_time,omitempty"`
}

func (x *Folder) Reset() {
	*x = Folder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_folder_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Folder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Folder) ProtoMessage() {}

func (x *Folder) ProtoReflect() protoreflect.Message {
	mi := &file_folder_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Folder.ProtoReflect.Descriptor instead.
func (*Folder) Descriptor() ([]byte, []int) {
	return file_folder_proto_rawDescGZIP(), []int{0}
}

func (x *Folder) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Folder) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Folder) GetParentId() int64 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *Folder) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *Folder) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *Folder) GetUpdateTime() int64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

func (x *Folder) GetDeleteTime() int64 {
	if x != nil {
		return x.DeleteTime
	}
	return 0
}

var File_folder_proto protoreflect.FileDescriptor

var file_folder_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04,
	0x6e, 0x6f, 0x74, 0x65, 0x22, 0xc7, 0x01, 0x0a, 0x06, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0xe9, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0xea, 0x07, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0xeb, 0x07, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x2a, 0x2c,
	0x0a, 0x0c, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x08,
	0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x54, 0x45,
	0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x49, 0x4c, 0x45, 0x10, 0x02, 0x42, 0x13, 0x5a, 0x11,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_folder_proto_rawDescOnce sync.Once
	file_folder_proto_rawDescData = file_folder_proto_rawDesc
)

func file_folder_proto_rawDescGZIP() []byte {
	file_folder_proto_rawDescOnce.Do(func() {
		file_folder_proto_rawDescData = protoimpl.X.CompressGZIP(file_folder_proto_rawDescData)
	})
	return file_folder_proto_rawDescData
}

var file_folder_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_folder_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_folder_proto_goTypes = []interface{}{
	(DomainFolder)(0), // 0: note.DomainFolder
	(*Folder)(nil),    // 1: note.Folder
}
var file_folder_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_folder_proto_init() }
func file_folder_proto_init() {
	if File_folder_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_folder_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Folder); i {
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
			RawDescriptor: file_folder_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_folder_proto_goTypes,
		DependencyIndexes: file_folder_proto_depIdxs,
		EnumInfos:         file_folder_proto_enumTypes,
		MessageInfos:      file_folder_proto_msgTypes,
	}.Build()
	File_folder_proto = out.File
	file_folder_proto_rawDesc = nil
	file_folder_proto_goTypes = nil
	file_folder_proto_depIdxs = nil
}
