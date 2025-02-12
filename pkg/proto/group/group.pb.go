// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.12.4
// source: group/group.proto

package group

import (
	public "github.com/qingw1230/study-im-server/pkg/proto/public"
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

type CreateGroupReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupInfo   *public.GroupInfo `protobuf:"bytes,1,opt,name=GroupInfo,proto3" json:"GroupInfo,omitempty"`
	OpUserId    string            `protobuf:"bytes,2,opt,name=OpUserId,proto3" json:"OpUserId,omitempty"`
	OwnerUserId string            `protobuf:"bytes,3,opt,name=OwnerUserId,proto3" json:"OwnerUserId,omitempty"`
}

func (x *CreateGroupReq) Reset() {
	*x = CreateGroupReq{}
	mi := &file_group_group_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateGroupReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGroupReq) ProtoMessage() {}

func (x *CreateGroupReq) ProtoReflect() protoreflect.Message {
	mi := &file_group_group_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGroupReq.ProtoReflect.Descriptor instead.
func (*CreateGroupReq) Descriptor() ([]byte, []int) {
	return file_group_group_proto_rawDescGZIP(), []int{0}
}

func (x *CreateGroupReq) GetGroupInfo() *public.GroupInfo {
	if x != nil {
		return x.GroupInfo
	}
	return nil
}

func (x *CreateGroupReq) GetOpUserId() string {
	if x != nil {
		return x.OpUserId
	}
	return ""
}

func (x *CreateGroupReq) GetOwnerUserId() string {
	if x != nil {
		return x.OwnerUserId
	}
	return ""
}

type CreateGroupResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommonResp *public.CommonResp `protobuf:"bytes,1,opt,name=CommonResp,proto3" json:"CommonResp,omitempty"`
}

func (x *CreateGroupResp) Reset() {
	*x = CreateGroupResp{}
	mi := &file_group_group_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateGroupResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGroupResp) ProtoMessage() {}

func (x *CreateGroupResp) ProtoReflect() protoreflect.Message {
	mi := &file_group_group_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGroupResp.ProtoReflect.Descriptor instead.
func (*CreateGroupResp) Descriptor() ([]byte, []int) {
	return file_group_group_proto_rawDescGZIP(), []int{1}
}

func (x *CreateGroupResp) GetCommonResp() *public.CommonResp {
	if x != nil {
		return x.CommonResp
	}
	return nil
}

type DeleteGroupReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId  string `protobuf:"bytes,1,opt,name=GroupId,proto3" json:"GroupId,omitempty"`
	OpUserId string `protobuf:"bytes,2,opt,name=OpUserId,proto3" json:"OpUserId,omitempty"`
}

func (x *DeleteGroupReq) Reset() {
	*x = DeleteGroupReq{}
	mi := &file_group_group_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteGroupReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGroupReq) ProtoMessage() {}

func (x *DeleteGroupReq) ProtoReflect() protoreflect.Message {
	mi := &file_group_group_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGroupReq.ProtoReflect.Descriptor instead.
func (*DeleteGroupReq) Descriptor() ([]byte, []int) {
	return file_group_group_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteGroupReq) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *DeleteGroupReq) GetOpUserId() string {
	if x != nil {
		return x.OpUserId
	}
	return ""
}

type DeleteGroupResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommonResp *public.CommonResp `protobuf:"bytes,1,opt,name=CommonResp,proto3" json:"CommonResp,omitempty"`
}

func (x *DeleteGroupResp) Reset() {
	*x = DeleteGroupResp{}
	mi := &file_group_group_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteGroupResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGroupResp) ProtoMessage() {}

func (x *DeleteGroupResp) ProtoReflect() protoreflect.Message {
	mi := &file_group_group_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGroupResp.ProtoReflect.Descriptor instead.
func (*DeleteGroupResp) Descriptor() ([]byte, []int) {
	return file_group_group_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteGroupResp) GetCommonResp() *public.CommonResp {
	if x != nil {
		return x.CommonResp
	}
	return nil
}

type QuitGroupReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId  string `protobuf:"bytes,1,opt,name=GroupId,proto3" json:"GroupId,omitempty"`
	OpUserId string `protobuf:"bytes,2,opt,name=opUserId,proto3" json:"opUserId,omitempty"`
}

