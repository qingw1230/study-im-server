package base_info

import "github.com/qingw1230/study-im-server/pkg/proto/public"

type SendMsgReq struct {
	SendId         string `json:"sendId" binding:"required"`
	RecvId         string `json:"recvId"`
	GroupId        string `json:"groupId"`
	SenderNickName string `json:"senderNickName" binding:"required"`
	SenderFaceUrl  string `json:"senderFaceUrl"`
	SessionType    int32  `json:"sessionType" binding:"required"`
	ContentType    int32  `json:"contentType" binding:"required"`
	Content        string `json:"content" binding:"required"`
	MsgFrom        int32  `json:"msgFrom" binding:"required"`
	CreateTime     int64  `json:"createTime" binding:"required"`
}

type SendMsgRespData struct {
	Seq         uint64 `json:"seq" binding:"required"`
	ServerMsgId string `json:"serverMsgId" binding:"required"`
	SendTime    int64  `json:"sendTime" binding:"required"`
}

type SendMsgResp struct {
	CommonResp
}

type UserGetNewestSeqReq struct {
	ReqIdentifier int    `json:"reqIdentifier" binding:"required"`
	UserId        string `json:"userId" binding:"required"`
}

type UserGetNewestSeqResp struct {
	CommonResp
	ReqIdentifier int    `json:"reqIdentifier" binding:"required"`
	NewestSeq     uint32 `json:"-"`
}

type UserPullMsgBySeqListReq struct {
	ReqIdentifier int      `json:"reqIdentifier" binding:"required"`
	UserId        string   `json:"userId" binding:"required"`
	SeqList       []uint64 `json:"seqList"`
}

type UserPullMsgBySeqListResp struct {
	CommonResp
	ReqIdentifier int               `json:"reqIdentifier" binding:"required"`
	MsgDataList   []*public.MsgData `json:"-"`
}
