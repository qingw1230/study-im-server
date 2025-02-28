package base_info

import "github.com/qingw1230/study-im-server/pkg/proto/sdkws"

type GetConversationListReq struct {
	FromUserId string `json:"fromUserId" binding:"required"`
}

type GetConversationListResp struct {
	CommonResp
	ConversationList []*sdkws.ConversationInfo `json:"-"`
}

type CreateConversationReq struct {
	OwnerUserId      string `json:"ownerUserId" binding:"required"`
	ConversationId   string `json:"conversationId"`
	ConversationType int32  `json:"conversationType" binding:"required"`
	ConversationName string `json:"conversationName" binding:"required"`
	UserId           string `json:"userId"`
	GroupId          string `json:"groupId"`
}

type CreateConversationResp struct {
	CommonResp
}
