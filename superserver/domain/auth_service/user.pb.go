// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.2
// source: auth/user.proto

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
	return file_auth_user_proto_enumTypes[0].Descriptor()
}

func (UserStatus) Type() protoreflect.EnumType {
	return &file_auth_user_proto_enumTypes[0]
}

func (x UserStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserStatus.Descriptor instead.
func (UserStatus) EnumDescriptor() ([]byte, []int) {
	return file_auth_user_proto_rawDescGZIP(), []int{0}
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username   string     `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password   string     `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Status     UserStatus `protobuf:"varint,4,opt,name=status,proto3,enum=auth_service.UserStatus" json:"status,omitempty"`
	DeleteTime int64      `protobuf:"varint,1001,opt,name=delete_time,json=deleteTime,proto3" json:"delete_time,omitempty"`
	CreateTime int64      `protobuf:"varint,1002,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime int64      `protobuf:"varint,1003,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_auth_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_auth_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *User) GetStatus() UserStatus {
	if x != nil {
		return x.Status
	}
	return UserStatus_USER_NORMAL
}

func (x *User) GetDeleteTime() int64 {
	if x != nil {
		return x.DeleteTime
	}
	return 0
}

func (x *User) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *User) GetUpdateTime() int64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

type UserFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids             []int32              `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	Usernames       []string             `protobuf:"bytes,2,rep,name=usernames,proto3" json:"usernames,omitempty"`
	States          []UserStatus         `protobuf:"varint,3,rep,packed,name=states,proto3,enum=auth_service.UserStatus" json:"states,omitempty"`
	DeleteTimeRange *common.BetweenInt64 `protobuf:"bytes,101,opt,name=delete_time_range,json=deleteTimeRange,proto3" json:"delete_time_range,omitempty"`
	CreateTimeRange *common.BetweenInt64 `protobuf:"bytes,102,opt,name=create_time_range,json=createTimeRange,proto3" json:"create_time_range,omitempty"`
	UpdateTimeRange *common.BetweenInt64 `protobuf:"bytes,103,opt,name=update_time_range,json=updateTimeRange,proto3" json:"update_time_range,omitempty"`
}

func (x *UserFilter) Reset() {
	*x = UserFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserFilter) ProtoMessage() {}

func (x *UserFilter) ProtoReflect() protoreflect.Message {
	mi := &file_auth_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserFilter.ProtoReflect.Descriptor instead.
func (*UserFilter) Descriptor() ([]byte, []int) {
	return file_auth_user_proto_rawDescGZIP(), []int{1}
}

func (x *UserFilter) GetIds() []int32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *UserFilter) GetUsernames() []string {
	if x != nil {
		return x.Usernames
	}
	return nil
}

func (x *UserFilter) GetStates() []UserStatus {
	if x != nil {
		return x.States
	}
	return nil
}

func (x *UserFilter) GetDeleteTimeRange() *common.BetweenInt64 {
	if x != nil {
		return x.DeleteTimeRange
	}
	return nil
}

func (x *UserFilter) GetCreateTimeRange() *common.BetweenInt64 {
	if x != nil {
		return x.CreateTimeRange
	}
	return nil
}

func (x *UserFilter) GetUpdateTimeRange() *common.BetweenInt64 {
	if x != nil {
		return x.UpdateTimeRange
	}
	return nil
}

type ListUserParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *common.Header `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Pager  *common.Pager  `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Sorts  []*common.Sort `protobuf:"bytes,3,rep,name=sorts,proto3" json:"sorts,omitempty"`
	Filter *UserFilter    `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *ListUserParams) Reset() {
	*x = ListUserParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUserParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserParams) ProtoMessage() {}

func (x *ListUserParams) ProtoReflect() protoreflect.Message {
	mi := &file_auth_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserParams.ProtoReflect.Descriptor instead.
func (*ListUserParams) Descriptor() ([]byte, []int) {
	return file_auth_user_proto_rawDescGZIP(), []int{2}
}

func (x *ListUserParams) GetHeader() *common.Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *ListUserParams) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *ListUserParams) GetSorts() []*common.Sort {
	if x != nil {
		return x.Sorts
	}
	return nil
}

func (x *ListUserParams) GetFilter() *UserFilter {
	if x != nil {
		return x.Filter
	}
	return nil
}

type ListUserReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *common.ReplyHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Pager  *common.Pager       `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Data   []*User             `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ListUserReply) Reset() {
	*x = ListUserReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUserReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserReply) ProtoMessage() {}

func (x *ListUserReply) ProtoReflect() protoreflect.Message {
	mi := &file_auth_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserReply.ProtoReflect.Descriptor instead.
func (*ListUserReply) Descriptor() ([]byte, []int) {
	return file_auth_user_proto_rawDescGZIP(), []int{3}
}

func (x *ListUserReply) GetHeader() *common.ReplyHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *ListUserReply) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *ListUserReply) GetData() []*User {
	if x != nil {
		return x.Data
	}
	return nil
}

type OperateUserParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header  *common.Header   `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Operate common.Operation `protobuf:"varint,2,opt,name=operate,proto3,enum=public.Operation" json:"operate,omitempty"`
	Data    *User            `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *OperateUserParams) Reset() {
	*x = OperateUserParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperateUserParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperateUserParams) ProtoMessage() {}

func (x *OperateUserParams) ProtoReflect() protoreflect.Message {
	mi := &file_auth_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperateUserParams.ProtoReflect.Descriptor instead.
func (*OperateUserParams) Descriptor() ([]byte, []int) {
	return file_auth_user_proto_rawDescGZIP(), []int{4}
}

func (x *OperateUserParams) GetHeader() *common.Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *OperateUserParams) GetOperate() common.Operation {
	if x != nil {
		return x.Operate
	}
	return common.Operation(0)
}

func (x *OperateUserParams) GetData() *User {
	if x != nil {
		return x.Data
	}
	return nil
}

type OperateUserReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *common.ReplyHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Data   *User               `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *OperateUserReply) Reset() {
	*x = OperateUserReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperateUserReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperateUserReply) ProtoMessage() {}

func (x *OperateUserReply) ProtoReflect() protoreflect.Message {
	mi := &file_auth_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperateUserReply.ProtoReflect.Descriptor instead.
func (*OperateUserReply) Descriptor() ([]byte, []int) {
	return file_auth_user_proto_rawDescGZIP(), []int{5}
}

func (x *OperateUserReply) GetHeader() *common.ReplyHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *OperateUserReply) GetData() *User {
	if x != nil {
		return x.Data
	}
	return nil
}

type BatchOperateUserParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header  *common.Header   `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Operate common.Operation `protobuf:"varint,2,opt,name=operate,proto3,enum=public.Operation" json:"operate,omitempty"`
	Data    []*User          `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *BatchOperateUserParams) Reset() {
	*x = BatchOperateUserParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchOperateUserParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchOperateUserParams) ProtoMessage() {}

func (x *BatchOperateUserParams) ProtoReflect() protoreflect.Message {
	mi := &file_auth_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchOperateUserParams.ProtoReflect.Descriptor instead.
func (*BatchOperateUserParams) Descriptor() ([]byte, []int) {
	return file_auth_user_proto_rawDescGZIP(), []int{6}
}

func (x *BatchOperateUserParams) GetHeader() *common.Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *BatchOperateUserParams) GetOperate() common.Operation {
	if x != nil {
		return x.Operate
	}
	return common.Operation(0)
}

func (x *BatchOperateUserParams) GetData() []*User {
	if x != nil {
		return x.Data
	}
	return nil
}

type BatchOperateUserReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *common.ReplyHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
}

func (x *BatchOperateUserReply) Reset() {
	*x = BatchOperateUserReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchOperateUserReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchOperateUserReply) ProtoMessage() {}

func (x *BatchOperateUserReply) ProtoReflect() protoreflect.Message {
	mi := &file_auth_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchOperateUserReply.ProtoReflect.Descriptor instead.
func (*BatchOperateUserReply) Descriptor() ([]byte, []int) {
	return file_auth_user_proto_rawDescGZIP(), []int{7}
}

func (x *BatchOperateUserReply) GetHeader() *common.ReplyHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

var File_auth_user_proto protoreflect.FileDescriptor

var file_auth_user_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a,
	0x13, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe6, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x30, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0xe9, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0xea, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0xeb, 0x07, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xb4, 0x02,
	0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03,
	0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x1c,
	0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x30, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x12, 0x40,
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
	0x61, 0x6e, 0x67, 0x65, 0x22, 0xb3, 0x01, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x26, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12,
	0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70,
	0x61, 0x67, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x05, 0x73, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x53, 0x6f, 0x72,
	0x74, 0x52, 0x05, 0x73, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x30, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x89, 0x01, 0x0a, 0x0d, 0x4c,
	0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2b, 0x0a, 0x06,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x26,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x90, 0x01, 0x0a, 0x11, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x26, 0x0a, 0x06,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x07, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x12, 0x26, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x67, 0x0a, 0x10, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2b, 0x0a,
	0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x95, 0x01, 0x0a, 0x16, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x26, 0x0a,
	0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x07, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x44, 0x0a, 0x15, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x2b, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x2a, 0x4e, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0f,
	0x0a, 0x0b, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10, 0x00, 0x12,
	0x18, 0x0a, 0x14, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x44,
	0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x55, 0x53, 0x45,
	0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x42, 0x4c, 0x41, 0x43, 0x4b, 0x10, 0x02,
	0x42, 0x21, 0x5a, 0x1f, 0x73, 0x75, 0x70, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auth_user_proto_rawDescOnce sync.Once
	file_auth_user_proto_rawDescData = file_auth_user_proto_rawDesc
)

func file_auth_user_proto_rawDescGZIP() []byte {
	file_auth_user_proto_rawDescOnce.Do(func() {
		file_auth_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_user_proto_rawDescData)
	})
	return file_auth_user_proto_rawDescData
}

var file_auth_user_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_auth_user_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_auth_user_proto_goTypes = []interface{}{
	(UserStatus)(0),                // 0: auth_service.UserStatus
	(*User)(nil),                   // 1: auth_service.User
	(*UserFilter)(nil),             // 2: auth_service.UserFilter
	(*ListUserParams)(nil),         // 3: auth_service.ListUserParams
	(*ListUserReply)(nil),          // 4: auth_service.ListUserReply
	(*OperateUserParams)(nil),      // 5: auth_service.OperateUserParams
	(*OperateUserReply)(nil),       // 6: auth_service.OperateUserReply
	(*BatchOperateUserParams)(nil), // 7: auth_service.BatchOperateUserParams
	(*BatchOperateUserReply)(nil),  // 8: auth_service.BatchOperateUserReply
	(*common.BetweenInt64)(nil),    // 9: public.BetweenInt64
	(*common.Header)(nil),          // 10: public.Header
	(*common.Pager)(nil),           // 11: public.Pager
	(*common.Sort)(nil),            // 12: public.Sort
	(*common.ReplyHeader)(nil),     // 13: public.ReplyHeader
	(common.Operation)(0),          // 14: public.Operation
}
var file_auth_user_proto_depIdxs = []int32{
	0,  // 0: auth_service.User.status:type_name -> auth_service.UserStatus
	0,  // 1: auth_service.UserFilter.states:type_name -> auth_service.UserStatus
	9,  // 2: auth_service.UserFilter.delete_time_range:type_name -> public.BetweenInt64
	9,  // 3: auth_service.UserFilter.create_time_range:type_name -> public.BetweenInt64
	9,  // 4: auth_service.UserFilter.update_time_range:type_name -> public.BetweenInt64
	10, // 5: auth_service.ListUserParams.header:type_name -> public.Header
	11, // 6: auth_service.ListUserParams.pager:type_name -> public.Pager
	12, // 7: auth_service.ListUserParams.sorts:type_name -> public.Sort
	2,  // 8: auth_service.ListUserParams.filter:type_name -> auth_service.UserFilter
	13, // 9: auth_service.ListUserReply.header:type_name -> public.ReplyHeader
	11, // 10: auth_service.ListUserReply.pager:type_name -> public.Pager
	1,  // 11: auth_service.ListUserReply.data:type_name -> auth_service.User
	10, // 12: auth_service.OperateUserParams.header:type_name -> public.Header
	14, // 13: auth_service.OperateUserParams.operate:type_name -> public.Operation
	1,  // 14: auth_service.OperateUserParams.data:type_name -> auth_service.User
	13, // 15: auth_service.OperateUserReply.header:type_name -> public.ReplyHeader
	1,  // 16: auth_service.OperateUserReply.data:type_name -> auth_service.User
	10, // 17: auth_service.BatchOperateUserParams.header:type_name -> public.Header
	14, // 18: auth_service.BatchOperateUserParams.operate:type_name -> public.Operation
	1,  // 19: auth_service.BatchOperateUserParams.data:type_name -> auth_service.User
	13, // 20: auth_service.BatchOperateUserReply.header:type_name -> public.ReplyHeader
	21, // [21:21] is the sub-list for method output_type
	21, // [21:21] is the sub-list for method input_type
	21, // [21:21] is the sub-list for extension type_name
	21, // [21:21] is the sub-list for extension extendee
	0,  // [0:21] is the sub-list for field type_name
}

func init() { file_auth_user_proto_init() }
func file_auth_user_proto_init() {
	if File_auth_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auth_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_auth_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserFilter); i {
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
		file_auth_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUserParams); i {
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
		file_auth_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUserReply); i {
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
		file_auth_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperateUserParams); i {
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
		file_auth_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperateUserReply); i {
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
		file_auth_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchOperateUserParams); i {
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
		file_auth_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchOperateUserReply); i {
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
			RawDescriptor: file_auth_user_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_auth_user_proto_goTypes,
		DependencyIndexes: file_auth_user_proto_depIdxs,
		EnumInfos:         file_auth_user_proto_enumTypes,
		MessageInfos:      file_auth_user_proto_msgTypes,
	}.Build()
	File_auth_user_proto = out.File
	file_auth_user_proto_rawDesc = nil
	file_auth_user_proto_goTypes = nil
	file_auth_user_proto_depIdxs = nil
}
