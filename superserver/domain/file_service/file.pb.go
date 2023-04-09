// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.2
// source: file/file.proto

package file_service

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

type File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                      // 主键
	FolderId   int32  `protobuf:"varint,2,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`          // 文件夹ID
	Sign       string `protobuf:"bytes,3,opt,name=sign,proto3" json:"sign,omitempty"`                                   // 内容签名
	Name       string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`                                   // 文件名
	Public     bool   `protobuf:"varint,5,opt,name=public,proto3" json:"public,omitempty"`                              // 公开的
	ShareUrl   string `protobuf:"bytes,6,opt,name=share_url,json=shareUrl,proto3" json:"share_url,omitempty"`           // 共享链接
	DeleteTime int64  `protobuf:"varint,1001,opt,name=delete_time,json=deleteTime,proto3" json:"delete_time,omitempty"` // 删除时间 时间戳：秒
	CreateTime int64  `protobuf:"varint,1002,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"` // 创建时间 时间戳：秒
	UpdateTime int64  `protobuf:"varint,1003,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"` // 最后更新时间 时间戳：秒
	Creator    int32  `protobuf:"varint,1004,opt,name=creator,proto3" json:"creator,omitempty"`                         // 创建人
	Updater    int32  `protobuf:"varint,1005,opt,name=updater,proto3" json:"updater,omitempty"`                         // 更新人
}

func (x *File) Reset() {
	*x = File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_file_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_file_file_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_file_file_proto_rawDescGZIP(), []int{0}
}

func (x *File) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *File) GetFolderId() int32 {
	if x != nil {
		return x.FolderId
	}
	return 0
}

func (x *File) GetSign() string {
	if x != nil {
		return x.Sign
	}
	return ""
}

func (x *File) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *File) GetPublic() bool {
	if x != nil {
		return x.Public
	}
	return false
}

func (x *File) GetShareUrl() string {
	if x != nil {
		return x.ShareUrl
	}
	return ""
}

func (x *File) GetDeleteTime() int64 {
	if x != nil {
		return x.DeleteTime
	}
	return 0
}

func (x *File) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *File) GetUpdateTime() int64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

func (x *File) GetCreator() int32 {
	if x != nil {
		return x.Creator
	}
	return 0
}

func (x *File) GetUpdater() int32 {
	if x != nil {
		return x.Updater
	}
	return 0
}

type FileFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids             []int32              `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	FolderIds       []int32              `protobuf:"varint,2,rep,packed,name=folder_ids,json=folderIds,proto3" json:"folder_ids,omitempty"`
	PublicSelect    common.BooleanScope  `protobuf:"varint,3,opt,name=public_select,json=publicSelect,proto3,enum=public.BooleanScope" json:"public_select,omitempty"`
	DeleteTimeRange *common.BetweenInt64 `protobuf:"bytes,101,opt,name=delete_time_range,json=deleteTimeRange,proto3" json:"delete_time_range,omitempty"`
	CreateTimeRange *common.BetweenInt64 `protobuf:"bytes,102,opt,name=create_time_range,json=createTimeRange,proto3" json:"create_time_range,omitempty"`
	UpdateTimeRange *common.BetweenInt64 `protobuf:"bytes,103,opt,name=update_time_range,json=updateTimeRange,proto3" json:"update_time_range,omitempty"`
	Creators        []int32              `protobuf:"varint,104,rep,packed,name=creators,proto3" json:"creators,omitempty"`
	Updaters        []int32              `protobuf:"varint,105,rep,packed,name=updaters,proto3" json:"updaters,omitempty"`
	Deleted         common.BooleanScope  `protobuf:"varint,500,opt,name=deleted,proto3,enum=public.BooleanScope" json:"deleted,omitempty"`
	Keywords        *FileFilter_Keywords `protobuf:"bytes,201,opt,name=keywords,proto3" json:"keywords,omitempty"`
}

func (x *FileFilter) Reset() {
	*x = FileFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_file_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileFilter) ProtoMessage() {}

func (x *FileFilter) ProtoReflect() protoreflect.Message {
	mi := &file_file_file_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileFilter.ProtoReflect.Descriptor instead.
func (*FileFilter) Descriptor() ([]byte, []int) {
	return file_file_file_proto_rawDescGZIP(), []int{1}
}

func (x *FileFilter) GetIds() []int32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *FileFilter) GetFolderIds() []int32 {
	if x != nil {
		return x.FolderIds
	}
	return nil
}

func (x *FileFilter) GetPublicSelect() common.BooleanScope {
	if x != nil {
		return x.PublicSelect
	}
	return common.BooleanScope(0)
}

func (x *FileFilter) GetDeleteTimeRange() *common.BetweenInt64 {
	if x != nil {
		return x.DeleteTimeRange
	}
	return nil
}