func (x *QuitGroupReq) Reset() {
	*x = QuitGroupReq{}
	mi := &file_group_group_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *QuitGroupReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuitGroupReq) ProtoMessage() {}

func (x *QuitGroupReq) ProtoReflect() protoreflect.Message {
	mi := &file_group_group_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuitGroupReq.ProtoReflect.Descriptor instead.
func (*QuitGroupReq) Descriptor() ([]byte, []int) {
	return file_group_group_proto_rawDescGZIP(), []int{4}
}

func (x *QuitGroupReq) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *QuitGroupReq) GetOpUserId() string {
	if x != nil {
		return x.OpUserId
	}
	return ""
}

type QuitGroupResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommonResp *public.CommonResp `protobuf:"bytes,1,opt,name=CommonResp,proto3" json:"CommonResp,omitempty"`
}

func (x *QuitGroupResp) Reset() {
	*x = QuitGroupResp{}
	mi := &file_group_group_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *QuitGroupResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuitGroupResp) ProtoMessage() {}

func (x *QuitGroupResp) ProtoReflect() protoreflect.Message {
	mi := &file_group_group_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuitGroupResp.ProtoReflect.Descriptor instead.
func (*QuitGroupResp) Descriptor() ([]byte, []int) {
	return file_group_group_proto_rawDescGZIP(), []int{5}
}

func (x *QuitGroupResp) GetCommonResp() *public.CommonResp {
	if x != nil {
		return x.CommonResp
	}
	return nil
}

type GetJoinedGroupListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromUserId string `protobuf:"bytes,1,opt,name=FromUserId,proto3" json:"FromUserId,omitempty"`
	RoleLevel  int32  `protobuf:"varint,2,opt,name=RoleLevel,proto3" json:"RoleLevel,omitempty"`
	OpUserId   string `protobuf:"bytes,3,opt,name=OpUserId,proto3" json:"OpUserId,omitempty"`
}

func (x *GetJoinedGroupListReq) Reset() {
	*x = GetJoinedGroupListReq{}
	mi := &file_group_group_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetJoinedGroupListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetJoinedGroupListReq) ProtoMessage() {}

func (x *GetJoinedGroupListReq) ProtoReflect() protoreflect.Message {
	mi := &file_group_group_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetJoinedGroupListReq.ProtoReflect.Descriptor instead.
func (*GetJoinedGroupListReq) Descriptor() ([]byte, []int) {
	return file_group_group_proto_rawDescGZIP(), []int{6}
}

func (x *GetJoinedGroupListReq) GetFromUserId() string {
	if x != nil {
		return x.FromUserId
	}
	return ""
}

func (x *GetJoinedGroupListReq) GetRoleLevel() int32 {
	if x != nil {
		return x.RoleLevel
	}
	return 0
}

func (x *GetJoinedGroupListReq) GetOpUserId() string {
	if x != nil {
		return x.OpUserId
	}
	return ""
}

type GetJoinedGroupListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommonResp    *public.CommonResp  `protobuf:"bytes,1,opt,name=CommonResp,proto3" json:"CommonResp,omitempty"`
	GroupInfoList []*public.GroupInfo `protobuf:"bytes,2,rep,name=GroupInfoList,proto3" json:"GroupInfoList,omitempty"`
}

func (x *GetJoinedGroupListResp) Reset() {
	*x = GetJoinedGroupListResp{}
	mi := &file_group_group_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetJoinedGroupListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetJoinedGroupListResp) ProtoMessage() {}

