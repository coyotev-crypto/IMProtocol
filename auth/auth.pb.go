// Copyright © 2023 OpenIM. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: auth/auth.proto

package auth

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetAdminTokenReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Secret        string                 `protobuf:"bytes,1,opt,name=secret,proto3" json:"secret"`
	UserID        string                 `protobuf:"bytes,3,opt,name=userID,proto3" json:"userID"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAdminTokenReq) Reset() {
	*x = GetAdminTokenReq{}
	mi := &file_auth_auth_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAdminTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAdminTokenReq) ProtoMessage() {}

func (x *GetAdminTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAdminTokenReq.ProtoReflect.Descriptor instead.
func (*GetAdminTokenReq) Descriptor() ([]byte, []int) {
	return file_auth_auth_proto_rawDescGZIP(), []int{0}
}

func (x *GetAdminTokenReq) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

func (x *GetAdminTokenReq) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type GetAdminTokenResp struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	Token             string                 `protobuf:"bytes,2,opt,name=token,proto3" json:"token"`
	ExpireTimeSeconds int64                  `protobuf:"varint,3,opt,name=expireTimeSeconds,proto3" json:"expireTimeSeconds"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *GetAdminTokenResp) Reset() {
	*x = GetAdminTokenResp{}
	mi := &file_auth_auth_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAdminTokenResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAdminTokenResp) ProtoMessage() {}

func (x *GetAdminTokenResp) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAdminTokenResp.ProtoReflect.Descriptor instead.
func (*GetAdminTokenResp) Descriptor() ([]byte, []int) {
	return file_auth_auth_proto_rawDescGZIP(), []int{1}
}

func (x *GetAdminTokenResp) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *GetAdminTokenResp) GetExpireTimeSeconds() int64 {
	if x != nil {
		return x.ExpireTimeSeconds
	}
	return 0
}

type ForceLogoutReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PlatformID    int32                  `protobuf:"varint,1,opt,name=platformID,proto3" json:"platformID"`
	UserID        string                 `protobuf:"bytes,2,opt,name=userID,proto3" json:"userID"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ForceLogoutReq) Reset() {
	*x = ForceLogoutReq{}
	mi := &file_auth_auth_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ForceLogoutReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForceLogoutReq) ProtoMessage() {}

func (x *ForceLogoutReq) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForceLogoutReq.ProtoReflect.Descriptor instead.
func (*ForceLogoutReq) Descriptor() ([]byte, []int) {
	return file_auth_auth_proto_rawDescGZIP(), []int{2}
}

func (x *ForceLogoutReq) GetPlatformID() int32 {
	if x != nil {
		return x.PlatformID
	}
	return 0
}

func (x *ForceLogoutReq) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type ForceLogoutResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ForceLogoutResp) Reset() {
	*x = ForceLogoutResp{}
	mi := &file_auth_auth_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ForceLogoutResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForceLogoutResp) ProtoMessage() {}

func (x *ForceLogoutResp) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForceLogoutResp.ProtoReflect.Descriptor instead.
func (*ForceLogoutResp) Descriptor() ([]byte, []int) {
	return file_auth_auth_proto_rawDescGZIP(), []int{3}
}

type ParseTokenReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ParseTokenReq) Reset() {
	*x = ParseTokenReq{}
	mi := &file_auth_auth_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ParseTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParseTokenReq) ProtoMessage() {}

func (x *ParseTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParseTokenReq.ProtoReflect.Descriptor instead.
func (*ParseTokenReq) Descriptor() ([]byte, []int) {
	return file_auth_auth_proto_rawDescGZIP(), []int{4}
}

func (x *ParseTokenReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type ParseTokenResp struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	UserID            string                 `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID"`
	PlatformID        int32                  `protobuf:"varint,2,opt,name=platformID,proto3" json:"platformID"`
	ExpireTimeSeconds int64                  `protobuf:"varint,4,opt,name=expireTimeSeconds,proto3" json:"expireTimeSeconds"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *ParseTokenResp) Reset() {
	*x = ParseTokenResp{}
	mi := &file_auth_auth_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ParseTokenResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParseTokenResp) ProtoMessage() {}

func (x *ParseTokenResp) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParseTokenResp.ProtoReflect.Descriptor instead.
func (*ParseTokenResp) Descriptor() ([]byte, []int) {
	return file_auth_auth_proto_rawDescGZIP(), []int{5}
}

func (x *ParseTokenResp) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *ParseTokenResp) GetPlatformID() int32 {
	if x != nil {
		return x.PlatformID
	}
	return 0
}

func (x *ParseTokenResp) GetExpireTimeSeconds() int64 {
	if x != nil {
		return x.ExpireTimeSeconds
	}
	return 0
}

type GetUserTokenReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PlatformID    int32                  `protobuf:"varint,1,opt,name=platformID,proto3" json:"platformID"`
	UserID        string                 `protobuf:"bytes,2,opt,name=userID,proto3" json:"userID"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserTokenReq) Reset() {
	*x = GetUserTokenReq{}
	mi := &file_auth_auth_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserTokenReq) ProtoMessage() {}

