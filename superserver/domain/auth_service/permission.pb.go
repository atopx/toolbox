// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: auth/permission.proto

package auth_service

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

type Permission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AccessId   int32 `protobuf:"varint,2,opt,name=access_id,json=accessId,proto3" json:"access_id,omitempty"`
	RoleId     int32 `protobuf:"varint,3,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	DeleteTime int64 `protobuf:"varint,1001,opt,name=delete_time,json=deleteTime,proto3" json:"delete_time,omitempty"` // 删除时间 时间戳：秒
	CreateTime int64 `protobuf:"varint,1002,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"` // 创建时间 时间戳：秒
	UpdateTime int64 `protobuf:"varint,1003,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"` // 最后更新时间 时间戳：秒
	Creator    int32 `protobuf:"varint,1004,opt,name=creator,proto3" json:"creator,omitempty"`                         // 创建人
	Updater    int32 `protobuf:"varint,1005,opt,name=updater,proto3" json:"updater,omitempty"`                         // 更新人
}

func (x *Permission) Reset() {
	*x = Permission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_permission_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Permission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Permission) ProtoMessage() {}

func (x *Permission) ProtoReflect() protoreflect.Message {
	mi := &file_auth_permission_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Permission.ProtoReflect.Descriptor instead.
func (*Permission) Descriptor() ([]byte, []int) {
	return file_auth_permission_proto_rawDescGZIP(), []int{0}
}

func (x *Permission) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Permission) GetAccessId() int32 {
	if x != nil {
		return x.AccessId
	}
	return 0
}

func (x *Permission) GetRoleId() int32 {
	if x != nil {
		return x.RoleId
	}
	return 0
}

func (x *Permission) GetDeleteTime() int64 {
	if x != nil {
		return x.DeleteTime
	}
	return 0
}

func (x *Permission) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *Permission) GetUpdateTime() int64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

func (x *Permission) GetCreator() int32 {
	if x != nil {
		return x.Creator
	}
	return 0
}

func (x *Permission) GetUpdater() int32 {
	if x != nil {
		return x.Updater
	}
	return 0
}

type PermissionFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids             []int32              `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	AccessIds       []int32              `protobuf:"varint,2,rep,packed,name=access_ids,json=accessIds,proto3" json:"access_ids,omitempty"`
	RoleIds         []int32              `protobuf:"varint,3,rep,packed,name=role_ids,json=roleIds,proto3" json:"role_ids,omitempty"`
	DeleteTimeRange *common.BetweenInt64 `protobuf:"bytes,101,opt,name=delete_time_range,json=deleteTimeRange,proto3" json:"delete_time_range,omitempty"`
	CreateTimeRange *common.BetweenInt64 `protobuf:"bytes,102,opt,name=create_time_range,json=createTimeRange,proto3" json:"create_time_range,omitempty"`
	UpdateTimeRange *common.BetweenInt64 `protobuf:"bytes,103,opt,name=update_time_range,json=updateTimeRange,proto3" json:"update_time_range,omitempty"`
	Creators        []int32              `protobuf:"varint,104,rep,packed,name=creators,proto3" json:"creators,omitempty"`
	Updaters        []int32              `protobuf:"varint,105,rep,packed,name=updaters,proto3" json:"updaters,omitempty"`
}

func (x *PermissionFilter) Reset() {
	*x = PermissionFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_permission_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PermissionFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionFilter) ProtoMessage() {}

func (x *PermissionFilter) ProtoReflect() protoreflect.Message {
	mi := &file_auth_permission_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionFilter.ProtoReflect.Descriptor instead.
func (*PermissionFilter) Descriptor() ([]byte, []int) {
	return file_auth_permission_proto_rawDescGZIP(), []int{1}
}

func (x *PermissionFilter) GetIds() []int32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *PermissionFilter) GetAccessIds() []int32 {
	if x != nil {
		return x.AccessIds
	}
	return nil
}

func (x *PermissionFilter) GetRoleIds() []int32 {
	if x != nil {
		return x.RoleIds
	}
	return nil
}

func (x *PermissionFilter) GetDeleteTimeRange() *common.BetweenInt64 {
	if x != nil {
		return x.DeleteTimeRange
	}
	return nil
}

func (x *PermissionFilter) GetCreateTimeRange() *common.BetweenInt64 {
	if x != nil {
		return x.CreateTimeRange
	}
	return nil
}

