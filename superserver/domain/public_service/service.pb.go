// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.2
// source: public/service.proto

package public_service

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

var File_public_service_proto protoreflect.FileDescriptor

var file_public_service_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x13, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x66,
	0x6f, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x2f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
	0xae, 0x04, 0x0a, 0x0d, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x4c, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x1f,
	0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a,
	0x1e, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x55, 0x0a, 0x0c, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12,
	0x22, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x1a, 0x21, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x64, 0x0a, 0x11, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x27, 0x2e, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x1a, 0x26, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x4f, 0x0a, 0x0a,
	0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x20, 0x2e, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x1f, 0x2e, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x58, 0x0a,
	0x0d, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x23,
	0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x1a, 0x22, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x6c, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x67, 0x0a, 0x12, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x28, 0x2e,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x6c, 0x64, 0x65,
	0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x27, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x42, 0x23, 0x5a, 0x21, 0x73, 0x75, 0x70, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_public_service_proto_goTypes = []interface{}{
	(*ListLabelParams)(nil),          // 0: public_service.ListLabelParams
	(*OperateLabelParams)(nil),       // 1: public_service.OperateLabelParams
	(*BatchOperateLabelParams)(nil),  // 2: public_service.BatchOperateLabelParams
	(*ListFolderParams)(nil),         // 3: public_service.ListFolderParams
	(*OperateFolderParams)(nil),      // 4: public_service.OperateFolderParams
	(*BatchOperateFolderParams)(nil), // 5: public_service.BatchOperateFolderParams
	(*ListLabelReply)(nil),           // 6: public_service.ListLabelReply
	(*OperateLabelReply)(nil),        // 7: public_service.OperateLabelReply
	(*BatchOperateLabelReply)(nil),   // 8: public_service.BatchOperateLabelReply
	(*ListFolderReply)(nil),          // 9: public_service.ListFolderReply
	(*OperateFolderReply)(nil),       // 10: public_service.OperateFolderReply
	(*BatchOperateFolderReply)(nil),  // 11: public_service.BatchOperateFolderReply
}
var file_public_service_proto_depIdxs = []int32{
	0,  // 0: public_service.PublicService.ListLabel:input_type -> public_service.ListLabelParams
	1,  // 1: public_service.PublicService.OperateLabel:input_type -> public_service.OperateLabelParams
	2,  // 2: public_service.PublicService.BatchOperateLabel:input_type -> public_service.BatchOperateLabelParams
	3,  // 3: public_service.PublicService.ListFolder:input_type -> public_service.ListFolderParams
	4,  // 4: public_service.PublicService.OperateFolder:input_type -> public_service.OperateFolderParams
	5,  // 5: public_service.PublicService.BatchOperateFolder:input_type -> public_service.BatchOperateFolderParams
	6,  // 6: public_service.PublicService.ListLabel:output_type -> public_service.ListLabelReply
	7,  // 7: public_service.PublicService.OperateLabel:output_type -> public_service.OperateLabelReply
	8,  // 8: public_service.PublicService.BatchOperateLabel:output_type -> public_service.BatchOperateLabelReply
	9,  // 9: public_service.PublicService.ListFolder:output_type -> public_service.ListFolderReply
	10, // 10: public_service.PublicService.OperateFolder:output_type -> public_service.OperateFolderReply
	11, // 11: public_service.PublicService.BatchOperateFolder:output_type -> public_service.BatchOperateFolderReply
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_public_service_proto_init() }
func file_public_service_proto_init() {
	if File_public_service_proto != nil {
		return
	}
	file_public_folder_proto_init()
	file_public_label_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_public_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_public_service_proto_goTypes,
		DependencyIndexes: file_public_service_proto_depIdxs,
	}.Build()
	File_public_service_proto = out.File
	file_public_service_proto_rawDesc = nil
	file_public_service_proto_goTypes = nil
	file_public_service_proto_depIdxs = nil
}
