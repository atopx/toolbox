// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.2
// source: note/topic.proto

package note_service

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

type NoteTopic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                      // 主键
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                                   // 名称
	DeleteTime int64  `protobuf:"varint,1001,opt,name=delete_time,json=deleteTime,proto3" json:"delete_time,omitempty"` // 删除时间 时间戳：秒
	CreateTime int64  `protobuf:"varint,1002,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"` // 创建时间 时间戳：秒
	UpdateTime int64  `protobuf:"varint,1003,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"` // 最后更新时间 时间戳：秒
	Creator    int32  `protobuf:"varint,1004,opt,name=creator,proto3" json:"creator,omitempty"`                         // 创建人
	Updater    int32  `protobuf:"varint,1005,opt,name=updater,proto3" json:"updater,omitempty"`                         // 更新人
}

func (x *NoteTopic) Reset() {
	*x = NoteTopic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_note_topic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoteTopic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoteTopic) ProtoMessage() {}

func (x *NoteTopic) ProtoReflect() protoreflect.Message {
	mi := &file_note_topic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoteTopic.ProtoReflect.Descriptor instead.
func (*NoteTopic) Descriptor() ([]byte, []int) {
	return file_note_topic_proto_rawDescGZIP(), []int{0}
}

func (x *NoteTopic) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *NoteTopic) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NoteTopic) GetDeleteTime() int64 {
	if x != nil {
		return x.DeleteTime
	}
	return 0
}

func (x *NoteTopic) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *NoteTopic) GetUpdateTime() int64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

func (x *NoteTopic) GetCreator() int32 {
	if x != nil {
		return x.Creator
	}
	return 0
}

func (x *NoteTopic) GetUpdater() int32 {
	if x != nil {
		return x.Updater
	}
	return 0
}

type NoteTopicFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids             []int32                   `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	DeleteTimeRange *common.BetweenInt64      `protobuf:"bytes,101,opt,name=delete_time_range,json=deleteTimeRange,proto3" json:"delete_time_range,omitempty"`
	CreateTimeRange *common.BetweenInt64      `protobuf:"bytes,102,opt,name=create_time_range,json=createTimeRange,proto3" json:"create_time_range,omitempty"`
	UpdateTimeRange *common.BetweenInt64      `protobuf:"bytes,103,opt,name=update_time_range,json=updateTimeRange,proto3" json:"update_time_range,omitempty"`
	Creators        []int32                   `protobuf:"varint,104,rep,packed,name=creators,proto3" json:"creators,omitempty"`
	Updaters        []int32                   `protobuf:"varint,105,rep,packed,name=updaters,proto3" json:"updaters,omitempty"`
	Keywords        *NoteTopicFilter_Keywords `protobuf:"bytes,201,opt,name=keywords,proto3" json:"keywords,omitempty"`
}

func (x *NoteTopicFilter) Reset() {
	*x = NoteTopicFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_note_topic_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoteTopicFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoteTopicFilter) ProtoMessage() {}

func (x *NoteTopicFilter) ProtoReflect() protoreflect.Message {
	mi := &file_note_topic_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoteTopicFilter.ProtoReflect.Descriptor instead.
func (*NoteTopicFilter) Descriptor() ([]byte, []int) {
	return file_note_topic_proto_rawDescGZIP(), []int{1}
}

func (x *NoteTopicFilter) GetIds() []int32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *NoteTopicFilter) GetDeleteTimeRange() *common.BetweenInt64 {
	if x != nil {
		return x.DeleteTimeRange
	}
	return nil
}

func (x *NoteTopicFilter) GetCreateTimeRange() *common.BetweenInt64 {
	if x != nil {
		return x.CreateTimeRange
	}
	return nil
}

func (x *NoteTopicFilter) GetUpdateTimeRange() *common.BetweenInt64 {
	if x != nil {
		return x.UpdateTimeRange
	}
	return nil
}

func (x *NoteTopicFilter) GetCreators() []int32 {
	if x != nil {
		return x.Creators
	}
	return nil
}

func (x *NoteTopicFilter) GetUpdaters() []int32 {
	if x != nil {
		return x.Updaters
	}
	return nil
}

func (x *NoteTopicFilter) GetKeywords() *NoteTopicFilter_Keywords {
	if x != nil {
		return x.Keywords
	}
	return nil
}

type ListNoteTopicParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *common.Header   `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Pager  *common.Pager    `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Sorts  []*common.Sort   `protobuf:"bytes,3,rep,name=sorts,proto3" json:"sorts,omitempty"`
	Filter *NoteTopicFilter `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *ListNoteTopicParams) Reset() {
	*x = ListNoteTopicParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_note_topic_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNoteTopicParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNoteTopicParams) ProtoMessage() {}