func (x *PermissionFilter) GetUpdateTimeRange() *common.BetweenInt64 {
	if x != nil {
		return x.UpdateTimeRange
	}
	return nil
}

func (x *PermissionFilter) GetCreators() []int32 {
	if x != nil {
		return x.Creators
	}
	return nil
}

func (x *PermissionFilter) GetUpdaters() []int32 {
	if x != nil {
		return x.Updaters
	}
	return nil
}

type ListPermissionParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *common.Header    `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Pager  *common.Pager     `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Sorts  []*common.Sort    `protobuf:"bytes,3,rep,name=sorts,proto3" json:"sorts,omitempty"`
	Filter *PermissionFilter `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *ListPermissionParams) Reset() {
	*x = ListPermissionParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_permission_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPermissionParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPermissionParams) ProtoMessage() {}

func (x *ListPermissionParams) ProtoReflect() protoreflect.Message {
	mi := &file_auth_permission_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPermissionParams.ProtoReflect.Descriptor instead.
func (*ListPermissionParams) Descriptor() ([]byte, []int) {
	return file_auth_permission_proto_rawDescGZIP(), []int{2}
}

func (x *ListPermissionParams) GetHeader() *common.Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *ListPermissionParams) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *ListPermissionParams) GetSorts() []*common.Sort {
	if x != nil {
		return x.Sorts
	}
	return nil
}

func (x *ListPermissionParams) GetFilter() *PermissionFilter {
	if x != nil {
		return x.Filter
	}
	return nil
}

type ListPermissionReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *common.ReplyHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Pager  *common.Pager       `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Data   []*Permission       `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ListPermissionReply) Reset() {
	*x = ListPermissionReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_permission_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPermissionReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPermissionReply) ProtoMessage() {}

func (x *ListPermissionReply) ProtoReflect() protoreflect.Message {
	mi := &file_auth_permission_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPermissionReply.ProtoReflect.Descriptor instead.
func (*ListPermissionReply) Descriptor() ([]byte, []int) {
	return file_auth_permission_proto_rawDescGZIP(), []int{3}
}

func (x *ListPermissionReply) GetHeader() *common.ReplyHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *ListPermissionReply) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *ListPermissionReply) GetData() []*Permission {
	if x != nil {
		return x.Data
	}
	return nil
}

type OperatePermissionParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header  *common.Header   `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Operate common.Operation `protobuf:"varint,2,opt,name=operate,proto3,enum=public.Operation" json:"operate,omitempty"`
	Data    *Permission      `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *OperatePermissionParams) Reset() {
	*x = OperatePermissionParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_permission_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperatePermissionParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperatePermissionParams) ProtoMessage() {}

func (x *OperatePermissionParams) ProtoReflect() protoreflect.Message {
	mi := &file_auth_permission_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperatePermissionParams.ProtoReflect.Descriptor instead.
func (*OperatePermissionParams) Descriptor() ([]byte, []int) {
	return file_auth_permission_proto_rawDescGZIP(), []int{4}
}

func (x *OperatePermissionParams) GetHeader() *common.Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *OperatePermissionParams) GetOperate() common.Operation {
	if x != nil {
		return x.Operate
	}
	return common.Operation(0)
}

func (x *OperatePermissionParams) GetData() *Permission {
	if x != nil {
		return x.Data
	}
	return nil
}

type OperatePermissionReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *common.ReplyHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Data   *Permission         `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *OperatePermissionReply) Reset() {
	*x = OperatePermissionReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_permission_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperatePermissionReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperatePermissionReply) ProtoMessage() {}

func (x *OperatePermissionReply) ProtoReflect() protoreflect.Message {
	mi := &file_auth_permission_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperatePermissionReply.ProtoReflect.Descriptor instead.
func (*OperatePermissionReply) Descriptor() ([]byte, []int) {
	return file_auth_permission_proto_rawDescGZIP(), []int{5}
}

func (x *OperatePermissionReply) GetHeader() *common.ReplyHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *OperatePermissionReply) GetData() *Permission {
	if x != nil {
		return x.Data
	}
	return nil
}

type BatchOperatePermissionParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header  *common.Header   `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Operate common.Operation `protobuf:"varint,2,opt,name=operate,proto3,enum=public.Operation" json:"operate,omitempty"`
	Data    []*Permission    `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *BatchOperatePermissionParams) Reset() {
	*x = BatchOperatePermissionParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_permission_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchOperatePermissionParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchOperatePermissionParams) ProtoMessage() {}

