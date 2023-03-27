// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: auth/token.proto

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

type AuthToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId       int32  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	AccessToken  string `protobuf:"bytes,3,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	RefreshToken string `protobuf:"bytes,4,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	IssuedTime   int64  `protobuf:"varint,5,opt,name=issued_time,json=issuedTime,proto3" json:"issued_time,omitempty"`
	ExpireTime   int64  `protobuf:"varint,6,opt,name=expire_time,json=expireTime,proto3" json:"expire_time,omitempty"`
	DeleteTime   int64  `protobuf:"varint,1001,opt,name=delete_time,json=deleteTime,proto3" json:"delete_time,omitempty"`
	CreateTime   int64  `protobuf:"varint,1002,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime   int64  `protobuf:"varint,1003,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *AuthToken) Reset() {
	*x = AuthToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_token_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthToken) ProtoMessage() {}

func (x *AuthToken) ProtoReflect() protoreflect.Message {
	mi := &file_auth_token_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthToken.ProtoReflect.Descriptor instead.
func (*AuthToken) Descriptor() ([]byte, []int) {
	return file_auth_token_proto_rawDescGZIP(), []int{0}
}

func (x *AuthToken) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AuthToken) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *AuthToken) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *AuthToken) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *AuthToken) GetIssuedTime() int64 {
	if x != nil {
		return x.IssuedTime
	}
	return 0
}

func (x *AuthToken) GetExpireTime() int64 {
	if x != nil {
		return x.ExpireTime
	}
	return 0
}

func (x *AuthToken) GetDeleteTime() int64 {
	if x != nil {
		return x.DeleteTime
	}
	return 0
}

func (x *AuthToken) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *AuthToken) GetUpdateTime() int64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

type AuthTokenFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids             []int32              `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	UserIds         []int32              `protobuf:"varint,2,rep,packed,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty"`
	AccessTokens    []string             `protobuf:"bytes,3,rep,name=access_tokens,json=accessTokens,proto3" json:"access_tokens,omitempty"`
	RefreshTokens   []string             `protobuf:"bytes,4,rep,name=refresh_tokens,json=refreshTokens,proto3" json:"refresh_tokens,omitempty"`
	IssuedTimeRange *common.BetweenInt64 `protobuf:"bytes,21,opt,name=issued_time_range,json=issuedTimeRange,proto3" json:"issued_time_range,omitempty"`
	ExpireTimeRange *common.BetweenInt64 `protobuf:"bytes,22,opt,name=expire_time_range,json=expireTimeRange,proto3" json:"expire_time_range,omitempty"`
	DeleteTimeRange *common.BetweenInt64 `protobuf:"bytes,101,opt,name=delete_time_range,json=deleteTimeRange,proto3" json:"delete_time_range,omitempty"`
	CreateTimeRange *common.BetweenInt64 `protobuf:"bytes,102,opt,name=create_time_range,json=createTimeRange,proto3" json:"create_time_range,omitempty"`
	UpdateTimeRange *common.BetweenInt64 `protobuf:"bytes,103,opt,name=update_time_range,json=updateTimeRange,proto3" json:"update_time_range,omitempty"`
}

func (x *AuthTokenFilter) Reset() {
	*x = AuthTokenFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_token_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthTokenFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthTokenFilter) ProtoMessage() {}

func (x *AuthTokenFilter) ProtoReflect() protoreflect.Message {
	mi := &file_auth_token_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthTokenFilter.ProtoReflect.Descriptor instead.
func (*AuthTokenFilter) Descriptor() ([]byte, []int) {
	return file_auth_token_proto_rawDescGZIP(), []int{1}
}

func (x *AuthTokenFilter) GetIds() []int32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *AuthTokenFilter) GetUserIds() []int32 {
	if x != nil {
		return x.UserIds
	}
	return nil
}

func (x *AuthTokenFilter) GetAccessTokens() []string {
	if x != nil {
		return x.AccessTokens
	}
	return nil
}

func (x *AuthTokenFilter) GetRefreshTokens() []string {
	if x != nil {
		return x.RefreshTokens
	}
	return nil
}

func (x *AuthTokenFilter) GetIssuedTimeRange() *common.BetweenInt64 {
	if x != nil {
		return x.IssuedTimeRange
	}
	return nil
}

func (x *AuthTokenFilter) GetExpireTimeRange() *common.BetweenInt64 {
	if x != nil {
		return x.ExpireTimeRange
	}
	return nil
}

func (x *AuthTokenFilter) GetDeleteTimeRange() *common.BetweenInt64 {
	if x != nil {
		return x.DeleteTimeRange
	}
	return nil
}

func (x *AuthTokenFilter) GetCreateTimeRange() *common.BetweenInt64 {
	if x != nil {
		return x.CreateTimeRange
	}
	return nil
}

func (x *AuthTokenFilter) GetUpdateTimeRange() *common.BetweenInt64 {
	if x != nil {
		return x.UpdateTimeRange
	}
	return nil
}

type ListAuthTokenParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *common.Header   `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Pager  *common.Pager    `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Sorts  []*common.Sort   `protobuf:"bytes,3,rep,name=sorts,proto3" json:"sorts,omitempty"`
	Filter *AuthTokenFilter `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *ListAuthTokenParams) Reset() {
	*x = ListAuthTokenParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_token_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAuthTokenParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAuthTokenParams) ProtoMessage() {}

