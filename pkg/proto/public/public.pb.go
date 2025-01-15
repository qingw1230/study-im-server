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

type UserRegisterInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	NickName string `protobuf:"bytes,3,opt,name=nickName,proto3" json:"nickName,omitempty"`
}

func (x *UserRegisterInfo) Reset() {
	*x = UserRegisterInfo{}
	mi := &file_public_public_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserRegisterInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRegisterInfo) ProtoMessage() {}

func (x *UserRegisterInfo) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use UserRegisterInfo.ProtoReflect.Descriptor instead.
func (*UserRegisterInfo) Descriptor() ([]byte, []int) {
	return file_public_public_proto_rawDescGZIP(), []int{1}
}

func (x *UserRegisterInfo) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserRegisterInfo) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserRegisterInfo) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

type UserLoginInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *UserLoginInfo) Reset() {
	*x = UserLoginInfo{}
	mi := &file_public_public_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserLoginInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLoginInfo) ProtoMessage() {}

func (x *UserLoginInfo) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use UserLoginInfo.ProtoReflect.Descriptor instead.
func (*UserLoginInfo) Descriptor() ([]byte, []int) {
	return file_public_public_proto_rawDescGZIP(), []int{2}
}

func (x *UserLoginInfo) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserLoginInfo) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type UserLoginSuccessInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token             string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Admin             bool   `protobuf:"varint,2,opt,name=admin,proto3" json:"admin,omitempty"`
	UserID            string `protobuf:"bytes,3,opt,name=userID,proto3" json:"userID,omitempty"`
	NickName          string `protobuf:"bytes,4,opt,name=nickName,proto3" json:"nickName,omitempty"`
	PersonalSignature string `protobuf:"bytes,5,opt,name=personalSignature,proto3" json:"personalSignature,omitempty"`
	JoinType          int32  `protobuf:"varint,6,opt,name=joinType,proto3" json:"joinType,omitempty"`
	Sex               int32  `protobuf:"varint,7,opt,name=sex,proto3" json:"sex,omitempty"`
	AreaName          string `protobuf:"bytes,8,opt,name=areaName,proto3" json:"areaName,omitempty"`
	AreaCode          string `protobuf:"bytes,9,opt,name=areaCode,proto3" json:"areaCode,omitempty"`
}

func (x *UserLoginSuccessInfo) Reset() {
	*x = UserLoginSuccessInfo{}
	mi := &file_public_public_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserLoginSuccessInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLoginSuccessInfo) ProtoMessage() {}

func (x *UserLoginSuccessInfo) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use UserLoginSuccessInfo.ProtoReflect.Descriptor instead.
func (*UserLoginSuccessInfo) Descriptor() ([]byte, []int) {
	return file_public_public_proto_rawDescGZIP(), []int{3}
}

func (x *UserLoginSuccessInfo) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *UserLoginSuccessInfo) GetAdmin() bool {
	if x != nil {
		return x.Admin
	}
	return false
}

func (x *UserLoginSuccessInfo) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *UserLoginSuccessInfo) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *UserLoginSuccessInfo) GetPersonalSignature() string {
	if x != nil {
		return x.PersonalSignature
	}
	return ""
}

func (x *UserLoginSuccessInfo) GetJoinType() int32 {
	if x != nil {
		return x.JoinType
	}
	return 0
}

func (x *UserLoginSuccessInfo) GetSex() int32 {
	if x != nil {
		return x.Sex
	}
	return 0
}

func (x *UserLoginSuccessInfo) GetAreaName() string {
	if x != nil {
		return x.AreaName
	}
	return ""
}

func (x *UserLoginSuccessInfo) GetAreaCode() string {
	if x != nil {
		return x.AreaCode
	}
	return ""
}

type PublicUserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID   string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	NickName string `protobuf:"bytes,2,opt,name=nickName,proto3" json:"nickName,omitempty"`
	Sex      int32  `protobuf:"varint,3,opt,name=sex,proto3" json:"sex,omitempty"`
	AreaName string `protobuf:"bytes,4,opt,name=areaName,proto3" json:"areaName,omitempty"`
	AreaCode string `protobuf:"bytes,5,opt,name=areaCode,proto3" json:"areaCode,omitempty"`
}

func (x *PublicUserInfo) Reset() {
	*x = PublicUserInfo{}
	mi := &file_public_public_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PublicUserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicUserInfo) ProtoMessage() {}

func (x *PublicUserInfo) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use PublicUserInfo.ProtoReflect.Descriptor instead.
func (*PublicUserInfo) Descriptor() ([]byte, []int) {
	return file_public_public_proto_rawDescGZIP(), []int{4}
}

func (x *PublicUserInfo) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *PublicUserInfo) GetNickName() string {
	if x != nil {
		return x.NickName
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

var File_public_public_proto protoreflect.FileDescriptor

var file_public_public_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x70,
	0x69, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x4c, 0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x60, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x41, 0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x8a, 0x02, 0x0a, 0x14,
	0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c,
	0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x11, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6a, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6a, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x73, 0x65, 0x78, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x73, 0x65, 0x78,
	0x12, 0x1a, 0x0a, 0x08, 0x61, 0x72, 0x65, 0x61, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x61, 0x72, 0x65, 0x61, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x61, 0x72, 0x65, 0x61, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x61, 0x72, 0x65, 0x61, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x8e, 0x01, 0x0a, 0x0e, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x73, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x73, 0x65,
	0x78, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x72, 0x65, 0x61, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x72, 0x65, 0x61, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x61, 0x72, 0x65, 0x61, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x61, 0x72, 0x65, 0x61, 0x43, 0x6f, 0x64, 0x65, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x71, 0x69, 0x6e, 0x67, 0x77, 0x31, 0x32, 0x33,
	0x30, 0x2f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x2d, 0x69, 0x6d, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_public_public_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_public_public_proto_goTypes = []any{
	(*CommonResp)(nil),           // 0: server_api_params.CommonResp
	(*UserRegisterInfo)(nil),     // 1: server_api_params.UserRegisterInfo
	(*UserLoginInfo)(nil),        // 2: server_api_params.UserLoginInfo
	(*UserLoginSuccessInfo)(nil), // 3: server_api_params.UserLoginSuccessInfo
	(*PublicUserInfo)(nil),       // 4: server_api_params.PublicUserInfo
}
var file_public_public_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
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
			NumMessages:   5,
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
