// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.12.4
// source: public/public.proto

package public

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

type CommonResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Code   int32  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Info   string `protobuf:"bytes,3,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *CommonResp) Reset() {
	*x = CommonResp{}
	mi := &file_public_public_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CommonResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonResp) ProtoMessage() {}

func (x *CommonResp) ProtoReflect() protoreflect.Message {
	mi := &file_public_public_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonResp.ProtoReflect.Descriptor instead.
func (*CommonResp) Descriptor() ([]byte, []int) {
	return file_public_public_proto_rawDescGZIP(), []int{0}
}

func (x *CommonResp) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *CommonResp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CommonResp) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token             string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Admin             bool   `protobuf:"varint,2,opt,name=admin,proto3" json:"admin,omitempty"`
	UserId            string `protobuf:"bytes,3,opt,name=userId,proto3" json:"userId,omitempty"`
	Password          string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	NickName          string `protobuf:"bytes,5,opt,name=nickName,proto3" json:"nickName,omitempty"`
	PersonalSignature string `protobuf:"bytes,6,opt,name=personalSignature,proto3" json:"personalSignature,omitempty"`
	JoinType          int32  `protobuf:"varint,7,opt,name=joinType,proto3" json:"joinType,omitempty"`
	Sex               int32  `protobuf:"varint,8,opt,name=sex,proto3" json:"sex,omitempty"`
	AreaName          string `protobuf:"bytes,9,opt,name=areaName,proto3" json:"areaName,omitempty"`
	AreaCode          string `protobuf:"bytes,10,opt,name=areaCode,proto3" json:"areaCode,omitempty"`
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	mi := &file_public_public_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_public_public_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_public_public_proto_rawDescGZIP(), []int{1}
}

func (x *UserInfo) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *UserInfo) GetAdmin() bool {
	if x != nil {
		return x.Admin
	}
	return false
}

func (x *UserInfo) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserInfo) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserInfo) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *UserInfo) GetPersonalSignature() string {
	if x != nil {
		return x.PersonalSignature
	}
	return ""
}

func (x *UserInfo) GetJoinType() int32 {
	if x != nil {
		return x.JoinType
	}
	return 0
}

func (x *UserInfo) GetSex() int32 {
	if x != nil {
		return x.Sex
	}
	return 0
}

func (x *UserInfo) GetAreaName() string {
	if x != nil {
		return x.AreaName
	}
	return ""
}

func (x *UserInfo) GetAreaCode() string {
	if x != nil {
		return x.AreaCode
	}
	return ""
}

type PublicUserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId            string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	NickName          string `protobuf:"bytes,2,opt,name=nickName,proto3" json:"nickName,omitempty"`
	PersonalSignature string `protobuf:"bytes,3,opt,name=personalSignature,proto3" json:"personalSignature,omitempty"`
	Sex               int32  `protobuf:"varint,4,opt,name=sex,proto3" json:"sex,omitempty"`
	AreaName          string `protobuf:"bytes,5,opt,name=areaName,proto3" json:"areaName,omitempty"`
	AreaCode          string `protobuf:"bytes,6,opt,name=areaCode,proto3" json:"areaCode,omitempty"`
	IsFriend          bool   `protobuf:"varint,7,opt,name=isFriend,proto3" json:"isFriend,omitempty"`
}

func (x *PublicUserInfo) Reset() {
	*x = PublicUserInfo{}
	mi := &file_public_public_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PublicUserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicUserInfo) ProtoMessage() {}

func (x *PublicUserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_public_public_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicUserInfo.ProtoReflect.Descriptor instead.
func (*PublicUserInfo) Descriptor() ([]byte, []int) {
	return file_public_public_proto_rawDescGZIP(), []int{2}
}

func (x *PublicUserInfo) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *PublicUserInfo) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *PublicUserInfo) GetPersonalSignature() string {
	if x != nil {
		return x.PersonalSignature
	}
	return ""
}

func (x *PublicUserInfo) GetSex() int32 {
	if x != nil {
		return x.Sex
	}
	return 0
}

func (x *PublicUserInfo) GetAreaName() string {
	if x != nil {
		return x.AreaName
	}
	return ""
}

func (x *PublicUserInfo) GetAreaCode() string {
	if x != nil {
		return x.AreaCode
	}
	return ""
}

func (x *PublicUserInfo) GetIsFriend() bool {
	if x != nil {
		return x.IsFriend
	}
	return false
}

type FriendInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerUserId string          `protobuf:"bytes,1,opt,name=ownerUserId,proto3" json:"ownerUserId,omitempty"`
	Remark      string          `protobuf:"bytes,2,opt,name=remark,proto3" json:"remark,omitempty"`
	CreateTime  uint32          `protobuf:"varint,3,opt,name=createTime,proto3" json:"createTime,omitempty"`
	FriendInfo  *PublicUserInfo `protobuf:"bytes,4,opt,name=friendInfo,proto3" json:"friendInfo,omitempty"`
	AddSource   int32           `protobuf:"varint,5,opt,name=addSource,proto3" json:"addSource,omitempty"`
	OpUserId    string          `protobuf:"bytes,6,opt,name=opUserId,proto3" json:"opUserId,omitempty"`
}

func (x *FriendInfo) Reset() {
	*x = FriendInfo{}
	mi := &file_public_public_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FriendInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendInfo) ProtoMessage() {}

func (x *FriendInfo) ProtoReflect() protoreflect.Message {
	mi := &file_public_public_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendInfo.ProtoReflect.Descriptor instead.
func (*FriendInfo) Descriptor() ([]byte, []int) {
	return file_public_public_proto_rawDescGZIP(), []int{3}
}

func (x *FriendInfo) GetOwnerUserId() string {
	if x != nil {
		return x.OwnerUserId
	}
	return ""
}

func (x *FriendInfo) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *FriendInfo) GetCreateTime() uint32 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *FriendInfo) GetFriendInfo() *PublicUserInfo {
	if x != nil {
		return x.FriendInfo
	}
	return nil
}

func (x *FriendInfo) GetAddSource() int32 {
	if x != nil {
		return x.AddSource
	}
	return 0
}

func (x *FriendInfo) GetOpUserId() string {
	if x != nil {
		return x.OpUserId
	}
	return ""
}

type GroupInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId      string `protobuf:"bytes,1,opt,name=groupId,proto3" json:"groupId,omitempty"`
	GroupName    string `protobuf:"bytes,2,opt,name=groupName,proto3" json:"groupName,omitempty"`
	FaceUrl      string `protobuf:"bytes,3,opt,name=faceUrl,proto3" json:"faceUrl,omitempty"`
	Introduction string `protobuf:"bytes,4,opt,name=introduction,proto3" json:"introduction,omitempty"`
	Notifacation string `protobuf:"bytes,5,opt,name=notifacation,proto3" json:"notifacation,omitempty"`
	OwnerUserId  string `protobuf:"bytes,6,opt,name=ownerUserId,proto3" json:"ownerUserId,omitempty"`
	CreateTime   uint32 `protobuf:"varint,7,opt,name=createTime,proto3" json:"createTime,omitempty"`
	MemberCount  uint32 `protobuf:"varint,8,opt,name=memberCount,proto3" json:"memberCount,omitempty"`
	Ex           string `protobuf:"bytes,9,opt,name=ex,proto3" json:"ex,omitempty"`
	Status       int32  `protobuf:"varint,10,opt,name=status,proto3" json:"status,omitempty"`
	CreateUserId string `protobuf:"bytes,11,opt,name=createUserId,proto3" json:"createUserId,omitempty"`
	GroupType    int32  `protobuf:"varint,12,opt,name=groupType,proto3" json:"groupType,omitempty"`
}

func (x *GroupInfo) Reset() {
	*x = GroupInfo{}
	mi := &file_public_public_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GroupInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupInfo) ProtoMessage() {}

func (x *GroupInfo) ProtoReflect() protoreflect.Message {
	mi := &file_public_public_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupInfo.ProtoReflect.Descriptor instead.
func (*GroupInfo) Descriptor() ([]byte, []int) {
	return file_public_public_proto_rawDescGZIP(), []int{4}
}

func (x *GroupInfo) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *GroupInfo) GetGroupName() string {
	if x != nil {
		return x.GroupName
	}
	return ""
}

func (x *GroupInfo) GetFaceUrl() string {
	if x != nil {
		return x.FaceUrl
	}
	return ""
}

func (x *GroupInfo) GetIntroduction() string {
	if x != nil {
		return x.Introduction
	}
	return ""
}

func (x *GroupInfo) GetNotifacation() string {
	if x != nil {
		return x.Notifacation
	}
	return ""
}

func (x *GroupInfo) GetOwnerUserId() string {
	if x != nil {
		return x.OwnerUserId
	}
	return ""
}

func (x *GroupInfo) GetCreateTime() uint32 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *GroupInfo) GetMemberCount() uint32 {
	if x != nil {
		return x.MemberCount
	}
	return 0
}

func (x *GroupInfo) GetEx() string {
	if x != nil {
		return x.Ex
	}
	return ""
}

func (x *GroupInfo) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *GroupInfo) GetCreateUserId() string {
	if x != nil {
		return x.CreateUserId
	}
	return ""
}

func (x *GroupInfo) GetGroupType() int32 {
	if x != nil {
		return x.GroupType
	}
	return 0
}

type FriendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromUserId    string `protobuf:"bytes,1,opt,name=fromUserId,proto3" json:"fromUserId,omitempty"`
	FromNickName  string `protobuf:"bytes,2,opt,name=fromNickName,proto3" json:"fromNickName,omitempty"`
	FromFaceURL   string `protobuf:"bytes,3,opt,name=fromFaceURL,proto3" json:"fromFaceURL,omitempty"`
	FromSex       int32  `protobuf:"varint,4,opt,name=fromSex,proto3" json:"fromSex,omitempty"`
	ToUserId      string `protobuf:"bytes,5,opt,name=toUserId,proto3" json:"toUserId,omitempty"`
	ToNickName    string `protobuf:"bytes,6,opt,name=toNickName,proto3" json:"toNickName,omitempty"`
	ToFaceURL     string `protobuf:"bytes,7,opt,name=toFaceURL,proto3" json:"toFaceURL,omitempty"`
	ToSex         int32  `protobuf:"varint,8,opt,name=toSex,proto3" json:"toSex,omitempty"`
	HandleResult  int32  `protobuf:"varint,9,opt,name=handleResult,proto3" json:"handleResult,omitempty"`
	ReqMsg        string `protobuf:"bytes,10,opt,name=reqMsg,proto3" json:"reqMsg,omitempty"`
	CreateTime    uint32 `protobuf:"varint,11,opt,name=createTime,proto3" json:"createTime,omitempty"`
	HandlerUserId string `protobuf:"bytes,12,opt,name=handlerUserId,proto3" json:"handlerUserId,omitempty"`
	HandleMsg     string `protobuf:"bytes,13,opt,name=handleMsg,proto3" json:"handleMsg,omitempty"`
	HandleTime    uint32 `protobuf:"varint,14,opt,name=handleTime,proto3" json:"handleTime,omitempty"`
	Ex            string `protobuf:"bytes,15,opt,name=ex,proto3" json:"ex,omitempty"`
}

func (x *FriendRequest) Reset() {
	*x = FriendRequest{}
	mi := &file_public_public_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FriendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendRequest) ProtoMessage() {}

func (x *FriendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_public_public_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendRequest.ProtoReflect.Descriptor instead.
func (*FriendRequest) Descriptor() ([]byte, []int) {
	return file_public_public_proto_rawDescGZIP(), []int{5}
}

func (x *FriendRequest) GetFromUserId() string {
	if x != nil {
		return x.FromUserId
	}
	return ""
}

func (x *FriendRequest) GetFromNickName() string {
	if x != nil {
		return x.FromNickName
	}
	return ""
}

func (x *FriendRequest) GetFromFaceURL() string {
	if x != nil {
		return x.FromFaceURL
	}
	return ""
}

func (x *FriendRequest) GetFromSex() int32 {
	if x != nil {
		return x.FromSex
	}
	return 0
}

func (x *FriendRequest) GetToUserId() string {
	if x != nil {
		return x.ToUserId
	}
	return ""
}

func (x *FriendRequest) GetToNickName() string {
	if x != nil {
		return x.ToNickName
	}
	return ""
}

func (x *FriendRequest) GetToFaceURL() string {
	if x != nil {
		return x.ToFaceURL
	}
	return ""
}

func (x *FriendRequest) GetToSex() int32 {
	if x != nil {
		return x.ToSex
	}
	return 0
}

func (x *FriendRequest) GetHandleResult() int32 {
	if x != nil {
		return x.HandleResult
	}
	return 0
}

func (x *FriendRequest) GetReqMsg() string {
	if x != nil {
		return x.ReqMsg
	}
	return ""
}

func (x *FriendRequest) GetCreateTime() uint32 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *FriendRequest) GetHandlerUserId() string {
	if x != nil {
		return x.HandlerUserId
	}
	return ""
}

func (x *FriendRequest) GetHandleMsg() string {
	if x != nil {
		return x.HandleMsg
	}
	return ""
}

func (x *FriendRequest) GetHandleTime() uint32 {
	if x != nil {
		return x.HandleTime
	}
	return 0
}

func (x *FriendRequest) GetEx() string {
	if x != nil {
		return x.Ex
	}
	return ""
}

var File_public_public_proto protoreflect.FileDescriptor

var file_public_public_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x70,
	0x69, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x4c, 0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x9a, 0x02, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x12,
	0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x2c, 0x0a, 0x11, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x70, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x6a, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x6a, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x78,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x73, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x61,
	0x72, 0x65, 0x61, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61,
	0x72, 0x65, 0x61, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x72, 0x65, 0x61, 0x43,
	0x6f, 0x64, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x72, 0x65, 0x61, 0x43,
	0x6f, 0x64, 0x65, 0x22, 0xd8, 0x01, 0x0a, 0x0e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x70, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x53,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x78, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x73, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x72,
	0x65, 0x61, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x72,
	0x65, 0x61, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x72, 0x65, 0x61, 0x43, 0x6f,
	0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x72, 0x65, 0x61, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x22, 0xe3,
	0x01, 0x0a, 0x0a, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20, 0x0a,
	0x0b, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x41, 0x0a, 0x0a, 0x66, 0x72, 0x69, 0x65, 0x6e,
	0x64, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x2e,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a,
	0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x64,
	0x64, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x61,
	0x64, 0x64, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x70, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x70, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x22, 0xf3, 0x02, 0x0a, 0x09, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x61,
	0x63, 0x65, 0x55, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x61, 0x63,
	0x65, 0x55, 0x72, 0x6c, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x74, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x61, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x61, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0b, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x65, 0x78, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x65, 0x78,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x54, 0x79, 0x70, 0x65, 0x22, 0xcf, 0x03, 0x0a, 0x0d, 0x46,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a,
	0x66, 0x72, 0x6f, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x66, 0x72, 0x6f, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c,
	0x66, 0x72, 0x6f, 0x6d, 0x4e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x4e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x66, 0x72, 0x6f, 0x6d, 0x46, 0x61, 0x63, 0x65, 0x55, 0x52, 0x4c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x72, 0x6f, 0x6d, 0x46, 0x61, 0x63, 0x65, 0x55,
	0x52, 0x4c, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x72, 0x6f, 0x6d, 0x53, 0x65, 0x78, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x66, 0x72, 0x6f, 0x6d, 0x53, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08,
	0x74, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x74, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x4e, 0x69,
	0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x6f,
	0x4e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x6f, 0x46, 0x61,
	0x63, 0x65, 0x55, 0x52, 0x4c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x46,
	0x61, 0x63, 0x65, 0x55, 0x52, 0x4c, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x53, 0x65, 0x78, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x53, 0x65, 0x78, 0x12, 0x22, 0x0a, 0x0c,
	0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0c, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x71, 0x4d, 0x73, 0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x72, 0x65, 0x71, 0x4d, 0x73, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x68, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c,
	0x0a, 0x09, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x4d, 0x73, 0x67, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x4d, 0x73, 0x67, 0x12, 0x1e, 0x0a, 0x0a,
	0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0a, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x65, 0x78, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x65, 0x78, 0x42, 0x37, 0x5a, 0x35,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x71, 0x69, 0x6e, 0x67, 0x77,
	0x31, 0x32, 0x33, 0x30, 0x2f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x2d, 0x69, 0x6d, 0x2d, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_public_public_proto_rawDescOnce sync.Once
	file_public_public_proto_rawDescData = file_public_public_proto_rawDesc
)

func file_public_public_proto_rawDescGZIP() []byte {
	file_public_public_proto_rawDescOnce.Do(func() {
		file_public_public_proto_rawDescData = protoimpl.X.CompressGZIP(file_public_public_proto_rawDescData)
	})
	return file_public_public_proto_rawDescData
}

var file_public_public_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_public_public_proto_goTypes = []any{
	(*CommonResp)(nil),     // 0: server_api_params.CommonResp
	(*UserInfo)(nil),       // 1: server_api_params.UserInfo
	(*PublicUserInfo)(nil), // 2: server_api_params.PublicUserInfo
	(*FriendInfo)(nil),     // 3: server_api_params.FriendInfo
	(*GroupInfo)(nil),      // 4: server_api_params.GroupInfo
	(*FriendRequest)(nil),  // 5: server_api_params.FriendRequest
}
var file_public_public_proto_depIdxs = []int32{
	2, // 0: server_api_params.FriendInfo.friendInfo:type_name -> server_api_params.PublicUserInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_public_public_proto_init() }
func file_public_public_proto_init() {
	if File_public_public_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_public_public_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_public_public_proto_goTypes,
		DependencyIndexes: file_public_public_proto_depIdxs,
		MessageInfos:      file_public_public_proto_msgTypes,
	}.Build()
	File_public_public_proto = out.File
	file_public_public_proto_rawDesc = nil
	file_public_public_proto_goTypes = nil
	file_public_public_proto_depIdxs = nil
}
