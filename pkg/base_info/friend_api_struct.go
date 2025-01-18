package base_info

type AddFriendReq struct {
	FromUserID string `json:"fromUserID" binding:"required"`
	ToUserID   string `json:"toUserID" binding:"required"`
	ReqMsg     string `json:"reqMsg"`
}

type AddFriendResp struct {
	CommonResp
}

type AddFriendResponseReq struct {
	FromUserID   string `json:"fromUserID" binding:"required"`
	ToUserID     string `json:"toUserID" binding:"required"`
	HandleResult int32  `json:"handleResult" binding:"required,oneof=-1 0 1"`
	HandleMsg    string `json:"handleMsg"`
}

type AddFriendResponseResp struct {
	CommonResp
}