func (x *GetUserTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserTokenReq.ProtoReflect.Descriptor instead.
func (*GetUserTokenReq) Descriptor() ([]byte, []int) {
	return file_auth_auth_proto_rawDescGZIP(), []int{6}
}

func (x *GetUserTokenReq) GetPlatformID() int32 {
	if x != nil {
		return x.PlatformID
	}
	return 0
}

func (x *GetUserTokenReq) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type GetUserTokenResp struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	Token             string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token"`
	ExpireTimeSeconds int64                  `protobuf:"varint,2,opt,name=expireTimeSeconds,proto3" json:"expireTimeSeconds"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *GetUserTokenResp) Reset() {
	*x = GetUserTokenResp{}
	mi := &file_auth_auth_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserTokenResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserTokenResp) ProtoMessage() {}

func (x *GetUserTokenResp) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserTokenResp.ProtoReflect.Descriptor instead.
func (*GetUserTokenResp) Descriptor() ([]byte, []int) {
	return file_auth_auth_proto_rawDescGZIP(), []int{7}
}

func (x *GetUserTokenResp) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *GetUserTokenResp) GetExpireTimeSeconds() int64 {
	if x != nil {
		return x.ExpireTimeSeconds
	}
	return 0
}

type InvalidateTokenReq struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	PreservedToken string                 `protobuf:"bytes,1,opt,name=preservedToken,proto3" json:"preservedToken"`
	UserID         string                 `protobuf:"bytes,2,opt,name=userID,proto3" json:"userID"`
	PlatformID     int32                  `protobuf:"varint,3,opt,name=platformID,proto3" json:"platformID"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *InvalidateTokenReq) Reset() {
	*x = InvalidateTokenReq{}
	mi := &file_auth_auth_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InvalidateTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvalidateTokenReq) ProtoMessage() {}

func (x *InvalidateTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvalidateTokenReq.ProtoReflect.Descriptor instead.
func (*InvalidateTokenReq) Descriptor() ([]byte, []int) {
	return file_auth_auth_proto_rawDescGZIP(), []int{8}
}

func (x *InvalidateTokenReq) GetPreservedToken() string {
	if x != nil {
		return x.PreservedToken
	}
	return ""
}

func (x *InvalidateTokenReq) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *InvalidateTokenReq) GetPlatformID() int32 {
	if x != nil {
		return x.PlatformID
	}
	return 0
}

type InvalidateTokenResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InvalidateTokenResp) Reset() {
	*x = InvalidateTokenResp{}
	mi := &file_auth_auth_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InvalidateTokenResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvalidateTokenResp) ProtoMessage() {}

func (x *InvalidateTokenResp) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvalidateTokenResp.ProtoReflect.Descriptor instead.
func (*InvalidateTokenResp) Descriptor() ([]byte, []int) {
	return file_auth_auth_proto_rawDescGZIP(), []int{9}
}

type KickTokensReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Tokens        []string               `protobuf:"bytes,1,rep,name=tokens,proto3" json:"tokens"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *KickTokensReq) Reset() {
	*x = KickTokensReq{}
	mi := &file_auth_auth_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KickTokensReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KickTokensReq) ProtoMessage() {}

func (x *KickTokensReq) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KickTokensReq.ProtoReflect.Descriptor instead.
func (*KickTokensReq) Descriptor() ([]byte, []int) {
	return file_auth_auth_proto_rawDescGZIP(), []int{10}
}

func (x *KickTokensReq) GetTokens() []string {
	if x != nil {
		return x.Tokens
	}
	return nil
}

type KickTokensResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *KickTokensResp) Reset() {
	*x = KickTokensResp{}
	mi := &file_auth_auth_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KickTokensResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KickTokensResp) ProtoMessage() {}

func (x *KickTokensResp) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KickTokensResp.ProtoReflect.Descriptor instead.
func (*KickTokensResp) Descriptor() ([]byte, []int) {
	return file_auth_auth_proto_rawDescGZIP(), []int{11}
}

var File_auth_auth_proto protoreflect.FileDescriptor

const file_auth_auth_proto_rawDesc = "" +
	"\n" +
	"\x0fauth/auth.proto\x12\vopenim.auth\"B\n" +
	"\x10getAdminTokenReq\x12\x16\n" +
	"\x06secret\x18\x01 \x01(\tR\x06secret\x12\x16\n" +
	"\x06userID\x18\x03 \x01(\tR\x06userID\"W\n" +
	"\x11getAdminTokenResp\x12\x14\n" +
	"\x05token\x18\x02 \x01(\tR\x05token\x12,\n" +
	"\x11expireTimeSeconds\x18\x03 \x01(\x03R\x11expireTimeSeconds\"H\n" +
	"\x0eforceLogoutReq\x12\x1e\n" +
	"\n" +
	"platformID\x18\x01 \x01(\x05R\n" +
	"platformID\x12\x16\n" +
	"\x06userID\x18\x02 \x01(\tR\x06userID\"\x11\n" +
	"\x0fforceLogoutResp\"%\n" +
	"\rparseTokenReq\x12\x14\n" +
	"\x05token\x18\x01 \x01(\tR\x05token\"v\n" +
	"\x0eparseTokenResp\x12\x16\n" +
	"\x06userID\x18\x01 \x01(\tR\x06userID\x12\x1e\n" +
	"\n" +
	"platformID\x18\x02 \x01(\x05R\n" +
	"platformID\x12,\n" +
	"\x11expireTimeSeconds\x18\x04 \x01(\x03R\x11expireTimeSeconds\"I\n" +
	"\x0fgetUserTokenReq\x12\x1e\n" +
	"\n" +
	"platformID\x18\x01 \x01(\x05R\n" +
	"platformID\x12\x16\n" +
	"\x06userID\x18\x02 \x01(\tR\x06userID\"V\n" +
	"\x10getUserTokenResp\x12\x14\n" +
	"\x05token\x18\x01 \x01(\tR\x05token\x12,\n" +
	"\x11expireTimeSeconds\x18\x02 \x01(\x03R\x11expireTimeSeconds\"t\n" +
	"\x12invalidateTokenReq\x12&\n" +
	"\x0epreservedToken\x18\x01 \x01(\tR\x0epreservedToken\x12\x16\n" +
	"\x06userID\x18\x02 \x01(\tR\x06userID\x12\x1e\n" +
	"\n" +
	"platformID\x18\x03 \x01(\x05R\n" +
	"platformID\"\x15\n" +
	"\x13invalidateTokenResp\"'\n" +
	"\rkickTokensReq\x12\x16\n" +
	"\x06tokens\x18\x01 \x03(\tR\x06tokens\"\x10\n" +
	"\x0ekickTokensResp2\xd1\x03\n" +
	"\x04Auth\x12N\n" +
	"\rgetAdminToken\x12\x1d.openim.auth.getAdminTokenReq\x1a\x1e.openim.auth.getAdminTokenResp\x12K\n" +
	"\fgetUserToken\x12\x1c.openim.auth.getUserTokenReq\x1a\x1d.openim.auth.getUserTokenResp\x12H\n" +
	"\vforceLogout\x12\x1b.openim.auth.forceLogoutReq\x1a\x1c.openim.auth.forceLogoutResp\x12E\n" +
	"\n" +
	"parseToken\x12\x1a.openim.auth.parseTokenReq\x1a\x1b.openim.auth.parseTokenResp\x12T\n" +
	"\x0finvalidateToken\x12\x1f.openim.auth.invalidateTokenReq\x1a .openim.auth.invalidateTokenResp\x12E\n" +
	"\n" +
	"kickTokens\x12\x1a.openim.auth.kickTokensReq\x1a\x1b.openim.auth.kickTokensRespB+Z)github.com/coyotev-crypto/IMProtocol/authb\x06proto3"

var (
	file_auth_auth_proto_rawDescOnce sync.Once
	file_auth_auth_proto_rawDescData []byte
)

func file_auth_auth_proto_rawDescGZIP() []byte {
	file_auth_auth_proto_rawDescOnce.Do(func() {
		file_auth_auth_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_auth_auth_proto_rawDesc), len(file_auth_auth_proto_rawDesc)))
	})
	return file_auth_auth_proto_rawDescData
}

var file_auth_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_auth_auth_proto_goTypes = []any{
	(*GetAdminTokenReq)(nil),    // 0: openim.auth.getAdminTokenReq
	(*GetAdminTokenResp)(nil),   // 1: openim.auth.getAdminTokenResp
	(*ForceLogoutReq)(nil),      // 2: openim.auth.forceLogoutReq
	(*ForceLogoutResp)(nil),     // 3: openim.auth.forceLogoutResp
	(*ParseTokenReq)(nil),       // 4: openim.auth.parseTokenReq
	(*ParseTokenResp)(nil),      // 5: openim.auth.parseTokenResp
	(*GetUserTokenReq)(nil),     // 6: openim.auth.getUserTokenReq
	(*GetUserTokenResp)(nil),    // 7: openim.auth.getUserTokenResp
	(*InvalidateTokenReq)(nil),  // 8: openim.auth.invalidateTokenReq
	(*InvalidateTokenResp)(nil), // 9: openim.auth.invalidateTokenResp
	(*KickTokensReq)(nil),       // 10: openim.auth.kickTokensReq
	(*KickTokensResp)(nil),      // 11: openim.auth.kickTokensResp
}
var file_auth_auth_proto_depIdxs = []int32{
	0,  // 0: openim.auth.Auth.getAdminToken:input_type -> openim.auth.getAdminTokenReq
	6,  // 1: openim.auth.Auth.getUserToken:input_type -> openim.auth.getUserTokenReq
	2,  // 2: openim.auth.Auth.forceLogout:input_type -> openim.auth.forceLogoutReq
	4,  // 3: openim.auth.Auth.parseToken:input_type -> openim.auth.parseTokenReq
	8,  // 4: openim.auth.Auth.invalidateToken:input_type -> openim.auth.invalidateTokenReq
	10, // 5: openim.auth.Auth.kickTokens:input_type -> openim.auth.kickTokensReq
	1,  // 6: openim.auth.Auth.getAdminToken:output_type -> openim.auth.getAdminTokenResp
	7,  // 7: openim.auth.Auth.getUserToken:output_type -> openim.auth.getUserTokenResp
	3,  // 8: openim.auth.Auth.forceLogout:output_type -> openim.auth.forceLogoutResp
	5,  // 9: openim.auth.Auth.parseToken:output_type -> openim.auth.parseTokenResp
	9,  // 10: openim.auth.Auth.invalidateToken:output_type -> openim.auth.invalidateTokenResp
	11, // 11: openim.auth.Auth.kickTokens:output_type -> openim.auth.kickTokensResp
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_auth_auth_proto_init() }
func file_auth_auth_proto_init() {
	if File_auth_auth_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_auth_auth_proto_rawDesc), len(file_auth_auth_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_auth_proto_goTypes,
		DependencyIndexes: file_auth_auth_proto_depIdxs,
		MessageInfos:      file_auth_auth_proto_msgTypes,
	}.Build()
	File_auth_auth_proto = out.File
	file_auth_auth_proto_goTypes = nil
	file_auth_auth_proto_depIdxs = nil
}