func (x *BatchOperatePermissionParams) ProtoReflect() protoreflect.Message {
	mi := &file_auth_permission_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchOperatePermissionParams.ProtoReflect.Descriptor instead.
func (*BatchOperatePermissionParams) Descriptor() ([]byte, []int) {
	return file_auth_permission_proto_rawDescGZIP(), []int{6}
}

func (x *BatchOperatePermissionParams) GetHeader() *common.Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *BatchOperatePermissionParams) GetOperate() common.Operation {
	if x != nil {
		return x.Operate
	}
	return common.Operation(0)
}

func (x *BatchOperatePermissionParams) GetData() []*Permission {
	if x != nil {
		return x.Data
	}
	return nil
}

type BatchOperatePermissionReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *common.ReplyHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
}

func (x *BatchOperatePermissionReply) Reset() {
	*x = BatchOperatePermissionReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_permission_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchOperatePermissionReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchOperatePermissionReply) ProtoMessage() {}

func (x *BatchOperatePermissionReply) ProtoReflect() protoreflect.Message {
	mi := &file_auth_permission_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchOperatePermissionReply.ProtoReflect.Descriptor instead.
func (*BatchOperatePermissionReply) Descriptor() ([]byte, []int) {
	return file_auth_permission_proto_rawDescGZIP(), []int{7}
}

func (x *BatchOperatePermissionReply) GetHeader() *common.ReplyHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

var File_auth_permission_proto protoreflect.FileDescriptor

var file_auth_permission_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x13, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xee, 0x01, 0x0a, 0x0a, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0xe9,
	0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0xea, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0xeb, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72,
	0x18, 0xec, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72,
	0x12, 0x19, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x18, 0xed, 0x07, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x22, 0xdc, 0x02, 0x0a, 0x10,
	0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x03, 0x69,
	0x64, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x09, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x49, 0x64,
	0x73, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x05, 0x52, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x73, 0x12, 0x40, 0x0a, 0x11,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x72, 0x61, 0x6e, 0x67,
	0x65, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x2e, 0x42, 0x65, 0x74, 0x77, 0x65, 0x65, 0x6e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x52, 0x0f, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x40,
	0x0a, 0x11, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x72, 0x61,
	0x6e, 0x67, 0x65, 0x18, 0x66, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x2e, 0x42, 0x65, 0x74, 0x77, 0x65, 0x65, 0x6e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x52,
	0x0f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65,
	0x12, 0x40, 0x0a, 0x11, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f,
	0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x67, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x2e, 0x42, 0x65, 0x74, 0x77, 0x65, 0x65, 0x6e, 0x49, 0x6e, 0x74, 0x36,
	0x34, 0x52, 0x0f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e,
	0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x68,
	0x20, 0x03, 0x28, 0x05, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x73, 0x18, 0x69, 0x20, 0x03, 0x28, 0x05,
	0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x73, 0x22, 0xbf, 0x01, 0x0a, 0x14, 0x4c,
	0x69, 0x73, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x12, 0x26, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x05, 0x70,
	0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72,
	0x12, 0x22, 0x0a, 0x05, 0x73, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x52, 0x05, 0x73,
	0x6f, 0x72, 0x74, 0x73, 0x12, 0x36, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x95, 0x01, 0x0a,
	0x13, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x2b, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52,
	0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x9c, 0x01, 0x0a, 0x17, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x12, 0x26, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x07, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x12, 0x2c, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x73, 0x0a, 0x16, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2b, 0x0a,
	0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xa1, 0x01, 0x0a, 0x1c, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x26, 0x0a, 0x06, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x12, 0x2b, 0x0a, 0x07, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x11, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x12, 0x2c,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x4a, 0x0a, 0x1b,
	0x42, 0x61, 0x74, 0x63, 0x68, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2b, 0x0a, 0x06, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x42, 0x21, 0x5a, 0x1f, 0x73, 0x75, 0x70, 0x65,
	0x72, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_auth_permission_proto_rawDescOnce sync.Once
	file_auth_permission_proto_rawDescData = file_auth_permission_proto_rawDesc
)

func file_auth_permission_proto_rawDescGZIP() []byte {
	file_auth_permission_proto_rawDescOnce.Do(func() {
		file_auth_permission_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_permission_proto_rawDescData)
	})
	return file_auth_permission_proto_rawDescData
}