func (x *GetJoinedGroupListResp) ProtoReflect() protoreflect.Message {
	mi := &file_group_group_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetJoinedGroupListResp.ProtoReflect.Descriptor instead.
func (*GetJoinedGroupListResp) Descriptor() ([]byte, []int) {
	return file_group_group_proto_rawDescGZIP(), []int{7}
}

func (x *GetJoinedGroupListResp) GetCommonResp() *public.CommonResp {
	if x != nil {
		return x.CommonResp
	}
	return nil
}

func (x *GetJoinedGroupListResp) GetGroupInfoList() []*public.GroupInfo {
	if x != nil {
		return x.GroupInfoList
	}
	return nil
}

type GetGroupInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId  string `protobuf:"bytes,1,opt,name=GroupId,proto3" json:"GroupId,omitempty"`
	OpUserId string `protobuf:"bytes,2,opt,name=OpUserId,proto3" json:"OpUserId,omitempty"`
}

func (x *GetGroupInfoReq) Reset() {
	*x = GetGroupInfoReq{}
	mi := &file_group_group_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetGroupInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupInfoReq) ProtoMessage() {}

func (x *GetGroupInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_group_group_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupInfoReq.ProtoReflect.Descriptor instead.
func (*GetGroupInfoReq) Descriptor() ([]byte, []int) {
	return file_group_group_proto_rawDescGZIP(), []int{8}
}

func (x *GetGroupInfoReq) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *GetGroupInfoReq) GetOpUserId() string {
	if x != nil {
		return x.OpUserId
	}
	return ""
}

type GetGroupInfoResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommonResp *public.CommonResp `protobuf:"bytes,1,opt,name=CommonResp,proto3" json:"CommonResp,omitempty"`
	GroupInfo  *public.GroupInfo  `protobuf:"bytes,2,opt,name=GroupInfo,proto3" json:"GroupInfo,omitempty"`
}

func (x *GetGroupInfoResp) Reset() {
	*x = GetGroupInfoResp{}
	mi := &file_group_group_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetGroupInfoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupInfoResp) ProtoMessage() {}

func (x *GetGroupInfoResp) ProtoReflect() protoreflect.Message {
	mi := &file_group_group_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupInfoResp.ProtoReflect.Descriptor instead.
func (*GetGroupInfoResp) Descriptor() ([]byte, []int) {
	return file_group_group_proto_rawDescGZIP(), []int{9}
}

func (x *GetGroupInfoResp) GetCommonResp() *public.CommonResp {
	if x != nil {
		return x.CommonResp
	}
	return nil
}

func (x *GetGroupInfoResp) GetGroupInfo() *public.GroupInfo {
	if x != nil {
		return x.GroupInfo
	}
	return nil
}

type SetGroupInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupInfo *public.GroupInfo `protobuf:"bytes,1,opt,name=GroupInfo,proto3" json:"GroupInfo,omitempty"`
	OpUserId  string            `protobuf:"bytes,2,opt,name=OpUserId,proto3" json:"OpUserId,omitempty"`
}

func (x *SetGroupInfoReq) Reset() {
	*x = SetGroupInfoReq{}
	mi := &file_group_group_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetGroupInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetGroupInfoReq) ProtoMessage() {}

func (x *SetGroupInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_group_group_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetGroupInfoReq.ProtoReflect.Descriptor instead.
func (*SetGroupInfoReq) Descriptor() ([]byte, []int) {
	return file_group_group_proto_rawDescGZIP(), []int{10}
}

func (x *SetGroupInfoReq) GetGroupInfo() *public.GroupInfo {
	if x != nil {
		return x.GroupInfo
	}
	return nil
}

func (x *SetGroupInfoReq) GetOpUserId() string {
	if x != nil {
		return x.OpUserId
	}
	return ""
}

type SetGroupInfoResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommonResp *public.CommonResp `protobuf:"bytes,1,opt,name=CommonResp,proto3" json:"CommonResp,omitempty"`
}

func (x *SetGroupInfoResp) Reset() {
	*x = SetGroupInfoResp{}
	mi := &file_group_group_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetGroupInfoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetGroupInfoResp) ProtoMessage() {}