func (x *ListNoteTopicParams) ProtoReflect() protoreflect.Message {
	mi := &file_note_topic_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNoteTopicParams.ProtoReflect.Descriptor instead.
func (*ListNoteTopicParams) Descriptor() ([]byte, []int) {
	return file_note_topic_proto_rawDescGZIP(), []int{2}
}

func (x *ListNoteTopicParams) GetHeader() *common.Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *ListNoteTopicParams) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *ListNoteTopicParams) GetSorts() []*common.Sort {
	if x != nil {
		return x.Sorts
	}
	return nil
}

func (x *ListNoteTopicParams) GetFilter() *NoteTopicFilter {
	if x != nil {
		return x.Filter
	}
	return nil
}

type ListNoteTopicReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *common.ReplyHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Pager  *common.Pager       `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Data   []*NoteTopic        `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ListNoteTopicReply) Reset() {
	*x = ListNoteTopicReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_note_topic_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNoteTopicReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNoteTopicReply) ProtoMessage() {}

func (x *ListNoteTopicReply) ProtoReflect() protoreflect.Message {
	mi := &file_note_topic_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNoteTopicReply.ProtoReflect.Descriptor instead.
func (*ListNoteTopicReply) Descriptor() ([]byte, []int) {
	return file_note_topic_proto_rawDescGZIP(), []int{3}
}

func (x *ListNoteTopicReply) GetHeader() *common.ReplyHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *ListNoteTopicReply) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *ListNoteTopicReply) GetData() []*NoteTopic {
	if x != nil {
		return x.Data
	}
	return nil
}

type OperateNoteTopicParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header  *common.Header   `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Operate common.Operation `protobuf:"varint,2,opt,name=operate,proto3,enum=public.Operation" json:"operate,omitempty"`
	Data    *NoteTopic       `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *OperateNoteTopicParams) Reset() {
	*x = OperateNoteTopicParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_note_topic_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperateNoteTopicParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperateNoteTopicParams) ProtoMessage() {}

func (x *OperateNoteTopicParams) ProtoReflect() protoreflect.Message {
	mi := &file_note_topic_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperateNoteTopicParams.ProtoReflect.Descriptor instead.
func (*OperateNoteTopicParams) Descriptor() ([]byte, []int) {
	return file_note_topic_proto_rawDescGZIP(), []int{4}
}

func (x *OperateNoteTopicParams) GetHeader() *common.Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *OperateNoteTopicParams) GetOperate() common.Operation {
	if x != nil {
		return x.Operate
	}
	return common.Operation(0)
}

func (x *OperateNoteTopicParams) GetData() *NoteTopic {
	if x != nil {
		return x.Data
	}
	return nil
}

type OperateNoteTopicReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *common.ReplyHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Data   *NoteTopic          `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *OperateNoteTopicReply) Reset() {
	*x = OperateNoteTopicReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_note_topic_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperateNoteTopicReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperateNoteTopicReply) ProtoMessage() {}

func (x *OperateNoteTopicReply) ProtoReflect() protoreflect.Message {
	mi := &file_note_topic_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperateNoteTopicReply.ProtoReflect.Descriptor instead.
func (*OperateNoteTopicReply) Descriptor() ([]byte, []int) {
	return file_note_topic_proto_rawDescGZIP(), []int{5}
}

func (x *OperateNoteTopicReply) GetHeader() *common.ReplyHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *OperateNoteTopicReply) GetData() *NoteTopic {
	if x != nil {
		return x.Data
	}
	return nil
}

type BatchOperateNoteTopicParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header  *common.Header   `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Operate common.Operation `protobuf:"varint,2,opt,name=operate,proto3,enum=public.Operation" json:"operate,omitempty"`
	Data    []*NoteTopic     `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *BatchOperateNoteTopicParams) Reset() {
	*x = BatchOperateNoteTopicParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_note_topic_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchOperateNoteTopicParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchOperateNoteTopicParams) ProtoMessage() {}

func (x *BatchOperateNoteTopicParams) ProtoReflect() protoreflect.Message {
	mi := &file_note_topic_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchOperateNoteTopicParams.ProtoReflect.Descriptor instead.
func (*BatchOperateNoteTopicParams) Descriptor() ([]byte, []int) {
	return file_note_topic_proto_rawDescGZIP(), []int{6}
}

func (x *BatchOperateNoteTopicParams) GetHeader() *common.Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *BatchOperateNoteTopicParams) GetOperate() common.Operation {
	if x != nil {
		return x.Operate
	}
	return common.Operation(0)
}

func (x *BatchOperateNoteTopicParams) GetData() []*NoteTopic {
	if x != nil {
		return x.Data
	}
	return nil
}

type BatchOperateNoteTopicReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *common.ReplyHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
}

func (x *BatchOperateNoteTopicReply) Reset() {
	*x = BatchOperateNoteTopicReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_note_topic_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchOperateNoteTopicReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchOperateNoteTopicReply) ProtoMessage() {}

func (x *BatchOperateNoteTopicReply) ProtoReflect() protoreflect.Message {
	mi := &file_note_topic_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchOperateNoteTopicReply.ProtoReflect.Descriptor instead.
func (*BatchOperateNoteTopicReply) Descriptor() ([]byte, []int) {
	return file_note_topic_proto_rawDescGZIP(), []int{7}
}

func (x *BatchOperateNoteTopicReply) GetHeader() *common.ReplyHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

type NoteTopicFilter_Keywords struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *NoteTopicFilter_Keywords) Reset() {
	*x = NoteTopicFilter_Keywords{}
	if protoimpl.UnsafeEnabled {
		mi := &file_note_topic_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoteTopicFilter_Keywords) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoteTopicFilter_Keywords) ProtoMessage() {}

func (x *NoteTopicFilter_Keywords) ProtoReflect() protoreflect.Message {
	mi := &file_note_topic_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoteTopicFilter_Keywords.ProtoReflect.Descriptor instead.
func (*NoteTopicFilter_Keywords) Descriptor() ([]byte, []int) {
	return file_note_topic_proto_rawDescGZIP(), []int{1, 0}
}

func (x *NoteTopicFilter_Keywords) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_note_topic_proto protoreflect.FileDescriptor

var file_note_topic_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6e, 0x6f, 0x74, 0x65, 0x2f, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0c, 0x6e, 0x6f, 0x74, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x1a, 0x13, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcb, 0x01, 0x0a, 0x09, 0x4e, 0x6f, 0x74, 0x65, 0x54, 0x6f,
	0x70, 0x69, 0x63, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0xe9, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0xea, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0xeb, 0x07, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a,
	0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0xec, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x19, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x72, 0x18, 0xed, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x72, 0x22, 0x86, 0x03, 0x0a, 0x0f, 0x4e, 0x6f, 0x74, 0x65, 0x54, 0x6f, 0x70, 0x69,
	0x63, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x05, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x40, 0x0a, 0x11, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x65,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x42, 0x65,
	0x74, 0x77, 0x65, 0x65, 0x6e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x52, 0x0f, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x40, 0x0a, 0x11, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65,
	0x18, 0x66, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e,
	0x42, 0x65, 0x74, 0x77, 0x65, 0x65, 0x6e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x52, 0x0f, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x40, 0x0a,
	0x11, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x72, 0x61, 0x6e,
	0x67, 0x65, 0x18, 0x67, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x2e, 0x42, 0x65, 0x74, 0x77, 0x65, 0x65, 0x6e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x52, 0x0f,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x68, 0x20, 0x03, 0x28,
	0x05, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x73, 0x18, 0x69, 0x20, 0x03, 0x28, 0x05, 0x52, 0x08, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x73, 0x12, 0x43, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f,
	0x72, 0x64, 0x73, 0x18, 0xc9, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6e, 0x6f, 0x74,
	0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x54, 0x6f,
	0x70, 0x69, 0x63, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72,
	0x64, 0x73, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x1a, 0x1e, 0x0a, 0x08,
	0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xbd, 0x01, 0x0a,
	0x13, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x12, 0x26, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x05,
	0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65,
	0x72, 0x12, 0x22, 0x0a, 0x05, 0x73, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x52, 0x05,
	0x73, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x35, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x93, 0x01, 0x0a,
	0x12, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x2b, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05,
	0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x9a, 0x01, 0x0a, 0x16, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4e, 0x6f,
	0x74, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x26, 0x0a,
	0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x07, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x4e, 0x6f, 0x74, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x71, 0x0a, 0x15, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x54, 0x6f,
	0x70, 0x69, 0x63, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2b, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x9f, 0x01, 0x0a, 0x1b, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x12, 0x26, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x07, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x49, 0x0a, 0x1a, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x2b, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x42,
	0x21, 0x5a, 0x1f, 0x73, 0x75, 0x70, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_note_topic_proto_rawDescOnce sync.Once
	file_note_topic_proto_rawDescData = file_note_topic_proto_rawDesc
)

func file_note_topic_proto_rawDescGZIP() []byte {
	file_note_topic_proto_rawDescOnce.Do(func() {
		file_note_topic_proto_rawDescData = protoimpl.X.CompressGZIP(file_note_topic_proto_rawDescData)
	})
	return file_note_topic_proto_rawDescData
}

var file_note_topic_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_note_topic_proto_goTypes = []interface{}{
	(*NoteTopic)(nil),                   // 0: note_service.NoteTopic
	(*NoteTopicFilter)(nil),             // 1: note_service.NoteTopicFilter
	(*ListNoteTopicParams)(nil),         // 2: note_service.ListNoteTopicParams
	(*ListNoteTopicReply)(nil),          // 3: note_service.ListNoteTopicReply
	(*OperateNoteTopicParams)(nil),      // 4: note_service.OperateNoteTopicParams
	(*OperateNoteTopicReply)(nil),       // 5: note_service.OperateNoteTopicReply
	(*BatchOperateNoteTopicParams)(nil), // 6: note_service.BatchOperateNoteTopicParams
	(*BatchOperateNoteTopicReply)(nil),  // 7: note_service.BatchOperateNoteTopicReply
	(*NoteTopicFilter_Keywords)(nil),    // 8: note_service.NoteTopicFilter.Keywords
	(*common.BetweenInt64)(nil),         // 9: public.BetweenInt64
	(*common.Header)(nil),               // 10: public.Header
	(*common.Pager)(nil),                // 11: public.Pager
	(*common.Sort)(nil),                 // 12: public.Sort
	(*common.ReplyHeader)(nil),          // 13: public.ReplyHeader
	(common.Operation)(0),               // 14: public.Operation
}
var file_note_topic_proto_depIdxs = []int32{
	9,  // 0: note_service.NoteTopicFilter.delete_time_range:type_name -> public.BetweenInt64
	9,  // 1: note_service.NoteTopicFilter.create_time_range:type_name -> public.BetweenInt64
	9,  // 2: note_service.NoteTopicFilter.update_time_range:type_name -> public.BetweenInt64
	8,  // 3: note_service.NoteTopicFilter.keywords:type_name -> note_service.NoteTopicFilter.Keywords
	10, // 4: note_service.ListNoteTopicParams.header:type_name -> public.Header
	11, // 5: note_service.ListNoteTopicParams.pager:type_name -> public.Pager
	12, // 6: note_service.ListNoteTopicParams.sorts:type_name -> public.Sort
	1,  // 7: note_service.ListNoteTopicParams.filter:type_name -> note_service.NoteTopicFilter
	13, // 8: note_service.ListNoteTopicReply.header:type_name -> public.ReplyHeader
	11, // 9: note_service.ListNoteTopicReply.pager:type_name -> public.Pager
	0,  // 10: note_service.ListNoteTopicReply.data:type_name -> note_service.NoteTopic
	10, // 11: note_service.OperateNoteTopicParams.header:type_name -> public.Header
	14, // 12: note_service.OperateNoteTopicParams.operate:type_name -> public.Operation
	0,  // 13: note_service.OperateNoteTopicParams.data:type_name -> note_service.NoteTopic
	13, // 14: note_service.OperateNoteTopicReply.header:type_name -> public.ReplyHeader
	0,  // 15: note_service.OperateNoteTopicReply.data:type_name -> note_service.NoteTopic
	10, // 16: note_service.BatchOperateNoteTopicParams.header:type_name -> public.Header
	14, // 17: note_service.BatchOperateNoteTopicParams.operate:type_name -> public.Operation
	0,  // 18: note_service.BatchOperateNoteTopicParams.data:type_name -> note_service.NoteTopic
	13, // 19: note_service.BatchOperateNoteTopicReply.header:type_name -> public.ReplyHeader
	20, // [20:20] is the sub-list for method output_type
	20, // [20:20] is the sub-list for method input_type
	20, // [20:20] is the sub-list for extension type_name
	20, // [20:20] is the sub-list for extension extendee
	0,  // [0:20] is the sub-list for field type_name
}

func init() { file_note_topic_proto_init() }
func file_note_topic_proto_init() {
	if File_note_topic_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_note_topic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoteTopic); i {
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
		file_note_topic_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoteTopicFilter); i {
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
		file_note_topic_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNoteTopicParams); i {
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
		file_note_topic_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNoteTopicReply); i {
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
		file_note_topic_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperateNoteTopicParams); i {
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
		file_note_topic_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperateNoteTopicReply); i {
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
		file_note_topic_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchOperateNoteTopicParams); i {
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
		file_note_topic_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchOperateNoteTopicReply); i {
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
		file_note_topic_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoteTopicFilter_Keywords); i {
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
			RawDescriptor: file_note_topic_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_note_topic_proto_goTypes,
		DependencyIndexes: file_note_topic_proto_depIdxs,
		MessageInfos:      file_note_topic_proto_msgTypes,
	}.Build()
	File_note_topic_proto = out.File
	file_note_topic_proto_rawDesc = nil
	file_note_topic_proto_goTypes = nil
	file_note_topic_proto_depIdxs = nil
}