func (x *FileFilter) GetCreateTimeRange() *common.BetweenInt64 {
	if x != nil {
		return x.CreateTimeRange
	}
	return nil
}

func (x *FileFilter) GetUpdateTimeRange() *common.BetweenInt64 {
	if x != nil {
		return x.UpdateTimeRange
	}
	return nil
}

func (x *FileFilter) GetCreators() []int32 {
	if x != nil {
		return x.Creators
	}
	return nil
}

func (x *FileFilter) GetUpdaters() []int32 {
	if x != nil {
		return x.Updaters
	}
	return nil
}

func (x *FileFilter) GetDeleted() common.BooleanScope {
	if x != nil {
		return x.Deleted
	}
	return common.BooleanScope(0)
}

func (x *FileFilter) GetKeywords() *FileFilter_Keywords {
	if x != nil {
		return x.Keywords
	}
	return nil
}

type ListFileParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *common.Header `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Pager  *common.Pager  `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Sorts  []*common.Sort `protobuf:"bytes,3,rep,name=sorts,proto3" json:"sorts,omitempty"`
	Filter *FileFilter    `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *ListFileParams) Reset() {
	*x = ListFileParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_file_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFileParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFileParams) ProtoMessage() {}

func (x *ListFileParams) ProtoReflect() protoreflect.Message {
	mi := &file_file_file_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFileParams.ProtoReflect.Descriptor instead.
func (*ListFileParams) Descriptor() ([]byte, []int) {
	return file_file_file_proto_rawDescGZIP(), []int{2}
}

func (x *ListFileParams) GetHeader() *common.Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *ListFileParams) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *ListFileParams) GetSorts() []*common.Sort {
	if x != nil {
		return x.Sorts
	}
	return nil
}

func (x *ListFileParams) GetFilter() *FileFilter {
	if x != nil {
		return x.Filter
	}
	return nil
}

type ListFileReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *common.ReplyHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Pager  *common.Pager       `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Data   []*File             `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ListFileReply) Reset() {
	*x = ListFileReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_file_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFileReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFileReply) ProtoMessage() {}

func (x *ListFileReply) ProtoReflect() protoreflect.Message {
	mi := &file_file_file_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFileReply.ProtoReflect.Descriptor instead.
func (*ListFileReply) Descriptor() ([]byte, []int) {
	return file_file_file_proto_rawDescGZIP(), []int{3}
}

func (x *ListFileReply) GetHeader() *common.ReplyHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *ListFileReply) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *ListFileReply) GetData() []*File {
	if x != nil {
		return x.Data
	}
	return nil
}

type OperateFileParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header  *common.Header   `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Operate common.Operation `protobuf:"varint,2,opt,name=operate,proto3,enum=public.Operation" json:"operate,omitempty"`
	Data    *File            `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *OperateFileParams) Reset() {
	*x = OperateFileParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_file_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperateFileParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperateFileParams) ProtoMessage() {}

func (x *OperateFileParams) ProtoReflect() protoreflect.Message {
	mi := &file_file_file_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperateFileParams.ProtoReflect.Descriptor instead.
func (*OperateFileParams) Descriptor() ([]byte, []int) {
	return file_file_file_proto_rawDescGZIP(), []int{4}
}

func (x *OperateFileParams) GetHeader() *common.Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *OperateFileParams) GetOperate() common.Operation {
	if x != nil {
		return x.Operate
	}
	return common.Operation(0)
}

func (x *OperateFileParams) GetData() *File {
	if x != nil {
		return x.Data
	}
	return nil
}

type OperateFileReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *common.ReplyHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Data   *File               `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *OperateFileReply) Reset() {
	*x = OperateFileReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_file_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperateFileReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperateFileReply) ProtoMessage() {}

func (x *OperateFileReply) ProtoReflect() protoreflect.Message {
	mi := &file_file_file_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperateFileReply.ProtoReflect.Descriptor instead.
func (*OperateFileReply) Descriptor() ([]byte, []int) {
	return file_file_file_proto_rawDescGZIP(), []int{5}
}

func (x *OperateFileReply) GetHeader() *common.ReplyHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *OperateFileReply) GetData() *File {
	if x != nil {
		return x.Data
	}
	return nil
}

type BatchOperateFileParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header  *common.Header   `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Operate common.Operation `protobuf:"varint,2,opt,name=operate,proto3,enum=public.Operation" json:"operate,omitempty"`
	Data    []*File          `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *BatchOperateFileParams) Reset() {
	*x = BatchOperateFileParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_file_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchOperateFileParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchOperateFileParams) ProtoMessage() {}