func (x *SetGroupInfoResp) ProtoReflect() protoreflect.Message {
	mi := &file_group_group_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetGroupInfoResp.ProtoReflect.Descriptor instead.
func (*SetGroupInfoResp) Descriptor() ([]byte, []int) {
	return file_group_group_proto_rawDescGZIP(), []int{11}
}

func (x *SetGroupInfoResp) GetCommonResp() *public.CommonResp {
	if x != nil {
		return x.CommonResp
	}
	return nil
}

var File_group_group_proto protoreflect.FileDescriptor

var file_group_group_proto_rawDesc = []byte{
	0x0a, 0x11, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x62, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x1a, 0x13, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x8a, 0x01, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x65, 0x71, 0x12, 0x3a, 0x0a, 0x09, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66,
	0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x5f, 0x61, 0x70, 0x69, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x2e, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x1a, 0x0a, 0x08, 0x4f, 0x70, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x4f, 0x70, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b,
	0x4f, 0x77, 0x6e, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x50,
	0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x3d, 0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61,
	0x70, 0x69, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x52, 0x0a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x22, 0x46, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52,
	0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x4f, 0x70, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x4f, 0x70, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x50, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3d, 0x0a, 0x0a, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x52, 0x0a,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x44, 0x0a, 0x0c, 0x51, 0x75,
	0x69, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x70, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x70, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x4e, 0x0a, 0x0d, 0x51, 0x75, 0x69, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x3d, 0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61,
	0x70, 0x69, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x52, 0x0a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x22, 0x71, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x69, 0x6e, 0x65, 0x64, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1e, 0x0a, 0x0a, 0x46, 0x72, 0x6f,
	0x6d, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x46,
	0x72, 0x6f, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x52, 0x6f, 0x6c,
	0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x52, 0x6f,
	0x6c, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x4f, 0x70, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4f, 0x70, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x9b, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x69, 0x6e, 0x65,
	0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3d,
	0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x5f,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x52, 0x0a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x42, 0x0a,
	0x0d, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x70,
	0x69, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x0d, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73,
	0x74, 0x22, 0x47, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x4f, 0x70, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x4f, 0x70, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x8d, 0x01, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x3d, 0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69,
	0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x52, 0x0a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3a,
	0x0a, 0x09, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x09, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x69, 0x0a, 0x0f, 0x53, 0x65,
	0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x3a, 0x0a,
	0x09, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x4f, 0x70, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4f, 0x70, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x51, 0x0a, 0x10, 0x53, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3d, 0x0a, 0x0a, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x52, 0x0a, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x32, 0xa8, 0x03, 0x0a, 0x05, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x12, 0x40, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x70, 0x62, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x40, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x70,
	0x62, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3a, 0x0a, 0x09, 0x51, 0x75, 0x69, 0x74, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x51, 0x75,
	0x69, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x2e, 0x51, 0x75, 0x69, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x55, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x69, 0x6e, 0x65, 0x64, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1e, 0x2e, 0x70, 0x62, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x2e, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x69, 0x6e, 0x65, 0x64, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1f, 0x2e, 0x70, 0x62, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x2e, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x69, 0x6e, 0x65, 0x64, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x43, 0x0a, 0x0c, 0x47, 0x65, 0x74,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x2e, 0x70, 0x62, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x47, 0x65,
	0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x12, 0x43,
	0x0a, 0x0c, 0x53, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18,
	0x2e, 0x70, 0x62, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x53, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x2e, 0x53, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x71, 0x69, 0x6e, 0x67, 0x77, 0x31, 0x32, 0x33, 0x30, 0x2f, 0x73, 0x74, 0x75, 0x64,
	0x79, 0x2d, 0x69, 0x6d, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_group_group_proto_rawDescOnce sync.Once
	file_group_group_proto_rawDescData = file_group_group_proto_rawDesc
)

func file_group_group_proto_rawDescGZIP() []byte {
	file_group_group_proto_rawDescOnce.Do(func() {
		file_group_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_group_group_proto_rawDescData)
	})
	return file_group_group_proto_rawDescData
}