var file_auth_permission_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_auth_permission_proto_goTypes = []interface{}{
	(*Permission)(nil),                   // 0: auth_service.Permission
	(*PermissionFilter)(nil),             // 1: auth_service.PermissionFilter
	(*ListPermissionParams)(nil),         // 2: auth_service.ListPermissionParams
	(*ListPermissionReply)(nil),          // 3: auth_service.ListPermissionReply
	(*OperatePermissionParams)(nil),      // 4: auth_service.OperatePermissionParams
	(*OperatePermissionReply)(nil),       // 5: auth_service.OperatePermissionReply
	(*BatchOperatePermissionParams)(nil), // 6: auth_service.BatchOperatePermissionParams
	(*BatchOperatePermissionReply)(nil),  // 7: auth_service.BatchOperatePermissionReply
	(*common.BetweenInt64)(nil),          // 8: public.BetweenInt64
	(*common.Header)(nil),                // 9: public.Header
	(*common.Pager)(nil),                 // 10: public.Pager
	(*common.Sort)(nil),                  // 11: public.Sort
	(*common.ReplyHeader)(nil),           // 12: public.ReplyHeader
	(common.Operation)(0),                // 13: public.Operation
}
var file_auth_permission_proto_depIdxs = []int32{
	8,  // 0: auth_service.PermissionFilter.delete_time_range:type_name -> public.BetweenInt64
	8,  // 1: auth_service.PermissionFilter.create_time_range:type_name -> public.BetweenInt64
	8,  // 2: auth_service.PermissionFilter.update_time_range:type_name -> public.BetweenInt64
	9,  // 3: auth_service.ListPermissionParams.header:type_name -> public.Header
	10, // 4: auth_service.ListPermissionParams.pager:type_name -> public.Pager
	11, // 5: auth_service.ListPermissionParams.sorts:type_name -> public.Sort
	1,  // 6: auth_service.ListPermissionParams.filter:type_name -> auth_service.PermissionFilter
	12, // 7: auth_service.ListPermissionReply.header:type_name -> public.ReplyHeader
	10, // 8: auth_service.ListPermissionReply.pager:type_name -> public.Pager
	0,  // 9: auth_service.ListPermissionReply.data:type_name -> auth_service.Permission
	9,  // 10: auth_service.OperatePermissionParams.header:type_name -> public.Header
	13, // 11: auth_service.OperatePermissionParams.operate:type_name -> public.Operation
	0,  // 12: auth_service.OperatePermissionParams.data:type_name -> auth_service.Permission
	12, // 13: auth_service.OperatePermissionReply.header:type_name -> public.ReplyHeader
	0,  // 14: auth_service.OperatePermissionReply.data:type_name -> auth_service.Permission
	9,  // 15: auth_service.BatchOperatePermissionParams.header:type_name -> public.Header
	13, // 16: auth_service.BatchOperatePermissionParams.operate:type_name -> public.Operation
	0,  // 17: auth_service.BatchOperatePermissionParams.data:type_name -> auth_service.Permission
	12, // 18: auth_service.BatchOperatePermissionReply.header:type_name -> public.ReplyHeader
	19, // [19:19] is the sub-list for method output_type
	19, // [19:19] is the sub-list for method input_type
	19, // [19:19] is the sub-list for extension type_name
	19, // [19:19] is the sub-list for extension extendee
	0,  // [0:19] is the sub-list for field type_name
}

func init() { file_auth_permission_proto_init() }
func file_auth_permission_proto_init() {
	if File_auth_permission_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auth_permission_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Permission); i {
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
		file_auth_permission_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PermissionFilter); i {
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
		file_auth_permission_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPermissionParams); i {
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
		file_auth_permission_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPermissionReply); i {
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
		file_auth_permission_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperatePermissionParams); i {
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
		file_auth_permission_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperatePermissionReply); i {
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
		file_auth_permission_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchOperatePermissionParams); i {
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
		file_auth_permission_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchOperatePermissionReply); i {
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
			RawDescriptor: file_auth_permission_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_auth_permission_proto_goTypes,
		DependencyIndexes: file_auth_permission_proto_depIdxs,
		MessageInfos:      file_auth_permission_proto_msgTypes,
	}.Build()
	File_auth_permission_proto = out.File
	file_auth_permission_proto_rawDesc = nil
	file_auth_permission_proto_goTypes = nil
	file_auth_permission_proto_depIdxs = nil
}
