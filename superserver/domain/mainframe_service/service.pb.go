// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.2
// source: mainframe/service.proto

package mainframe_service

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

var File_mainframe_service_proto protoreflect.FileDescriptor

var file_mainframe_service_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6d, 0x61, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x6d, 0x61, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x18, 0x6d, 0x61,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xa7, 0x03, 0x0a, 0x10, 0x4d, 0x61, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5b, 0x0a, 0x0c, 0x44,
	0x65, 0x61, 0x6c, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x72, 0x12, 0x25, 0x2e, 0x6d, 0x61,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x44, 0x65, 0x61, 0x6c, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x72, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x1a, 0x24, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x61, 0x6c, 0x43, 0x6f, 0x6d, 0x70, 0x75,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x5b, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74,
	0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x72, 0x12, 0x25, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a,
	0x24, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x64, 0x0a, 0x0f, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x72, 0x12, 0x28, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x72, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x1a, 0x27, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x43, 0x6f,
	0x6d, 0x70, 0x75, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x73, 0x0a, 0x14, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x75,
	0x74, 0x65, 0x72, 0x12, 0x2d, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x72, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x1a, 0x2c, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x42, 0x26, 0x5a, 0x24, 0x73, 0x75, 0x70, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x6d, 0x65,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_mainframe_service_proto_goTypes = []interface{}{
	(*DealComputerParams)(nil),         // 0: mainframe_service.DealComputerParams
	(*ListComputerParams)(nil),         // 1: mainframe_service.ListComputerParams
	(*OperateComputerParams)(nil),      // 2: mainframe_service.OperateComputerParams
	(*BatchOperateComputerParams)(nil), // 3: mainframe_service.BatchOperateComputerParams
	(*DealComputerReply)(nil),          // 4: mainframe_service.DealComputerReply
	(*ListComputerReply)(nil),          // 5: mainframe_service.ListComputerReply
	(*OperateComputerReply)(nil),       // 6: mainframe_service.OperateComputerReply
	(*BatchOperateComputerReply)(nil),  // 7: mainframe_service.BatchOperateComputerReply
}
var file_mainframe_service_proto_depIdxs = []int32{
	0, // 0: mainframe_service.MainframeService.DealComputer:input_type -> mainframe_service.DealComputerParams
	1, // 1: mainframe_service.MainframeService.ListComputer:input_type -> mainframe_service.ListComputerParams
	2, // 2: mainframe_service.MainframeService.OperateComputer:input_type -> mainframe_service.OperateComputerParams
	3, // 3: mainframe_service.MainframeService.BatchOperateComputer:input_type -> mainframe_service.BatchOperateComputerParams
	4, // 4: mainframe_service.MainframeService.DealComputer:output_type -> mainframe_service.DealComputerReply
	5, // 5: mainframe_service.MainframeService.ListComputer:output_type -> mainframe_service.ListComputerReply
	6, // 6: mainframe_service.MainframeService.OperateComputer:output_type -> mainframe_service.OperateComputerReply
	7, // 7: mainframe_service.MainframeService.BatchOperateComputer:output_type -> mainframe_service.BatchOperateComputerReply
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mainframe_service_proto_init() }
func file_mainframe_service_proto_init() {
	if File_mainframe_service_proto != nil {
		return
	}
	file_mainframe_computer_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mainframe_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mainframe_service_proto_goTypes,
		DependencyIndexes: file_mainframe_service_proto_depIdxs,
	}.Build()
	File_mainframe_service_proto = out.File
	file_mainframe_service_proto_rawDesc = nil
	file_mainframe_service_proto_goTypes = nil
	file_mainframe_service_proto_depIdxs = nil
}
