// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.12.4
// source: sdkws/sdkws.proto

package sdkws

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
	mi := &file_sdkws_sdkws_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CommonResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonResp) ProtoMessage() {}

func (x *CommonResp) ProtoReflect() protoreflect.Message {
	mi := &file_sdkws_sdkws_proto_msgTypes[0]
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
	return file_sdkws_sdkws_proto_rawDescGZIP(), []int{0}
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
	mi := &file_sdkws_sdkws_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_sdkws_sdkws_proto_msgTypes[1]
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
	return file_sdkws_sdkws_proto_rawDescGZIP(), []int{1}
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
	FaceUrl           string `protobuf:"bytes,3,opt,name=faceUrl,proto3" json:"faceUrl,omitempty"`
	PersonalSignature string `protobuf:"bytes,4,opt,name=personalSignature,proto3" json:"personalSignature,omitempty"`
	Sex               int32  `protobuf:"varint,5,opt,name=sex,proto3" json:"sex,omitempty"`
	AreaName          string `protobuf:"bytes,6,opt,name=areaName,proto3" json:"areaName,omitempty"`
	AreaCode          string `protobuf:"bytes,7,opt,name=areaCode,proto3" json:"areaCode,omitempty"`
	IsFriend          bool   `protobuf:"varint,8,opt,name=isFriend,proto3" json:"isFriend,omitempty"`
}

func (x *PublicUserInfo) Reset() {
	*x = PublicUserInfo{}
	mi := &file_sdkws_sdkws_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PublicUserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicUserInfo) ProtoMessage() {}