func (x *ListAuthTokenParams) ProtoReflect() protoreflect.Message {
	mi := &file_auth_token_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAuthTokenParams.ProtoReflect.Descriptor instead.
func (*ListAuthTokenParams) Descriptor() ([]byte, []int) {
	return file_auth_token_proto_rawDescGZIP(), []int{2}
}

func (x *ListAuthTokenParams) GetHeader() *common.Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *ListAuthTokenParams) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *ListAuthTokenParams) GetSorts() []*common.Sort {
	if x != nil {
		return x.Sorts
	}
	return nil
}

func (x *ListAuthTokenParams) GetFilter() *AuthTokenFilter {
	if x != nil {
		return x.Filter
	}
	return nil
}

type ListAuthTokenReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *common.ReplyHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Pager  *common.Pager       `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	Data   []*AuthToken        `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ListAuthTokenReply) Reset() {
	*x = ListAuthTokenReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_token_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAuthTokenReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAuthTokenReply) ProtoMessage() {}

func (x *ListAuthTokenReply) ProtoReflect() protoreflect.Message {
	mi := &file_auth_token_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAuthTokenReply.ProtoReflect.Descriptor instead.
func (*ListAuthTokenReply) Descriptor() ([]byte, []int) {
	return file_auth_token_proto_rawDescGZIP(), []int{3}
}

func (x *ListAuthTokenReply) GetHeader() *common.ReplyHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *ListAuthTokenReply) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *ListAuthTokenReply) GetData() []*AuthToken {
	if x != nil {
		return x.Data
	}
	return nil
}

type OperateAuthTokenParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header  *common.Header   `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Operate common.Operation `protobuf:"varint,2,opt,name=operate,proto3,enum=public.Operation" json:"operate,omitempty"`
	Data    *AuthToken       `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *OperateAuthTokenParams) Reset() {
	*x = OperateAuthTokenParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_token_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperateAuthTokenParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperateAuthTokenParams) ProtoMessage() {}

func (x *OperateAuthTokenParams) ProtoReflect() protoreflect.Message {
	mi := &file_auth_token_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperateAuthTokenParams.ProtoReflect.Descriptor instead.
func (*OperateAuthTokenParams) Descriptor() ([]byte, []int) {
	return file_auth_token_proto_rawDescGZIP(), []int{4}
}

func (x *OperateAuthTokenParams) GetHeader() *common.Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *OperateAuthTokenParams) GetOperate() common.Operation {
	if x != nil {
		return x.Operate
	}
	return common.Operation(0)
}

func (x *OperateAuthTokenParams) GetData() *AuthToken {
	if x != nil {
		return x.Data
	}
	return nil
}

type OperateAuthTokenReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *common.ReplyHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Data   *AuthToken          `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *OperateAuthTokenReply) Reset() {
	*x = OperateAuthTokenReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_token_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperateAuthTokenReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperateAuthTokenReply) ProtoMessage() {}

func (x *OperateAuthTokenReply) ProtoReflect() protoreflect.Message {
	mi := &file_auth_token_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperateAuthTokenReply.ProtoReflect.Descriptor instead.
func (*OperateAuthTokenReply) Descriptor() ([]byte, []int) {
	return file_auth_token_proto_rawDescGZIP(), []int{5}
}

func (x *OperateAuthTokenReply) GetHeader() *common.ReplyHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *OperateAuthTokenReply) GetData() *AuthToken {
	if x != nil {
		return x.Data
	}
	return nil
}

type BatchOperateAuthTokenParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header  *common.Header   `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Operate common.Operation `protobuf:"varint,2,opt,name=operate,proto3,enum=public.Operation" json:"operate,omitempty"`
	Data    []*AuthToken     `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *BatchOperateAuthTokenParams) Reset() {
	*x = BatchOperateAuthTokenParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_token_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchOperateAuthTokenParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchOperateAuthTokenParams) ProtoMessage() {}

func (x *BatchOperateAuthTokenParams) ProtoReflect() protoreflect.Message {
	mi := &file_auth_token_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchOperateAuthTokenParams.ProtoReflect.Descriptor instead.
func (*BatchOperateAuthTokenParams) Descriptor() ([]byte, []int) {
	return file_auth_token_proto_rawDescGZIP(), []int{6}
}

func (x *BatchOperateAuthTokenParams) GetHeader() *common.Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *BatchOperateAuthTokenParams) GetOperate() common.Operation {
	if x != nil {
		return x.Operate
	}
	return common.Operation(0)
}

func (x *BatchOperateAuthTokenParams) GetData() []*AuthToken {
	if x != nil {
		return x.Data
	}
	return nil
}

type BatchOperateAuthTokenReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *common.ReplyHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
}

func (x *BatchOperateAuthTokenReply) Reset() {
	*x = BatchOperateAuthTokenReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_token_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchOperateAuthTokenReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchOperateAuthTokenReply) ProtoMessage() {}

func (x *BatchOperateAuthTokenReply) ProtoReflect() protoreflect.Message {
	mi := &file_auth_token_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchOperateAuthTokenReply.ProtoReflect.Descriptor instead.
func (*BatchOperateAuthTokenReply) Descriptor() ([]byte, []int) {
	return file_auth_token_proto_rawDescGZIP(), []int{7}
}

func (x *BatchOperateAuthTokenReply) GetHeader() *common.ReplyHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

var File_auth_token_proto protoreflect.FileDescriptor

var file_auth_token_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0c, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x1a, 0x13, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa4, 0x02, 0x0a, 0x09, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x73, 0x75, 0x65, 0x64, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x69, 0x73, 0x73, 0x75, 0x65,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0xe9, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0xea, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0xeb, 0x07, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xd4, 0x03, 0x0a,
	0x0f, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x03, 0x69,
	0x64, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x05, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x12, 0x23, 0x0a,
	0x0d, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x40, 0x0a, 0x11, 0x69, 0x73, 0x73,
	0x75, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x15,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x42, 0x65,
	0x74, 0x77, 0x65, 0x65, 0x6e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x52, 0x0f, 0x69, 0x73, 0x73, 0x75,
	0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x40, 0x0a, 0x11, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65,
	0x18, 0x16, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e,
	0x42, 0x65, 0x74, 0x77, 0x65, 0x65, 0x6e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x52, 0x0f, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x40, 0x0a,
	0x11, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x72, 0x61, 0x6e,
	0x67, 0x65, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x2e, 0x42, 0x65, 0x74, 0x77, 0x65, 0x65, 0x6e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x52, 0x0f,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12,
	0x40, 0x0a, 0x11, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x72,
	0x61, 0x6e, 0x67, 0x65, 0x18, 0x66, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x2e, 0x42, 0x65, 0x74, 0x77, 0x65, 0x65, 0x6e, 0x49, 0x6e, 0x74, 0x36, 0x34,
	0x52, 0x0f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67,
	0x65, 0x12, 0x40, 0x0a, 0x11, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x67, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x42, 0x65, 0x74, 0x77, 0x65, 0x65, 0x6e, 0x49, 0x6e, 0x74,
	0x36, 0x34, 0x52, 0x0f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61,
	0x6e, 0x67, 0x65, 0x22, 0xbd, 0x01, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x74, 0x68,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x26, 0x0a, 0x06, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x50, 0x61, 0x67, 0x65,
	0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x05, 0x73, 0x6f, 0x72, 0x74,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x2e, 0x53, 0x6f, 0x72, 0x74, 0x52, 0x05, 0x73, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x35, 0x0a, 0x06,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x22, 0x93, 0x01, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x74, 0x68,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2b, 0x0a, 0x06, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52,
	0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e,
	0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x9a, 0x01, 0x0a, 0x16, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x12, 0x26, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x07,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x07, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x71, 0x0a, 0x15, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x2b, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x9f, 0x01, 0x0a, 0x1b, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x26, 0x0a, 0x06, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x12, 0x2b, 0x0a, 0x07, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x11, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x12, 0x2b,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x49, 0x0a, 0x1a, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2b, 0x0a, 0x06, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x42, 0x21, 0x5a, 0x1f, 0x73, 0x75, 0x70, 0x65, 0x72, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_auth_token_proto_rawDescOnce sync.Once
	file_auth_token_proto_rawDescData = file_auth_token_proto_rawDesc
)

func file_auth_token_proto_rawDescGZIP() []byte {
	file_auth_token_proto_rawDescOnce.Do(func() {
		file_auth_token_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_token_proto_rawDescData)
	})
	return file_auth_token_proto_rawDescData
}

var file_auth_token_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_auth_token_proto_goTypes = []interface{}{
	(*AuthToken)(nil),                   // 0: auth_service.AuthToken
	(*AuthTokenFilter)(nil),             // 1: auth_service.AuthTokenFilter
	(*ListAuthTokenParams)(nil),         // 2: auth_service.ListAuthTokenParams
	(*ListAuthTokenReply)(nil),          // 3: auth_service.ListAuthTokenReply
	(*OperateAuthTokenParams)(nil),      // 4: auth_service.OperateAuthTokenParams
	(*OperateAuthTokenReply)(nil),       // 5: auth_service.OperateAuthTokenReply
	(*BatchOperateAuthTokenParams)(nil), // 6: auth_service.BatchOperateAuthTokenParams
	(*BatchOperateAuthTokenReply)(nil),  // 7: auth_service.BatchOperateAuthTokenReply
	(*common.BetweenInt64)(nil),         // 8: public.BetweenInt64
	(*common.Header)(nil),               // 9: public.Header
	(*common.Pager)(nil),                // 10: public.Pager
	(*common.Sort)(nil),                 // 11: public.Sort
	(*common.ReplyHeader)(nil),          // 12: public.ReplyHeader
	(common.Operation)(0),               // 13: public.Operation
}
var file_auth_token_proto_depIdxs = []int32{
	8,  // 0: auth_service.AuthTokenFilter.issued_time_range:type_name -> public.BetweenInt64
	8,  // 1: auth_service.AuthTokenFilter.expire_time_range:type_name -> public.BetweenInt64
	8,  // 2: auth_service.AuthTokenFilter.delete_time_range:type_name -> public.BetweenInt64
	8,  // 3: auth_service.AuthTokenFilter.create_time_range:type_name -> public.BetweenInt64
	8,  // 4: auth_service.AuthTokenFilter.update_time_range:type_name -> public.BetweenInt64
	9,  // 5: auth_service.ListAuthTokenParams.header:type_name -> public.Header
	10, // 6: auth_service.ListAuthTokenParams.pager:type_name -> public.Pager
	11, // 7: auth_service.ListAuthTokenParams.sorts:type_name -> public.Sort
	1,  // 8: auth_service.ListAuthTokenParams.filter:type_name -> auth_service.AuthTokenFilter
	12, // 9: auth_service.ListAuthTokenReply.header:type_name -> public.ReplyHeader
	10, // 10: auth_service.ListAuthTokenReply.pager:type_name -> public.Pager
	0,  // 11: auth_service.ListAuthTokenReply.data:type_name -> auth_service.AuthToken
	9,  // 12: auth_service.OperateAuthTokenParams.header:type_name -> public.Header
	13, // 13: auth_service.OperateAuthTokenParams.operate:type_name -> public.Operation
	0,  // 14: auth_service.OperateAuthTokenParams.data:type_name -> auth_service.AuthToken
	12, // 15: auth_service.OperateAuthTokenReply.header:type_name -> public.ReplyHeader
	0,  // 16: auth_service.OperateAuthTokenReply.data:type_name -> auth_service.AuthToken
	9,  // 17: auth_service.BatchOperateAuthTokenParams.header:type_name -> public.Header
	13, // 18: auth_service.BatchOperateAuthTokenParams.operate:type_name -> public.Operation
	0,  // 19: auth_service.BatchOperateAuthTokenParams.data:type_name -> auth_service.AuthToken
	12, // 20: auth_service.BatchOperateAuthTokenReply.header:type_name -> public.ReplyHeader
	21, // [21:21] is the sub-list for method output_type
	21, // [21:21] is the sub-list for method input_type
	21, // [21:21] is the sub-list for extension type_name
	21, // [21:21] is the sub-list for extension extendee
	0,  // [0:21] is the sub-list for field type_name
}

func init() { file_auth_token_proto_init() }
func file_auth_token_proto_init() {
	if File_auth_token_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auth_token_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthToken); i {
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
		file_auth_token_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthTokenFilter); i {
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
		file_auth_token_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAuthTokenParams); i {
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
		file_auth_token_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAuthTokenReply); i {
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
		file_auth_token_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperateAuthTokenParams); i {
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
		file_auth_token_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperateAuthTokenReply); i {
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
		file_auth_token_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchOperateAuthTokenParams); i {
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
		file_auth_token_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchOperateAuthTokenReply); i {
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
			RawDescriptor: file_auth_token_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_auth_token_proto_goTypes,
		DependencyIndexes: file_auth_token_proto_depIdxs,
		MessageInfos:      file_auth_token_proto_msgTypes,
	}.Build()
	File_auth_token_proto = out.File
	file_auth_token_proto_rawDesc = nil
	file_auth_token_proto_goTypes = nil
	file_auth_token_proto_depIdxs = nil
}