func (x *BatchOperateFileParams) ProtoReflect() protoreflect.Message {
	mi := &file_file_file_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchOperateFileParams.ProtoReflect.Descriptor instead.
func (*BatchOperateFileParams) Descriptor() ([]byte, []int) {
	return file_file_file_proto_rawDescGZIP(), []int{6}
}

func (x *BatchOperateFileParams) GetHeader() *common.Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *BatchOperateFileParams) GetOperate() common.Operation {
	if x != nil {
		return x.Operate
	}
	return common.Operation(0)
}

func (x *BatchOperateFileParams) GetData() []*File {
	if x != nil {
		return x.Data
	}
	return nil
}

type BatchOperateFileReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *common.ReplyHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
}

func (x *BatchOperateFileReply) Reset() {
	*x = BatchOperateFileReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_file_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchOperateFileReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchOperateFileReply) ProtoMessage() {}

func (x *BatchOperateFileReply) ProtoReflect() protoreflect.Message {
	mi := &file_file_file_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchOperateFileReply.ProtoReflect.Descriptor instead.
func (*BatchOperateFileReply) Descriptor() ([]byte, []int) {
	return file_file_file_proto_rawDescGZIP(), []int{7}
}

func (x *BatchOperateFileReply) GetHeader() *common.ReplyHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

type FileFilter_Keywords struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *FileFilter_Keywords) Reset() {
	*x = FileFilter_Keywords{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_file_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileFilter_Keywords) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileFilter_Keywords) ProtoMessage() {}

func (x *FileFilter_Keywords) ProtoReflect() protoreflect.Message {
	mi := &file_file_file_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileFilter_Keywords.ProtoReflect.Descriptor instead.
func (*FileFilter_Keywords) Descriptor() ([]byte, []int) {
	return file_file_file_proto_rawDescGZIP(), []int{1, 0}
}

func (x *FileFilter_Keywords) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_file_file_proto protoreflect.FileDescriptor

var file_file_file_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a,
	0x13, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xac, 0x02, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69,
	0x67, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x69, 0x67, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0xe9, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0xea, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0xeb, 0x07, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a,
	0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0xec, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x19, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x72, 0x18, 0xed, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x72, 0x22, 0x87, 0x04, 0x0a, 0x0a, 0x46, 0x69, 0x6c, 0x65, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52,
	0x03, 0x69, 0x64, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x09, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x73, 0x12, 0x39, 0x0a, 0x0d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x73, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x53, 0x63, 0x6f, 0x70, 0x65,
	0x52, 0x0c, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x12, 0x40,
	0x0a, 0x11, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x72, 0x61,
	0x6e, 0x67, 0x65, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x2e, 0x42, 0x65, 0x74, 0x77, 0x65, 0x65, 0x6e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x52,
	0x0f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65,
	0x12, 0x40, 0x0a, 0x11, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f,
	0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x66, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x2e, 0x42, 0x65, 0x74, 0x77, 0x65, 0x65, 0x6e, 0x49, 0x6e, 0x74, 0x36,
	0x34, 0x52, 0x0f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e,
	0x67, 0x65, 0x12, 0x40, 0x0a, 0x11, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x67, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x42, 0x65, 0x74, 0x77, 0x65, 0x65, 0x6e, 0x49, 0x6e,
	0x74, 0x36, 0x34, 0x52, 0x0f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x52,
	0x61, 0x6e, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x73,
	0x18, 0x68, 0x20, 0x03, 0x28, 0x05, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x73,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x73, 0x18, 0x69, 0x20, 0x03,
	0x28, 0x05, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x73, 0x12, 0x2f, 0x0a, 0x07,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0xf4, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14,
	0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x53,
	0x63, 0x6f, 0x70, 0x65, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x3e, 0x0a,
	0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0xc9, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x46, 0x69, 0x6c, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x4b, 0x65, 0x79, 0x77, 0x6f,
	0x72, 0x64, 0x73, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x1a, 0x1e, 0x0a,
	0x08, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xb3, 0x01,
	0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x12, 0x26, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x22, 0x0a,
	0x05, 0x73, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x52, 0x05, 0x73, 0x6f, 0x72, 0x74,
	0x73, 0x12, 0x30, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x22, 0x89, 0x01, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2b, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72,
	0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x90, 0x01, 0x0a, 0x11, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x26, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x2b, 0x0a,
	0x07, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11,
	0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x07, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x67, 0x0a, 0x10, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2b, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x95, 0x01, 0x0a, 0x16,
	0x42, 0x61, 0x74, 0x63, 0x68, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x26, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x2b,
	0x0a, 0x07, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x11, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x07, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x44, 0x0a, 0x15, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2b, 0x0a, 0x06,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x42, 0x21, 0x5a, 0x1f, 0x73, 0x75, 0x70,
	0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f,
	0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_file_file_proto_rawDescOnce sync.Once
	file_file_file_proto_rawDescData = file_file_file_proto_rawDesc
)

