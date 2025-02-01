package base_info

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
