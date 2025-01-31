// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.12.4
// source: reply/reply.proto

package reply

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

type OnlinePushMsgReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MsgData      *public.MsgData `protobuf:"bytes,1,opt,name=msgData,proto3" json:"msgData,omitempty"`
	PushToUserId string          `protobuf:"bytes,2,opt,name=pushToUserId,proto3" json:"pushToUserId,omitempty"`
}

func (x *OnlinePushMsgReq) Reset() {
	*x = OnlinePushMsgReq{}
	mi := &file_reply_reply_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OnlinePushMsgReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnlinePushMsgReq) ProtoMessage() {}

func (x *OnlinePushMsgReq) ProtoReflect() protoreflect.Message {
	mi := &file_reply_reply_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnlinePushMsgReq.ProtoReflect.Descriptor instead.
func (*OnlinePushMsgReq) Descriptor() ([]byte, []int) {
	return file_reply_reply_proto_rawDescGZIP(), []int{0}
}

func (x *OnlinePushMsgReq) GetMsgData() *public.MsgData {
	if x != nil {
		return x.MsgData
	}
	return nil
}

func (x *OnlinePushMsgReq) GetPushToUserId() string {
	if x != nil {
		return x.PushToUserId
	}
	return ""
}

type OnlinePushMsgResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resp []*SingleMsgToUser `protobuf:"bytes,1,rep,name=resp,proto3" json:"resp,omitempty"`
}

func (x *OnlinePushMsgResp) Reset() {
	*x = OnlinePushMsgResp{}
	mi := &file_reply_reply_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OnlinePushMsgResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnlinePushMsgResp) ProtoMessage() {}

func (x *OnlinePushMsgResp) ProtoReflect() protoreflect.Message {
	mi := &file_reply_reply_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnlinePushMsgResp.ProtoReflect.Descriptor instead.
func (*OnlinePushMsgResp) Descriptor() ([]byte, []int) {
	return file_reply_reply_proto_rawDescGZIP(), []int{1}
}

func (x *OnlinePushMsgResp) GetResp() []*SingleMsgToUser {
	if x != nil {
		return x.Resp
	}
	return nil
}

type SingleMsgToUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResultCode int64  `protobuf:"varint,1,opt,name=ResultCode,proto3" json:"ResultCode,omitempty"`
	RecvId     string `protobuf:"bytes,2,opt,name=RecvId,proto3" json:"RecvId,omitempty"`
}

func (x *SingleMsgToUser) Reset() {
	*x = SingleMsgToUser{}
	mi := &file_reply_reply_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SingleMsgToUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SingleMsgToUser) ProtoMessage() {}

func (x *SingleMsgToUser) ProtoReflect() protoreflect.Message {
	mi := &file_reply_reply_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SingleMsgToUser.ProtoReflect.Descriptor instead.
func (*SingleMsgToUser) Descriptor() ([]byte, []int) {
	return file_reply_reply_proto_rawDescGZIP(), []int{2}
}

func (x *SingleMsgToUser) GetResultCode() int64 {
	if x != nil {
		return x.ResultCode
	}
	return 0
}

func (x *SingleMsgToUser) GetRecvId() string {
	if x != nil {
		return x.RecvId
	}
	return ""
}

var File_reply_reply_proto protoreflect.FileDescriptor

var file_reply_reply_proto_rawDesc = []byte{
	0x0a, 0x11, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x2f, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x62, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x1a, 0x13, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x6c, 0x0a, 0x10, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x50, 0x75, 0x73, 0x68, 0x4d,
	0x73, 0x67, 0x52, 0x65, 0x71, 0x12, 0x34, 0x0a, 0x07, 0x6d, 0x73, 0x67, 0x44, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f,
	0x61, 0x70, 0x69, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x2e, 0x4d, 0x73, 0x67, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x07, 0x6d, 0x73, 0x67, 0x44, 0x61, 0x74, 0x61, 0x12, 0x22, 0x0a, 0x0c, 0x70,
	0x75, 0x73, 0x68, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x70, 0x75, 0x73, 0x68, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x41, 0x0a, 0x11, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x73, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x2c, 0x0a, 0x04, 0x72, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x62, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x53, 0x69, 0x6e,
	0x67, 0x6c, 0x65, 0x4d, 0x73, 0x67, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x72, 0x65,
	0x73, 0x70, 0x22, 0x49, 0x0a, 0x0f, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x4d, 0x73, 0x67, 0x54,
	0x6f, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x63, 0x76, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x65, 0x63, 0x76, 0x49, 0x64, 0x32, 0x63, 0x0a,
	0x19, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x6c, 0x61, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x0d, 0x4f, 0x6e,
	0x6c, 0x69, 0x6e, 0x65, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x73, 0x67, 0x12, 0x19, 0x2e, 0x70, 0x62,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x50, 0x75, 0x73, 0x68,
	0x4d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x70, 0x62, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x2e, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x73, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x71, 0x69, 0x6e, 0x67, 0x77, 0x31, 0x32, 0x33, 0x30, 0x2f, 0x73, 0x74, 0x75, 0x64, 0x79,
	0x2d, 0x69, 0x6d, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_reply_reply_proto_rawDescOnce sync.Once
	file_reply_reply_proto_rawDescData = file_reply_reply_proto_rawDesc
)

func file_reply_reply_proto_rawDescGZIP() []byte {
	file_reply_reply_proto_rawDescOnce.Do(func() {
		file_reply_reply_proto_rawDescData = protoimpl.X.CompressGZIP(file_reply_reply_proto_rawDescData)
	})
	return file_reply_reply_proto_rawDescData
}

var file_reply_reply_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_reply_reply_proto_goTypes = []any{
	(*OnlinePushMsgReq)(nil),  // 0: pbReply.OnlinePushMsgReq
	(*OnlinePushMsgResp)(nil), // 1: pbReply.OnlinePushMsgResp
	(*SingleMsgToUser)(nil),   // 2: pbReply.SingleMsgToUser
	(*public.MsgData)(nil),    // 3: server_api_params.MsgData
}
var file_reply_reply_proto_depIdxs = []int32{
	3, // 0: pbReply.OnlinePushMsgReq.msgData:type_name -> server_api_params.MsgData
	2, // 1: pbReply.OnlinePushMsgResp.resp:type_name -> pbReply.SingleMsgToUser
	0, // 2: pbReply.OnlineMessageRelayService.OnlinePushMsg:input_type -> pbReply.OnlinePushMsgReq
	1, // 3: pbReply.OnlineMessageRelayService.OnlinePushMsg:output_type -> pbReply.OnlinePushMsgResp
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_reply_reply_proto_init() }
func file_reply_reply_proto_init() {
	if File_reply_reply_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_reply_reply_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_reply_reply_proto_goTypes,
		DependencyIndexes: file_reply_reply_proto_depIdxs,
		MessageInfos:      file_reply_reply_proto_msgTypes,
	}.Build()
	File_reply_reply_proto = out.File
	file_reply_reply_proto_rawDesc = nil
	file_reply_reply_proto_goTypes = nil
	file_reply_reply_proto_depIdxs = nil
}