func file_file_file_proto_rawDescGZIP() []byte {
	file_file_file_proto_rawDescOnce.Do(func() {
		file_file_file_proto_rawDescData = protoimpl.X.CompressGZIP(file_file_file_proto_rawDescData)
	})
	return file_file_file_proto_rawDescData
}

var file_file_file_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_file_file_proto_goTypes = []interface{}{
	(*File)(nil),                   // 0: file_service.File
	(*FileFilter)(nil),             // 1: file_service.FileFilter
	(*ListFileParams)(nil),         // 2: file_service.ListFileParams
	(*ListFileReply)(nil),          // 3: file_service.ListFileReply
	(*OperateFileParams)(nil),      // 4: file_service.OperateFileParams
	(*OperateFileReply)(nil),       // 5: file_service.OperateFileReply
	(*BatchOperateFileParams)(nil), // 6: file_service.BatchOperateFileParams
	(*BatchOperateFileReply)(nil),  // 7: file_service.BatchOperateFileReply
	(*FileFilter_Keywords)(nil),    // 8: file_service.FileFilter.Keywords
	(common.BooleanScope)(0),       // 9: public.BooleanScope
	(*common.BetweenInt64)(nil),    // 10: public.BetweenInt64
	(*common.Header)(nil),          // 11: public.Header
	(*common.Pager)(nil),           // 12: public.Pager
	(*common.Sort)(nil),            // 13: public.Sort
	(*common.ReplyHeader)(nil),     // 14: public.ReplyHeader
	(common.Operation)(0),          // 15: public.Operation
}
var file_file_file_proto_depIdxs = []int32{
	9,  // 0: file_service.FileFilter.public_select:type_name -> public.BooleanScope
	10, // 1: file_service.FileFilter.delete_time_range:type_name -> public.BetweenInt64
	10, // 2: file_service.FileFilter.create_time_range:type_name -> public.BetweenInt64
	10, // 3: file_service.FileFilter.update_time_range:type_name -> public.BetweenInt64
	9,  // 4: file_service.FileFilter.deleted:type_name -> public.BooleanScope
	8,  // 5: file_service.FileFilter.keywords:type_name -> file_service.FileFilter.Keywords
	11, // 6: file_service.ListFileParams.header:type_name -> public.Header
	12, // 7: file_service.ListFileParams.pager:type_name -> public.Pager
	13, // 8: file_service.ListFileParams.sorts:type_name -> public.Sort
	1,  // 9: file_service.ListFileParams.filter:type_name -> file_service.FileFilter
	14, // 10: file_service.ListFileReply.header:type_name -> public.ReplyHeader
	12, // 11: file_service.ListFileReply.pager:type_name -> public.Pager
	0,  // 12: file_service.ListFileReply.data:type_name -> file_service.File
	11, // 13: file_service.OperateFileParams.header:type_name -> public.Header
	15, // 14: file_service.OperateFileParams.operate:type_name -> public.Operation
	0,  // 15: file_service.OperateFileParams.data:type_name -> file_service.File
	14, // 16: file_service.OperateFileReply.header:type_name -> public.ReplyHeader
	0,  // 17: file_service.OperateFileReply.data:type_name -> file_service.File
	11, // 18: file_service.BatchOperateFileParams.header:type_name -> public.Header
	15, // 19: file_service.BatchOperateFileParams.operate:type_name -> public.Operation
	0,  // 20: file_service.BatchOperateFileParams.data:type_name -> file_service.File
	14, // 21: file_service.BatchOperateFileReply.header:type_name -> public.ReplyHeader
	22, // [22:22] is the sub-list for method output_type
	22, // [22:22] is the sub-list for method input_type
	22, // [22:22] is the sub-list for extension type_name
	22, // [22:22] is the sub-list for extension extendee
	0,  // [0:22] is the sub-list for field type_name
}

func init() { file_file_file_proto_init() }
func file_file_file_proto_init() {
	if File_file_file_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_file_file_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*File); i {
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
		file_file_file_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileFilter); i {
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
		file_file_file_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFileParams); i {
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
		file_file_file_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFileReply); i {
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
		file_file_file_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperateFileParams); i {
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
		file_file_file_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperateFileReply); i {
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
		file_file_file_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchOperateFileParams); i {
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
		file_file_file_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchOperateFileReply); i {
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
		file_file_file_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileFilter_Keywords); i {
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
			RawDescriptor: file_file_file_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_file_file_proto_goTypes,
		DependencyIndexes: file_file_file_proto_depIdxs,
		MessageInfos:      file_file_file_proto_msgTypes,
	}.Build()
	File_file_file_proto = out.File
	file_file_file_proto_rawDesc = nil
	file_file_file_proto_goTypes = nil
	file_file_file_proto_depIdxs = nil
}
