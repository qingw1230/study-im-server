package base_info

import "github.com/qingw1230/study-im-server/pkg/proto/public"

type UserSendMsgReq struct {
	SendId         string `json:"sendId" binding:"required"`
	SenderNickName string `json:"senderNickName"`
	SenderFaceUrl  string `json:"senderFaceUrl"`
	Data           struct {
		RecvId      string `json:"recvId" `
		GroupId     string `json:"groupId" `
		SessionType int32  `json:"sessionType" binding:"required"`
		MsgFrom     int32  `json:"msgFrom" binding:"required"`
		ContentType int32  `json:"contentType" binding:"required"`
		Content     string `json:"content" binding:"required"`
		CreateTime  int64  `json:"createTime" binding:"required"`
	}
}

type UserPullMsgBySeqListReq struct {
	ReqIdentifier int      `json:"reqIdentifier" binding:"required"`
	UserId        string   `json:"userId" binding:"required"`
	SeqList       []uint32 `json:"seqList"`
}

type UserPullMsgBySeqListResp struct {
	CommonResp
	ReqIdentifier int               `json:"reqIdentifier" binding:"required"`
	MsgDataList   []*public.MsgData `json:"-"`
}
