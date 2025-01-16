package base_info

type AddFriendReq struct {
	FromUserID string `json:"fromUserID" binding:"required"`
	ToUserID   string `json:"toUserID" binding:"required"`
	ReqMsg     string `json:"reqMsg"`
}