func (x *PublicUserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_sdkws_sdkws_proto_msgTypes[2]
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
	return file_sdkws_sdkws_proto_rawDescGZIP(), []int{2}
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

func (x *PublicUserInfo) GetFaceUrl() string {
	if x != nil {
		return x.FaceUrl
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

type GroupInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId                string `protobuf:"bytes,1,opt,name=groupId,proto3" json:"groupId,omitempty"`
	GroupName              string `protobuf:"bytes,2,opt,name=groupName,proto3" json:"groupName,omitempty"`
	Notification           string `protobuf:"bytes,3,opt,name=notification,proto3" json:"notification,omitempty"`
	Introduction           string `protobuf:"bytes,4,opt,name=introduction,proto3" json:"introduction,omitempty"`
	FaceUrl                string `protobuf:"bytes,5,opt,name=faceUrl,proto3" json:"faceUrl,omitempty"`
	OwnerUserId            string `protobuf:"bytes,6,opt,name=ownerUserId,proto3" json:"ownerUserId,omitempty"`
	CreateTime             int64  `protobuf:"varint,7,opt,name=createTime,proto3" json:"createTime,omitempty"`
	MemberCount            uint32 `protobuf:"varint,8,opt,name=memberCount,proto3" json:"memberCount,omitempty"`
	Status                 int32  `protobuf:"varint,9,opt,name=status,proto3" json:"status,omitempty"`
	CreateUserId           string `protobuf:"bytes,10,opt,name=createUserId,proto3" json:"createUserId,omitempty"`
	GroupType              int32  `protobuf:"varint,11,opt,name=groupType,proto3" json:"groupType,omitempty"`
	NeedVerification       int32  `protobuf:"varint,12,opt,name=needVerification,proto3" json:"needVerification,omitempty"`
	NotificationUpdateTime int64  `protobuf:"varint,13,opt,name=notificationUpdateTime,proto3" json:"notificationUpdateTime,omitempty"`
	NotificationUserId     string `protobuf:"bytes,14,opt,name=notificationUserId,proto3" json:"notificationUserId,omitempty"`
}

func (x *GroupInfo) Reset() {
	*x = GroupInfo{}
	mi := &file_sdkws_sdkws_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GroupInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupInfo) ProtoMessage() {}

func (x *GroupInfo) ProtoReflect() protoreflect.Message {
	mi := &file_sdkws_sdkws_proto_msgTypes[3]
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
	return file_sdkws_sdkws_proto_rawDescGZIP(), []int{3}
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

func (x *GroupInfo) GetNotification() string {
	if x != nil {
		return x.Notification
	}
	return ""
}

func (x *GroupInfo) GetIntroduction() string {
	if x != nil {
		return x.Introduction
	}
	return ""
}

func (x *GroupInfo) GetFaceUrl() string {
	if x != nil {
		return x.FaceUrl
	}
	return ""
}

func (x *GroupInfo) GetOwnerUserId() string {
	if x != nil {
		return x.OwnerUserId
	}
	return ""
}

func (x *GroupInfo) GetCreateTime() int64 {
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

func (x *GroupInfo) GetNeedVerification() int32 {
	if x != nil {
		return x.NeedVerification
	}
	return 0
}

func (x *GroupInfo) GetNotificationUpdateTime() int64 {
	if x != nil {
		return x.NotificationUpdateTime
	}
	return 0
}

func (x *GroupInfo) GetNotificationUserId() string {
	if x != nil {
		return x.NotificationUserId
	}
	return ""
}

type UserSendMsgResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerMsgId string `protobuf:"bytes,1,opt,name=serverMsgId,proto3" json:"serverMsgId,omitempty"`
	ClientMsgId string `protobuf:"bytes,2,opt,name=clientMsgId,proto3" json:"clientMsgId,omitempty"`
	SendTime    int64  `protobuf:"varint,3,opt,name=sendTime,proto3" json:"sendTime,omitempty"`
}

func (x *UserSendMsgResp) Reset() {
	*x = UserSendMsgResp{}
	mi := &file_sdkws_sdkws_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserSendMsgResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserSendMsgResp) ProtoMessage() {}

func (x *UserSendMsgResp) ProtoReflect() protoreflect.Message {
	mi := &file_sdkws_sdkws_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserSendMsgResp.ProtoReflect.Descriptor instead.
func (*UserSendMsgResp) Descriptor() ([]byte, []int) {
	return file_sdkws_sdkws_proto_rawDescGZIP(), []int{4}
}

func (x *UserSendMsgResp) GetServerMsgId() string {
	if x != nil {
		return x.ServerMsgId
	}
	return ""
}

func (x *UserSendMsgResp) GetClientMsgId() string {
	if x != nil {
		return x.ClientMsgId
	}
	return ""
}

func (x *UserSendMsgResp) GetSendTime() int64 {
	if x != nil {
		return x.SendTime
	}
	return 0
}

type MsgData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SendId         string `protobuf:"bytes,1,opt,name=sendId,proto3" json:"sendId,omitempty"`                 // client 设置
	RecvId         string `protobuf:"bytes,2,opt,name=recvId,proto3" json:"recvId,omitempty"`                 // client 设置
	GroupId        string `protobuf:"bytes,3,opt,name=groupId,proto3" json:"groupId,omitempty"`               // client 设置
	Seq            uint64 `protobuf:"varint,4,opt,name=seq,proto3" json:"seq,omitempty"`                      // rpc 设置
	ServerMsgId    string `protobuf:"bytes,5,opt,name=serverMsgId,proto3" json:"serverMsgId,omitempty"`       // rpc 设置 timestamp + sendId + randNum
	SenderNickName string `protobuf:"bytes,6,opt,name=senderNickName,proto3" json:"senderNickName,omitempty"` // client 设置
	SenderFaceUrl  string `protobuf:"bytes,7,opt,name=senderFaceUrl,proto3" json:"senderFaceUrl,omitempty"`   // client 设置
	SessionType    int32  `protobuf:"varint,9,opt,name=sessionType,proto3" json:"sessionType,omitempty"`      // client 设置 Single Group ...
	ContentType    int32  `protobuf:"varint,10,opt,name=contentType,proto3" json:"contentType,omitempty"`     // client 设置 Text Picture Voice Video File ...
	Content        string `protobuf:"bytes,11,opt,name=content,proto3" json:"content,omitempty"`              // client 设置
	MsgFrom        int32  `protobuf:"varint,8,opt,name=msgFrom,proto3" json:"msgFrom,omitempty"`              // client 设置 User Sys
	CreateTime     int64  `protobuf:"varint,13,opt,name=createTime,proto3" json:"createTime,omitempty"`       // client 设置
	SendTime       int64  `protobuf:"varint,12,opt,name=sendTime,proto3" json:"sendTime,omitempty"`           // rpc 设置 timestamp
	Status         int32  `protobuf:"varint,14,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *MsgData) Reset() {
	*x = MsgData{}
	mi := &file_sdkws_sdkws_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MsgData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgData) ProtoMessage() {}

func (x *MsgData) ProtoReflect() protoreflect.Message {
	mi := &file_sdkws_sdkws_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgData.ProtoReflect.Descriptor instead.
func (*MsgData) Descriptor() ([]byte, []int) {
	return file_sdkws_sdkws_proto_rawDescGZIP(), []int{5}
}

func (x *MsgData) GetSendId() string {
	if x != nil {
		return x.SendId
	}
	return ""
}

func (x *MsgData) GetRecvId() string {
	if x != nil {
		return x.RecvId
	}
	return ""
}

func (x *MsgData) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *MsgData) GetSeq() uint64 {
	if x != nil {
		return x.Seq
	}
	return 0
}

func (x *MsgData) GetServerMsgId() string {
	if x != nil {
		return x.ServerMsgId
	}
	return ""
}

func (x *MsgData) GetSenderNickName() string {
	if x != nil {
		return x.SenderNickName
	}
	return ""
}

func (x *MsgData) GetSenderFaceUrl() string {
	if x != nil {
		return x.SenderFaceUrl
	}
	return ""
}

func (x *MsgData) GetSessionType() int32 {
	if x != nil {
		return x.SessionType
	}
	return 0
}

func (x *MsgData) GetContentType() int32 {
	if x != nil {
		return x.ContentType
	}
	return 0
}

func (x *MsgData) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *MsgData) GetMsgFrom() int32 {
	if x != nil {
		return x.MsgFrom
	}
	return 0
}

func (x *MsgData) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *MsgData) GetSendTime() int64 {
	if x != nil {
		return x.SendTime
	}
	return 0
}

func (x *MsgData) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type ConversationInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerUserId      string `protobuf:"bytes,1,opt,name=ownerUserId,proto3" json:"ownerUserId,omitempty"`
	ConversationId   string `protobuf:"bytes,2,opt,name=conversationId,proto3" json:"conversationId,omitempty"`
	ConversationType int32  `protobuf:"varint,3,opt,name=conversationType,proto3" json:"conversationType,omitempty"`
	ConversationName string `protobuf:"bytes,4,opt,name=conversationName,proto3" json:"conversationName,omitempty"`
	UserId           string `protobuf:"bytes,5,opt,name=userId,proto3" json:"userId,omitempty"`
	GroupId          string `protobuf:"bytes,6,opt,name=groupId,proto3" json:"groupId,omitempty"`
	MemberCount      int32  `protobuf:"varint,7,opt,name=memberCount,proto3" json:"memberCount,omitempty"`
	NoReadCount      int32  `protobuf:"varint,8,opt,name=noReadCount,proto3" json:"noReadCount,omitempty"`
	TopType          int32  `protobuf:"varint,9,opt,name=topType,proto3" json:"topType,omitempty"`
	LastMessage      string `protobuf:"bytes,10,opt,name=lastMessage,proto3" json:"lastMessage,omitempty"`
	LastMessageTime  int64  `protobuf:"varint,11,opt,name=lastMessageTime,proto3" json:"lastMessageTime,omitempty"`
	Status           int32  `protobuf:"varint,12,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ConversationInfo) Reset() {
	*x = ConversationInfo{}
	mi := &file_sdkws_sdkws_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConversationInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConversationInfo) ProtoMessage() {}

func (x *ConversationInfo) ProtoReflect() protoreflect.Message {
	mi := &file_sdkws_sdkws_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConversationInfo.ProtoReflect.Descriptor instead.
func (*ConversationInfo) Descriptor() ([]byte, []int) {
	return file_sdkws_sdkws_proto_rawDescGZIP(), []int{6}
}

func (x *ConversationInfo) GetOwnerUserId() string {
	if x != nil {
		return x.OwnerUserId
	}
	return ""
}

func (x *ConversationInfo) GetConversationId() string {
	if x != nil {
		return x.ConversationId
	}
	return ""
}

func (x *ConversationInfo) GetConversationType() int32 {
	if x != nil {
		return x.ConversationType
	}
	return 0
}

func (x *ConversationInfo) GetConversationName() string {
	if x != nil {
		return x.ConversationName
	}
	return ""
}

func (x *ConversationInfo) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ConversationInfo) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *ConversationInfo) GetMemberCount() int32 {
	if x != nil {
		return x.MemberCount
	}
	return 0
}

func (x *ConversationInfo) GetNoReadCount() int32 {
	if x != nil {
		return x.NoReadCount
	}
	return 0
}

func (x *ConversationInfo) GetTopType() int32 {
	if x != nil {
		return x.TopType
	}
	return 0
}

func (x *ConversationInfo) GetLastMessage() string {
	if x != nil {
		return x.LastMessage
	}
	return ""
}

func (x *ConversationInfo) GetLastMessageTime() int64 {
	if x != nil {
		return x.LastMessageTime
	}
	return 0
}

func (x *ConversationInfo) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

var File_sdkws_sdkws_proto protoreflect.FileDescriptor

var file_sdkws_sdkws_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x64, 0x6b, 0x77, 0x73, 0x2f, 0x73, 0x64, 0x6b, 0x77, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x11, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x5f,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x4c, 0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x69, 0x6e, 0x66, 0x6f, 0x22, 0x9a, 0x02, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a,
	0x11, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x61, 0x6c, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6a,
	0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6a,
	0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x78, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x73, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x72, 0x65,
	0x61, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x72, 0x65,
	0x61, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x72, 0x65, 0x61, 0x43, 0x6f, 0x64,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x72, 0x65, 0x61, 0x43, 0x6f, 0x64,
	0x65, 0x22, 0xf2, 0x01, 0x0a, 0x0e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x61, 0x63, 0x65,
	0x55, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x61, 0x63, 0x65, 0x55,
	0x72, 0x6c, 0x12, 0x2c, 0x0a, 0x11, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x70,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x73,
	0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x72, 0x65, 0x61, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x72, 0x65, 0x61, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x61, 0x72, 0x65, 0x61, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x61, 0x72, 0x65, 0x61, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73,
	0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73,
	0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x22, 0xf7, 0x03, 0x0a, 0x09, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x1c,
	0x0a, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x22, 0x0a, 0x0c, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x61, 0x63, 0x65, 0x55, 0x72, 0x6c, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x61, 0x63, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x20,
	0x0a, 0x0b, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c,
	0x0a, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2a, 0x0a, 0x10,
	0x6e, 0x65, 0x65, 0x64, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x6e, 0x65, 0x65, 0x64, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x16, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x16, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x2e, 0x0a, 0x12, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x71, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x73, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4d, 0x73, 0x67,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x4d, 0x73, 0x67, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d,
	0x73, 0x67, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x4d, 0x73, 0x67, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x22, 0xa1, 0x03, 0x0a, 0x07, 0x4d, 0x73, 0x67, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x65, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x63, 0x76, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x63, 0x76, 0x49, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x71,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x73, 0x65, 0x71, 0x12, 0x20, 0x0a, 0x0b, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x4d, 0x73, 0x67, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4d, 0x73, 0x67, 0x49, 0x64, 0x12, 0x26, 0x0a,
	0x0e, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x4e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x4e, 0x69, 0x63,
	0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x46,
	0x61, 0x63, 0x65, 0x55, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x46, 0x61, 0x63, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0b, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x73, 0x67,
	0x46, 0x72, 0x6f, 0x6d, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d, 0x73, 0x67, 0x46,
	0x72, 0x6f, 0x6d, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xa8, 0x03, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20, 0x0a, 0x0b,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x26,
	0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72,
	0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x10, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x6f,
	0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64,
	0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x6e, 0x6f, 0x52, 0x65, 0x61, 0x64, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6e, 0x6f, 0x52, 0x65, 0x61, 0x64, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x6f, 0x70, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x74, 0x6f, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x28, 0x0a, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x71, 0x69, 0x6e, 0x67, 0x77, 0x31, 0x32, 0x33, 0x30, 0x2f, 0x73, 0x74, 0x75, 0x64, 0x79,
	0x2d, 0x69, 0x6d, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x64, 0x6b, 0x77, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_sdkws_sdkws_proto_rawDescOnce sync.Once
	file_sdkws_sdkws_proto_rawDescData = file_sdkws_sdkws_proto_rawDesc
)

func file_sdkws_sdkws_proto_rawDescGZIP() []byte {
	file_sdkws_sdkws_proto_rawDescOnce.Do(func() {
		file_sdkws_sdkws_proto_rawDescData = protoimpl.X.CompressGZIP(file_sdkws_sdkws_proto_rawDescData)
	})
	return file_sdkws_sdkws_proto_rawDescData
}

var file_sdkws_sdkws_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_sdkws_sdkws_proto_goTypes = []any{
	(*CommonResp)(nil),       // 0: server_api_params.CommonResp
	(*UserInfo)(nil),         // 1: server_api_params.UserInfo
	(*PublicUserInfo)(nil),   // 2: server_api_params.PublicUserInfo
	(*GroupInfo)(nil),        // 3: server_api_params.GroupInfo
	(*UserSendMsgResp)(nil),  // 4: server_api_params.UserSendMsgResp
	(*MsgData)(nil),          // 5: server_api_params.MsgData
	(*ConversationInfo)(nil), // 6: server_api_params.ConversationInfo
}
var file_sdkws_sdkws_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sdkws_sdkws_proto_init() }
func file_sdkws_sdkws_proto_init() {
	if File_sdkws_sdkws_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sdkws_sdkws_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sdkws_sdkws_proto_goTypes,
		DependencyIndexes: file_sdkws_sdkws_proto_depIdxs,
		MessageInfos:      file_sdkws_sdkws_proto_msgTypes,
	}.Build()
	File_sdkws_sdkws_proto = out.File
	file_sdkws_sdkws_proto_rawDesc = nil
	file_sdkws_sdkws_proto_goTypes = nil
	file_sdkws_sdkws_proto_depIdxs = nil
}
