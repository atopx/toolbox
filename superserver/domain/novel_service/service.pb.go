// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.2
// source: novel/service.proto

package novel_service

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

var File_novel_service_proto protoreflect.FileDescriptor

var file_novel_service_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6e, 0x6f, 0x76, 0x65, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6e, 0x6f, 0x76, 0x65, 0x6c, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x1a, 0x10, 0x6e, 0x6f, 0x76, 0x65, 0x6c, 0x2f, 0x62, 0x6f, 0x6f, 0x6b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x6e, 0x6f, 0x76, 0x65, 0x6c, 0x2f, 0x63, 0x68,
	0x61, 0x70, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xa9, 0x01, 0x0a, 0x0c,
	0x4e, 0x6f, 0x76, 0x65, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x08,
	0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x1d, 0x2e, 0x6e, 0x6f, 0x76, 0x65, 0x6c,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f,
	0x6b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x1c, 0x2e, 0x6e, 0x6f, 0x76, 0x65, 0x6c, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6b,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x50, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x61,
	0x70, 0x74, 0x65, 0x72, 0x12, 0x20, 0x2e, 0x6e, 0x6f, 0x76, 0x65, 0x6c, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x1f, 0x2e, 0x6e, 0x6f, 0x76, 0x65, 0x6c, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x61, 0x70, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x22, 0x5a, 0x20, 0x73, 0x75, 0x70, 0x65, 0x72,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x6e, 0x6f,
	0x76, 0x65, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var file_novel_service_proto_goTypes = []interface{}{
	(*ListBookParams)(nil),    // 0: novel_service.ListBookParams
	(*ListChapterParams)(nil), // 1: novel_service.ListChapterParams
	(*ListBookReply)(nil),     // 2: novel_service.ListBookReply
	(*ListChapterReply)(nil),  // 3: novel_service.ListChapterReply
}
var file_novel_service_proto_depIdxs = []int32{
	0, // 0: novel_service.NovelService.ListBook:input_type -> novel_service.ListBookParams
	1, // 1: novel_service.NovelService.ListChapter:input_type -> novel_service.ListChapterParams
	2, // 2: novel_service.NovelService.ListBook:output_type -> novel_service.ListBookReply
	3, // 3: novel_service.NovelService.ListChapter:output_type -> novel_service.ListChapterReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_novel_service_proto_init() }
func file_novel_service_proto_init() {
	if File_novel_service_proto != nil {
		return
	}
	file_novel_book_proto_init()
	file_novel_chapter_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_novel_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_novel_service_proto_goTypes,
		DependencyIndexes: file_novel_service_proto_depIdxs,
	}.Build()
	File_novel_service_proto = out.File
	file_novel_service_proto_rawDesc = nil
	file_novel_service_proto_goTypes = nil
	file_novel_service_proto_depIdxs = nil
}