var file_group_group_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_group_group_proto_goTypes = []any{
	(*CreateGroupReq)(nil),         // 0: pbGroup.CreateGroupReq
	(*CreateGroupResp)(nil),        // 1: pbGroup.CreateGroupResp
	(*DeleteGroupReq)(nil),         // 2: pbGroup.DeleteGroupReq
	(*DeleteGroupResp)(nil),        // 3: pbGroup.DeleteGroupResp
	(*QuitGroupReq)(nil),           // 4: pbGroup.QuitGroupReq
	(*QuitGroupResp)(nil),          // 5: pbGroup.QuitGroupResp
	(*GetJoinedGroupListReq)(nil),  // 6: pbGroup.GetJoinedGroupListReq
	(*GetJoinedGroupListResp)(nil), // 7: pbGroup.GetJoinedGroupListResp
	(*GetGroupInfoReq)(nil),        // 8: pbGroup.GetGroupInfoReq
	(*GetGroupInfoResp)(nil),       // 9: pbGroup.GetGroupInfoResp
	(*SetGroupInfoReq)(nil),        // 10: pbGroup.SetGroupInfoReq
	(*SetGroupInfoResp)(nil),       // 11: pbGroup.SetGroupInfoResp
	(*public.GroupInfo)(nil),       // 12: server_api_params.GroupInfo
	(*public.CommonResp)(nil),      // 13: server_api_params.CommonResp
}
var file_group_group_proto_depIdxs = []int32{
	12, // 0: pbGroup.CreateGroupReq.GroupInfo:type_name -> server_api_params.GroupInfo
	13, // 1: pbGroup.CreateGroupResp.CommonResp:type_name -> server_api_params.CommonResp
	13, // 2: pbGroup.DeleteGroupResp.CommonResp:type_name -> server_api_params.CommonResp
	13, // 3: pbGroup.QuitGroupResp.CommonResp:type_name -> server_api_params.CommonResp
	13, // 4: pbGroup.GetJoinedGroupListResp.CommonResp:type_name -> server_api_params.CommonResp
	12, // 5: pbGroup.GetJoinedGroupListResp.GroupInfoList:type_name -> server_api_params.GroupInfo
	13, // 6: pbGroup.GetGroupInfoResp.CommonResp:type_name -> server_api_params.CommonResp
	12, // 7: pbGroup.GetGroupInfoResp.GroupInfo:type_name -> server_api_params.GroupInfo
	12, // 8: pbGroup.SetGroupInfoReq.GroupInfo:type_name -> server_api_params.GroupInfo
	13, // 9: pbGroup.SetGroupInfoResp.CommonResp:type_name -> server_api_params.CommonResp
	0,  // 10: pbGroup.Group.CreateGroup:input_type -> pbGroup.CreateGroupReq
	2,  // 11: pbGroup.Group.DeleteGroup:input_type -> pbGroup.DeleteGroupReq
	4,  // 12: pbGroup.Group.QuitGroup:input_type -> pbGroup.QuitGroupReq
	6,  // 13: pbGroup.Group.GetJoinedGroupList:input_type -> pbGroup.GetJoinedGroupListReq
	8,  // 14: pbGroup.Group.GetGroupInfo:input_type -> pbGroup.GetGroupInfoReq
	10, // 15: pbGroup.Group.SetGroupInfo:input_type -> pbGroup.SetGroupInfoReq
	1,  // 16: pbGroup.Group.CreateGroup:output_type -> pbGroup.CreateGroupResp
	3,  // 17: pbGroup.Group.DeleteGroup:output_type -> pbGroup.DeleteGroupResp
	5,  // 18: pbGroup.Group.QuitGroup:output_type -> pbGroup.QuitGroupResp
	7,  // 19: pbGroup.Group.GetJoinedGroupList:output_type -> pbGroup.GetJoinedGroupListResp
	9,  // 20: pbGroup.Group.GetGroupInfo:output_type -> pbGroup.GetGroupInfoResp
	11, // 21: pbGroup.Group.SetGroupInfo:output_type -> pbGroup.SetGroupInfoResp
	16, // [16:22] is the sub-list for method output_type
	10, // [10:16] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_group_group_proto_init() }
func file_group_group_proto_init() {
	if File_group_group_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_group_group_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_group_group_proto_goTypes,
		DependencyIndexes: file_group_group_proto_depIdxs,
		MessageInfos:      file_group_group_proto_msgTypes,
	}.Build()
	File_group_group_proto = out.File
	file_group_group_proto_rawDesc = nil
	file_group_group_proto_goTypes = nil
	file_group_group_proto_depIdxs = nil
}
